//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityConfiguration) DeepCopyInto(out *SecurityConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityConfiguration.
func (in *SecurityConfiguration) DeepCopy() *SecurityConfiguration {
	if in == nil {
		return nil
	}
	out := new(SecurityConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecurityConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityConfigurationInitParameters) DeepCopyInto(out *SecurityConfigurationInitParameters) {
	*out = *in
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityConfigurationInitParameters.
func (in *SecurityConfigurationInitParameters) DeepCopy() *SecurityConfigurationInitParameters {
	if in == nil {
		return nil
	}
	out := new(SecurityConfigurationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityConfigurationList) DeepCopyInto(out *SecurityConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SecurityConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityConfigurationList.
func (in *SecurityConfigurationList) DeepCopy() *SecurityConfigurationList {
	if in == nil {
		return nil
	}
	out := new(SecurityConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecurityConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityConfigurationObservation) DeepCopyInto(out *SecurityConfigurationObservation) {
	*out = *in
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = new(string)
		**out = **in
	}
	if in.CreationDate != nil {
		in, out := &in.CreationDate, &out.CreationDate
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityConfigurationObservation.
func (in *SecurityConfigurationObservation) DeepCopy() *SecurityConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(SecurityConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityConfigurationParameters) DeepCopyInto(out *SecurityConfigurationParameters) {
	*out = *in
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityConfigurationParameters.
func (in *SecurityConfigurationParameters) DeepCopy() *SecurityConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(SecurityConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityConfigurationSpec) DeepCopyInto(out *SecurityConfigurationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityConfigurationSpec.
func (in *SecurityConfigurationSpec) DeepCopy() *SecurityConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(SecurityConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityConfigurationStatus) DeepCopyInto(out *SecurityConfigurationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityConfigurationStatus.
func (in *SecurityConfigurationStatus) DeepCopy() *SecurityConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(SecurityConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}
