//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Grant) DeepCopyInto(out *Grant) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Grant.
func (in *Grant) DeepCopy() *Grant {
	if in == nil {
		return nil
	}
	out := new(Grant)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Grant) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrantInitParameters) DeepCopyInto(out *GrantInitParameters) {
	*out = *in
	if in.EnableMultipleGrants != nil {
		in, out := &in.EnableMultipleGrants, &out.EnableMultipleGrants
		*out = new(bool)
		**out = **in
	}
	if in.Privilege != nil {
		in, out := &in.Privilege, &out.Privilege
		*out = new(string)
		**out = **in
	}
	if in.RevertOwnershipToRoleName != nil {
		in, out := &in.RevertOwnershipToRoleName, &out.RevertOwnershipToRoleName
		*out = new(string)
		**out = **in
	}
	if in.WithGrantOption != nil {
		in, out := &in.WithGrantOption, &out.WithGrantOption
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrantInitParameters.
func (in *GrantInitParameters) DeepCopy() *GrantInitParameters {
	if in == nil {
		return nil
	}
	out := new(GrantInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrantList) DeepCopyInto(out *GrantList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Grant, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrantList.
func (in *GrantList) DeepCopy() *GrantList {
	if in == nil {
		return nil
	}
	out := new(GrantList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GrantList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrantObservation) DeepCopyInto(out *GrantObservation) {
	*out = *in
	if in.DatabaseName != nil {
		in, out := &in.DatabaseName, &out.DatabaseName
		*out = new(string)
		**out = **in
	}
	if in.EnableMultipleGrants != nil {
		in, out := &in.EnableMultipleGrants, &out.EnableMultipleGrants
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Privilege != nil {
		in, out := &in.Privilege, &out.Privilege
		*out = new(string)
		**out = **in
	}
	if in.RevertOwnershipToRoleName != nil {
		in, out := &in.RevertOwnershipToRoleName, &out.RevertOwnershipToRoleName
		*out = new(string)
		**out = **in
	}
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Shares != nil {
		in, out := &in.Shares, &out.Shares
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.WithGrantOption != nil {
		in, out := &in.WithGrantOption, &out.WithGrantOption
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrantObservation.
func (in *GrantObservation) DeepCopy() *GrantObservation {
	if in == nil {
		return nil
	}
	out := new(GrantObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrantParameters) DeepCopyInto(out *GrantParameters) {
	*out = *in
	if in.DatabaseName != nil {
		in, out := &in.DatabaseName, &out.DatabaseName
		*out = new(string)
		**out = **in
	}
	if in.DatabaseRef != nil {
		in, out := &in.DatabaseRef, &out.DatabaseRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.DatabaseSelector != nil {
		in, out := &in.DatabaseSelector, &out.DatabaseSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.EnableMultipleGrants != nil {
		in, out := &in.EnableMultipleGrants, &out.EnableMultipleGrants
		*out = new(bool)
		**out = **in
	}
	if in.Privilege != nil {
		in, out := &in.Privilege, &out.Privilege
		*out = new(string)
		**out = **in
	}
	if in.RevertOwnershipToRoleName != nil {
		in, out := &in.RevertOwnershipToRoleName, &out.RevertOwnershipToRoleName
		*out = new(string)
		**out = **in
	}
	if in.RoleRefs != nil {
		in, out := &in.RoleRefs, &out.RoleRefs
		*out = make([]v1.Reference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RoleSelector != nil {
		in, out := &in.RoleSelector, &out.RoleSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ShareRefs != nil {
		in, out := &in.ShareRefs, &out.ShareRefs
		*out = make([]v1.Reference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ShareSelector != nil {
		in, out := &in.ShareSelector, &out.ShareSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Shares != nil {
		in, out := &in.Shares, &out.Shares
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.WithGrantOption != nil {
		in, out := &in.WithGrantOption, &out.WithGrantOption
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrantParameters.
func (in *GrantParameters) DeepCopy() *GrantParameters {
	if in == nil {
		return nil
	}
	out := new(GrantParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrantSpec) DeepCopyInto(out *GrantSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrantSpec.
func (in *GrantSpec) DeepCopy() *GrantSpec {
	if in == nil {
		return nil
	}
	out := new(GrantSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrantStatus) DeepCopyInto(out *GrantStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrantStatus.
func (in *GrantStatus) DeepCopy() *GrantStatus {
	if in == nil {
		return nil
	}
	out := new(GrantStatus)
	in.DeepCopyInto(out)
	return out
}
