// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SQLStoredProcedureInitParameters struct {

	// The body of the stored procedure.
	Body *string `json:"body,omitempty" tf:"body,omitempty"`
}

type SQLStoredProcedureObservation struct {

	// The name of the Cosmos DB Account to create the stored procedure within. Changing this forces a new resource to be created.
	AccountName *string `json:"accountName,omitempty" tf:"account_name,omitempty"`

	// The body of the stored procedure.
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// The name of the Cosmos DB SQL Container to create the stored procedure within. Changing this forces a new resource to be created.
	ContainerName *string `json:"containerName,omitempty" tf:"container_name,omitempty"`

	// The name of the Cosmos DB SQL Database to create the stored procedure within. Changing this forces a new resource to be created.
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// The ID of the Cosmos DB SQL Stored Procedure.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the resource group in which the Cosmos DB SQL Database is created. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`
}

type SQLStoredProcedureParameters struct {

	// The name of the Cosmos DB Account to create the stored procedure within. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/cosmosdb/v1beta1.Account
	// +kubebuilder:validation:Optional
	AccountName *string `json:"accountName,omitempty" tf:"account_name,omitempty"`

	// Reference to a Account in cosmosdb to populate accountName.
	// +kubebuilder:validation:Optional
	AccountNameRef *v1.Reference `json:"accountNameRef,omitempty" tf:"-"`

	// Selector for a Account in cosmosdb to populate accountName.
	// +kubebuilder:validation:Optional
	AccountNameSelector *v1.Selector `json:"accountNameSelector,omitempty" tf:"-"`

	// The body of the stored procedure.
	// +kubebuilder:validation:Optional
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// The name of the Cosmos DB SQL Container to create the stored procedure within. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/cosmosdb/v1beta1.SQLContainer
	// +kubebuilder:validation:Optional
	ContainerName *string `json:"containerName,omitempty" tf:"container_name,omitempty"`

	// Reference to a SQLContainer in cosmosdb to populate containerName.
	// +kubebuilder:validation:Optional
	ContainerNameRef *v1.Reference `json:"containerNameRef,omitempty" tf:"-"`

	// Selector for a SQLContainer in cosmosdb to populate containerName.
	// +kubebuilder:validation:Optional
	ContainerNameSelector *v1.Selector `json:"containerNameSelector,omitempty" tf:"-"`

	// The name of the Cosmos DB SQL Database to create the stored procedure within. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/cosmosdb/v1beta1.SQLDatabase
	// +kubebuilder:validation:Optional
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// Reference to a SQLDatabase in cosmosdb to populate databaseName.
	// +kubebuilder:validation:Optional
	DatabaseNameRef *v1.Reference `json:"databaseNameRef,omitempty" tf:"-"`

	// Selector for a SQLDatabase in cosmosdb to populate databaseName.
	// +kubebuilder:validation:Optional
	DatabaseNameSelector *v1.Selector `json:"databaseNameSelector,omitempty" tf:"-"`

	// The name of the resource group in which the Cosmos DB SQL Database is created. Changing this forces a new resource to be created.
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

// SQLStoredProcedureSpec defines the desired state of SQLStoredProcedure
type SQLStoredProcedureSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SQLStoredProcedureParameters `json:"forProvider"`
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
	InitProvider SQLStoredProcedureInitParameters `json:"initProvider,omitempty"`
}

// SQLStoredProcedureStatus defines the observed state of SQLStoredProcedure.
type SQLStoredProcedureStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SQLStoredProcedureObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SQLStoredProcedure is the Schema for the SQLStoredProcedures API. Manages a SQL Stored Procedure within a Cosmos DB Account SQL Database.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SQLStoredProcedure struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.body) || (has(self.initProvider) && has(self.initProvider.body))",message="spec.forProvider.body is a required parameter"
	Spec   SQLStoredProcedureSpec   `json:"spec"`
	Status SQLStoredProcedureStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SQLStoredProcedureList contains a list of SQLStoredProcedures
type SQLStoredProcedureList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SQLStoredProcedure `json:"items"`
}

// Repository type metadata.
var (
	SQLStoredProcedure_Kind             = "SQLStoredProcedure"
	SQLStoredProcedure_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SQLStoredProcedure_Kind}.String()
	SQLStoredProcedure_KindAPIVersion   = SQLStoredProcedure_Kind + "." + CRDGroupVersion.String()
	SQLStoredProcedure_GroupVersionKind = CRDGroupVersion.WithKind(SQLStoredProcedure_Kind)
)

func init() {
	SchemeBuilder.Register(&SQLStoredProcedure{}, &SQLStoredProcedureList{})
}
