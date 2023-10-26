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

type LogAnalyticsLinkedServiceInitParameters struct {

	// The ID of the writable Resource that will be linked to the workspace. This should be used for linking to a Log Analytics Cluster resource.
	WriteAccessID *string `json:"writeAccessId,omitempty" tf:"write_access_id,omitempty"`
}

type LogAnalyticsLinkedServiceObservation struct {

	// The Log Analytics Linked Service ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The generated name of the Linked Service. The format for this attribute is always <workspace name>/<linked service type>(e.g. workspace1/Automation or workspace1/Cluster)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the readable Resource that will be linked to the workspace. This should be used for linking to an Automation Account resource.
	ReadAccessID *string `json:"readAccessId,omitempty" tf:"read_access_id,omitempty"`

	// The name of the resource group in which the Log Analytics Linked Service is created. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The ID of the Log Analytics Workspace that will contain the Log Analytics Linked Service resource.
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`

	// The ID of the writable Resource that will be linked to the workspace. This should be used for linking to a Log Analytics Cluster resource.
	WriteAccessID *string `json:"writeAccessId,omitempty" tf:"write_access_id,omitempty"`
}

type LogAnalyticsLinkedServiceParameters struct {

	// The ID of the readable Resource that will be linked to the workspace. This should be used for linking to an Automation Account resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/automation/v1beta1.Account
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ReadAccessID *string `json:"readAccessId,omitempty" tf:"read_access_id,omitempty"`

	// Reference to a Account in automation to populate readAccessId.
	// +kubebuilder:validation:Optional
	ReadAccessIDRef *v1.Reference `json:"readAccessIdRef,omitempty" tf:"-"`

	// Selector for a Account in automation to populate readAccessId.
	// +kubebuilder:validation:Optional
	ReadAccessIDSelector *v1.Selector `json:"readAccessIdSelector,omitempty" tf:"-"`

	// The name of the resource group in which the Log Analytics Linked Service is created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The ID of the Log Analytics Workspace that will contain the Log Analytics Linked Service resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/operationalinsights/v1beta1.Workspace
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`

	// Reference to a Workspace in operationalinsights to populate workspaceId.
	// +kubebuilder:validation:Optional
	WorkspaceIDRef *v1.Reference `json:"workspaceIdRef,omitempty" tf:"-"`

	// Selector for a Workspace in operationalinsights to populate workspaceId.
	// +kubebuilder:validation:Optional
	WorkspaceIDSelector *v1.Selector `json:"workspaceIdSelector,omitempty" tf:"-"`

	// The ID of the writable Resource that will be linked to the workspace. This should be used for linking to a Log Analytics Cluster resource.
	// +kubebuilder:validation:Optional
	WriteAccessID *string `json:"writeAccessId,omitempty" tf:"write_access_id,omitempty"`
}

// LogAnalyticsLinkedServiceSpec defines the desired state of LogAnalyticsLinkedService
type LogAnalyticsLinkedServiceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LogAnalyticsLinkedServiceParameters `json:"forProvider"`
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
	InitProvider LogAnalyticsLinkedServiceInitParameters `json:"initProvider,omitempty"`
}

// LogAnalyticsLinkedServiceStatus defines the observed state of LogAnalyticsLinkedService.
type LogAnalyticsLinkedServiceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LogAnalyticsLinkedServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LogAnalyticsLinkedService is the Schema for the LogAnalyticsLinkedServices API. Manages a Log Analytics Linked Service.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LogAnalyticsLinkedService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LogAnalyticsLinkedServiceSpec   `json:"spec"`
	Status            LogAnalyticsLinkedServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LogAnalyticsLinkedServiceList contains a list of LogAnalyticsLinkedServices
type LogAnalyticsLinkedServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LogAnalyticsLinkedService `json:"items"`
}

// Repository type metadata.
var (
	LogAnalyticsLinkedService_Kind             = "LogAnalyticsLinkedService"
	LogAnalyticsLinkedService_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LogAnalyticsLinkedService_Kind}.String()
	LogAnalyticsLinkedService_KindAPIVersion   = LogAnalyticsLinkedService_Kind + "." + CRDGroupVersion.String()
	LogAnalyticsLinkedService_GroupVersionKind = CRDGroupVersion.WithKind(LogAnalyticsLinkedService_Kind)
)

func init() {
	SchemeBuilder.Register(&LogAnalyticsLinkedService{}, &LogAnalyticsLinkedServiceList{})
}
