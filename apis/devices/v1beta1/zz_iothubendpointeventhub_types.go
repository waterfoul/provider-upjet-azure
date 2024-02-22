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

type IOTHubEndpointEventHubInitParameters struct {

	// Type used to authenticate against the Event Hub endpoint. Possible values are keyBased and identityBased. Defaults to keyBased.
	AuthenticationType *string `json:"authenticationType,omitempty" tf:"authentication_type,omitempty"`

	// URI of the Event Hubs Namespace endpoint. This attribute can only be specified and is mandatory when authentication_type is identityBased.
	EndpointURI *string `json:"endpointUri,omitempty" tf:"endpoint_uri,omitempty"`

	// Name of the Event Hub. This attribute can only be specified and is mandatory when authentication_type is identityBased.
	EntityPath *string `json:"entityPath,omitempty" tf:"entity_path,omitempty"`

	// ID of the User Managed Identity used to authenticate against the Event Hub endpoint.
	IdentityID *string `json:"identityId,omitempty" tf:"identity_id,omitempty"`

	// The name of the resource group under which the Event Hub has been created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`
}

type IOTHubEndpointEventHubObservation struct {

	// Type used to authenticate against the Event Hub endpoint. Possible values are keyBased and identityBased. Defaults to keyBased.
	AuthenticationType *string `json:"authenticationType,omitempty" tf:"authentication_type,omitempty"`

	// URI of the Event Hubs Namespace endpoint. This attribute can only be specified and is mandatory when authentication_type is identityBased.
	EndpointURI *string `json:"endpointUri,omitempty" tf:"endpoint_uri,omitempty"`

	// Name of the Event Hub. This attribute can only be specified and is mandatory when authentication_type is identityBased.
	EntityPath *string `json:"entityPath,omitempty" tf:"entity_path,omitempty"`

	// The ID of the IoTHub EventHub Endpoint.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The IoTHub ID for the endpoint. Changing this forces a new resource to be created.
	IOTHubID *string `json:"iothubId,omitempty" tf:"iothub_id,omitempty"`

	// ID of the User Managed Identity used to authenticate against the Event Hub endpoint.
	IdentityID *string `json:"identityId,omitempty" tf:"identity_id,omitempty"`

	// The name of the resource group under which the Event Hub has been created. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`
}

type IOTHubEndpointEventHubParameters struct {

	// Type used to authenticate against the Event Hub endpoint. Possible values are keyBased and identityBased. Defaults to keyBased.
	// +kubebuilder:validation:Optional
	AuthenticationType *string `json:"authenticationType,omitempty" tf:"authentication_type,omitempty"`

	// The connection string for the endpoint. This attribute can only be specified and is mandatory when authentication_type is keyBased.
	// +kubebuilder:validation:Optional
	ConnectionStringSecretRef *v1.SecretKeySelector `json:"connectionStringSecretRef,omitempty" tf:"-"`

	// URI of the Event Hubs Namespace endpoint. This attribute can only be specified and is mandatory when authentication_type is identityBased.
	// +kubebuilder:validation:Optional
	EndpointURI *string `json:"endpointUri,omitempty" tf:"endpoint_uri,omitempty"`

	// Name of the Event Hub. This attribute can only be specified and is mandatory when authentication_type is identityBased.
	// +kubebuilder:validation:Optional
	EntityPath *string `json:"entityPath,omitempty" tf:"entity_path,omitempty"`

	// The IoTHub ID for the endpoint. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/devices/v1beta1.IOTHub
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	IOTHubID *string `json:"iothubId,omitempty" tf:"iothub_id,omitempty"`

	// Reference to a IOTHub in devices to populate iothubId.
	// +kubebuilder:validation:Optional
	IOTHubIDRef *v1.Reference `json:"iothubIdRef,omitempty" tf:"-"`

	// Selector for a IOTHub in devices to populate iothubId.
	// +kubebuilder:validation:Optional
	IOTHubIDSelector *v1.Selector `json:"iothubIdSelector,omitempty" tf:"-"`

	// ID of the User Managed Identity used to authenticate against the Event Hub endpoint.
	// +kubebuilder:validation:Optional
	IdentityID *string `json:"identityId,omitempty" tf:"identity_id,omitempty"`

	// The name of the resource group under which the Event Hub has been created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`
}

// IOTHubEndpointEventHubSpec defines the desired state of IOTHubEndpointEventHub
type IOTHubEndpointEventHubSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IOTHubEndpointEventHubParameters `json:"forProvider"`
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
	InitProvider IOTHubEndpointEventHubInitParameters `json:"initProvider,omitempty"`
}

// IOTHubEndpointEventHubStatus defines the observed state of IOTHubEndpointEventHub.
type IOTHubEndpointEventHubStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IOTHubEndpointEventHubObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// IOTHubEndpointEventHub is the Schema for the IOTHubEndpointEventHubs API. Manages an IotHub EventHub Endpoint
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type IOTHubEndpointEventHub struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IOTHubEndpointEventHubSpec   `json:"spec"`
	Status            IOTHubEndpointEventHubStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IOTHubEndpointEventHubList contains a list of IOTHubEndpointEventHubs
type IOTHubEndpointEventHubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IOTHubEndpointEventHub `json:"items"`
}

// Repository type metadata.
var (
	IOTHubEndpointEventHub_Kind             = "IOTHubEndpointEventHub"
	IOTHubEndpointEventHub_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IOTHubEndpointEventHub_Kind}.String()
	IOTHubEndpointEventHub_KindAPIVersion   = IOTHubEndpointEventHub_Kind + "." + CRDGroupVersion.String()
	IOTHubEndpointEventHub_GroupVersionKind = CRDGroupVersion.WithKind(IOTHubEndpointEventHub_Kind)
)

func init() {
	SchemeBuilder.Register(&IOTHubEndpointEventHub{}, &IOTHubEndpointEventHubList{})
}
