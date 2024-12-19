/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	apiv1alpha1 "inference.networking.x-k8s.io/llm-instance-gateway/api/v1alpha1"
	versioned "inference.networking.x-k8s.io/llm-instance-gateway/client-go/clientset/versioned"
	internalinterfaces "inference.networking.x-k8s.io/llm-instance-gateway/client-go/informers/externalversions/internalinterfaces"
	v1alpha1 "inference.networking.x-k8s.io/llm-instance-gateway/client-go/listers/api/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// InferencePoolInformer provides access to a shared informer and lister for
// InferencePools.
type InferencePoolInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.InferencePoolLister
}

type inferencePoolInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewInferencePoolInformer constructs a new informer for InferencePool type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewInferencePoolInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredInferencePoolInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredInferencePoolInformer constructs a new informer for InferencePool type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredInferencePoolInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ApiV1alpha1().InferencePools(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ApiV1alpha1().InferencePools(namespace).Watch(context.TODO(), options)
			},
		},
		&apiv1alpha1.InferencePool{},
		resyncPeriod,
		indexers,
	)
}

func (f *inferencePoolInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredInferencePoolInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *inferencePoolInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apiv1alpha1.InferencePool{}, f.defaultInformer)
}

func (f *inferencePoolInformer) Lister() v1alpha1.InferencePoolLister {
	return v1alpha1.NewInferencePoolLister(f.Informer().GetIndexer())
}