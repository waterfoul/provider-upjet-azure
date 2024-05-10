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

	// Specifies the name of the PostgreSQL Configuration, which needs to be a valid PostgreSQL configuration name. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the PostgreSQL Flexible Server where we want to change configuration. Changing this forces a new PostgreSQL Flexible Server Configuration resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/dbforpostgresql/v1beta1.FlexibleServer
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// Reference to a FlexibleServer in dbforpostgresql to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDRef *v1.Reference `json:"serverIdRef,omitempty" tf:"-"`

	// Selector for a FlexibleServer in dbforpostgresql to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDSelector *v1.Selector `json:"serverIdSelector,omitempty" tf:"-"`

	// Specifies the value of the PostgreSQL Configuration. See the PostgreSQL documentation for valid values.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type FlexibleServerConfigurationObservation struct {

	// The ID of the PostgreSQL Configuration.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the name of the PostgreSQL Configuration, which needs to be a valid PostgreSQL configuration name. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the PostgreSQL Flexible Server where we want to change configuration. Changing this forces a new PostgreSQL Flexible Server Configuration resource.
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// Specifies the value of the PostgreSQL Configuration. See the PostgreSQL documentation for valid values.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type FlexibleServerConfigurationParameters struct {

	// Specifies the name of the PostgreSQL Configuration, which needs to be a valid PostgreSQL configuration name. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the PostgreSQL Flexible Server where we want to change configuration. Changing this forces a new PostgreSQL Flexible Server Configuration resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/dbforpostgresql/v1beta1.FlexibleServer
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// Reference to a FlexibleServer in dbforpostgresql to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDRef *v1.Reference `json:"serverIdRef,omitempty" tf:"-"`

	// Selector for a FlexibleServer in dbforpostgresql to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDSelector *v1.Selector `json:"serverIdSelector,omitempty" tf:"-"`

	// Specifies the value of the PostgreSQL Configuration. See the PostgreSQL documentation for valid values.
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

// FlexibleServerConfiguration is the Schema for the FlexibleServerConfigurations API. Sets a PostgreSQL Configuration value on a Azure PostgreSQL Flexible Server.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type FlexibleServerConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
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
