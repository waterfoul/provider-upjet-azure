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

type AuthorizationServerInitParameters struct {

	// The OAUTH Authorization Endpoint.
	AuthorizationEndpoint *string `json:"authorizationEndpoint,omitempty" tf:"authorization_endpoint,omitempty"`

	// The HTTP Verbs supported by the Authorization Endpoint. Possible values are DELETE, GET, HEAD, OPTIONS, PATCH, POST, PUT and TRACE.
	// +listType=set
	AuthorizationMethods []*string `json:"authorizationMethods,omitempty" tf:"authorization_methods,omitempty"`

	// The mechanism by which Access Tokens are passed to the API. Possible values are authorizationHeader and query.
	// +listType=set
	BearerTokenSendingMethods []*string `json:"bearerTokenSendingMethods,omitempty" tf:"bearer_token_sending_methods,omitempty"`

	// The Authentication Methods supported by the Token endpoint of this Authorization Server.. Possible values are Basic and Body.
	// +listType=set
	ClientAuthenticationMethod []*string `json:"clientAuthenticationMethod,omitempty" tf:"client_authentication_method,omitempty"`

	// The Client/App ID registered with this Authorization Server.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// The URI of page where Client/App Registration is performed for this Authorization Server.
	ClientRegistrationEndpoint *string `json:"clientRegistrationEndpoint,omitempty" tf:"client_registration_endpoint,omitempty"`

	// The Default Scope used when requesting an Access Token, specified as a string containing space-delimited values.
	DefaultScope *string `json:"defaultScope,omitempty" tf:"default_scope,omitempty"`

	// A description of the Authorization Server, which may contain HTML formatting tags.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The user-friendly name of this Authorization Server.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Form of Authorization Grants required when requesting an Access Token. Possible values are authorizationCode, clientCredentials, implicit and resourceOwnerPassword.
	// +listType=set
	GrantTypes []*string `json:"grantTypes,omitempty" tf:"grant_types,omitempty"`

	// The username associated with the Resource Owner.
	ResourceOwnerUsername *string `json:"resourceOwnerUsername,omitempty" tf:"resource_owner_username,omitempty"`

	// Does this Authorization Server support State? If this is set to true the client may use the state parameter to raise protocol security.
	SupportState *bool `json:"supportState,omitempty" tf:"support_state,omitempty"`

	// A token_body_parameter block as defined below.
	TokenBodyParameter []TokenBodyParameterInitParameters `json:"tokenBodyParameter,omitempty" tf:"token_body_parameter,omitempty"`

	// The OAUTH Token Endpoint.
	TokenEndpoint *string `json:"tokenEndpoint,omitempty" tf:"token_endpoint,omitempty"`
}

type AuthorizationServerObservation struct {

	// The name of the API Management Service in which this Authorization Server should be created. Changing this forces a new resource to be created.
	APIManagementName *string `json:"apiManagementName,omitempty" tf:"api_management_name,omitempty"`

	// The OAUTH Authorization Endpoint.
	AuthorizationEndpoint *string `json:"authorizationEndpoint,omitempty" tf:"authorization_endpoint,omitempty"`

	// The HTTP Verbs supported by the Authorization Endpoint. Possible values are DELETE, GET, HEAD, OPTIONS, PATCH, POST, PUT and TRACE.
	// +listType=set
	AuthorizationMethods []*string `json:"authorizationMethods,omitempty" tf:"authorization_methods,omitempty"`

	// The mechanism by which Access Tokens are passed to the API. Possible values are authorizationHeader and query.
	// +listType=set
	BearerTokenSendingMethods []*string `json:"bearerTokenSendingMethods,omitempty" tf:"bearer_token_sending_methods,omitempty"`

	// The Authentication Methods supported by the Token endpoint of this Authorization Server.. Possible values are Basic and Body.
	// +listType=set
	ClientAuthenticationMethod []*string `json:"clientAuthenticationMethod,omitempty" tf:"client_authentication_method,omitempty"`

	// The Client/App ID registered with this Authorization Server.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// The URI of page where Client/App Registration is performed for this Authorization Server.
	ClientRegistrationEndpoint *string `json:"clientRegistrationEndpoint,omitempty" tf:"client_registration_endpoint,omitempty"`

	// The Default Scope used when requesting an Access Token, specified as a string containing space-delimited values.
	DefaultScope *string `json:"defaultScope,omitempty" tf:"default_scope,omitempty"`

	// A description of the Authorization Server, which may contain HTML formatting tags.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The user-friendly name of this Authorization Server.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Form of Authorization Grants required when requesting an Access Token. Possible values are authorizationCode, clientCredentials, implicit and resourceOwnerPassword.
	// +listType=set
	GrantTypes []*string `json:"grantTypes,omitempty" tf:"grant_types,omitempty"`

	// The ID of the API Management Authorization Server.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The username associated with the Resource Owner.
	ResourceOwnerUsername *string `json:"resourceOwnerUsername,omitempty" tf:"resource_owner_username,omitempty"`

	// Does this Authorization Server support State? If this is set to true the client may use the state parameter to raise protocol security.
	SupportState *bool `json:"supportState,omitempty" tf:"support_state,omitempty"`

	// A token_body_parameter block as defined below.
	TokenBodyParameter []TokenBodyParameterObservation `json:"tokenBodyParameter,omitempty" tf:"token_body_parameter,omitempty"`

	// The OAUTH Token Endpoint.
	TokenEndpoint *string `json:"tokenEndpoint,omitempty" tf:"token_endpoint,omitempty"`
}

type AuthorizationServerParameters struct {

	// The name of the API Management Service in which this Authorization Server should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/apimanagement/v1beta1.Management
	// +kubebuilder:validation:Optional
	APIManagementName *string `json:"apiManagementName,omitempty" tf:"api_management_name,omitempty"`

	// Reference to a Management in apimanagement to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameRef *v1.Reference `json:"apiManagementNameRef,omitempty" tf:"-"`

	// Selector for a Management in apimanagement to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameSelector *v1.Selector `json:"apiManagementNameSelector,omitempty" tf:"-"`

	// The OAUTH Authorization Endpoint.
	// +kubebuilder:validation:Optional
	AuthorizationEndpoint *string `json:"authorizationEndpoint,omitempty" tf:"authorization_endpoint,omitempty"`

	// The HTTP Verbs supported by the Authorization Endpoint. Possible values are DELETE, GET, HEAD, OPTIONS, PATCH, POST, PUT and TRACE.
	// +kubebuilder:validation:Optional
	// +listType=set
	AuthorizationMethods []*string `json:"authorizationMethods,omitempty" tf:"authorization_methods,omitempty"`

	// The mechanism by which Access Tokens are passed to the API. Possible values are authorizationHeader and query.
	// +kubebuilder:validation:Optional
	// +listType=set
	BearerTokenSendingMethods []*string `json:"bearerTokenSendingMethods,omitempty" tf:"bearer_token_sending_methods,omitempty"`

	// The Authentication Methods supported by the Token endpoint of this Authorization Server.. Possible values are Basic and Body.
	// +kubebuilder:validation:Optional
	// +listType=set
	ClientAuthenticationMethod []*string `json:"clientAuthenticationMethod,omitempty" tf:"client_authentication_method,omitempty"`

	// The Client/App ID registered with this Authorization Server.
	// +kubebuilder:validation:Optional
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// The URI of page where Client/App Registration is performed for this Authorization Server.
	// +kubebuilder:validation:Optional
	ClientRegistrationEndpoint *string `json:"clientRegistrationEndpoint,omitempty" tf:"client_registration_endpoint,omitempty"`

	// The Client/App Secret registered with this Authorization Server.
	// +kubebuilder:validation:Optional
	ClientSecretSecretRef *v1.SecretKeySelector `json:"clientSecretSecretRef,omitempty" tf:"-"`

	// The Default Scope used when requesting an Access Token, specified as a string containing space-delimited values.
	// +kubebuilder:validation:Optional
	DefaultScope *string `json:"defaultScope,omitempty" tf:"default_scope,omitempty"`

	// A description of the Authorization Server, which may contain HTML formatting tags.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The user-friendly name of this Authorization Server.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Form of Authorization Grants required when requesting an Access Token. Possible values are authorizationCode, clientCredentials, implicit and resourceOwnerPassword.
	// +kubebuilder:validation:Optional
	// +listType=set
	GrantTypes []*string `json:"grantTypes,omitempty" tf:"grant_types,omitempty"`

	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The password associated with the Resource Owner.
	// +kubebuilder:validation:Optional
	ResourceOwnerPasswordSecretRef *v1.SecretKeySelector `json:"resourceOwnerPasswordSecretRef,omitempty" tf:"-"`

	// The username associated with the Resource Owner.
	// +kubebuilder:validation:Optional
	ResourceOwnerUsername *string `json:"resourceOwnerUsername,omitempty" tf:"resource_owner_username,omitempty"`

	// Does this Authorization Server support State? If this is set to true the client may use the state parameter to raise protocol security.
	// +kubebuilder:validation:Optional
	SupportState *bool `json:"supportState,omitempty" tf:"support_state,omitempty"`

	// A token_body_parameter block as defined below.
	// +kubebuilder:validation:Optional
	TokenBodyParameter []TokenBodyParameterParameters `json:"tokenBodyParameter,omitempty" tf:"token_body_parameter,omitempty"`

	// The OAUTH Token Endpoint.
	// +kubebuilder:validation:Optional
	TokenEndpoint *string `json:"tokenEndpoint,omitempty" tf:"token_endpoint,omitempty"`
}

type TokenBodyParameterInitParameters struct {

	// The Name of the Parameter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The Value of the Parameter.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TokenBodyParameterObservation struct {

	// The Name of the Parameter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The Value of the Parameter.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TokenBodyParameterParameters struct {

	// The Name of the Parameter.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The Value of the Parameter.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

// AuthorizationServerSpec defines the desired state of AuthorizationServer
type AuthorizationServerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AuthorizationServerParameters `json:"forProvider"`
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
	InitProvider AuthorizationServerInitParameters `json:"initProvider,omitempty"`
}

// AuthorizationServerStatus defines the observed state of AuthorizationServer.
type AuthorizationServerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AuthorizationServerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AuthorizationServer is the Schema for the AuthorizationServers API. Manages an Authorization Server within an API Management Service.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type AuthorizationServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.authorizationEndpoint) || (has(self.initProvider) && has(self.initProvider.authorizationEndpoint))",message="spec.forProvider.authorizationEndpoint is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.authorizationMethods) || (has(self.initProvider) && has(self.initProvider.authorizationMethods))",message="spec.forProvider.authorizationMethods is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.clientId) || (has(self.initProvider) && has(self.initProvider.clientId))",message="spec.forProvider.clientId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.clientRegistrationEndpoint) || (has(self.initProvider) && has(self.initProvider.clientRegistrationEndpoint))",message="spec.forProvider.clientRegistrationEndpoint is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.displayName) || (has(self.initProvider) && has(self.initProvider.displayName))",message="spec.forProvider.displayName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.grantTypes) || (has(self.initProvider) && has(self.initProvider.grantTypes))",message="spec.forProvider.grantTypes is a required parameter"
	Spec   AuthorizationServerSpec   `json:"spec"`
	Status AuthorizationServerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AuthorizationServerList contains a list of AuthorizationServers
type AuthorizationServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AuthorizationServer `json:"items"`
}

// Repository type metadata.
var (
	AuthorizationServer_Kind             = "AuthorizationServer"
	AuthorizationServer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AuthorizationServer_Kind}.String()
	AuthorizationServer_KindAPIVersion   = AuthorizationServer_Kind + "." + CRDGroupVersion.String()
	AuthorizationServer_GroupVersionKind = CRDGroupVersion.WithKind(AuthorizationServer_Kind)
)

func init() {
	SchemeBuilder.Register(&AuthorizationServer{}, &AuthorizationServerList{})
}
