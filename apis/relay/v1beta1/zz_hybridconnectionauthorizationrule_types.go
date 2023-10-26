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

type HybridConnectionAuthorizationRuleInitParameters struct {

	// Grants listen access to this Authorization Rule. Defaults to false.
	Listen *bool `json:"listen,omitempty" tf:"listen,omitempty"`

	// Grants manage access to this Authorization Rule. When this property is true - both listen and send must be set to true too. Defaults to false.
	Manage *bool `json:"manage,omitempty" tf:"manage,omitempty"`

	// Grants send access to this Authorization Rule. Defaults to false.
	Send *bool `json:"send,omitempty" tf:"send,omitempty"`
}

type HybridConnectionAuthorizationRuleObservation struct {

	// Name of the Azure Relay Hybrid Connection for which this Azure Relay Hybrid Connection Authorization Rule will be created. Changing this forces a new Azure Relay Hybrid Connection Authorization Rule to be created.
	HybridConnectionName *string `json:"hybridConnectionName,omitempty" tf:"hybrid_connection_name,omitempty"`

	// The ID of the Azure Relay Hybrid Connection Authorization Rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Grants listen access to this Authorization Rule. Defaults to false.
	Listen *bool `json:"listen,omitempty" tf:"listen,omitempty"`

	// Grants manage access to this Authorization Rule. When this property is true - both listen and send must be set to true too. Defaults to false.
	Manage *bool `json:"manage,omitempty" tf:"manage,omitempty"`

	// Name of the Azure Relay Namespace for which this Azure Relay Hybrid Connection Authorization Rule will be created. Changing this forces a new Azure Relay Hybrid Connection Authorization Rule to be created.
	NamespaceName *string `json:"namespaceName,omitempty" tf:"namespace_name,omitempty"`

	// The name of the Resource Group where the Azure Relay Hybrid Connection Authorization Rule should exist. Changing this forces a new Azure Relay Hybrid Connection Authorization Rule to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Grants send access to this Authorization Rule. Defaults to false.
	Send *bool `json:"send,omitempty" tf:"send,omitempty"`
}

type HybridConnectionAuthorizationRuleParameters struct {

	// Name of the Azure Relay Hybrid Connection for which this Azure Relay Hybrid Connection Authorization Rule will be created. Changing this forces a new Azure Relay Hybrid Connection Authorization Rule to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/relay/v1beta1.HybridConnection
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",false)
	// +kubebuilder:validation:Optional
	HybridConnectionName *string `json:"hybridConnectionName,omitempty" tf:"hybrid_connection_name,omitempty"`

	// Reference to a HybridConnection in relay to populate hybridConnectionName.
	// +kubebuilder:validation:Optional
	HybridConnectionNameRef *v1.Reference `json:"hybridConnectionNameRef,omitempty" tf:"-"`

	// Selector for a HybridConnection in relay to populate hybridConnectionName.
	// +kubebuilder:validation:Optional
	HybridConnectionNameSelector *v1.Selector `json:"hybridConnectionNameSelector,omitempty" tf:"-"`

	// Grants listen access to this Authorization Rule. Defaults to false.
	// +kubebuilder:validation:Optional
	Listen *bool `json:"listen,omitempty" tf:"listen,omitempty"`

	// Grants manage access to this Authorization Rule. When this property is true - both listen and send must be set to true too. Defaults to false.
	// +kubebuilder:validation:Optional
	Manage *bool `json:"manage,omitempty" tf:"manage,omitempty"`

	// Name of the Azure Relay Namespace for which this Azure Relay Hybrid Connection Authorization Rule will be created. Changing this forces a new Azure Relay Hybrid Connection Authorization Rule to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/relay/v1beta1.EventRelayNamespace
	// +kubebuilder:validation:Optional
	NamespaceName *string `json:"namespaceName,omitempty" tf:"namespace_name,omitempty"`

	// Reference to a EventRelayNamespace in relay to populate namespaceName.
	// +kubebuilder:validation:Optional
	NamespaceNameRef *v1.Reference `json:"namespaceNameRef,omitempty" tf:"-"`

	// Selector for a EventRelayNamespace in relay to populate namespaceName.
	// +kubebuilder:validation:Optional
	NamespaceNameSelector *v1.Selector `json:"namespaceNameSelector,omitempty" tf:"-"`

	// The name of the Resource Group where the Azure Relay Hybrid Connection Authorization Rule should exist. Changing this forces a new Azure Relay Hybrid Connection Authorization Rule to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Grants send access to this Authorization Rule. Defaults to false.
	// +kubebuilder:validation:Optional
	Send *bool `json:"send,omitempty" tf:"send,omitempty"`
}

// HybridConnectionAuthorizationRuleSpec defines the desired state of HybridConnectionAuthorizationRule
type HybridConnectionAuthorizationRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HybridConnectionAuthorizationRuleParameters `json:"forProvider"`
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
	InitProvider HybridConnectionAuthorizationRuleInitParameters `json:"initProvider,omitempty"`
}

// HybridConnectionAuthorizationRuleStatus defines the observed state of HybridConnectionAuthorizationRule.
type HybridConnectionAuthorizationRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HybridConnectionAuthorizationRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HybridConnectionAuthorizationRule is the Schema for the HybridConnectionAuthorizationRules API. Manages an Azure Relay Hybrid Connection Authorization Rule.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type HybridConnectionAuthorizationRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HybridConnectionAuthorizationRuleSpec   `json:"spec"`
	Status            HybridConnectionAuthorizationRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HybridConnectionAuthorizationRuleList contains a list of HybridConnectionAuthorizationRules
type HybridConnectionAuthorizationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HybridConnectionAuthorizationRule `json:"items"`
}

// Repository type metadata.
var (
	HybridConnectionAuthorizationRule_Kind             = "HybridConnectionAuthorizationRule"
	HybridConnectionAuthorizationRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HybridConnectionAuthorizationRule_Kind}.String()
	HybridConnectionAuthorizationRule_KindAPIVersion   = HybridConnectionAuthorizationRule_Kind + "." + CRDGroupVersion.String()
	HybridConnectionAuthorizationRule_GroupVersionKind = CRDGroupVersion.WithKind(HybridConnectionAuthorizationRule_Kind)
)

func init() {
	SchemeBuilder.Register(&HybridConnectionAuthorizationRule{}, &HybridConnectionAuthorizationRuleList{})
}
