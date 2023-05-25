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

type SecurityRuleObservation_2 struct {

	// Specifies whether network traffic is allowed or denied. Possible values are Allow and Deny.
	Access *string `json:"access,omitempty" tf:"access,omitempty"`

	// A description for this rule. Restricted to 140 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// CIDR or destination IP range or * to match any IP. Tags such as VirtualNetwork, AzureLoadBalancer and Internet can also be used. Besides, it also supports all available Service Tags like ‘Sql.WestEurope‘, ‘Storage.EastUS‘, etc. You can list the available service tags with the CLI: shell az network list-service-tags --location westcentralus. For further information please see Azure CLI - az network list-service-tags. This is required if destination_address_prefixes is not specified.
	DestinationAddressPrefix *string `json:"destinationAddressPrefix,omitempty" tf:"destination_address_prefix,omitempty"`

	// List of destination address prefixes. Tags may not be used. This is required if destination_address_prefix is not specified.
	DestinationAddressPrefixes []*string `json:"destinationAddressPrefixes,omitempty" tf:"destination_address_prefixes,omitempty"`

	// A List of destination Application Security Group IDs
	DestinationApplicationSecurityGroupIds []*string `json:"destinationApplicationSecurityGroupIds,omitempty" tf:"destination_application_security_group_ids,omitempty"`

	// Destination Port or Range. Integer or range between 0 and 65535 or * to match any. This is required if destination_port_ranges is not specified.
	DestinationPortRange *string `json:"destinationPortRange,omitempty" tf:"destination_port_range,omitempty"`

	// List of destination ports or port ranges. This is required if destination_port_range is not specified.
	DestinationPortRanges []*string `json:"destinationPortRanges,omitempty" tf:"destination_port_ranges,omitempty"`

	// The direction specifies if rule will be evaluated on incoming or outgoing traffic. Possible values are Inbound and Outbound.
	Direction *string `json:"direction,omitempty" tf:"direction,omitempty"`

	// The ID of the Network Security Rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the Network Security Group that we want to attach the rule to. Changing this forces a new resource to be created.
	NetworkSecurityGroupName *string `json:"networkSecurityGroupName,omitempty" tf:"network_security_group_name,omitempty"`

	// Specifies the priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// Network protocol this rule applies to. Possible values include Tcp, Udp, Icmp, Esp, Ah or * (which matches all).
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The name of the resource group in which to create the Network Security Rule. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// CIDR or source IP range or * to match any IP. Tags such as VirtualNetwork, AzureLoadBalancer and Internet can also be used. This is required if source_address_prefixes is not specified.
	SourceAddressPrefix *string `json:"sourceAddressPrefix,omitempty" tf:"source_address_prefix,omitempty"`

	// List of source address prefixes. Tags may not be used. This is required if source_address_prefix is not specified.
	SourceAddressPrefixes []*string `json:"sourceAddressPrefixes,omitempty" tf:"source_address_prefixes,omitempty"`

	// A List of source Application Security Group IDs
	SourceApplicationSecurityGroupIds []*string `json:"sourceApplicationSecurityGroupIds,omitempty" tf:"source_application_security_group_ids,omitempty"`

	// Source Port or Range. Integer or range between 0 and 65535 or * to match any. This is required if source_port_ranges is not specified.
	SourcePortRange *string `json:"sourcePortRange,omitempty" tf:"source_port_range,omitempty"`

	// List of source ports or port ranges. This is required if source_port_range is not specified.
	SourcePortRanges []*string `json:"sourcePortRanges,omitempty" tf:"source_port_ranges,omitempty"`
}

type SecurityRuleParameters_2 struct {

	// Specifies whether network traffic is allowed or denied. Possible values are Allow and Deny.
	// +kubebuilder:validation:Optional
	Access *string `json:"access,omitempty" tf:"access,omitempty"`

	// A description for this rule. Restricted to 140 characters.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// CIDR or destination IP range or * to match any IP. Tags such as VirtualNetwork, AzureLoadBalancer and Internet can also be used. Besides, it also supports all available Service Tags like ‘Sql.WestEurope‘, ‘Storage.EastUS‘, etc. You can list the available service tags with the CLI: shell az network list-service-tags --location westcentralus. For further information please see Azure CLI - az network list-service-tags. This is required if destination_address_prefixes is not specified.
	// +kubebuilder:validation:Optional
	DestinationAddressPrefix *string `json:"destinationAddressPrefix,omitempty" tf:"destination_address_prefix,omitempty"`

	// List of destination address prefixes. Tags may not be used. This is required if destination_address_prefix is not specified.
	// +kubebuilder:validation:Optional
	DestinationAddressPrefixes []*string `json:"destinationAddressPrefixes,omitempty" tf:"destination_address_prefixes,omitempty"`

	// A List of destination Application Security Group IDs
	// +kubebuilder:validation:Optional
	DestinationApplicationSecurityGroupIds []*string `json:"destinationApplicationSecurityGroupIds,omitempty" tf:"destination_application_security_group_ids,omitempty"`

	// Destination Port or Range. Integer or range between 0 and 65535 or * to match any. This is required if destination_port_ranges is not specified.
	// +kubebuilder:validation:Optional
	DestinationPortRange *string `json:"destinationPortRange,omitempty" tf:"destination_port_range,omitempty"`

	// List of destination ports or port ranges. This is required if destination_port_range is not specified.
	// +kubebuilder:validation:Optional
	DestinationPortRanges []*string `json:"destinationPortRanges,omitempty" tf:"destination_port_ranges,omitempty"`

	// The direction specifies if rule will be evaluated on incoming or outgoing traffic. Possible values are Inbound and Outbound.
	// +kubebuilder:validation:Optional
	Direction *string `json:"direction,omitempty" tf:"direction,omitempty"`

	// The name of the Network Security Group that we want to attach the rule to. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=SecurityGroup
	// +kubebuilder:validation:Optional
	NetworkSecurityGroupName *string `json:"networkSecurityGroupName,omitempty" tf:"network_security_group_name,omitempty"`

	// Reference to a SecurityGroup to populate networkSecurityGroupName.
	// +kubebuilder:validation:Optional
	NetworkSecurityGroupNameRef *v1.Reference `json:"networkSecurityGroupNameRef,omitempty" tf:"-"`

	// Selector for a SecurityGroup to populate networkSecurityGroupName.
	// +kubebuilder:validation:Optional
	NetworkSecurityGroupNameSelector *v1.Selector `json:"networkSecurityGroupNameSelector,omitempty" tf:"-"`

	// Specifies the priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// Network protocol this rule applies to. Possible values include Tcp, Udp, Icmp, Esp, Ah or * (which matches all).
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The name of the resource group in which to create the Network Security Rule. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// CIDR or source IP range or * to match any IP. Tags such as VirtualNetwork, AzureLoadBalancer and Internet can also be used. This is required if source_address_prefixes is not specified.
	// +kubebuilder:validation:Optional
	SourceAddressPrefix *string `json:"sourceAddressPrefix,omitempty" tf:"source_address_prefix,omitempty"`

	// List of source address prefixes. Tags may not be used. This is required if source_address_prefix is not specified.
	// +kubebuilder:validation:Optional
	SourceAddressPrefixes []*string `json:"sourceAddressPrefixes,omitempty" tf:"source_address_prefixes,omitempty"`

	// A List of source Application Security Group IDs
	// +kubebuilder:validation:Optional
	SourceApplicationSecurityGroupIds []*string `json:"sourceApplicationSecurityGroupIds,omitempty" tf:"source_application_security_group_ids,omitempty"`

	// Source Port or Range. Integer or range between 0 and 65535 or * to match any. This is required if source_port_ranges is not specified.
	// +kubebuilder:validation:Optional
	SourcePortRange *string `json:"sourcePortRange,omitempty" tf:"source_port_range,omitempty"`

	// List of source ports or port ranges. This is required if source_port_range is not specified.
	// +kubebuilder:validation:Optional
	SourcePortRanges []*string `json:"sourcePortRanges,omitempty" tf:"source_port_ranges,omitempty"`
}

// SecurityRuleSpec defines the desired state of SecurityRule
type SecurityRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecurityRuleParameters_2 `json:"forProvider"`
}

// SecurityRuleStatus defines the observed state of SecurityRule.
type SecurityRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecurityRuleObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityRule is the Schema for the SecurityRules API. Manages a Network Security Rule.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SecurityRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.access)",message="access is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.direction)",message="direction is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.priority)",message="priority is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.protocol)",message="protocol is a required parameter"
	Spec   SecurityRuleSpec   `json:"spec"`
	Status SecurityRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityRuleList contains a list of SecurityRules
type SecurityRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityRule `json:"items"`
}

// Repository type metadata.
var (
	SecurityRule_Kind             = "SecurityRule"
	SecurityRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecurityRule_Kind}.String()
	SecurityRule_KindAPIVersion   = SecurityRule_Kind + "." + CRDGroupVersion.String()
	SecurityRule_GroupVersionKind = CRDGroupVersion.WithKind(SecurityRule_Kind)
)

func init() {
	SchemeBuilder.Register(&SecurityRule{}, &SecurityRuleList{})
}
