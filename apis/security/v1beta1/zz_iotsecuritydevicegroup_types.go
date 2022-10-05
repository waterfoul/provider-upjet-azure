/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AllowRuleObservation struct {
}

type AllowRuleParameters struct {

	// Specifies which IP is not allowed to be connected to in current device group for inbound connection.
	// +kubebuilder:validation:Optional
	ConnectionFromIpsNotAllowed []*string `json:"connectionFromIpsNotAllowed,omitempty" tf:"connection_from_ips_not_allowed,omitempty"`

	// Specifies which IP is not allowed to be connected to in current device group for outbound connection.
	// +kubebuilder:validation:Optional
	ConnectionToIpsNotAllowed []*string `json:"connectionToIpsNotAllowed,omitempty" tf:"connection_to_ips_not_allowed,omitempty"`

	// Specifies which local user is not allowed to login in current device group.
	// +kubebuilder:validation:Optional
	LocalUsersNotAllowed []*string `json:"localUsersNotAllowed,omitempty" tf:"local_users_not_allowed,omitempty"`

	// Specifies which process is not allowed to be executed in current device group.
	// +kubebuilder:validation:Optional
	ProcessesNotAllowed []*string `json:"processesNotAllowed,omitempty" tf:"processes_not_allowed,omitempty"`
}

type IOTSecurityDeviceGroupObservation struct {

	// The ID of the Iot Security Device Group resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type IOTSecurityDeviceGroupParameters struct {

	// an allow_rule blocks as defined below.
	// +kubebuilder:validation:Optional
	AllowRule []AllowRuleParameters `json:"allowRule,omitempty" tf:"allow_rule,omitempty"`

	// The ID of the IoT Hub which to link the Security Device Group to. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/devices/v1beta1.IOTHub
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	IOTHubID *string `json:"iothubId,omitempty" tf:"iothub_id,omitempty"`

	// Reference to a IOTHub in devices to populate iothubId.
	// +kubebuilder:validation:Optional
	IOTHubIDRef *v1.Reference `json:"iothubIdRef,omitempty" tf:"-"`

	// Selector for a IOTHub in devices to populate iothubId.
	// +kubebuilder:validation:Optional
	IOTHubIDSelector *v1.Selector `json:"iothubIdSelector,omitempty" tf:"-"`

	// Specifies the name of the Device Security Group. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// One or more range_rule blocks as defined below.
	// +kubebuilder:validation:Optional
	RangeRule []RangeRuleParameters `json:"rangeRule,omitempty" tf:"range_rule,omitempty"`
}

type RangeRuleObservation struct {
}

type RangeRuleParameters struct {

	// Specifies the time range. represented in ISO 8601 duration format.
	// +kubebuilder:validation:Required
	Duration *string `json:"duration" tf:"duration,omitempty"`

	// The maximum threshold in the given time window.
	// +kubebuilder:validation:Required
	Max *float64 `json:"max" tf:"max,omitempty"`

	// The minimum threshold in the given time window.
	// +kubebuilder:validation:Required
	Min *float64 `json:"min" tf:"min,omitempty"`

	// The type of supported rule type. Possible Values are ActiveConnectionsNotInAllowedRange, AmqpC2DMessagesNotInAllowedRange, MqttC2DMessagesNotInAllowedRange, HttpC2DMessagesNotInAllowedRange, AmqpC2DRejectedMessagesNotInAllowedRange, MqttC2DRejectedMessagesNotInAllowedRange, HttpC2DRejectedMessagesNotInAllowedRange, AmqpD2CMessagesNotInAllowedRange, MqttD2CMessagesNotInAllowedRange, HttpD2CMessagesNotInAllowedRange, DirectMethodInvokesNotInAllowedRange, FailedLocalLoginsNotInAllowedRange, FileUploadsNotInAllowedRange, QueuePurgesNotInAllowedRange, TwinUpdatesNotInAllowedRange and UnauthorizedOperationsNotInAllowedRange.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// IOTSecurityDeviceGroupSpec defines the desired state of IOTSecurityDeviceGroup
type IOTSecurityDeviceGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IOTSecurityDeviceGroupParameters `json:"forProvider"`
}

// IOTSecurityDeviceGroupStatus defines the observed state of IOTSecurityDeviceGroup.
type IOTSecurityDeviceGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IOTSecurityDeviceGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IOTSecurityDeviceGroup is the Schema for the IOTSecurityDeviceGroups API. Manages a Iot Security Device Group.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type IOTSecurityDeviceGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IOTSecurityDeviceGroupSpec   `json:"spec"`
	Status            IOTSecurityDeviceGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IOTSecurityDeviceGroupList contains a list of IOTSecurityDeviceGroups
type IOTSecurityDeviceGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IOTSecurityDeviceGroup `json:"items"`
}

// Repository type metadata.
var (
	IOTSecurityDeviceGroup_Kind             = "IOTSecurityDeviceGroup"
	IOTSecurityDeviceGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IOTSecurityDeviceGroup_Kind}.String()
	IOTSecurityDeviceGroup_KindAPIVersion   = IOTSecurityDeviceGroup_Kind + "." + CRDGroupVersion.String()
	IOTSecurityDeviceGroup_GroupVersionKind = CRDGroupVersion.WithKind(IOTSecurityDeviceGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&IOTSecurityDeviceGroup{}, &IOTSecurityDeviceGroupList{})
}