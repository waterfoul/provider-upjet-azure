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

type AdditionalWorkspaceInitParameters struct {

	// A list of data types which sent to workspace. Possible values are Alerts and RawEvents.
	// +listType=set
	DataTypes []*string `json:"dataTypes,omitempty" tf:"data_types,omitempty"`

	// The resource ID of the Log Analytics Workspace.
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`
}

type AdditionalWorkspaceObservation struct {

	// A list of data types which sent to workspace. Possible values are Alerts and RawEvents.
	// +listType=set
	DataTypes []*string `json:"dataTypes,omitempty" tf:"data_types,omitempty"`

	// The resource ID of the Log Analytics Workspace.
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`
}

type AdditionalWorkspaceParameters struct {

	// A list of data types which sent to workspace. Possible values are Alerts and RawEvents.
	// +kubebuilder:validation:Optional
	// +listType=set
	DataTypes []*string `json:"dataTypes" tf:"data_types,omitempty"`

	// The resource ID of the Log Analytics Workspace.
	// +kubebuilder:validation:Optional
	WorkspaceID *string `json:"workspaceId" tf:"workspace_id,omitempty"`
}

type IOTSecuritySolutionInitParameters struct {

	// A additional_workspace block as defined below.
	AdditionalWorkspace []AdditionalWorkspaceInitParameters `json:"additionalWorkspace,omitempty" tf:"additional_workspace,omitempty"`

	// A list of disabled data sources for the Iot Security Solution. Possible value is TwinData.
	// +listType=set
	DisabledDataSources []*string `json:"disabledDataSources,omitempty" tf:"disabled_data_sources,omitempty"`

	// Specifies the Display Name for this Iot Security Solution.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Is the Iot Security Solution enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A list of data which is to exported to analytic workspace. Valid values include RawEvents.
	// +listType=set
	EventsToExport []*string `json:"eventsToExport,omitempty" tf:"events_to_export,omitempty"`

	// Specifies the IoT Hub resource IDs to which this Iot Security Solution is applied.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/devices/v1beta1.IOTHub
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("id",true)
	// +listType=set
	IOTHubIds []*string `json:"iothubIds,omitempty" tf:"iothub_ids,omitempty"`

	// References to IOTHub in devices to populate iothubIds.
	// +kubebuilder:validation:Optional
	IOTHubIdsRefs []v1.Reference `json:"iothubIdsRefs,omitempty" tf:"-"`

	// Selector for a list of IOTHub in devices to populate iothubIds.
	// +kubebuilder:validation:Optional
	IOTHubIdsSelector *v1.Selector `json:"iothubIdsSelector,omitempty" tf:"-"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the Log Analytics Workspace ID to which the security data will be sent.
	LogAnalyticsWorkspaceID *string `json:"logAnalyticsWorkspaceId,omitempty" tf:"log_analytics_workspace_id,omitempty"`

	// Should IP addressed be unmasked in the log? Defaults to false.
	LogUnmaskedIpsEnabled *bool `json:"logUnmaskedIpsEnabled,omitempty" tf:"log_unmasked_ips_enabled,omitempty"`

	// Specifies the name of the Iot Security Solution. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// An Azure Resource Graph query used to set the resources monitored.
	QueryForResources *string `json:"queryForResources,omitempty" tf:"query_for_resources,omitempty"`

	// A list of subscription Ids on which the user defined resources query should be executed.
	// +listType=set
	QuerySubscriptionIds []*string `json:"querySubscriptionIds,omitempty" tf:"query_subscription_ids,omitempty"`

	// A recommendations_enabled block of options to enable or disable as defined below.
	RecommendationsEnabled []RecommendationsEnabledInitParameters `json:"recommendationsEnabled,omitempty" tf:"recommendations_enabled,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type IOTSecuritySolutionObservation struct {

	// A additional_workspace block as defined below.
	AdditionalWorkspace []AdditionalWorkspaceObservation `json:"additionalWorkspace,omitempty" tf:"additional_workspace,omitempty"`

	// A list of disabled data sources for the Iot Security Solution. Possible value is TwinData.
	// +listType=set
	DisabledDataSources []*string `json:"disabledDataSources,omitempty" tf:"disabled_data_sources,omitempty"`

	// Specifies the Display Name for this Iot Security Solution.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Is the Iot Security Solution enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A list of data which is to exported to analytic workspace. Valid values include RawEvents.
	// +listType=set
	EventsToExport []*string `json:"eventsToExport,omitempty" tf:"events_to_export,omitempty"`

	// The ID of the Iot Security Solution resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the IoT Hub resource IDs to which this Iot Security Solution is applied.
	// +listType=set
	IOTHubIds []*string `json:"iothubIds,omitempty" tf:"iothub_ids,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the Log Analytics Workspace ID to which the security data will be sent.
	LogAnalyticsWorkspaceID *string `json:"logAnalyticsWorkspaceId,omitempty" tf:"log_analytics_workspace_id,omitempty"`

	// Should IP addressed be unmasked in the log? Defaults to false.
	LogUnmaskedIpsEnabled *bool `json:"logUnmaskedIpsEnabled,omitempty" tf:"log_unmasked_ips_enabled,omitempty"`

	// Specifies the name of the Iot Security Solution. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// An Azure Resource Graph query used to set the resources monitored.
	QueryForResources *string `json:"queryForResources,omitempty" tf:"query_for_resources,omitempty"`

	// A list of subscription Ids on which the user defined resources query should be executed.
	// +listType=set
	QuerySubscriptionIds []*string `json:"querySubscriptionIds,omitempty" tf:"query_subscription_ids,omitempty"`

	// A recommendations_enabled block of options to enable or disable as defined below.
	RecommendationsEnabled []RecommendationsEnabledObservation `json:"recommendationsEnabled,omitempty" tf:"recommendations_enabled,omitempty"`

	// Specifies the name of the resource group in which to create the Iot Security Solution. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type IOTSecuritySolutionParameters struct {

	// A additional_workspace block as defined below.
	// +kubebuilder:validation:Optional
	AdditionalWorkspace []AdditionalWorkspaceParameters `json:"additionalWorkspace,omitempty" tf:"additional_workspace,omitempty"`

	// A list of disabled data sources for the Iot Security Solution. Possible value is TwinData.
	// +kubebuilder:validation:Optional
	// +listType=set
	DisabledDataSources []*string `json:"disabledDataSources,omitempty" tf:"disabled_data_sources,omitempty"`

	// Specifies the Display Name for this Iot Security Solution.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Is the Iot Security Solution enabled? Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A list of data which is to exported to analytic workspace. Valid values include RawEvents.
	// +kubebuilder:validation:Optional
	// +listType=set
	EventsToExport []*string `json:"eventsToExport,omitempty" tf:"events_to_export,omitempty"`

	// Specifies the IoT Hub resource IDs to which this Iot Security Solution is applied.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/devices/v1beta1.IOTHub
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("id",true)
	// +kubebuilder:validation:Optional
	// +listType=set
	IOTHubIds []*string `json:"iothubIds,omitempty" tf:"iothub_ids,omitempty"`

	// References to IOTHub in devices to populate iothubIds.
	// +kubebuilder:validation:Optional
	IOTHubIdsRefs []v1.Reference `json:"iothubIdsRefs,omitempty" tf:"-"`

	// Selector for a list of IOTHub in devices to populate iothubIds.
	// +kubebuilder:validation:Optional
	IOTHubIdsSelector *v1.Selector `json:"iothubIdsSelector,omitempty" tf:"-"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the Log Analytics Workspace ID to which the security data will be sent.
	// +kubebuilder:validation:Optional
	LogAnalyticsWorkspaceID *string `json:"logAnalyticsWorkspaceId,omitempty" tf:"log_analytics_workspace_id,omitempty"`

	// Should IP addressed be unmasked in the log? Defaults to false.
	// +kubebuilder:validation:Optional
	LogUnmaskedIpsEnabled *bool `json:"logUnmaskedIpsEnabled,omitempty" tf:"log_unmasked_ips_enabled,omitempty"`

	// Specifies the name of the Iot Security Solution. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// An Azure Resource Graph query used to set the resources monitored.
	// +kubebuilder:validation:Optional
	QueryForResources *string `json:"queryForResources,omitempty" tf:"query_for_resources,omitempty"`

	// A list of subscription Ids on which the user defined resources query should be executed.
	// +kubebuilder:validation:Optional
	// +listType=set
	QuerySubscriptionIds []*string `json:"querySubscriptionIds,omitempty" tf:"query_subscription_ids,omitempty"`

	// A recommendations_enabled block of options to enable or disable as defined below.
	// +kubebuilder:validation:Optional
	RecommendationsEnabled []RecommendationsEnabledParameters `json:"recommendationsEnabled,omitempty" tf:"recommendations_enabled,omitempty"`

	// Specifies the name of the resource group in which to create the Iot Security Solution. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RecommendationsEnabledInitParameters struct {

	// Is Principal Authentication enabled for the ACR repository? Defaults to true.
	AcrAuthentication *bool `json:"acrAuthentication,omitempty" tf:"acr_authentication,omitempty"`

	// Is Agent send underutilized messages enabled? Defaults to true.
	AgentSendUnutilizedMsg *bool `json:"agentSendUnutilizedMsg,omitempty" tf:"agent_send_unutilized_msg,omitempty"`

	// Is Security related system configuration issues identified? Defaults to true.
	Baseline *bool `json:"baseline,omitempty" tf:"baseline,omitempty"`

	// Is IoT Edge Hub memory optimized? Defaults to true.
	EdgeHubMemOptimize *bool `json:"edgeHubMemOptimize,omitempty" tf:"edge_hub_mem_optimize,omitempty"`

	// Is logging configured for IoT Edge module? Defaults to true.
	EdgeLoggingOption *bool `json:"edgeLoggingOption,omitempty" tf:"edge_logging_option,omitempty"`

	// Is Default IP filter policy denied? Defaults to true.
	IPFilterDenyAll *bool `json:"ipFilterDenyAll,omitempty" tf:"ip_filter_deny_all,omitempty"`

	// Is IP filter rule source allowable IP range too large? Defaults to true.
	IPFilterPermissiveRule *bool `json:"ipFilterPermissiveRule,omitempty" tf:"ip_filter_permissive_rule,omitempty"`

	// Is inconsistent module settings enabled for SecurityGroup? Defaults to true.
	InconsistentModuleSettings *bool `json:"inconsistentModuleSettings,omitempty" tf:"inconsistent_module_settings,omitempty"`

	// is Azure IoT Security agent installed? Defaults to true.
	InstallAgent *bool `json:"installAgent,omitempty" tf:"install_agent,omitempty"`

	// Is any ports open on the device? Defaults to true.
	OpenPorts *bool `json:"openPorts,omitempty" tf:"open_ports,omitempty"`

	// Does firewall policy exist which allow necessary communication to/from the device? Defaults to true.
	PermissiveFirewallPolicy *bool `json:"permissiveFirewallPolicy,omitempty" tf:"permissive_firewall_policy,omitempty"`

	// Is only necessary addresses or ports are permitted in? Defaults to true.
	PermissiveInputFirewallRules *bool `json:"permissiveInputFirewallRules,omitempty" tf:"permissive_input_firewall_rules,omitempty"`

	// Is only necessary addresses or ports are permitted out? Defaults to true.
	PermissiveOutputFirewallRules *bool `json:"permissiveOutputFirewallRules,omitempty" tf:"permissive_output_firewall_rules,omitempty"`

	// Is high level permissions are needed for the module? Defaults to true.
	PrivilegedDockerOptions *bool `json:"privilegedDockerOptions,omitempty" tf:"privileged_docker_options,omitempty"`

	// Is any credentials shared among devices? Defaults to true.
	SharedCredentials *bool `json:"sharedCredentials,omitempty" tf:"shared_credentials,omitempty"`

	// Does TLS cipher suite need to be updated? Defaults to true.
	VulnerableTLSCipherSuite *bool `json:"vulnerableTlsCipherSuite,omitempty" tf:"vulnerable_tls_cipher_suite,omitempty"`
}

type RecommendationsEnabledObservation struct {

	// Is Principal Authentication enabled for the ACR repository? Defaults to true.
	AcrAuthentication *bool `json:"acrAuthentication,omitempty" tf:"acr_authentication,omitempty"`

	// Is Agent send underutilized messages enabled? Defaults to true.
	AgentSendUnutilizedMsg *bool `json:"agentSendUnutilizedMsg,omitempty" tf:"agent_send_unutilized_msg,omitempty"`

	// Is Security related system configuration issues identified? Defaults to true.
	Baseline *bool `json:"baseline,omitempty" tf:"baseline,omitempty"`

	// Is IoT Edge Hub memory optimized? Defaults to true.
	EdgeHubMemOptimize *bool `json:"edgeHubMemOptimize,omitempty" tf:"edge_hub_mem_optimize,omitempty"`

	// Is logging configured for IoT Edge module? Defaults to true.
	EdgeLoggingOption *bool `json:"edgeLoggingOption,omitempty" tf:"edge_logging_option,omitempty"`

	// Is Default IP filter policy denied? Defaults to true.
	IPFilterDenyAll *bool `json:"ipFilterDenyAll,omitempty" tf:"ip_filter_deny_all,omitempty"`

	// Is IP filter rule source allowable IP range too large? Defaults to true.
	IPFilterPermissiveRule *bool `json:"ipFilterPermissiveRule,omitempty" tf:"ip_filter_permissive_rule,omitempty"`

	// Is inconsistent module settings enabled for SecurityGroup? Defaults to true.
	InconsistentModuleSettings *bool `json:"inconsistentModuleSettings,omitempty" tf:"inconsistent_module_settings,omitempty"`

	// is Azure IoT Security agent installed? Defaults to true.
	InstallAgent *bool `json:"installAgent,omitempty" tf:"install_agent,omitempty"`

	// Is any ports open on the device? Defaults to true.
	OpenPorts *bool `json:"openPorts,omitempty" tf:"open_ports,omitempty"`

	// Does firewall policy exist which allow necessary communication to/from the device? Defaults to true.
	PermissiveFirewallPolicy *bool `json:"permissiveFirewallPolicy,omitempty" tf:"permissive_firewall_policy,omitempty"`

	// Is only necessary addresses or ports are permitted in? Defaults to true.
	PermissiveInputFirewallRules *bool `json:"permissiveInputFirewallRules,omitempty" tf:"permissive_input_firewall_rules,omitempty"`

	// Is only necessary addresses or ports are permitted out? Defaults to true.
	PermissiveOutputFirewallRules *bool `json:"permissiveOutputFirewallRules,omitempty" tf:"permissive_output_firewall_rules,omitempty"`

	// Is high level permissions are needed for the module? Defaults to true.
	PrivilegedDockerOptions *bool `json:"privilegedDockerOptions,omitempty" tf:"privileged_docker_options,omitempty"`

	// Is any credentials shared among devices? Defaults to true.
	SharedCredentials *bool `json:"sharedCredentials,omitempty" tf:"shared_credentials,omitempty"`

	// Does TLS cipher suite need to be updated? Defaults to true.
	VulnerableTLSCipherSuite *bool `json:"vulnerableTlsCipherSuite,omitempty" tf:"vulnerable_tls_cipher_suite,omitempty"`
}

type RecommendationsEnabledParameters struct {

	// Is Principal Authentication enabled for the ACR repository? Defaults to true.
	// +kubebuilder:validation:Optional
	AcrAuthentication *bool `json:"acrAuthentication,omitempty" tf:"acr_authentication,omitempty"`

	// Is Agent send underutilized messages enabled? Defaults to true.
	// +kubebuilder:validation:Optional
	AgentSendUnutilizedMsg *bool `json:"agentSendUnutilizedMsg,omitempty" tf:"agent_send_unutilized_msg,omitempty"`

	// Is Security related system configuration issues identified? Defaults to true.
	// +kubebuilder:validation:Optional
	Baseline *bool `json:"baseline,omitempty" tf:"baseline,omitempty"`

	// Is IoT Edge Hub memory optimized? Defaults to true.
	// +kubebuilder:validation:Optional
	EdgeHubMemOptimize *bool `json:"edgeHubMemOptimize,omitempty" tf:"edge_hub_mem_optimize,omitempty"`

	// Is logging configured for IoT Edge module? Defaults to true.
	// +kubebuilder:validation:Optional
	EdgeLoggingOption *bool `json:"edgeLoggingOption,omitempty" tf:"edge_logging_option,omitempty"`

	// Is Default IP filter policy denied? Defaults to true.
	// +kubebuilder:validation:Optional
	IPFilterDenyAll *bool `json:"ipFilterDenyAll,omitempty" tf:"ip_filter_deny_all,omitempty"`

	// Is IP filter rule source allowable IP range too large? Defaults to true.
	// +kubebuilder:validation:Optional
	IPFilterPermissiveRule *bool `json:"ipFilterPermissiveRule,omitempty" tf:"ip_filter_permissive_rule,omitempty"`

	// Is inconsistent module settings enabled for SecurityGroup? Defaults to true.
	// +kubebuilder:validation:Optional
	InconsistentModuleSettings *bool `json:"inconsistentModuleSettings,omitempty" tf:"inconsistent_module_settings,omitempty"`

	// is Azure IoT Security agent installed? Defaults to true.
	// +kubebuilder:validation:Optional
	InstallAgent *bool `json:"installAgent,omitempty" tf:"install_agent,omitempty"`

	// Is any ports open on the device? Defaults to true.
	// +kubebuilder:validation:Optional
	OpenPorts *bool `json:"openPorts,omitempty" tf:"open_ports,omitempty"`

	// Does firewall policy exist which allow necessary communication to/from the device? Defaults to true.
	// +kubebuilder:validation:Optional
	PermissiveFirewallPolicy *bool `json:"permissiveFirewallPolicy,omitempty" tf:"permissive_firewall_policy,omitempty"`

	// Is only necessary addresses or ports are permitted in? Defaults to true.
	// +kubebuilder:validation:Optional
	PermissiveInputFirewallRules *bool `json:"permissiveInputFirewallRules,omitempty" tf:"permissive_input_firewall_rules,omitempty"`

	// Is only necessary addresses or ports are permitted out? Defaults to true.
	// +kubebuilder:validation:Optional
	PermissiveOutputFirewallRules *bool `json:"permissiveOutputFirewallRules,omitempty" tf:"permissive_output_firewall_rules,omitempty"`

	// Is high level permissions are needed for the module? Defaults to true.
	// +kubebuilder:validation:Optional
	PrivilegedDockerOptions *bool `json:"privilegedDockerOptions,omitempty" tf:"privileged_docker_options,omitempty"`

	// Is any credentials shared among devices? Defaults to true.
	// +kubebuilder:validation:Optional
	SharedCredentials *bool `json:"sharedCredentials,omitempty" tf:"shared_credentials,omitempty"`

	// Does TLS cipher suite need to be updated? Defaults to true.
	// +kubebuilder:validation:Optional
	VulnerableTLSCipherSuite *bool `json:"vulnerableTlsCipherSuite,omitempty" tf:"vulnerable_tls_cipher_suite,omitempty"`
}

// IOTSecuritySolutionSpec defines the desired state of IOTSecuritySolution
type IOTSecuritySolutionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IOTSecuritySolutionParameters `json:"forProvider"`
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
	InitProvider IOTSecuritySolutionInitParameters `json:"initProvider,omitempty"`
}

// IOTSecuritySolutionStatus defines the observed state of IOTSecuritySolution.
type IOTSecuritySolutionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IOTSecuritySolutionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// IOTSecuritySolution is the Schema for the IOTSecuritySolutions API. Manages an iot security solution.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type IOTSecuritySolution struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.displayName) || (has(self.initProvider) && has(self.initProvider.displayName))",message="spec.forProvider.displayName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   IOTSecuritySolutionSpec   `json:"spec"`
	Status IOTSecuritySolutionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IOTSecuritySolutionList contains a list of IOTSecuritySolutions
type IOTSecuritySolutionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IOTSecuritySolution `json:"items"`
}

// Repository type metadata.
var (
	IOTSecuritySolution_Kind             = "IOTSecuritySolution"
	IOTSecuritySolution_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IOTSecuritySolution_Kind}.String()
	IOTSecuritySolution_KindAPIVersion   = IOTSecuritySolution_Kind + "." + CRDGroupVersion.String()
	IOTSecuritySolution_GroupVersionKind = CRDGroupVersion.WithKind(IOTSecuritySolution_Kind)
)

func init() {
	SchemeBuilder.Register(&IOTSecuritySolution{}, &IOTSecuritySolutionList{})
}
