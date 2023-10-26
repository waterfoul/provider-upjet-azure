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

type ConfigServerGitSettingInitParameters struct {

	// A http_basic_auth block as defined below.
	HTTPBasicAuth []HTTPBasicAuthInitParameters `json:"httpBasicAuth,omitempty" tf:"http_basic_auth,omitempty"`

	// The default label of the Git repository, should be the branch name, tag name, or commit-id of the repository.
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// One or more repository blocks as defined below.
	Repository []ConfigServerGitSettingRepositoryInitParameters `json:"repository,omitempty" tf:"repository,omitempty"`

	// A ssh_auth block as defined below.
	SSHAuth []ConfigServerGitSettingSSHAuthInitParameters `json:"sshAuth,omitempty" tf:"ssh_auth,omitempty"`

	// An array of strings used to search subdirectories of the Git repository.
	SearchPaths []*string `json:"searchPaths,omitempty" tf:"search_paths,omitempty"`

	// The URI of the default Git repository used as the Config Server back end, should be started with http://, https://, git@, or ssh://.
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type ConfigServerGitSettingObservation struct {

	// A http_basic_auth block as defined below.
	HTTPBasicAuth []HTTPBasicAuthObservation `json:"httpBasicAuth,omitempty" tf:"http_basic_auth,omitempty"`

	// The default label of the Git repository, should be the branch name, tag name, or commit-id of the repository.
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// One or more repository blocks as defined below.
	Repository []ConfigServerGitSettingRepositoryObservation `json:"repository,omitempty" tf:"repository,omitempty"`

	// A ssh_auth block as defined below.
	SSHAuth []ConfigServerGitSettingSSHAuthObservation `json:"sshAuth,omitempty" tf:"ssh_auth,omitempty"`

	// An array of strings used to search subdirectories of the Git repository.
	SearchPaths []*string `json:"searchPaths,omitempty" tf:"search_paths,omitempty"`

	// The URI of the default Git repository used as the Config Server back end, should be started with http://, https://, git@, or ssh://.
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type ConfigServerGitSettingParameters struct {

	// A http_basic_auth block as defined below.
	// +kubebuilder:validation:Optional
	HTTPBasicAuth []HTTPBasicAuthParameters `json:"httpBasicAuth,omitempty" tf:"http_basic_auth,omitempty"`

	// The default label of the Git repository, should be the branch name, tag name, or commit-id of the repository.
	// +kubebuilder:validation:Optional
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// One or more repository blocks as defined below.
	// +kubebuilder:validation:Optional
	Repository []ConfigServerGitSettingRepositoryParameters `json:"repository,omitempty" tf:"repository,omitempty"`

	// A ssh_auth block as defined below.
	// +kubebuilder:validation:Optional
	SSHAuth []ConfigServerGitSettingSSHAuthParameters `json:"sshAuth,omitempty" tf:"ssh_auth,omitempty"`

	// An array of strings used to search subdirectories of the Git repository.
	// +kubebuilder:validation:Optional
	SearchPaths []*string `json:"searchPaths,omitempty" tf:"search_paths,omitempty"`

	// The URI of the default Git repository used as the Config Server back end, should be started with http://, https://, git@, or ssh://.
	// +kubebuilder:validation:Optional
	URI *string `json:"uri" tf:"uri,omitempty"`
}

type ConfigServerGitSettingRepositoryInitParameters struct {

	// A http_basic_auth block as defined below.
	HTTPBasicAuth []RepositoryHTTPBasicAuthInitParameters `json:"httpBasicAuth,omitempty" tf:"http_basic_auth,omitempty"`

	// The default label of the Git repository, should be the branch name, tag name, or commit-id of the repository.
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// A name to identify on the Git repository, required only if repos exists.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// An array of strings used to match an application name. For each pattern, use the {application}/{profile} format with wildcards.
	Pattern []*string `json:"pattern,omitempty" tf:"pattern,omitempty"`

	// A ssh_auth block as defined below.
	SSHAuth []RepositorySSHAuthInitParameters `json:"sshAuth,omitempty" tf:"ssh_auth,omitempty"`

	// An array of strings used to search subdirectories of the Git repository.
	SearchPaths []*string `json:"searchPaths,omitempty" tf:"search_paths,omitempty"`

	// The URI of the Git repository that's used as the Config Server back end should be started with http://, https://, git@, or ssh://.
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type ConfigServerGitSettingRepositoryObservation struct {

	// A http_basic_auth block as defined below.
	HTTPBasicAuth []RepositoryHTTPBasicAuthObservation `json:"httpBasicAuth,omitempty" tf:"http_basic_auth,omitempty"`

	// The default label of the Git repository, should be the branch name, tag name, or commit-id of the repository.
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// A name to identify on the Git repository, required only if repos exists.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// An array of strings used to match an application name. For each pattern, use the {application}/{profile} format with wildcards.
	Pattern []*string `json:"pattern,omitempty" tf:"pattern,omitempty"`

	// A ssh_auth block as defined below.
	SSHAuth []RepositorySSHAuthObservation `json:"sshAuth,omitempty" tf:"ssh_auth,omitempty"`

	// An array of strings used to search subdirectories of the Git repository.
	SearchPaths []*string `json:"searchPaths,omitempty" tf:"search_paths,omitempty"`

	// The URI of the Git repository that's used as the Config Server back end should be started with http://, https://, git@, or ssh://.
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type ConfigServerGitSettingRepositoryParameters struct {

	// A http_basic_auth block as defined below.
	// +kubebuilder:validation:Optional
	HTTPBasicAuth []RepositoryHTTPBasicAuthParameters `json:"httpBasicAuth,omitempty" tf:"http_basic_auth,omitempty"`

	// The default label of the Git repository, should be the branch name, tag name, or commit-id of the repository.
	// +kubebuilder:validation:Optional
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// A name to identify on the Git repository, required only if repos exists.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// An array of strings used to match an application name. For each pattern, use the {application}/{profile} format with wildcards.
	// +kubebuilder:validation:Optional
	Pattern []*string `json:"pattern,omitempty" tf:"pattern,omitempty"`

	// A ssh_auth block as defined below.
	// +kubebuilder:validation:Optional
	SSHAuth []RepositorySSHAuthParameters `json:"sshAuth,omitempty" tf:"ssh_auth,omitempty"`

	// An array of strings used to search subdirectories of the Git repository.
	// +kubebuilder:validation:Optional
	SearchPaths []*string `json:"searchPaths,omitempty" tf:"search_paths,omitempty"`

	// The URI of the Git repository that's used as the Config Server back end should be started with http://, https://, git@, or ssh://.
	// +kubebuilder:validation:Optional
	URI *string `json:"uri" tf:"uri,omitempty"`
}

type ConfigServerGitSettingSSHAuthInitParameters struct {

	// The host key algorithm, should be ssh-dss, ssh-rsa, ecdsa-sha2-nistp256, ecdsa-sha2-nistp384, or ecdsa-sha2-nistp521. Required only if host-key exists.
	HostKeyAlgorithm *string `json:"hostKeyAlgorithm,omitempty" tf:"host_key_algorithm,omitempty"`

	// Indicates whether the Config Server instance will fail to start if the host_key does not match. Defaults to true.
	StrictHostKeyCheckingEnabled *bool `json:"strictHostKeyCheckingEnabled,omitempty" tf:"strict_host_key_checking_enabled,omitempty"`
}

type ConfigServerGitSettingSSHAuthObservation struct {

	// The host key algorithm, should be ssh-dss, ssh-rsa, ecdsa-sha2-nistp256, ecdsa-sha2-nistp384, or ecdsa-sha2-nistp521. Required only if host-key exists.
	HostKeyAlgorithm *string `json:"hostKeyAlgorithm,omitempty" tf:"host_key_algorithm,omitempty"`

	// Indicates whether the Config Server instance will fail to start if the host_key does not match. Defaults to true.
	StrictHostKeyCheckingEnabled *bool `json:"strictHostKeyCheckingEnabled,omitempty" tf:"strict_host_key_checking_enabled,omitempty"`
}

type ConfigServerGitSettingSSHAuthParameters struct {

	// The host key algorithm, should be ssh-dss, ssh-rsa, ecdsa-sha2-nistp256, ecdsa-sha2-nistp384, or ecdsa-sha2-nistp521. Required only if host-key exists.
	// +kubebuilder:validation:Optional
	HostKeyAlgorithm *string `json:"hostKeyAlgorithm,omitempty" tf:"host_key_algorithm,omitempty"`

	// The host key of the Git repository server, should not include the algorithm prefix as covered by host-key-algorithm.
	// +kubebuilder:validation:Optional
	HostKeySecretRef *v1.SecretKeySelector `json:"hostKeySecretRef,omitempty" tf:"-"`

	// The SSH private key to access the Git repository, required when the URI starts with git@ or ssh://.
	// +kubebuilder:validation:Required
	PrivateKeySecretRef v1.SecretKeySelector `json:"privateKeySecretRef" tf:"-"`

	// Indicates whether the Config Server instance will fail to start if the host_key does not match. Defaults to true.
	// +kubebuilder:validation:Optional
	StrictHostKeyCheckingEnabled *bool `json:"strictHostKeyCheckingEnabled,omitempty" tf:"strict_host_key_checking_enabled,omitempty"`
}

type HTTPBasicAuthInitParameters struct {

	// The username that's used to access the Git repository server, required when the Git repository server supports HTTP Basic Authentication.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type HTTPBasicAuthObservation struct {

	// The username that's used to access the Git repository server, required when the Git repository server supports HTTP Basic Authentication.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type HTTPBasicAuthParameters struct {

	// The password used to access the Git repository server, required when the Git repository server supports HTTP Basic Authentication.
	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// The username that's used to access the Git repository server, required when the Git repository server supports HTTP Basic Authentication.
	// +kubebuilder:validation:Optional
	Username *string `json:"username" tf:"username,omitempty"`
}

type NetworkInitParameters struct {

	// Specifies the Name of the resource group containing network resources of Azure Spring Cloud Apps. Changing this forces a new resource to be created.
	AppNetworkResourceGroup *string `json:"appNetworkResourceGroup,omitempty" tf:"app_network_resource_group,omitempty"`

	// A list of (at least 3) CIDR ranges (at least /16) which are used to host the Spring Cloud infrastructure, which must not overlap with any existing CIDR ranges in the Subnet. Changing this forces a new resource to be created.
	CidrRanges []*string `json:"cidrRanges,omitempty" tf:"cidr_ranges,omitempty"`

	// Ingress read time out in seconds.
	ReadTimeoutSeconds *float64 `json:"readTimeoutSeconds,omitempty" tf:"read_timeout_seconds,omitempty"`

	// Specifies the Name of the resource group containing network resources of Azure Spring Cloud Service Runtime. Changing this forces a new resource to be created.
	ServiceRuntimeNetworkResourceGroup *string `json:"serviceRuntimeNetworkResourceGroup,omitempty" tf:"service_runtime_network_resource_group,omitempty"`
}

type NetworkObservation struct {

	// Specifies the Name of the resource group containing network resources of Azure Spring Cloud Apps. Changing this forces a new resource to be created.
	AppNetworkResourceGroup *string `json:"appNetworkResourceGroup,omitempty" tf:"app_network_resource_group,omitempty"`

	// Specifies the ID of the Subnet which should host the Spring Boot Applications deployed in this Spring Cloud Service. Changing this forces a new resource to be created.
	AppSubnetID *string `json:"appSubnetId,omitempty" tf:"app_subnet_id,omitempty"`

	// A list of (at least 3) CIDR ranges (at least /16) which are used to host the Spring Cloud infrastructure, which must not overlap with any existing CIDR ranges in the Subnet. Changing this forces a new resource to be created.
	CidrRanges []*string `json:"cidrRanges,omitempty" tf:"cidr_ranges,omitempty"`

	// Ingress read time out in seconds.
	ReadTimeoutSeconds *float64 `json:"readTimeoutSeconds,omitempty" tf:"read_timeout_seconds,omitempty"`

	// Specifies the Name of the resource group containing network resources of Azure Spring Cloud Service Runtime. Changing this forces a new resource to be created.
	ServiceRuntimeNetworkResourceGroup *string `json:"serviceRuntimeNetworkResourceGroup,omitempty" tf:"service_runtime_network_resource_group,omitempty"`

	// Specifies the ID of the Subnet where the Service Runtime components of the Spring Cloud Service will exist. Changing this forces a new resource to be created.
	ServiceRuntimeSubnetID *string `json:"serviceRuntimeSubnetId,omitempty" tf:"service_runtime_subnet_id,omitempty"`
}

type NetworkParameters struct {

	// Specifies the Name of the resource group containing network resources of Azure Spring Cloud Apps. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	AppNetworkResourceGroup *string `json:"appNetworkResourceGroup,omitempty" tf:"app_network_resource_group,omitempty"`

	// Specifies the ID of the Subnet which should host the Spring Boot Applications deployed in this Spring Cloud Service. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	AppSubnetID *string `json:"appSubnetId,omitempty" tf:"app_subnet_id,omitempty"`

	// Reference to a Subnet in network to populate appSubnetId.
	// +kubebuilder:validation:Optional
	AppSubnetIDRef *v1.Reference `json:"appSubnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate appSubnetId.
	// +kubebuilder:validation:Optional
	AppSubnetIDSelector *v1.Selector `json:"appSubnetIdSelector,omitempty" tf:"-"`

	// A list of (at least 3) CIDR ranges (at least /16) which are used to host the Spring Cloud infrastructure, which must not overlap with any existing CIDR ranges in the Subnet. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	CidrRanges []*string `json:"cidrRanges" tf:"cidr_ranges,omitempty"`

	// Ingress read time out in seconds.
	// +kubebuilder:validation:Optional
	ReadTimeoutSeconds *float64 `json:"readTimeoutSeconds,omitempty" tf:"read_timeout_seconds,omitempty"`

	// Specifies the Name of the resource group containing network resources of Azure Spring Cloud Service Runtime. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	ServiceRuntimeNetworkResourceGroup *string `json:"serviceRuntimeNetworkResourceGroup,omitempty" tf:"service_runtime_network_resource_group,omitempty"`

	// Specifies the ID of the Subnet where the Service Runtime components of the Spring Cloud Service will exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ServiceRuntimeSubnetID *string `json:"serviceRuntimeSubnetId,omitempty" tf:"service_runtime_subnet_id,omitempty"`

	// Reference to a Subnet in network to populate serviceRuntimeSubnetId.
	// +kubebuilder:validation:Optional
	ServiceRuntimeSubnetIDRef *v1.Reference `json:"serviceRuntimeSubnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate serviceRuntimeSubnetId.
	// +kubebuilder:validation:Optional
	ServiceRuntimeSubnetIDSelector *v1.Selector `json:"serviceRuntimeSubnetIdSelector,omitempty" tf:"-"`
}

type RepositoryHTTPBasicAuthInitParameters struct {

	// The username that's used to access the Git repository server, required when the Git repository server supports HTTP Basic Authentication.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type RepositoryHTTPBasicAuthObservation struct {

	// The username that's used to access the Git repository server, required when the Git repository server supports HTTP Basic Authentication.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type RepositoryHTTPBasicAuthParameters struct {

	// The password used to access the Git repository server, required when the Git repository server supports HTTP Basic Authentication.
	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// The username that's used to access the Git repository server, required when the Git repository server supports HTTP Basic Authentication.
	// +kubebuilder:validation:Optional
	Username *string `json:"username" tf:"username,omitempty"`
}

type RepositorySSHAuthInitParameters struct {

	// The host key algorithm, should be ssh-dss, ssh-rsa, ecdsa-sha2-nistp256, ecdsa-sha2-nistp384, or ecdsa-sha2-nistp521. Required only if host-key exists.
	HostKeyAlgorithm *string `json:"hostKeyAlgorithm,omitempty" tf:"host_key_algorithm,omitempty"`

	// Indicates whether the Config Server instance will fail to start if the host_key does not match. Defaults to true.
	StrictHostKeyCheckingEnabled *bool `json:"strictHostKeyCheckingEnabled,omitempty" tf:"strict_host_key_checking_enabled,omitempty"`
}

type RepositorySSHAuthObservation struct {

	// The host key algorithm, should be ssh-dss, ssh-rsa, ecdsa-sha2-nistp256, ecdsa-sha2-nistp384, or ecdsa-sha2-nistp521. Required only if host-key exists.
	HostKeyAlgorithm *string `json:"hostKeyAlgorithm,omitempty" tf:"host_key_algorithm,omitempty"`

	// Indicates whether the Config Server instance will fail to start if the host_key does not match. Defaults to true.
	StrictHostKeyCheckingEnabled *bool `json:"strictHostKeyCheckingEnabled,omitempty" tf:"strict_host_key_checking_enabled,omitempty"`
}

type RepositorySSHAuthParameters struct {

	// The host key algorithm, should be ssh-dss, ssh-rsa, ecdsa-sha2-nistp256, ecdsa-sha2-nistp384, or ecdsa-sha2-nistp521. Required only if host-key exists.
	// +kubebuilder:validation:Optional
	HostKeyAlgorithm *string `json:"hostKeyAlgorithm,omitempty" tf:"host_key_algorithm,omitempty"`

	// The host key of the Git repository server, should not include the algorithm prefix as covered by host-key-algorithm.
	// +kubebuilder:validation:Optional
	HostKeySecretRef *v1.SecretKeySelector `json:"hostKeySecretRef,omitempty" tf:"-"`

	// The SSH private key to access the Git repository, required when the URI starts with git@ or ssh://.
	// +kubebuilder:validation:Required
	PrivateKeySecretRef v1.SecretKeySelector `json:"privateKeySecretRef" tf:"-"`

	// Indicates whether the Config Server instance will fail to start if the host_key does not match. Defaults to true.
	// +kubebuilder:validation:Optional
	StrictHostKeyCheckingEnabled *bool `json:"strictHostKeyCheckingEnabled,omitempty" tf:"strict_host_key_checking_enabled,omitempty"`
}

type RequiredNetworkTrafficRulesInitParameters struct {
}

type RequiredNetworkTrafficRulesObservation struct {

	// The direction of required traffic. Possible values are Inbound, Outbound.
	Direction *string `json:"direction,omitempty" tf:"direction,omitempty"`

	// The FQDN list of required traffic.
	Fqdns []*string `json:"fqdns,omitempty" tf:"fqdns,omitempty"`

	// The IP list of required traffic.
	IPAddresses []*string `json:"ipAddresses,omitempty" tf:"ip_addresses,omitempty"`

	// The port of required traffic.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The protocol of required traffic.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type RequiredNetworkTrafficRulesParameters struct {
}

type SpringCloudServiceInitParameters struct {

	// Specifies the size for this Spring Cloud Service's default build agent pool. Possible values are S1, S2, S3, S4 and S5. This field is applicable only for Spring Cloud Service with enterprise tier.
	BuildAgentPoolSize *string `json:"buildAgentPoolSize,omitempty" tf:"build_agent_pool_size,omitempty"`

	// A config_server_git_setting block as defined below. This field is applicable only for Spring Cloud Service with basic and standard tier.
	ConfigServerGitSetting []ConfigServerGitSettingInitParameters `json:"configServerGitSetting,omitempty" tf:"config_server_git_setting,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Should the log stream in vnet injection instance could be accessed from Internet?
	LogStreamPublicEndpointEnabled *bool `json:"logStreamPublicEndpointEnabled,omitempty" tf:"log_stream_public_endpoint_enabled,omitempty"`

	// A network block as defined below. Changing this forces a new resource to be created.
	Network []NetworkInitParameters `json:"network,omitempty" tf:"network,omitempty"`

	// Whether enable the default Service Registry. This field is applicable only for Spring Cloud Service with enterprise tier.
	ServiceRegistryEnabled *bool `json:"serviceRegistryEnabled,omitempty" tf:"service_registry_enabled,omitempty"`

	// Specifies the SKU Name for this Spring Cloud Service. Possible values are B0, S0 and E0. Defaults to S0. Changing this forces a new resource to be created.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A trace block as defined below.
	Trace []TraceInitParameters `json:"trace,omitempty" tf:"trace,omitempty"`

	// Whether zone redundancy is enabled for this Spring Cloud Service. Defaults to false.
	ZoneRedundant *bool `json:"zoneRedundant,omitempty" tf:"zone_redundant,omitempty"`
}

type SpringCloudServiceObservation struct {

	// Specifies the size for this Spring Cloud Service's default build agent pool. Possible values are S1, S2, S3, S4 and S5. This field is applicable only for Spring Cloud Service with enterprise tier.
	BuildAgentPoolSize *string `json:"buildAgentPoolSize,omitempty" tf:"build_agent_pool_size,omitempty"`

	// A config_server_git_setting block as defined below. This field is applicable only for Spring Cloud Service with basic and standard tier.
	ConfigServerGitSetting []ConfigServerGitSettingObservation `json:"configServerGitSetting,omitempty" tf:"config_server_git_setting,omitempty"`

	// The ID of the Spring Cloud Service.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Should the log stream in vnet injection instance could be accessed from Internet?
	LogStreamPublicEndpointEnabled *bool `json:"logStreamPublicEndpointEnabled,omitempty" tf:"log_stream_public_endpoint_enabled,omitempty"`

	// A network block as defined below. Changing this forces a new resource to be created.
	Network []NetworkObservation `json:"network,omitempty" tf:"network,omitempty"`

	// A list of the outbound Public IP Addresses used by this Spring Cloud Service.
	OutboundPublicIPAddresses []*string `json:"outboundPublicIpAddresses,omitempty" tf:"outbound_public_ip_addresses,omitempty"`

	// A list of required_network_traffic_rules blocks as defined below.
	RequiredNetworkTrafficRules []RequiredNetworkTrafficRulesObservation `json:"requiredNetworkTrafficRules,omitempty" tf:"required_network_traffic_rules,omitempty"`

	// Specifies The name of the resource group in which to create the Spring Cloud Service. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Whether enable the default Service Registry. This field is applicable only for Spring Cloud Service with enterprise tier.
	ServiceRegistryEnabled *bool `json:"serviceRegistryEnabled,omitempty" tf:"service_registry_enabled,omitempty"`

	// The ID of the Spring Cloud Service Registry.
	ServiceRegistryID *string `json:"serviceRegistryId,omitempty" tf:"service_registry_id,omitempty"`

	// Specifies the SKU Name for this Spring Cloud Service. Possible values are B0, S0 and E0. Defaults to S0. Changing this forces a new resource to be created.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A trace block as defined below.
	Trace []TraceObservation `json:"trace,omitempty" tf:"trace,omitempty"`

	// Whether zone redundancy is enabled for this Spring Cloud Service. Defaults to false.
	ZoneRedundant *bool `json:"zoneRedundant,omitempty" tf:"zone_redundant,omitempty"`
}

type SpringCloudServiceParameters struct {

	// Specifies the size for this Spring Cloud Service's default build agent pool. Possible values are S1, S2, S3, S4 and S5. This field is applicable only for Spring Cloud Service with enterprise tier.
	// +kubebuilder:validation:Optional
	BuildAgentPoolSize *string `json:"buildAgentPoolSize,omitempty" tf:"build_agent_pool_size,omitempty"`

	// A config_server_git_setting block as defined below. This field is applicable only for Spring Cloud Service with basic and standard tier.
	// +kubebuilder:validation:Optional
	ConfigServerGitSetting []ConfigServerGitSettingParameters `json:"configServerGitSetting,omitempty" tf:"config_server_git_setting,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Should the log stream in vnet injection instance could be accessed from Internet?
	// +kubebuilder:validation:Optional
	LogStreamPublicEndpointEnabled *bool `json:"logStreamPublicEndpointEnabled,omitempty" tf:"log_stream_public_endpoint_enabled,omitempty"`

	// A network block as defined below. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Network []NetworkParameters `json:"network,omitempty" tf:"network,omitempty"`

	// Specifies The name of the resource group in which to create the Spring Cloud Service. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Whether enable the default Service Registry. This field is applicable only for Spring Cloud Service with enterprise tier.
	// +kubebuilder:validation:Optional
	ServiceRegistryEnabled *bool `json:"serviceRegistryEnabled,omitempty" tf:"service_registry_enabled,omitempty"`

	// Specifies the SKU Name for this Spring Cloud Service. Possible values are B0, S0 and E0. Defaults to S0. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A trace block as defined below.
	// +kubebuilder:validation:Optional
	Trace []TraceParameters `json:"trace,omitempty" tf:"trace,omitempty"`

	// Whether zone redundancy is enabled for this Spring Cloud Service. Defaults to false.
	// +kubebuilder:validation:Optional
	ZoneRedundant *bool `json:"zoneRedundant,omitempty" tf:"zone_redundant,omitempty"`
}

type TraceInitParameters struct {

	// The sampling rate of Application Insights Agent. Must be between 0.0 and 100.0. Defaults to 10.0.
	SampleRate *float64 `json:"sampleRate,omitempty" tf:"sample_rate,omitempty"`
}

type TraceObservation struct {

	// The connection string used for Application Insights.
	ConnectionString *string `json:"connectionString,omitempty" tf:"connection_string,omitempty"`

	// The sampling rate of Application Insights Agent. Must be between 0.0 and 100.0. Defaults to 10.0.
	SampleRate *float64 `json:"sampleRate,omitempty" tf:"sample_rate,omitempty"`
}

type TraceParameters struct {

	// The connection string used for Application Insights.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/insights/v1beta1.ApplicationInsights
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("connection_string",true)
	// +kubebuilder:validation:Optional
	ConnectionString *string `json:"connectionString,omitempty" tf:"connection_string,omitempty"`

	// Reference to a ApplicationInsights in insights to populate connectionString.
	// +kubebuilder:validation:Optional
	ConnectionStringRef *v1.Reference `json:"connectionStringRef,omitempty" tf:"-"`

	// Selector for a ApplicationInsights in insights to populate connectionString.
	// +kubebuilder:validation:Optional
	ConnectionStringSelector *v1.Selector `json:"connectionStringSelector,omitempty" tf:"-"`

	// The sampling rate of Application Insights Agent. Must be between 0.0 and 100.0. Defaults to 10.0.
	// +kubebuilder:validation:Optional
	SampleRate *float64 `json:"sampleRate,omitempty" tf:"sample_rate,omitempty"`
}

// SpringCloudServiceSpec defines the desired state of SpringCloudService
type SpringCloudServiceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SpringCloudServiceParameters `json:"forProvider"`
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
	InitProvider SpringCloudServiceInitParameters `json:"initProvider,omitempty"`
}

// SpringCloudServiceStatus defines the observed state of SpringCloudService.
type SpringCloudServiceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SpringCloudServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SpringCloudService is the Schema for the SpringCloudServices API. Manages an Azure Spring Cloud Service.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SpringCloudService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	Spec   SpringCloudServiceSpec   `json:"spec"`
	Status SpringCloudServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpringCloudServiceList contains a list of SpringCloudServices
type SpringCloudServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpringCloudService `json:"items"`
}

// Repository type metadata.
var (
	SpringCloudService_Kind             = "SpringCloudService"
	SpringCloudService_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SpringCloudService_Kind}.String()
	SpringCloudService_KindAPIVersion   = SpringCloudService_Kind + "." + CRDGroupVersion.String()
	SpringCloudService_GroupVersionKind = CRDGroupVersion.WithKind(SpringCloudService_Kind)
)

func init() {
	SchemeBuilder.Register(&SpringCloudService{}, &SpringCloudServiceList{})
}
