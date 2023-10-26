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

type BotChannelsRegistrationInitParameters struct {

	// The CMK Key Vault Key URL to encrypt the Bot Channels Registration with the Customer Managed Encryption Key.
	CmkKeyVaultURL *string `json:"cmkKeyVaultUrl,omitempty" tf:"cmk_key_vault_url,omitempty"`

	// The description of the Bot Channels Registration.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The Application Insights Application ID to associate with the Bot Channels Registration.
	DeveloperAppInsightsApplicationID *string `json:"developerAppInsightsApplicationId,omitempty" tf:"developer_app_insights_application_id,omitempty"`

	// The Application Insights Key to associate with the Bot Channels Registration.
	DeveloperAppInsightsKey *string `json:"developerAppInsightsKey,omitempty" tf:"developer_app_insights_key,omitempty"`

	// The name of the Bot Channels Registration will be displayed as. This defaults to name if not specified.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The Bot Channels Registration endpoint.
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// The icon URL to visually identify the Bot Channels Registration.
	IconURL *string `json:"iconUrl,omitempty" tf:"icon_url,omitempty"`

	// Is the Bot Channels Registration in an isolated network?
	IsolatedNetworkEnabled *bool `json:"isolatedNetworkEnabled,omitempty" tf:"isolated_network_enabled,omitempty"`

	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The Microsoft Application ID for the Bot Channels Registration. Changing this forces a new resource to be created.
	MicrosoftAppID *string `json:"microsoftAppId,omitempty" tf:"microsoft_app_id,omitempty"`

	// Specifies the name of the Bot Channels Registration. Changing this forces a new resource to be created. Must be globally unique.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Is the Bot Channels Registration in an isolated network?
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// The SKU of the Bot Channels Registration. Valid values include F0 or S1. Changing this forces a new resource to be created.
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// Is the streaming endpoint enabled for the Bot Channels Registration. Defaults to false.
	StreamingEndpointEnabled *bool `json:"streamingEndpointEnabled,omitempty" tf:"streaming_endpoint_enabled,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type BotChannelsRegistrationObservation struct {

	// The CMK Key Vault Key URL to encrypt the Bot Channels Registration with the Customer Managed Encryption Key.
	CmkKeyVaultURL *string `json:"cmkKeyVaultUrl,omitempty" tf:"cmk_key_vault_url,omitempty"`

	// The description of the Bot Channels Registration.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The Application Insights Application ID to associate with the Bot Channels Registration.
	DeveloperAppInsightsApplicationID *string `json:"developerAppInsightsApplicationId,omitempty" tf:"developer_app_insights_application_id,omitempty"`

	// The Application Insights Key to associate with the Bot Channels Registration.
	DeveloperAppInsightsKey *string `json:"developerAppInsightsKey,omitempty" tf:"developer_app_insights_key,omitempty"`

	// The name of the Bot Channels Registration will be displayed as. This defaults to name if not specified.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The Bot Channels Registration endpoint.
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// The ID of the Bot Channels Registration.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The icon URL to visually identify the Bot Channels Registration.
	IconURL *string `json:"iconUrl,omitempty" tf:"icon_url,omitempty"`

	// Is the Bot Channels Registration in an isolated network?
	IsolatedNetworkEnabled *bool `json:"isolatedNetworkEnabled,omitempty" tf:"isolated_network_enabled,omitempty"`

	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The Microsoft Application ID for the Bot Channels Registration. Changing this forces a new resource to be created.
	MicrosoftAppID *string `json:"microsoftAppId,omitempty" tf:"microsoft_app_id,omitempty"`

	// Specifies the name of the Bot Channels Registration. Changing this forces a new resource to be created. Must be globally unique.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Is the Bot Channels Registration in an isolated network?
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// The name of the resource group in which to create the Bot Channels Registration. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The SKU of the Bot Channels Registration. Valid values include F0 or S1. Changing this forces a new resource to be created.
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// Is the streaming endpoint enabled for the Bot Channels Registration. Defaults to false.
	StreamingEndpointEnabled *bool `json:"streamingEndpointEnabled,omitempty" tf:"streaming_endpoint_enabled,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type BotChannelsRegistrationParameters struct {

	// The CMK Key Vault Key URL to encrypt the Bot Channels Registration with the Customer Managed Encryption Key.
	// +kubebuilder:validation:Optional
	CmkKeyVaultURL *string `json:"cmkKeyVaultUrl,omitempty" tf:"cmk_key_vault_url,omitempty"`

	// The description of the Bot Channels Registration.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The Application Insights API Key to associate with the Bot Channels Registration.
	// +kubebuilder:validation:Optional
	DeveloperAppInsightsAPIKeySecretRef *v1.SecretKeySelector `json:"developerAppInsightsApiKeySecretRef,omitempty" tf:"-"`

	// The Application Insights Application ID to associate with the Bot Channels Registration.
	// +kubebuilder:validation:Optional
	DeveloperAppInsightsApplicationID *string `json:"developerAppInsightsApplicationId,omitempty" tf:"developer_app_insights_application_id,omitempty"`

	// The Application Insights Key to associate with the Bot Channels Registration.
	// +kubebuilder:validation:Optional
	DeveloperAppInsightsKey *string `json:"developerAppInsightsKey,omitempty" tf:"developer_app_insights_key,omitempty"`

	// The name of the Bot Channels Registration will be displayed as. This defaults to name if not specified.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The Bot Channels Registration endpoint.
	// +kubebuilder:validation:Optional
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// The icon URL to visually identify the Bot Channels Registration.
	// +kubebuilder:validation:Optional
	IconURL *string `json:"iconUrl,omitempty" tf:"icon_url,omitempty"`

	// Is the Bot Channels Registration in an isolated network?
	// +kubebuilder:validation:Optional
	IsolatedNetworkEnabled *bool `json:"isolatedNetworkEnabled,omitempty" tf:"isolated_network_enabled,omitempty"`

	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The Microsoft Application ID for the Bot Channels Registration. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	MicrosoftAppID *string `json:"microsoftAppId,omitempty" tf:"microsoft_app_id,omitempty"`

	// Specifies the name of the Bot Channels Registration. Changing this forces a new resource to be created. Must be globally unique.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Is the Bot Channels Registration in an isolated network?
	// +kubebuilder:validation:Optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// The name of the resource group in which to create the Bot Channels Registration. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The SKU of the Bot Channels Registration. Valid values include F0 or S1. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// Is the streaming endpoint enabled for the Bot Channels Registration. Defaults to false.
	// +kubebuilder:validation:Optional
	StreamingEndpointEnabled *bool `json:"streamingEndpointEnabled,omitempty" tf:"streaming_endpoint_enabled,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// BotChannelsRegistrationSpec defines the desired state of BotChannelsRegistration
type BotChannelsRegistrationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BotChannelsRegistrationParameters `json:"forProvider"`
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
	InitProvider BotChannelsRegistrationInitParameters `json:"initProvider,omitempty"`
}

// BotChannelsRegistrationStatus defines the observed state of BotChannelsRegistration.
type BotChannelsRegistrationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BotChannelsRegistrationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BotChannelsRegistration is the Schema for the BotChannelsRegistrations API. Manages a Bot Channels Registration.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type BotChannelsRegistration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.microsoftAppId) || (has(self.initProvider) && has(self.initProvider.microsoftAppId))",message="spec.forProvider.microsoftAppId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.sku) || (has(self.initProvider) && has(self.initProvider.sku))",message="spec.forProvider.sku is a required parameter"
	Spec   BotChannelsRegistrationSpec   `json:"spec"`
	Status BotChannelsRegistrationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BotChannelsRegistrationList contains a list of BotChannelsRegistrations
type BotChannelsRegistrationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BotChannelsRegistration `json:"items"`
}

// Repository type metadata.
var (
	BotChannelsRegistration_Kind             = "BotChannelsRegistration"
	BotChannelsRegistration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BotChannelsRegistration_Kind}.String()
	BotChannelsRegistration_KindAPIVersion   = BotChannelsRegistration_Kind + "." + CRDGroupVersion.String()
	BotChannelsRegistration_GroupVersionKind = CRDGroupVersion.WithKind(BotChannelsRegistration_Kind)
)

func init() {
	SchemeBuilder.Register(&BotChannelsRegistration{}, &BotChannelsRegistrationList{})
}
