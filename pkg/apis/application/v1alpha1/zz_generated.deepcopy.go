// +build !ignore_autogenerated

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Application) DeepCopyInto(out *Application) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Application.
func (in *Application) DeepCopy() *Application {
	if in == nil {
		return nil
	}
	out := new(Application)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Application) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationDestination) DeepCopyInto(out *ApplicationDestination) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationDestination.
func (in *ApplicationDestination) DeepCopy() *ApplicationDestination {
	if in == nil {
		return nil
	}
	out := new(ApplicationDestination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationList) DeepCopyInto(out *ApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Application, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationList.
func (in *ApplicationList) DeepCopy() *ApplicationList {
	if in == nil {
		return nil
	}
	out := new(ApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSource) DeepCopyInto(out *ApplicationSource) {
	*out = *in
	if in.ComponentParameterOverrides != nil {
		in, out := &in.ComponentParameterOverrides, &out.ComponentParameterOverrides
		*out = make([]ComponentParameter, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSource.
func (in *ApplicationSource) DeepCopy() *ApplicationSource {
	if in == nil {
		return nil
	}
	out := new(ApplicationSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSpec) DeepCopyInto(out *ApplicationSpec) {
	*out = *in
	in.Source.DeepCopyInto(&out.Source)
	if in.Destination != nil {
		in, out := &in.Destination, &out.Destination
		if *in == nil {
			*out = nil
		} else {
			*out = new(ApplicationDestination)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSpec.
func (in *ApplicationSpec) DeepCopy() *ApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(ApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationStatus) DeepCopyInto(out *ApplicationStatus) {
	*out = *in
	in.ComparisonResult.DeepCopyInto(&out.ComparisonResult)
	if in.RecentDeployments != nil {
		in, out := &in.RecentDeployments, &out.RecentDeployments
		*out = make([]DeploymentInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationStatus.
func (in *ApplicationStatus) DeepCopy() *ApplicationStatus {
	if in == nil {
		return nil
	}
	out := new(ApplicationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationWatchEvent) DeepCopyInto(out *ApplicationWatchEvent) {
	*out = *in
	in.Application.DeepCopyInto(&out.Application)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationWatchEvent.
func (in *ApplicationWatchEvent) DeepCopy() *ApplicationWatchEvent {
	if in == nil {
		return nil
	}
	out := new(ApplicationWatchEvent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	in.Config.DeepCopyInto(&out.Config)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConfig) DeepCopyInto(out *ClusterConfig) {
	*out = *in
	in.TLSClientConfig.DeepCopyInto(&out.TLSClientConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConfig.
func (in *ClusterConfig) DeepCopy() *ClusterConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterList) DeepCopyInto(out *ClusterList) {
	*out = *in
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterList.
func (in *ClusterList) DeepCopy() *ClusterList {
	if in == nil {
		return nil
	}
	out := new(ClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComparisonResult) DeepCopyInto(out *ComparisonResult) {
	*out = *in
	in.ComparedAt.DeepCopyInto(&out.ComparedAt)
	in.ComparedTo.DeepCopyInto(&out.ComparedTo)
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]ResourceState, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComparisonResult.
func (in *ComparisonResult) DeepCopy() *ComparisonResult {
	if in == nil {
		return nil
	}
	out := new(ComparisonResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentParameter) DeepCopyInto(out *ComponentParameter) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentParameter.
func (in *ComponentParameter) DeepCopy() *ComponentParameter {
	if in == nil {
		return nil
	}
	out := new(ComponentParameter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentInfo) DeepCopyInto(out *DeploymentInfo) {
	*out = *in
	if in.Params != nil {
		in, out := &in.Params, &out.Params
		*out = make([]ComponentParameter, len(*in))
		copy(*out, *in)
	}
	if in.ComponentParameterOverrides != nil {
		in, out := &in.ComponentParameterOverrides, &out.ComponentParameterOverrides
		*out = make([]ComponentParameter, len(*in))
		copy(*out, *in)
	}
	in.DeployedAt.DeepCopyInto(&out.DeployedAt)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentInfo.
func (in *DeploymentInfo) DeepCopy() *DeploymentInfo {
	if in == nil {
		return nil
	}
	out := new(DeploymentInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Repository) DeepCopyInto(out *Repository) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Repository.
func (in *Repository) DeepCopy() *Repository {
	if in == nil {
		return nil
	}
	out := new(Repository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryList) DeepCopyInto(out *RepositoryList) {
	*out = *in
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Repository, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryList.
func (in *RepositoryList) DeepCopy() *RepositoryList {
	if in == nil {
		return nil
	}
	out := new(RepositoryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceNode) DeepCopyInto(out *ResourceNode) {
	*out = *in
	if in.Children != nil {
		in, out := &in.Children, &out.Children
		*out = make([]ResourceNode, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceNode.
func (in *ResourceNode) DeepCopy() *ResourceNode {
	if in == nil {
		return nil
	}
	out := new(ResourceNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceState) DeepCopyInto(out *ResourceState) {
	*out = *in
	if in.ChildLiveResources != nil {
		in, out := &in.ChildLiveResources, &out.ChildLiveResources
		*out = make([]ResourceNode, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceState.
func (in *ResourceState) DeepCopy() *ResourceState {
	if in == nil {
		return nil
	}
	out := new(ResourceState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSClientConfig) DeepCopyInto(out *TLSClientConfig) {
	*out = *in
	if in.CertData != nil {
		in, out := &in.CertData, &out.CertData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.KeyData != nil {
		in, out := &in.KeyData, &out.KeyData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.CAData != nil {
		in, out := &in.CAData, &out.CAData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSClientConfig.
func (in *TLSClientConfig) DeepCopy() *TLSClientConfig {
	if in == nil {
		return nil
	}
	out := new(TLSClientConfig)
	in.DeepCopyInto(out)
	return out
}
