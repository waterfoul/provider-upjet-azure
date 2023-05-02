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
func (in *ChannelsObservation) DeepCopyInto(out *ChannelsObservation) {
	*out = *in
	if in.BandwidthMhz != nil {
		in, out := &in.BandwidthMhz, &out.BandwidthMhz
		*out = new(float64)
		**out = **in
	}
	if in.CenterFrequencyMhz != nil {
		in, out := &in.CenterFrequencyMhz, &out.CenterFrequencyMhz
		*out = new(float64)
		**out = **in
	}
	if in.DemodulationConfiguration != nil {
		in, out := &in.DemodulationConfiguration, &out.DemodulationConfiguration
		*out = new(string)
		**out = **in
	}
	if in.EndPoint != nil {
		in, out := &in.EndPoint, &out.EndPoint
		*out = make([]EndPointObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ModulationConfiguration != nil {
		in, out := &in.ModulationConfiguration, &out.ModulationConfiguration
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChannelsObservation.
func (in *ChannelsObservation) DeepCopy() *ChannelsObservation {
	if in == nil {
		return nil
	}
	out := new(ChannelsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChannelsParameters) DeepCopyInto(out *ChannelsParameters) {
	*out = *in
	if in.BandwidthMhz != nil {
		in, out := &in.BandwidthMhz, &out.BandwidthMhz
		*out = new(float64)
		**out = **in
	}
	if in.CenterFrequencyMhz != nil {
		in, out := &in.CenterFrequencyMhz, &out.CenterFrequencyMhz
		*out = new(float64)
		**out = **in
	}
	if in.DemodulationConfiguration != nil {
		in, out := &in.DemodulationConfiguration, &out.DemodulationConfiguration
		*out = new(string)
		**out = **in
	}
	if in.EndPoint != nil {
		in, out := &in.EndPoint, &out.EndPoint
		*out = make([]EndPointParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ModulationConfiguration != nil {
		in, out := &in.ModulationConfiguration, &out.ModulationConfiguration
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChannelsParameters.
func (in *ChannelsParameters) DeepCopy() *ChannelsParameters {
	if in == nil {
		return nil
	}
	out := new(ChannelsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContactProfile) DeepCopyInto(out *ContactProfile) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContactProfile.
func (in *ContactProfile) DeepCopy() *ContactProfile {
	if in == nil {
		return nil
	}
	out := new(ContactProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ContactProfile) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContactProfileList) DeepCopyInto(out *ContactProfileList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ContactProfile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContactProfileList.
func (in *ContactProfileList) DeepCopy() *ContactProfileList {
	if in == nil {
		return nil
	}
	out := new(ContactProfileList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ContactProfileList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContactProfileObservation) DeepCopyInto(out *ContactProfileObservation) {
	*out = *in
	if in.AutoTracking != nil {
		in, out := &in.AutoTracking, &out.AutoTracking
		*out = new(string)
		**out = **in
	}
	if in.EventHubURI != nil {
		in, out := &in.EventHubURI, &out.EventHubURI
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Links != nil {
		in, out := &in.Links, &out.Links
		*out = make([]LinksObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MinimumElevationDegrees != nil {
		in, out := &in.MinimumElevationDegrees, &out.MinimumElevationDegrees
		*out = new(float64)
		**out = **in
	}
	if in.MinimumVariableContactDuration != nil {
		in, out := &in.MinimumVariableContactDuration, &out.MinimumVariableContactDuration
		*out = new(string)
		**out = **in
	}
	if in.NetworkConfigurationSubnetID != nil {
		in, out := &in.NetworkConfigurationSubnetID, &out.NetworkConfigurationSubnetID
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContactProfileObservation.
func (in *ContactProfileObservation) DeepCopy() *ContactProfileObservation {
	if in == nil {
		return nil
	}
	out := new(ContactProfileObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContactProfileParameters) DeepCopyInto(out *ContactProfileParameters) {
	*out = *in
	if in.AutoTracking != nil {
		in, out := &in.AutoTracking, &out.AutoTracking
		*out = new(string)
		**out = **in
	}
	if in.EventHubURI != nil {
		in, out := &in.EventHubURI, &out.EventHubURI
		*out = new(string)
		**out = **in
	}
	if in.Links != nil {
		in, out := &in.Links, &out.Links
		*out = make([]LinksParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MinimumElevationDegrees != nil {
		in, out := &in.MinimumElevationDegrees, &out.MinimumElevationDegrees
		*out = new(float64)
		**out = **in
	}
	if in.MinimumVariableContactDuration != nil {
		in, out := &in.MinimumVariableContactDuration, &out.MinimumVariableContactDuration
		*out = new(string)
		**out = **in
	}
	if in.NetworkConfigurationSubnetID != nil {
		in, out := &in.NetworkConfigurationSubnetID, &out.NetworkConfigurationSubnetID
		*out = new(string)
		**out = **in
	}
	if in.NetworkConfigurationSubnetIDRef != nil {
		in, out := &in.NetworkConfigurationSubnetIDRef, &out.NetworkConfigurationSubnetIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.NetworkConfigurationSubnetIDSelector != nil {
		in, out := &in.NetworkConfigurationSubnetIDSelector, &out.NetworkConfigurationSubnetIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContactProfileParameters.
func (in *ContactProfileParameters) DeepCopy() *ContactProfileParameters {
	if in == nil {
		return nil
	}
	out := new(ContactProfileParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContactProfileSpec) DeepCopyInto(out *ContactProfileSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContactProfileSpec.
func (in *ContactProfileSpec) DeepCopy() *ContactProfileSpec {
	if in == nil {
		return nil
	}
	out := new(ContactProfileSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContactProfileStatus) DeepCopyInto(out *ContactProfileStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContactProfileStatus.
func (in *ContactProfileStatus) DeepCopy() *ContactProfileStatus {
	if in == nil {
		return nil
	}
	out := new(ContactProfileStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndPointObservation) DeepCopyInto(out *EndPointObservation) {
	*out = *in
	if in.EndPointName != nil {
		in, out := &in.EndPointName, &out.EndPointName
		*out = new(string)
		**out = **in
	}
	if in.IPAddress != nil {
		in, out := &in.IPAddress, &out.IPAddress
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(string)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndPointObservation.
func (in *EndPointObservation) DeepCopy() *EndPointObservation {
	if in == nil {
		return nil
	}
	out := new(EndPointObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndPointParameters) DeepCopyInto(out *EndPointParameters) {
	*out = *in
	if in.EndPointName != nil {
		in, out := &in.EndPointName, &out.EndPointName
		*out = new(string)
		**out = **in
	}
	if in.IPAddress != nil {
		in, out := &in.IPAddress, &out.IPAddress
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(string)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndPointParameters.
func (in *EndPointParameters) DeepCopy() *EndPointParameters {
	if in == nil {
		return nil
	}
	out := new(EndPointParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinksObservation) DeepCopyInto(out *LinksObservation) {
	*out = *in
	if in.Channels != nil {
		in, out := &in.Channels, &out.Channels
		*out = make([]ChannelsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Direction != nil {
		in, out := &in.Direction, &out.Direction
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Polarization != nil {
		in, out := &in.Polarization, &out.Polarization
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinksObservation.
func (in *LinksObservation) DeepCopy() *LinksObservation {
	if in == nil {
		return nil
	}
	out := new(LinksObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinksParameters) DeepCopyInto(out *LinksParameters) {
	*out = *in
	if in.Channels != nil {
		in, out := &in.Channels, &out.Channels
		*out = make([]ChannelsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Direction != nil {
		in, out := &in.Direction, &out.Direction
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Polarization != nil {
		in, out := &in.Polarization, &out.Polarization
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinksParameters.
func (in *LinksParameters) DeepCopy() *LinksParameters {
	if in == nil {
		return nil
	}
	out := new(LinksParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Spacecraft) DeepCopyInto(out *Spacecraft) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Spacecraft.
func (in *Spacecraft) DeepCopy() *Spacecraft {
	if in == nil {
		return nil
	}
	out := new(Spacecraft)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Spacecraft) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpacecraftLinksObservation) DeepCopyInto(out *SpacecraftLinksObservation) {
	*out = *in
	if in.BandwidthMhz != nil {
		in, out := &in.BandwidthMhz, &out.BandwidthMhz
		*out = new(float64)
		**out = **in
	}
	if in.CenterFrequencyMhz != nil {
		in, out := &in.CenterFrequencyMhz, &out.CenterFrequencyMhz
		*out = new(float64)
		**out = **in
	}
	if in.Direction != nil {
		in, out := &in.Direction, &out.Direction
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Polarization != nil {
		in, out := &in.Polarization, &out.Polarization
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpacecraftLinksObservation.
func (in *SpacecraftLinksObservation) DeepCopy() *SpacecraftLinksObservation {
	if in == nil {
		return nil
	}
	out := new(SpacecraftLinksObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpacecraftLinksParameters) DeepCopyInto(out *SpacecraftLinksParameters) {
	*out = *in
	if in.BandwidthMhz != nil {
		in, out := &in.BandwidthMhz, &out.BandwidthMhz
		*out = new(float64)
		**out = **in
	}
	if in.CenterFrequencyMhz != nil {
		in, out := &in.CenterFrequencyMhz, &out.CenterFrequencyMhz
		*out = new(float64)
		**out = **in
	}
	if in.Direction != nil {
		in, out := &in.Direction, &out.Direction
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Polarization != nil {
		in, out := &in.Polarization, &out.Polarization
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpacecraftLinksParameters.
func (in *SpacecraftLinksParameters) DeepCopy() *SpacecraftLinksParameters {
	if in == nil {
		return nil
	}
	out := new(SpacecraftLinksParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpacecraftList) DeepCopyInto(out *SpacecraftList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Spacecraft, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpacecraftList.
func (in *SpacecraftList) DeepCopy() *SpacecraftList {
	if in == nil {
		return nil
	}
	out := new(SpacecraftList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpacecraftList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpacecraftObservation) DeepCopyInto(out *SpacecraftObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Links != nil {
		in, out := &in.Links, &out.Links
		*out = make([]SpacecraftLinksObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.NoradID != nil {
		in, out := &in.NoradID, &out.NoradID
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TitleLine != nil {
		in, out := &in.TitleLine, &out.TitleLine
		*out = new(string)
		**out = **in
	}
	if in.TwoLineElements != nil {
		in, out := &in.TwoLineElements, &out.TwoLineElements
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpacecraftObservation.
func (in *SpacecraftObservation) DeepCopy() *SpacecraftObservation {
	if in == nil {
		return nil
	}
	out := new(SpacecraftObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpacecraftParameters) DeepCopyInto(out *SpacecraftParameters) {
	*out = *in
	if in.Links != nil {
		in, out := &in.Links, &out.Links
		*out = make([]SpacecraftLinksParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.NoradID != nil {
		in, out := &in.NoradID, &out.NoradID
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TitleLine != nil {
		in, out := &in.TitleLine, &out.TitleLine
		*out = new(string)
		**out = **in
	}
	if in.TwoLineElements != nil {
		in, out := &in.TwoLineElements, &out.TwoLineElements
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpacecraftParameters.
func (in *SpacecraftParameters) DeepCopy() *SpacecraftParameters {
	if in == nil {
		return nil
	}
	out := new(SpacecraftParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpacecraftSpec) DeepCopyInto(out *SpacecraftSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpacecraftSpec.
func (in *SpacecraftSpec) DeepCopy() *SpacecraftSpec {
	if in == nil {
		return nil
	}
	out := new(SpacecraftSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpacecraftStatus) DeepCopyInto(out *SpacecraftStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpacecraftStatus.
func (in *SpacecraftStatus) DeepCopy() *SpacecraftStatus {
	if in == nil {
		return nil
	}
	out := new(SpacecraftStatus)
	in.DeepCopyInto(out)
	return out
}