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

type SiteRecoveryNetworkMappingInitParameters struct {

	// The name of the network mapping. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the vault that should be updated. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/recoveryservices/v1beta1.Vault
	RecoveryVaultName *string `json:"recoveryVaultName,omitempty" tf:"recovery_vault_name,omitempty"`

	// Reference to a Vault in recoveryservices to populate recoveryVaultName.
	// +kubebuilder:validation:Optional
	RecoveryVaultNameRef *v1.Reference `json:"recoveryVaultNameRef,omitempty" tf:"-"`

	// Selector for a Vault in recoveryservices to populate recoveryVaultName.
	// +kubebuilder:validation:Optional
	RecoveryVaultNameSelector *v1.Selector `json:"recoveryVaultNameSelector,omitempty" tf:"-"`

	// Name of the resource group where the vault that should be updated is located. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The id of the primary network. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.VirtualNetwork
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	SourceNetworkID *string `json:"sourceNetworkId,omitempty" tf:"source_network_id,omitempty"`

	// Reference to a VirtualNetwork in network to populate sourceNetworkId.
	// +kubebuilder:validation:Optional
	SourceNetworkIDRef *v1.Reference `json:"sourceNetworkIdRef,omitempty" tf:"-"`

	// Selector for a VirtualNetwork in network to populate sourceNetworkId.
	// +kubebuilder:validation:Optional
	SourceNetworkIDSelector *v1.Selector `json:"sourceNetworkIdSelector,omitempty" tf:"-"`

	// Specifies the ASR fabric where mapping should be created. Changing this forces a new resource to be created.
	SourceRecoveryFabricName *string `json:"sourceRecoveryFabricName,omitempty" tf:"source_recovery_fabric_name,omitempty"`

	// The id of the recovery network. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.VirtualNetwork
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	TargetNetworkID *string `json:"targetNetworkId,omitempty" tf:"target_network_id,omitempty"`

	// Reference to a VirtualNetwork in network to populate targetNetworkId.
	// +kubebuilder:validation:Optional
	TargetNetworkIDRef *v1.Reference `json:"targetNetworkIdRef,omitempty" tf:"-"`

	// Selector for a VirtualNetwork in network to populate targetNetworkId.
	// +kubebuilder:validation:Optional
	TargetNetworkIDSelector *v1.Selector `json:"targetNetworkIdSelector,omitempty" tf:"-"`

	// The Azure Site Recovery fabric object corresponding to the recovery Azure region. Changing this forces a new resource to be created.
	TargetRecoveryFabricName *string `json:"targetRecoveryFabricName,omitempty" tf:"target_recovery_fabric_name,omitempty"`
}

type SiteRecoveryNetworkMappingObservation struct {

	// The ID of the Site Recovery Network Mapping.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the network mapping. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the vault that should be updated. Changing this forces a new resource to be created.
	RecoveryVaultName *string `json:"recoveryVaultName,omitempty" tf:"recovery_vault_name,omitempty"`

	// Name of the resource group where the vault that should be updated is located. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The id of the primary network. Changing this forces a new resource to be created.
	SourceNetworkID *string `json:"sourceNetworkId,omitempty" tf:"source_network_id,omitempty"`

	// Specifies the ASR fabric where mapping should be created. Changing this forces a new resource to be created.
	SourceRecoveryFabricName *string `json:"sourceRecoveryFabricName,omitempty" tf:"source_recovery_fabric_name,omitempty"`

	// The id of the recovery network. Changing this forces a new resource to be created.
	TargetNetworkID *string `json:"targetNetworkId,omitempty" tf:"target_network_id,omitempty"`

	// The Azure Site Recovery fabric object corresponding to the recovery Azure region. Changing this forces a new resource to be created.
	TargetRecoveryFabricName *string `json:"targetRecoveryFabricName,omitempty" tf:"target_recovery_fabric_name,omitempty"`
}

type SiteRecoveryNetworkMappingParameters struct {

	// The name of the network mapping. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the vault that should be updated. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/recoveryservices/v1beta1.Vault
	// +kubebuilder:validation:Optional
	RecoveryVaultName *string `json:"recoveryVaultName,omitempty" tf:"recovery_vault_name,omitempty"`

	// Reference to a Vault in recoveryservices to populate recoveryVaultName.
	// +kubebuilder:validation:Optional
	RecoveryVaultNameRef *v1.Reference `json:"recoveryVaultNameRef,omitempty" tf:"-"`

	// Selector for a Vault in recoveryservices to populate recoveryVaultName.
	// +kubebuilder:validation:Optional
	RecoveryVaultNameSelector *v1.Selector `json:"recoveryVaultNameSelector,omitempty" tf:"-"`

	// Name of the resource group where the vault that should be updated is located. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The id of the primary network. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.VirtualNetwork
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SourceNetworkID *string `json:"sourceNetworkId,omitempty" tf:"source_network_id,omitempty"`

	// Reference to a VirtualNetwork in network to populate sourceNetworkId.
	// +kubebuilder:validation:Optional
	SourceNetworkIDRef *v1.Reference `json:"sourceNetworkIdRef,omitempty" tf:"-"`

	// Selector for a VirtualNetwork in network to populate sourceNetworkId.
	// +kubebuilder:validation:Optional
	SourceNetworkIDSelector *v1.Selector `json:"sourceNetworkIdSelector,omitempty" tf:"-"`

	// Specifies the ASR fabric where mapping should be created. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SourceRecoveryFabricName *string `json:"sourceRecoveryFabricName,omitempty" tf:"source_recovery_fabric_name,omitempty"`

	// The id of the recovery network. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.VirtualNetwork
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	TargetNetworkID *string `json:"targetNetworkId,omitempty" tf:"target_network_id,omitempty"`

	// Reference to a VirtualNetwork in network to populate targetNetworkId.
	// +kubebuilder:validation:Optional
	TargetNetworkIDRef *v1.Reference `json:"targetNetworkIdRef,omitempty" tf:"-"`

	// Selector for a VirtualNetwork in network to populate targetNetworkId.
	// +kubebuilder:validation:Optional
	TargetNetworkIDSelector *v1.Selector `json:"targetNetworkIdSelector,omitempty" tf:"-"`

	// The Azure Site Recovery fabric object corresponding to the recovery Azure region. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	TargetRecoveryFabricName *string `json:"targetRecoveryFabricName,omitempty" tf:"target_recovery_fabric_name,omitempty"`
}

// SiteRecoveryNetworkMappingSpec defines the desired state of SiteRecoveryNetworkMapping
type SiteRecoveryNetworkMappingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SiteRecoveryNetworkMappingParameters `json:"forProvider"`
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
	InitProvider SiteRecoveryNetworkMappingInitParameters `json:"initProvider,omitempty"`
}

// SiteRecoveryNetworkMappingStatus defines the observed state of SiteRecoveryNetworkMapping.
type SiteRecoveryNetworkMappingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SiteRecoveryNetworkMappingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SiteRecoveryNetworkMapping is the Schema for the SiteRecoveryNetworkMappings API. Manages a site recovery network mapping on Azure.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SiteRecoveryNetworkMapping struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.sourceRecoveryFabricName) || (has(self.initProvider) && has(self.initProvider.sourceRecoveryFabricName))",message="spec.forProvider.sourceRecoveryFabricName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.targetRecoveryFabricName) || (has(self.initProvider) && has(self.initProvider.targetRecoveryFabricName))",message="spec.forProvider.targetRecoveryFabricName is a required parameter"
	Spec   SiteRecoveryNetworkMappingSpec   `json:"spec"`
	Status SiteRecoveryNetworkMappingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SiteRecoveryNetworkMappingList contains a list of SiteRecoveryNetworkMappings
type SiteRecoveryNetworkMappingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SiteRecoveryNetworkMapping `json:"items"`
}

// Repository type metadata.
var (
	SiteRecoveryNetworkMapping_Kind             = "SiteRecoveryNetworkMapping"
	SiteRecoveryNetworkMapping_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SiteRecoveryNetworkMapping_Kind}.String()
	SiteRecoveryNetworkMapping_KindAPIVersion   = SiteRecoveryNetworkMapping_Kind + "." + CRDGroupVersion.String()
	SiteRecoveryNetworkMapping_GroupVersionKind = CRDGroupVersion.WithKind(SiteRecoveryNetworkMapping_Kind)
)

func init() {
	SchemeBuilder.Register(&SiteRecoveryNetworkMapping{}, &SiteRecoveryNetworkMappingList{})
}
