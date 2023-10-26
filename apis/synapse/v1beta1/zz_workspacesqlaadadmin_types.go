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

type WorkspaceSQLAADAdminInitParameters struct {

	// The login name of the Azure AD Administrator of this Synapse Workspace.
	Login *string `json:"login,omitempty" tf:"login,omitempty"`

	// The object id of the Azure AD Administrator of this Synapse Workspace.
	ObjectID *string `json:"objectId,omitempty" tf:"object_id,omitempty"`

	// The tenant id of the Azure AD Administrator of this Synapse Workspace.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type WorkspaceSQLAADAdminObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The login name of the Azure AD Administrator of this Synapse Workspace.
	Login *string `json:"login,omitempty" tf:"login,omitempty"`

	// The object id of the Azure AD Administrator of this Synapse Workspace.
	ObjectID *string `json:"objectId,omitempty" tf:"object_id,omitempty"`

	// The ID of the Synapse Workspace where the Azure AD Administrator should be configured.
	SynapseWorkspaceID *string `json:"synapseWorkspaceId,omitempty" tf:"synapse_workspace_id,omitempty"`

	// The tenant id of the Azure AD Administrator of this Synapse Workspace.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type WorkspaceSQLAADAdminParameters struct {

	// The login name of the Azure AD Administrator of this Synapse Workspace.
	// +kubebuilder:validation:Optional
	Login *string `json:"login,omitempty" tf:"login,omitempty"`

	// The object id of the Azure AD Administrator of this Synapse Workspace.
	// +kubebuilder:validation:Optional
	ObjectID *string `json:"objectId,omitempty" tf:"object_id,omitempty"`

	// The ID of the Synapse Workspace where the Azure AD Administrator should be configured.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/synapse/v1beta1.Workspace
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SynapseWorkspaceID *string `json:"synapseWorkspaceId,omitempty" tf:"synapse_workspace_id,omitempty"`

	// Reference to a Workspace in synapse to populate synapseWorkspaceId.
	// +kubebuilder:validation:Optional
	SynapseWorkspaceIDRef *v1.Reference `json:"synapseWorkspaceIdRef,omitempty" tf:"-"`

	// Selector for a Workspace in synapse to populate synapseWorkspaceId.
	// +kubebuilder:validation:Optional
	SynapseWorkspaceIDSelector *v1.Selector `json:"synapseWorkspaceIdSelector,omitempty" tf:"-"`

	// The tenant id of the Azure AD Administrator of this Synapse Workspace.
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

// WorkspaceSQLAADAdminSpec defines the desired state of WorkspaceSQLAADAdmin
type WorkspaceSQLAADAdminSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WorkspaceSQLAADAdminParameters `json:"forProvider"`
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
	InitProvider WorkspaceSQLAADAdminInitParameters `json:"initProvider,omitempty"`
}

// WorkspaceSQLAADAdminStatus defines the observed state of WorkspaceSQLAADAdmin.
type WorkspaceSQLAADAdminStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WorkspaceSQLAADAdminObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// WorkspaceSQLAADAdmin is the Schema for the WorkspaceSQLAADAdmins API. Manages Synapse Workspace AAD Admin
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type WorkspaceSQLAADAdmin struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.login) || (has(self.initProvider) && has(self.initProvider.login))",message="spec.forProvider.login is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.objectId) || (has(self.initProvider) && has(self.initProvider.objectId))",message="spec.forProvider.objectId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.tenantId) || (has(self.initProvider) && has(self.initProvider.tenantId))",message="spec.forProvider.tenantId is a required parameter"
	Spec   WorkspaceSQLAADAdminSpec   `json:"spec"`
	Status WorkspaceSQLAADAdminStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkspaceSQLAADAdminList contains a list of WorkspaceSQLAADAdmins
type WorkspaceSQLAADAdminList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorkspaceSQLAADAdmin `json:"items"`
}

// Repository type metadata.
var (
	WorkspaceSQLAADAdmin_Kind             = "WorkspaceSQLAADAdmin"
	WorkspaceSQLAADAdmin_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: WorkspaceSQLAADAdmin_Kind}.String()
	WorkspaceSQLAADAdmin_KindAPIVersion   = WorkspaceSQLAADAdmin_Kind + "." + CRDGroupVersion.String()
	WorkspaceSQLAADAdmin_GroupVersionKind = CRDGroupVersion.WithKind(WorkspaceSQLAADAdmin_Kind)
)

func init() {
	SchemeBuilder.Register(&WorkspaceSQLAADAdmin{}, &WorkspaceSQLAADAdminList{})
}
