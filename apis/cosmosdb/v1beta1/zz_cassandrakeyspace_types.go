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

type AutoscaleSettingsInitParameters struct {

	// The maximum throughput of the Cassandra KeySpace (RU/s). Must be between 1,000 and 1,000,000. Must be set in increments of 1,000. Conflicts with throughput.
	MaxThroughput *float64 `json:"maxThroughput,omitempty" tf:"max_throughput,omitempty"`
}

type AutoscaleSettingsObservation struct {

	// The maximum throughput of the Cassandra KeySpace (RU/s). Must be between 1,000 and 1,000,000. Must be set in increments of 1,000. Conflicts with throughput.
	MaxThroughput *float64 `json:"maxThroughput,omitempty" tf:"max_throughput,omitempty"`
}

type AutoscaleSettingsParameters struct {

	// The maximum throughput of the Cassandra KeySpace (RU/s). Must be between 1,000 and 1,000,000. Must be set in increments of 1,000. Conflicts with throughput.
	// +kubebuilder:validation:Optional
	MaxThroughput *float64 `json:"maxThroughput,omitempty" tf:"max_throughput,omitempty"`
}

type CassandraKeySpaceInitParameters struct {

	// An autoscale_settings block as defined below.
	AutoscaleSettings []AutoscaleSettingsInitParameters `json:"autoscaleSettings,omitempty" tf:"autoscale_settings,omitempty"`

	// The throughput of Cassandra KeySpace (RU/s). Must be set in increments of 100. The minimum value is 400.
	Throughput *float64 `json:"throughput,omitempty" tf:"throughput,omitempty"`
}

type CassandraKeySpaceObservation struct {

	// The name of the Cosmos DB Cassandra KeySpace to create the table within. Changing this forces a new resource to be created.
	AccountName *string `json:"accountName,omitempty" tf:"account_name,omitempty"`

	// An autoscale_settings block as defined below.
	AutoscaleSettings []AutoscaleSettingsObservation `json:"autoscaleSettings,omitempty" tf:"autoscale_settings,omitempty"`

	// the ID of the CosmosDB Cassandra KeySpace.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the resource group in which the Cosmos DB Cassandra KeySpace is created. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The throughput of Cassandra KeySpace (RU/s). Must be set in increments of 100. The minimum value is 400.
	Throughput *float64 `json:"throughput,omitempty" tf:"throughput,omitempty"`
}

type CassandraKeySpaceParameters struct {

	// The name of the Cosmos DB Cassandra KeySpace to create the table within. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/cosmosdb/v1beta1.Account
	// +kubebuilder:validation:Optional
	AccountName *string `json:"accountName,omitempty" tf:"account_name,omitempty"`

	// Reference to a Account in cosmosdb to populate accountName.
	// +kubebuilder:validation:Optional
	AccountNameRef *v1.Reference `json:"accountNameRef,omitempty" tf:"-"`

	// Selector for a Account in cosmosdb to populate accountName.
	// +kubebuilder:validation:Optional
	AccountNameSelector *v1.Selector `json:"accountNameSelector,omitempty" tf:"-"`

	// An autoscale_settings block as defined below.
	// +kubebuilder:validation:Optional
	AutoscaleSettings []AutoscaleSettingsParameters `json:"autoscaleSettings,omitempty" tf:"autoscale_settings,omitempty"`

	// The name of the resource group in which the Cosmos DB Cassandra KeySpace is created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The throughput of Cassandra KeySpace (RU/s). Must be set in increments of 100. The minimum value is 400.
	// +kubebuilder:validation:Optional
	Throughput *float64 `json:"throughput,omitempty" tf:"throughput,omitempty"`
}

// CassandraKeySpaceSpec defines the desired state of CassandraKeySpace
type CassandraKeySpaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CassandraKeySpaceParameters `json:"forProvider"`
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
	InitProvider CassandraKeySpaceInitParameters `json:"initProvider,omitempty"`
}

// CassandraKeySpaceStatus defines the observed state of CassandraKeySpace.
type CassandraKeySpaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CassandraKeySpaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CassandraKeySpace is the Schema for the CassandraKeySpaces API. Manages a Cassandra KeySpace within a Cosmos DB Account.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type CassandraKeySpace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CassandraKeySpaceSpec   `json:"spec"`
	Status            CassandraKeySpaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CassandraKeySpaceList contains a list of CassandraKeySpaces
type CassandraKeySpaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CassandraKeySpace `json:"items"`
}

// Repository type metadata.
var (
	CassandraKeySpace_Kind             = "CassandraKeySpace"
	CassandraKeySpace_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CassandraKeySpace_Kind}.String()
	CassandraKeySpace_KindAPIVersion   = CassandraKeySpace_Kind + "." + CRDGroupVersion.String()
	CassandraKeySpace_GroupVersionKind = CRDGroupVersion.WithKind(CassandraKeySpace_Kind)
)

func init() {
	SchemeBuilder.Register(&CassandraKeySpace{}, &CassandraKeySpaceList{})
}
