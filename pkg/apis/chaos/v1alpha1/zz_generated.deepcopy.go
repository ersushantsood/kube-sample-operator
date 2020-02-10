// +build !ignore_autogenerated

// Code generated by operator-sdk-v0.13.0-x86_64-linux-gnu. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChaosStatus) DeepCopyInto(out *ChaosStatus) {
	*out = *in
	in.Experiment.DeepCopyInto(&out.Experiment)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChaosStatus.
func (in *ChaosStatus) DeepCopy() *ChaosStatus {
	if in == nil {
		return nil
	}
	out := new(ChaosStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Chaosmaster) DeepCopyInto(out *Chaosmaster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Chaosmaster.
func (in *Chaosmaster) DeepCopy() *Chaosmaster {
	if in == nil {
		return nil
	}
	out := new(Chaosmaster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Chaosmaster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChaosmasterList) DeepCopyInto(out *ChaosmasterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Chaosmaster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChaosmasterList.
func (in *ChaosmasterList) DeepCopy() *ChaosmasterList {
	if in == nil {
		return nil
	}
	out := new(ChaosmasterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ChaosmasterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChaosmasterSpec) DeepCopyInto(out *ChaosmasterSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChaosmasterSpec.
func (in *ChaosmasterSpec) DeepCopy() *ChaosmasterSpec {
	if in == nil {
		return nil
	}
	out := new(ChaosmasterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChaosmasterStatus) DeepCopyInto(out *ChaosmasterStatus) {
	*out = *in
	in.ChaosStatus.DeepCopyInto(&out.ChaosStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChaosmasterStatus.
func (in *ChaosmasterStatus) DeepCopy() *ChaosmasterStatus {
	if in == nil {
		return nil
	}
	out := new(ChaosmasterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExperimentStatus) DeepCopyInto(out *ExperimentStatus) {
	*out = *in
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.EndTime != nil {
		in, out := &in.EndTime, &out.EndTime
		*out = (*in).DeepCopy()
	}
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make([]ChaosStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExperimentStatus.
func (in *ExperimentStatus) DeepCopy() *ExperimentStatus {
	if in == nil {
		return nil
	}
	out := new(ExperimentStatus)
	in.DeepCopyInto(out)
	return out
}
