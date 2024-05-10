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

type FlexibleServerConfigurationInitParameters struct {

	// Specifies the value of the MySQL Flexible Server Configuration. See the MySQL documentation for valid values.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type FlexibleServerConfigurationObservation struct {

	// The ID of the MySQL Flexible Server Configuration.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the resource group in which the MySQL Flexible Server exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Specifies the name of the MySQL Flexible Server. Changing this forces a new resource to be created.
	ServerName *string `json:"serverName,omitempty" tf:"server_name,omitempty"`

	// Specifies the value of the MySQL Flexible Server Configuration. See the MySQL documentation for valid values.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type FlexibleServerConfigurationParameters struct {

	// The name of the resource group in which the MySQL Flexible Server exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Specifies the name of the MySQL Flexible Server. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/dbformysql/v1beta1.FlexibleServer
	// +kubebuilder:validation:Optional
	ServerName *string `json:"serverName,omitempty" tf:"server_name,omitempty"`

	// Reference to a FlexibleServer in dbformysql to populate serverName.
	// +kubebuilder:validation:Optional
	ServerNameRef *v1.Reference `json:"serverNameRef,omitempty" tf:"-"`

	// Selector for a FlexibleServer in dbformysql to populate serverName.
	// +kubebuilder:validation:Optional
	ServerNameSelector *v1.Selector `json:"serverNameSelector,omitempty" tf:"-"`

	// Specifies the value of the MySQL Flexible Server Configuration. See the MySQL documentation for valid values.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

// FlexibleServerConfigurationSpec defines the desired state of FlexibleServerConfiguration
type FlexibleServerConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FlexibleServerConfigurationParameters `json:"forProvider"`
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
	InitProvider FlexibleServerConfigurationInitParameters `json:"initProvider,omitempty"`
}

// FlexibleServerConfigurationStatus defines the observed state of FlexibleServerConfiguration.
type FlexibleServerConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FlexibleServerConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// FlexibleServerConfiguration is the Schema for the FlexibleServerConfigurations API. Sets a MySQL Flexible Server Configuration value on a MySQL Flexible Server.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type FlexibleServerConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.value) || (has(self.initProvider) && has(self.initProvider.value))",message="spec.forProvider.value is a required parameter"
	Spec   FlexibleServerConfigurationSpec   `json:"spec"`
	Status FlexibleServerConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FlexibleServerConfigurationList contains a list of FlexibleServerConfigurations
type FlexibleServerConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlexibleServerConfiguration `json:"items"`
}

// Repository type metadata.
var (
	FlexibleServerConfiguration_Kind             = "FlexibleServerConfiguration"
	FlexibleServerConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FlexibleServerConfiguration_Kind}.String()
	FlexibleServerConfiguration_KindAPIVersion   = FlexibleServerConfiguration_Kind + "." + CRDGroupVersion.String()
	FlexibleServerConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(FlexibleServerConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&FlexibleServerConfiguration{}, &FlexibleServerConfigurationList{})
}
