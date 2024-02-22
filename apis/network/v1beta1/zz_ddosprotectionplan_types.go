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

type DDoSProtectionPlanInitParameters struct {

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type DDoSProtectionPlanObservation struct {

	// The ID of the DDoS Protection Plan
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A list of Virtual Network IDs associated with the DDoS Protection Plan.
	VirtualNetworkIds []*string `json:"virtualNetworkIds,omitempty" tf:"virtual_network_ids,omitempty"`
}

type DDoSProtectionPlanParameters struct {

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// DDoSProtectionPlanSpec defines the desired state of DDoSProtectionPlan
type DDoSProtectionPlanSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DDoSProtectionPlanParameters `json:"forProvider"`
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
	InitProvider DDoSProtectionPlanInitParameters `json:"initProvider,omitempty"`
}

// DDoSProtectionPlanStatus defines the observed state of DDoSProtectionPlan.
type DDoSProtectionPlanStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DDoSProtectionPlanObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DDoSProtectionPlan is the Schema for the DDoSProtectionPlans API. Manages an Azure Network DDoS Protection Plan.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DDoSProtectionPlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	Spec   DDoSProtectionPlanSpec   `json:"spec"`
	Status DDoSProtectionPlanStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DDoSProtectionPlanList contains a list of DDoSProtectionPlans
type DDoSProtectionPlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DDoSProtectionPlan `json:"items"`
}

// Repository type metadata.
var (
	DDoSProtectionPlan_Kind             = "DDoSProtectionPlan"
	DDoSProtectionPlan_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DDoSProtectionPlan_Kind}.String()
	DDoSProtectionPlan_KindAPIVersion   = DDoSProtectionPlan_Kind + "." + CRDGroupVersion.String()
	DDoSProtectionPlan_GroupVersionKind = CRDGroupVersion.WithKind(DDoSProtectionPlan_Kind)
)

func init() {
	SchemeBuilder.Register(&DDoSProtectionPlan{}, &DDoSProtectionPlanList{})
}
