// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type IOTHubFallbackRouteInitParameters struct {

	// The condition that is evaluated to apply the routing rule. If no condition is provided, it evaluates to true by default. For grammar, see: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-query-language.
	Condition *string `json:"condition,omitempty" tf:"condition,omitempty"`

	// Used to specify whether the fallback route is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The source that the routing rule is to be applied to. Possible values include: DeviceConnectionStateEvents, DeviceJobLifecycleEvents, DeviceLifecycleEvents, DeviceMessages, DigitalTwinChangeEvents, Invalid, TwinChangeEvents.
	Source *string `json:"source,omitempty" tf:"source,omitempty"`
}

type IOTHubFallbackRouteObservation struct {

	// The condition that is evaluated to apply the routing rule. If no condition is provided, it evaluates to true by default. For grammar, see: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-query-language.
	Condition *string `json:"condition,omitempty" tf:"condition,omitempty"`

	// Used to specify whether the fallback route is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The endpoints to which messages that satisfy the condition are routed. Currently only 1 endpoint is allowed.
	EndpointNames []*string `json:"endpointNames,omitempty" tf:"endpoint_names,omitempty"`

	// The ID of the IoTHub Fallback Route.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the IoTHub to which this Fallback Route belongs. Changing this forces a new resource to be created.
	IOTHubName *string `json:"iothubName,omitempty" tf:"iothub_name,omitempty"`

	// The name of the resource group under which the IotHub Storage Container Endpoint resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The source that the routing rule is to be applied to. Possible values include: DeviceConnectionStateEvents, DeviceJobLifecycleEvents, DeviceLifecycleEvents, DeviceMessages, DigitalTwinChangeEvents, Invalid, TwinChangeEvents.
	Source *string `json:"source,omitempty" tf:"source,omitempty"`
}

type IOTHubFallbackRouteParameters struct {

	// The condition that is evaluated to apply the routing rule. If no condition is provided, it evaluates to true by default. For grammar, see: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-query-language.
	// +kubebuilder:validation:Optional
	Condition *string `json:"condition,omitempty" tf:"condition,omitempty"`

	// Used to specify whether the fallback route is enabled.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The endpoints to which messages that satisfy the condition are routed. Currently only 1 endpoint is allowed.
	// +crossplane:generate:reference:type=IOTHubEndpointStorageContainer
	// +kubebuilder:validation:Optional
	EndpointNames []*string `json:"endpointNames,omitempty" tf:"endpoint_names,omitempty"`

	// References to IOTHubEndpointStorageContainer to populate endpointNames.
	// +kubebuilder:validation:Optional
	EndpointNamesRefs []v1.Reference `json:"endpointNamesRefs,omitempty" tf:"-"`

	// Selector for a list of IOTHubEndpointStorageContainer to populate endpointNames.
	// +kubebuilder:validation:Optional
	EndpointNamesSelector *v1.Selector `json:"endpointNamesSelector,omitempty" tf:"-"`

	// The name of the IoTHub to which this Fallback Route belongs. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=IOTHub
	// +kubebuilder:validation:Optional
	IOTHubName *string `json:"iothubName,omitempty" tf:"iothub_name,omitempty"`

	// Reference to a IOTHub to populate iothubName.
	// +kubebuilder:validation:Optional
	IOTHubNameRef *v1.Reference `json:"iothubNameRef,omitempty" tf:"-"`

	// Selector for a IOTHub to populate iothubName.
	// +kubebuilder:validation:Optional
	IOTHubNameSelector *v1.Selector `json:"iothubNameSelector,omitempty" tf:"-"`

	// The name of the resource group under which the IotHub Storage Container Endpoint resource has to be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The source that the routing rule is to be applied to. Possible values include: DeviceConnectionStateEvents, DeviceJobLifecycleEvents, DeviceLifecycleEvents, DeviceMessages, DigitalTwinChangeEvents, Invalid, TwinChangeEvents.
	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`
}

// IOTHubFallbackRouteSpec defines the desired state of IOTHubFallbackRoute
type IOTHubFallbackRouteSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IOTHubFallbackRouteParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider IOTHubFallbackRouteInitParameters `json:"initProvider,omitempty"`
}

// IOTHubFallbackRouteStatus defines the observed state of IOTHubFallbackRoute.
type IOTHubFallbackRouteStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IOTHubFallbackRouteObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IOTHubFallbackRoute is the Schema for the IOTHubFallbackRoutes API. Manages an IotHub Fallback Route
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type IOTHubFallbackRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.enabled) || (has(self.initProvider) && has(self.initProvider.enabled))",message="spec.forProvider.enabled is a required parameter"
	Spec   IOTHubFallbackRouteSpec   `json:"spec"`
	Status IOTHubFallbackRouteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IOTHubFallbackRouteList contains a list of IOTHubFallbackRoutes
type IOTHubFallbackRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IOTHubFallbackRoute `json:"items"`
}

// Repository type metadata.
var (
	IOTHubFallbackRoute_Kind             = "IOTHubFallbackRoute"
	IOTHubFallbackRoute_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IOTHubFallbackRoute_Kind}.String()
	IOTHubFallbackRoute_KindAPIVersion   = IOTHubFallbackRoute_Kind + "." + CRDGroupVersion.String()
	IOTHubFallbackRoute_GroupVersionKind = CRDGroupVersion.WithKind(IOTHubFallbackRoute_Kind)
)

func init() {
	SchemeBuilder.Register(&IOTHubFallbackRoute{}, &IOTHubFallbackRouteList{})
}
