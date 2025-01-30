//go:build !ignore_autogenerated

/*
Copyright 2024 The Kubernetes Authors.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointPickerConfig) DeepCopyInto(out *EndpointPickerConfig) {
	*out = *in
	if in.Extension != nil {
		in, out := &in.Extension, &out.Extension
		*out = new(ExtensionConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointPickerConfig.
func (in *EndpointPickerConfig) DeepCopy() *EndpointPickerConfig {
	if in == nil {
		return nil
	}
	out := new(EndpointPickerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtensionConfig) DeepCopyInto(out *ExtensionConfig) {
	*out = *in
	if in.ExtensionRef != nil {
		in, out := &in.ExtensionRef, &out.ExtensionRef
		*out = new(ExtensionReference)
		(*in).DeepCopyInto(*out)
	}
	in.ExtensionConnection.DeepCopyInto(&out.ExtensionConnection)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtensionConfig.
func (in *ExtensionConfig) DeepCopy() *ExtensionConfig {
	if in == nil {
		return nil
	}
	out := new(ExtensionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtensionConnection) DeepCopyInto(out *ExtensionConnection) {
	*out = *in
	if in.FailureMode != nil {
		in, out := &in.FailureMode, &out.FailureMode
		*out = new(ExtensionFailureMode)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtensionConnection.
func (in *ExtensionConnection) DeepCopy() *ExtensionConnection {
	if in == nil {
		return nil
	}
	out := new(ExtensionConnection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtensionReference) DeepCopyInto(out *ExtensionReference) {
	*out = *in
	if in.Group != nil {
		in, out := &in.Group, &out.Group
		*out = new(string)
		**out = **in
	}
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(string)
		**out = **in
	}
	if in.TargetPortNumber != nil {
		in, out := &in.TargetPortNumber, &out.TargetPortNumber
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtensionReference.
func (in *ExtensionReference) DeepCopy() *ExtensionReference {
	if in == nil {
		return nil
	}
	out := new(ExtensionReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InferenceModel) DeepCopyInto(out *InferenceModel) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InferenceModel.
func (in *InferenceModel) DeepCopy() *InferenceModel {
	if in == nil {
		return nil
	}
	out := new(InferenceModel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InferenceModel) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InferenceModelList) DeepCopyInto(out *InferenceModelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InferenceModel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InferenceModelList.
func (in *InferenceModelList) DeepCopy() *InferenceModelList {
	if in == nil {
		return nil
	}
	out := new(InferenceModelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InferenceModelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InferenceModelSpec) DeepCopyInto(out *InferenceModelSpec) {
	*out = *in
	if in.Criticality != nil {
		in, out := &in.Criticality, &out.Criticality
		*out = new(Criticality)
		**out = **in
	}
	if in.TargetModels != nil {
		in, out := &in.TargetModels, &out.TargetModels
		*out = make([]TargetModel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.PoolRef = in.PoolRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InferenceModelSpec.
func (in *InferenceModelSpec) DeepCopy() *InferenceModelSpec {
	if in == nil {
		return nil
	}
	out := new(InferenceModelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InferenceModelStatus) DeepCopyInto(out *InferenceModelStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InferenceModelStatus.
func (in *InferenceModelStatus) DeepCopy() *InferenceModelStatus {
	if in == nil {
		return nil
	}
	out := new(InferenceModelStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InferencePool) DeepCopyInto(out *InferencePool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InferencePool.
func (in *InferencePool) DeepCopy() *InferencePool {
	if in == nil {
		return nil
	}
	out := new(InferencePool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InferencePool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InferencePoolList) DeepCopyInto(out *InferencePoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InferencePool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InferencePoolList.
func (in *InferencePoolList) DeepCopy() *InferencePoolList {
	if in == nil {
		return nil
	}
	out := new(InferencePoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InferencePoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InferencePoolSpec) DeepCopyInto(out *InferencePoolSpec) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = make(map[LabelKey]LabelValue, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.EndpointPickerConfig.DeepCopyInto(&out.EndpointPickerConfig)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InferencePoolSpec.
func (in *InferencePoolSpec) DeepCopy() *InferencePoolSpec {
	if in == nil {
		return nil
	}
	out := new(InferencePoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InferencePoolStatus) DeepCopyInto(out *InferencePoolStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InferencePoolStatus.
func (in *InferencePoolStatus) DeepCopy() *InferencePoolStatus {
	if in == nil {
		return nil
	}
	out := new(InferencePoolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoolObjectReference) DeepCopyInto(out *PoolObjectReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolObjectReference.
func (in *PoolObjectReference) DeepCopy() *PoolObjectReference {
	if in == nil {
		return nil
	}
	out := new(PoolObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetModel) DeepCopyInto(out *TargetModel) {
	*out = *in
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetModel.
func (in *TargetModel) DeepCopy() *TargetModel {
	if in == nil {
		return nil
	}
	out := new(TargetModel)
	in.DeepCopyInto(out)
	return out
}
