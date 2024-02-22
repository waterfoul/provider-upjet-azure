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

type SubscriptionPolicyExemptionInitParameters struct {

	// A description to use for this Policy Exemption.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A friendly display name to use for this Policy Exemption.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The category of this policy exemption. Possible values are Waiver and Mitigated.
	ExemptionCategory *string `json:"exemptionCategory,omitempty" tf:"exemption_category,omitempty"`

	// The expiration date and time in UTC ISO 8601 format of this policy exemption.
	ExpiresOn *string `json:"expiresOn,omitempty" tf:"expires_on,omitempty"`

	// The metadata for this policy exemption. This is a JSON string representing additional metadata that should be stored with the policy exemption.
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The ID of the Policy Assignment to be exempted at the specified Scope. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/authorization/v1beta1.SubscriptionPolicyAssignment
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	PolicyAssignmentID *string `json:"policyAssignmentId,omitempty" tf:"policy_assignment_id,omitempty"`

	// Reference to a SubscriptionPolicyAssignment in authorization to populate policyAssignmentId.
	// +kubebuilder:validation:Optional
	PolicyAssignmentIDRef *v1.Reference `json:"policyAssignmentIdRef,omitempty" tf:"-"`

	// Selector for a SubscriptionPolicyAssignment in authorization to populate policyAssignmentId.
	// +kubebuilder:validation:Optional
	PolicyAssignmentIDSelector *v1.Selector `json:"policyAssignmentIdSelector,omitempty" tf:"-"`

	// The policy definition reference ID list when the associated policy assignment is an assignment of a policy set definition.
	PolicyDefinitionReferenceIds []*string `json:"policyDefinitionReferenceIds,omitempty" tf:"policy_definition_reference_ids,omitempty"`

	// The Subscription ID where the Policy Exemption should be applied. Changing this forces a new resource to be created.
	SubscriptionID *string `json:"subscriptionId,omitempty" tf:"subscription_id,omitempty"`
}

type SubscriptionPolicyExemptionObservation struct {

	// A description to use for this Policy Exemption.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A friendly display name to use for this Policy Exemption.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The category of this policy exemption. Possible values are Waiver and Mitigated.
	ExemptionCategory *string `json:"exemptionCategory,omitempty" tf:"exemption_category,omitempty"`

	// The expiration date and time in UTC ISO 8601 format of this policy exemption.
	ExpiresOn *string `json:"expiresOn,omitempty" tf:"expires_on,omitempty"`

	// The Policy Exemption id.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The metadata for this policy exemption. This is a JSON string representing additional metadata that should be stored with the policy exemption.
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The ID of the Policy Assignment to be exempted at the specified Scope. Changing this forces a new resource to be created.
	PolicyAssignmentID *string `json:"policyAssignmentId,omitempty" tf:"policy_assignment_id,omitempty"`

	// The policy definition reference ID list when the associated policy assignment is an assignment of a policy set definition.
	PolicyDefinitionReferenceIds []*string `json:"policyDefinitionReferenceIds,omitempty" tf:"policy_definition_reference_ids,omitempty"`

	// The Subscription ID where the Policy Exemption should be applied. Changing this forces a new resource to be created.
	SubscriptionID *string `json:"subscriptionId,omitempty" tf:"subscription_id,omitempty"`
}

type SubscriptionPolicyExemptionParameters struct {

	// A description to use for this Policy Exemption.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A friendly display name to use for this Policy Exemption.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The category of this policy exemption. Possible values are Waiver and Mitigated.
	// +kubebuilder:validation:Optional
	ExemptionCategory *string `json:"exemptionCategory,omitempty" tf:"exemption_category,omitempty"`

	// The expiration date and time in UTC ISO 8601 format of this policy exemption.
	// +kubebuilder:validation:Optional
	ExpiresOn *string `json:"expiresOn,omitempty" tf:"expires_on,omitempty"`

	// The metadata for this policy exemption. This is a JSON string representing additional metadata that should be stored with the policy exemption.
	// +kubebuilder:validation:Optional
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The ID of the Policy Assignment to be exempted at the specified Scope. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/authorization/v1beta1.SubscriptionPolicyAssignment
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PolicyAssignmentID *string `json:"policyAssignmentId,omitempty" tf:"policy_assignment_id,omitempty"`

	// Reference to a SubscriptionPolicyAssignment in authorization to populate policyAssignmentId.
	// +kubebuilder:validation:Optional
	PolicyAssignmentIDRef *v1.Reference `json:"policyAssignmentIdRef,omitempty" tf:"-"`

	// Selector for a SubscriptionPolicyAssignment in authorization to populate policyAssignmentId.
	// +kubebuilder:validation:Optional
	PolicyAssignmentIDSelector *v1.Selector `json:"policyAssignmentIdSelector,omitempty" tf:"-"`

	// The policy definition reference ID list when the associated policy assignment is an assignment of a policy set definition.
	// +kubebuilder:validation:Optional
	PolicyDefinitionReferenceIds []*string `json:"policyDefinitionReferenceIds,omitempty" tf:"policy_definition_reference_ids,omitempty"`

	// The Subscription ID where the Policy Exemption should be applied. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SubscriptionID *string `json:"subscriptionId,omitempty" tf:"subscription_id,omitempty"`
}

// SubscriptionPolicyExemptionSpec defines the desired state of SubscriptionPolicyExemption
type SubscriptionPolicyExemptionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubscriptionPolicyExemptionParameters `json:"forProvider"`
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
	InitProvider SubscriptionPolicyExemptionInitParameters `json:"initProvider,omitempty"`
}

// SubscriptionPolicyExemptionStatus defines the observed state of SubscriptionPolicyExemption.
type SubscriptionPolicyExemptionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubscriptionPolicyExemptionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SubscriptionPolicyExemption is the Schema for the SubscriptionPolicyExemptions API. Manages a Subscription Policy Exemption.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SubscriptionPolicyExemption struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.exemptionCategory) || (has(self.initProvider) && has(self.initProvider.exemptionCategory))",message="spec.forProvider.exemptionCategory is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.subscriptionId) || (has(self.initProvider) && has(self.initProvider.subscriptionId))",message="spec.forProvider.subscriptionId is a required parameter"
	Spec   SubscriptionPolicyExemptionSpec   `json:"spec"`
	Status SubscriptionPolicyExemptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubscriptionPolicyExemptionList contains a list of SubscriptionPolicyExemptions
type SubscriptionPolicyExemptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SubscriptionPolicyExemption `json:"items"`
}

// Repository type metadata.
var (
	SubscriptionPolicyExemption_Kind             = "SubscriptionPolicyExemption"
	SubscriptionPolicyExemption_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SubscriptionPolicyExemption_Kind}.String()
	SubscriptionPolicyExemption_KindAPIVersion   = SubscriptionPolicyExemption_Kind + "." + CRDGroupVersion.String()
	SubscriptionPolicyExemption_GroupVersionKind = CRDGroupVersion.WithKind(SubscriptionPolicyExemption_Kind)
)

func init() {
	SchemeBuilder.Register(&SubscriptionPolicyExemption{}, &SubscriptionPolicyExemptionList{})
}
