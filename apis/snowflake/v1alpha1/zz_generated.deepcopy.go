//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Warehouse) DeepCopyInto(out *Warehouse) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Warehouse.
func (in *Warehouse) DeepCopy() *Warehouse {
	if in == nil {
		return nil
	}
	out := new(Warehouse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Warehouse) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WarehouseInitParameters) DeepCopyInto(out *WarehouseInitParameters) {
	*out = *in
	if in.AutoResume != nil {
		in, out := &in.AutoResume, &out.AutoResume
		*out = new(bool)
		**out = **in
	}
	if in.AutoSuspend != nil {
		in, out := &in.AutoSuspend, &out.AutoSuspend
		*out = new(float64)
		**out = **in
	}
	if in.Comment != nil {
		in, out := &in.Comment, &out.Comment
		*out = new(string)
		**out = **in
	}
	if in.EnableQueryAcceleration != nil {
		in, out := &in.EnableQueryAcceleration, &out.EnableQueryAcceleration
		*out = new(bool)
		**out = **in
	}
	if in.InitiallySuspended != nil {
		in, out := &in.InitiallySuspended, &out.InitiallySuspended
		*out = new(bool)
		**out = **in
	}
	if in.MaxClusterCount != nil {
		in, out := &in.MaxClusterCount, &out.MaxClusterCount
		*out = new(float64)
		**out = **in
	}
	if in.MaxConcurrencyLevel != nil {
		in, out := &in.MaxConcurrencyLevel, &out.MaxConcurrencyLevel
		*out = new(float64)
		**out = **in
	}
	if in.MinClusterCount != nil {
		in, out := &in.MinClusterCount, &out.MinClusterCount
		*out = new(float64)
		**out = **in
	}
	if in.QueryAccelerationMaxScaleFactor != nil {
		in, out := &in.QueryAccelerationMaxScaleFactor, &out.QueryAccelerationMaxScaleFactor
		*out = new(float64)
		**out = **in
	}
	if in.ResourceMonitor != nil {
		in, out := &in.ResourceMonitor, &out.ResourceMonitor
		*out = new(string)
		**out = **in
	}
	if in.ScalingPolicy != nil {
		in, out := &in.ScalingPolicy, &out.ScalingPolicy
		*out = new(string)
		**out = **in
	}
	if in.StatementQueuedTimeoutInSeconds != nil {
		in, out := &in.StatementQueuedTimeoutInSeconds, &out.StatementQueuedTimeoutInSeconds
		*out = new(float64)
		**out = **in
	}
	if in.StatementTimeoutInSeconds != nil {
		in, out := &in.StatementTimeoutInSeconds, &out.StatementTimeoutInSeconds
		*out = new(float64)
		**out = **in
	}
	if in.WaitForProvisioning != nil {
		in, out := &in.WaitForProvisioning, &out.WaitForProvisioning
		*out = new(bool)
		**out = **in
	}
	if in.WarehouseSize != nil {
		in, out := &in.WarehouseSize, &out.WarehouseSize
		*out = new(string)
		**out = **in
	}
	if in.WarehouseType != nil {
		in, out := &in.WarehouseType, &out.WarehouseType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WarehouseInitParameters.
func (in *WarehouseInitParameters) DeepCopy() *WarehouseInitParameters {
	if in == nil {
		return nil
	}
	out := new(WarehouseInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WarehouseList) DeepCopyInto(out *WarehouseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Warehouse, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WarehouseList.
func (in *WarehouseList) DeepCopy() *WarehouseList {
	if in == nil {
		return nil
	}
	out := new(WarehouseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WarehouseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WarehouseObservation) DeepCopyInto(out *WarehouseObservation) {
	*out = *in
	if in.AutoResume != nil {
		in, out := &in.AutoResume, &out.AutoResume
		*out = new(bool)
		**out = **in
	}
	if in.AutoSuspend != nil {
		in, out := &in.AutoSuspend, &out.AutoSuspend
		*out = new(float64)
		**out = **in
	}
	if in.Comment != nil {
		in, out := &in.Comment, &out.Comment
		*out = new(string)
		**out = **in
	}
	if in.EnableQueryAcceleration != nil {
		in, out := &in.EnableQueryAcceleration, &out.EnableQueryAcceleration
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.InitiallySuspended != nil {
		in, out := &in.InitiallySuspended, &out.InitiallySuspended
		*out = new(bool)
		**out = **in
	}
	if in.MaxClusterCount != nil {
		in, out := &in.MaxClusterCount, &out.MaxClusterCount
		*out = new(float64)
		**out = **in
	}
	if in.MaxConcurrencyLevel != nil {
		in, out := &in.MaxConcurrencyLevel, &out.MaxConcurrencyLevel
		*out = new(float64)
		**out = **in
	}
	if in.MinClusterCount != nil {
		in, out := &in.MinClusterCount, &out.MinClusterCount
		*out = new(float64)
		**out = **in
	}
	if in.QueryAccelerationMaxScaleFactor != nil {
		in, out := &in.QueryAccelerationMaxScaleFactor, &out.QueryAccelerationMaxScaleFactor
		*out = new(float64)
		**out = **in
	}
	if in.ResourceMonitor != nil {
		in, out := &in.ResourceMonitor, &out.ResourceMonitor
		*out = new(string)
		**out = **in
	}
	if in.ScalingPolicy != nil {
		in, out := &in.ScalingPolicy, &out.ScalingPolicy
		*out = new(string)
		**out = **in
	}
	if in.StatementQueuedTimeoutInSeconds != nil {
		in, out := &in.StatementQueuedTimeoutInSeconds, &out.StatementQueuedTimeoutInSeconds
		*out = new(float64)
		**out = **in
	}
	if in.StatementTimeoutInSeconds != nil {
		in, out := &in.StatementTimeoutInSeconds, &out.StatementTimeoutInSeconds
		*out = new(float64)
		**out = **in
	}
	if in.WaitForProvisioning != nil {
		in, out := &in.WaitForProvisioning, &out.WaitForProvisioning
		*out = new(bool)
		**out = **in
	}
	if in.WarehouseSize != nil {
		in, out := &in.WarehouseSize, &out.WarehouseSize
		*out = new(string)
		**out = **in
	}
	if in.WarehouseType != nil {
		in, out := &in.WarehouseType, &out.WarehouseType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WarehouseObservation.
func (in *WarehouseObservation) DeepCopy() *WarehouseObservation {
	if in == nil {
		return nil
	}
	out := new(WarehouseObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WarehouseParameters) DeepCopyInto(out *WarehouseParameters) {
	*out = *in
	if in.AutoResume != nil {
		in, out := &in.AutoResume, &out.AutoResume
		*out = new(bool)
		**out = **in
	}
	if in.AutoSuspend != nil {
		in, out := &in.AutoSuspend, &out.AutoSuspend
		*out = new(float64)
		**out = **in
	}
	if in.Comment != nil {
		in, out := &in.Comment, &out.Comment
		*out = new(string)
		**out = **in
	}
	if in.EnableQueryAcceleration != nil {
		in, out := &in.EnableQueryAcceleration, &out.EnableQueryAcceleration
		*out = new(bool)
		**out = **in
	}
	if in.InitiallySuspended != nil {
		in, out := &in.InitiallySuspended, &out.InitiallySuspended
		*out = new(bool)
		**out = **in
	}
	if in.MaxClusterCount != nil {
		in, out := &in.MaxClusterCount, &out.MaxClusterCount
		*out = new(float64)
		**out = **in
	}
	if in.MaxConcurrencyLevel != nil {
		in, out := &in.MaxConcurrencyLevel, &out.MaxConcurrencyLevel
		*out = new(float64)
		**out = **in
	}
	if in.MinClusterCount != nil {
		in, out := &in.MinClusterCount, &out.MinClusterCount
		*out = new(float64)
		**out = **in
	}
	if in.QueryAccelerationMaxScaleFactor != nil {
		in, out := &in.QueryAccelerationMaxScaleFactor, &out.QueryAccelerationMaxScaleFactor
		*out = new(float64)
		**out = **in
	}
	if in.ResourceMonitor != nil {
		in, out := &in.ResourceMonitor, &out.ResourceMonitor
		*out = new(string)
		**out = **in
	}
	if in.ScalingPolicy != nil {
		in, out := &in.ScalingPolicy, &out.ScalingPolicy
		*out = new(string)
		**out = **in
	}
	if in.StatementQueuedTimeoutInSeconds != nil {
		in, out := &in.StatementQueuedTimeoutInSeconds, &out.StatementQueuedTimeoutInSeconds
		*out = new(float64)
		**out = **in
	}
	if in.StatementTimeoutInSeconds != nil {
		in, out := &in.StatementTimeoutInSeconds, &out.StatementTimeoutInSeconds
		*out = new(float64)
		**out = **in
	}
	if in.WaitForProvisioning != nil {
		in, out := &in.WaitForProvisioning, &out.WaitForProvisioning
		*out = new(bool)
		**out = **in
	}
	if in.WarehouseSize != nil {
		in, out := &in.WarehouseSize, &out.WarehouseSize
		*out = new(string)
		**out = **in
	}
	if in.WarehouseType != nil {
		in, out := &in.WarehouseType, &out.WarehouseType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WarehouseParameters.
func (in *WarehouseParameters) DeepCopy() *WarehouseParameters {
	if in == nil {
		return nil
	}
	out := new(WarehouseParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WarehouseSpec) DeepCopyInto(out *WarehouseSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WarehouseSpec.
func (in *WarehouseSpec) DeepCopy() *WarehouseSpec {
	if in == nil {
		return nil
	}
	out := new(WarehouseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WarehouseStatus) DeepCopyInto(out *WarehouseStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WarehouseStatus.
func (in *WarehouseStatus) DeepCopy() *WarehouseStatus {
	if in == nil {
		return nil
	}
	out := new(WarehouseStatus)
	in.DeepCopyInto(out)
	return out
}