package backend

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"go.uber.org/multierr"
	"inference.networking.x-k8s.io/gateway-api-inference-extension/api/v1alpha1"
	logutil "inference.networking.x-k8s.io/gateway-api-inference-extension/pkg/ext-proc/util/logging"
	corev1 "k8s.io/api/core/v1"
	klog "k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	fetchMetricsTimeout = 5 * time.Second
	poolNameKey         = "pool-name-key"
)

type podLister func() (map[string]string, error)

func NewProvider(pmc PodMetricsClient, datastore *K8sDatastore, client client.Client) *Provider {
	p := &Provider{
		podMetrics: sync.Map{},
		pmc:        pmc,
		datastore:  datastore,
		client:     client,
	}
	p.PodListerFunc = p.listPods
	return p
}

// Provider provides backend pods and information such as metrics.
type Provider struct {
	// key: Pod, value: *PodMetrics
	podMetrics    sync.Map
	pmc           PodMetricsClient
	datastore     *K8sDatastore
	client        client.Client
	PodListerFunc podLister
}

type PodMetricsClient interface {
	FetchMetrics(ctx context.Context, pod Pod, existing *PodMetrics) (*PodMetrics, error)
}

func (p *Provider) AllPodMetrics() []*PodMetrics {
	res := []*PodMetrics{}
	fn := func(k, v any) bool {
		res = append(res, v.(*PodMetrics))
		return true
	}
	p.podMetrics.Range(fn)
	return res
}

func (p *Provider) UpdatePodMetrics(pod Pod, pm *PodMetrics) {
	p.podMetrics.Store(pod, pm)
}

func (p *Provider) GetPodMetrics(pod Pod) (*PodMetrics, bool) {
	val, ok := p.podMetrics.Load(pod)
	if ok {
		return val.(*PodMetrics), true
	}
	return nil, false
}

func (p *Provider) Init(refreshPodsInterval, refreshMetricsInterval time.Duration) error {
	p.refreshPodsOnce()

	if err := p.refreshMetricsOnce(); err != nil {
		klog.Errorf("Failed to init metrics: %v", err)
	}

	klog.Infof("Initialized pods and metrics: %+v", p.AllPodMetrics())

	// periodically refresh pods
	go func() {
		for {
			time.Sleep(refreshPodsInterval)
			p.refreshPodsOnce()
		}
	}()

	// periodically refresh metrics
	go func() {
		for {
			time.Sleep(refreshMetricsInterval)
			if err := p.refreshMetricsOnce(); err != nil {
				klog.V(logutil.DEBUG).Infof("Failed to refresh metrics: %v", err)
			}
		}
	}()

	// Periodically print out the pods and metrics for DEBUGGING.
	if klog.V(logutil.DEBUG).Enabled() {
		go func() {
			for {
				time.Sleep(5 * time.Second)
				klog.Infof("===DEBUG: Current Pods and metrics: %+v", p.AllPodMetrics())
			}
		}()
	}

	return nil
}

// refreshPodsOnce lists pods and updates keys in the podMetrics map.
// Note this function doesn't update the PodMetrics value, it's done separately.
func (p *Provider) refreshPodsOnce() {
	listedPods, err := p.PodListerFunc()
	if err != nil {
		klog.Errorf("Failed to list pods: %v", err)
		return
	}
	// remove pods that don't exist any more.
	mergeFn := func(k, v any) bool {
		pod := k.(Pod)
		if _, ok := listedPods[pod.Name]; !ok {
			p.podMetrics.Delete(pod)
		}
		return true
	}
	p.podMetrics.Range(mergeFn)
	// add new pod to the map
	for name, address := range listedPods {
		pod := Pod{
			Name:    name,
			Address: address,
		}
		if _, ok := p.podMetrics.Load(pod); !ok {
			new := &PodMetrics{
				Pod: pod,
				Metrics: Metrics{
					ActiveModels: make(map[string]int),
				},
			}
			p.podMetrics.Store(pod, new)
		}
	}
}

func (p *Provider) refreshMetricsOnce() error {
	ctx, cancel := context.WithTimeout(context.Background(), fetchMetricsTimeout)
	defer cancel()
	start := time.Now()
	defer func() {
		d := time.Since(start)
		// TODO: add a metric instead of logging
		klog.V(logutil.DEBUG).Infof("Refreshed metrics in %v", d)
	}()
	var wg sync.WaitGroup
	errCh := make(chan error)
	processOnePod := func(key, value any) bool {
		klog.V(logutil.DEBUG).Infof("Processing pod %v and metric %v", key, value)
		pod := key.(Pod)
		existing := value.(*PodMetrics)
		wg.Add(1)
		go func() {
			defer wg.Done()
			updated, err := p.pmc.FetchMetrics(ctx, pod, existing)
			if err != nil {
				errCh <- fmt.Errorf("failed to parse metrics from %s: %v", pod, err)
				return
			}
			p.UpdatePodMetrics(pod, updated)
			klog.V(logutil.DEBUG).Infof("Updated metrics for pod %s: %v", pod, updated.Metrics)
		}()
		return true
	}
	p.podMetrics.Range(processOnePod)

	// Wait for metric collection for all pods to complete and close the error channel in a
	// goroutine so this is unblocking, allowing the code to proceed to the error collection code
	// below.
	// Note we couldn't use a buffered error channel with a size because the size of the podMetrics
	// sync.Map is unknown beforehand.
	go func() {
		wg.Wait()
		close(errCh)
	}()

	var errs error
	for err := range errCh {
		errs = multierr.Append(errs, err)
	}
	return errs
}

func (p *Provider) listPods() (map[string]string, error) {
	pool, err := p.datastore.getInferencePool()
	if err != nil {
		return nil, err
	}
	ps := make(map[string]string)
	var podList corev1.PodList
	if err := p.client.List(context.Background(), &podList, client.InNamespace(pool.Namespace), toMatchingLabels(pool.Spec.Selector)); err != nil {
		return nil, err
	}
	for i := range podList.Items {
		pod := &podList.Items[i]
		if podIsReady(pod) {
			ps[pod.Name] = pod.Status.PodIP + ":" + strconv.Itoa(int(pool.Spec.TargetPortNumber))
		}
	}
	return ps, nil
}

func toMatchingLabels(labels map[v1alpha1.LabelKey]v1alpha1.LabelValue) client.MatchingLabels {
	outMap := make(client.MatchingLabels)
	for k, v := range labels {
		outMap[string(k)] = string(v)
	}
	return outMap
}

func podIsReady(pod *corev1.Pod) bool {
	for _, condition := range pod.Status.Conditions {
		if condition.Type == corev1.PodReady {
			if condition.Status == corev1.ConditionTrue {
				return true
			}
			break
		}
	}
	return false
}
