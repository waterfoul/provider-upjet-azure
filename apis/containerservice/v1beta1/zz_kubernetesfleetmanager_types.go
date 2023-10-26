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

type HubProfileInitParameters struct {
	DNSPrefix *string `json:"dnsPrefix,omitempty" tf:"dns_prefix,omitempty"`
}

type HubProfileObservation struct {
	DNSPrefix *string `json:"dnsPrefix,omitempty" tf:"dns_prefix,omitempty"`

	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	KubernetesVersion *string `json:"kubernetesVersion,omitempty" tf:"kubernetes_version,omitempty"`
}

type HubProfileParameters struct {

	// +kubebuilder:validation:Optional
	DNSPrefix *string `json:"dnsPrefix" tf:"dns_prefix,omitempty"`
}

type KubernetesFleetManagerInitParameters struct {

	// A hub_profile block as defined below. The FleetHubProfile configures the Fleet's hub. Changing this forces a new Kubernetes Fleet Manager to be created.
	HubProfile []HubProfileInitParameters `json:"hubProfile,omitempty" tf:"hub_profile,omitempty"`

	// The Azure Region where the Kubernetes Fleet Manager should exist. Changing this forces a new Kubernetes Fleet Manager to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A mapping of tags which should be assigned to the Kubernetes Fleet Manager.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type KubernetesFleetManagerObservation struct {

	// A hub_profile block as defined below. The FleetHubProfile configures the Fleet's hub. Changing this forces a new Kubernetes Fleet Manager to be created.
	HubProfile []HubProfileObservation `json:"hubProfile,omitempty" tf:"hub_profile,omitempty"`

	// The ID of the Kubernetes Fleet Manager.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Azure Region where the Kubernetes Fleet Manager should exist. Changing this forces a new Kubernetes Fleet Manager to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the name of the Resource Group within which this Kubernetes Fleet Manager should exist. Changing this forces a new Kubernetes Fleet Manager to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A mapping of tags which should be assigned to the Kubernetes Fleet Manager.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type KubernetesFleetManagerParameters struct {

	// A hub_profile block as defined below. The FleetHubProfile configures the Fleet's hub. Changing this forces a new Kubernetes Fleet Manager to be created.
	// +kubebuilder:validation:Optional
	HubProfile []HubProfileParameters `json:"hubProfile,omitempty" tf:"hub_profile,omitempty"`

	// The Azure Region where the Kubernetes Fleet Manager should exist. Changing this forces a new Kubernetes Fleet Manager to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the name of the Resource Group within which this Kubernetes Fleet Manager should exist. Changing this forces a new Kubernetes Fleet Manager to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags which should be assigned to the Kubernetes Fleet Manager.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// KubernetesFleetManagerSpec defines the desired state of KubernetesFleetManager
type KubernetesFleetManagerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     KubernetesFleetManagerParameters `json:"forProvider"`
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
	InitProvider KubernetesFleetManagerInitParameters `json:"initProvider,omitempty"`
}

// KubernetesFleetManagerStatus defines the observed state of KubernetesFleetManager.
type KubernetesFleetManagerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        KubernetesFleetManagerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// KubernetesFleetManager is the Schema for the KubernetesFleetManagers API. Manages a Kubernetes Fleet Manager.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type KubernetesFleetManager struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	Spec   KubernetesFleetManagerSpec   `json:"spec"`
	Status KubernetesFleetManagerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KubernetesFleetManagerList contains a list of KubernetesFleetManagers
type KubernetesFleetManagerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KubernetesFleetManager `json:"items"`
}

// Repository type metadata.
var (
	KubernetesFleetManager_Kind             = "KubernetesFleetManager"
	KubernetesFleetManager_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: KubernetesFleetManager_Kind}.String()
	KubernetesFleetManager_KindAPIVersion   = KubernetesFleetManager_Kind + "." + CRDGroupVersion.String()
	KubernetesFleetManager_GroupVersionKind = CRDGroupVersion.WithKind(KubernetesFleetManager_Kind)
)

func init() {
	SchemeBuilder.Register(&KubernetesFleetManager{}, &KubernetesFleetManagerList{})
}
