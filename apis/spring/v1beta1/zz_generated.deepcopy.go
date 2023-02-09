//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudApplicationLiveView) DeepCopyInto(out *CloudApplicationLiveView) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudApplicationLiveView.
func (in *CloudApplicationLiveView) DeepCopy() *CloudApplicationLiveView {
	if in == nil {
		return nil
	}
	out := new(CloudApplicationLiveView)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudApplicationLiveView) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudApplicationLiveViewList) DeepCopyInto(out *CloudApplicationLiveViewList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudApplicationLiveView, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudApplicationLiveViewList.
func (in *CloudApplicationLiveViewList) DeepCopy() *CloudApplicationLiveViewList {
	if in == nil {
		return nil
	}
	out := new(CloudApplicationLiveViewList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudApplicationLiveViewList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudApplicationLiveViewObservation) DeepCopyInto(out *CloudApplicationLiveViewObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudApplicationLiveViewObservation.
func (in *CloudApplicationLiveViewObservation) DeepCopy() *CloudApplicationLiveViewObservation {
	if in == nil {
		return nil
	}
	out := new(CloudApplicationLiveViewObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudApplicationLiveViewParameters) DeepCopyInto(out *CloudApplicationLiveViewParameters) {
	*out = *in
	if in.SpringCloudServiceID != nil {
		in, out := &in.SpringCloudServiceID, &out.SpringCloudServiceID
		*out = new(string)
		**out = **in
	}
	if in.SpringCloudServiceIDRef != nil {
		in, out := &in.SpringCloudServiceIDRef, &out.SpringCloudServiceIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.SpringCloudServiceIDSelector != nil {
		in, out := &in.SpringCloudServiceIDSelector, &out.SpringCloudServiceIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudApplicationLiveViewParameters.
func (in *CloudApplicationLiveViewParameters) DeepCopy() *CloudApplicationLiveViewParameters {
	if in == nil {
		return nil
	}
	out := new(CloudApplicationLiveViewParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudApplicationLiveViewSpec) DeepCopyInto(out *CloudApplicationLiveViewSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudApplicationLiveViewSpec.
func (in *CloudApplicationLiveViewSpec) DeepCopy() *CloudApplicationLiveViewSpec {
	if in == nil {
		return nil
	}
	out := new(CloudApplicationLiveViewSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudApplicationLiveViewStatus) DeepCopyInto(out *CloudApplicationLiveViewStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudApplicationLiveViewStatus.
func (in *CloudApplicationLiveViewStatus) DeepCopy() *CloudApplicationLiveViewStatus {
	if in == nil {
		return nil
	}
	out := new(CloudApplicationLiveViewStatus)
	in.DeepCopyInto(out)
	return out
}
