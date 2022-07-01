/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/resource"
	"github.com/upbound/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this HPCCache
func (mg *HPCCache) GetTerraformResourceType() string {
	return "azurerm_hpc_cache"
}

// GetConnectionDetailsMapping for this HPCCache
func (tr *HPCCache) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"directory_active_directory[*].password": "spec.forProvider.directoryActiveDirectory[*].passwordSecretRef", "directory_ldap[*].bind[*].password": "spec.forProvider.directoryLdap[*].bind[*].passwordSecretRef"}
}

// GetObservation of this HPCCache
func (tr *HPCCache) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this HPCCache
func (tr *HPCCache) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this HPCCache
func (tr *HPCCache) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this HPCCache
func (tr *HPCCache) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this HPCCache
func (tr *HPCCache) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this HPCCache using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *HPCCache) LateInitialize(attrs []byte) (bool, error) {
	params := &HPCCacheParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *HPCCache) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this HPCCacheAccessPolicy
func (mg *HPCCacheAccessPolicy) GetTerraformResourceType() string {
	return "azurerm_hpc_cache_access_policy"
}

// GetConnectionDetailsMapping for this HPCCacheAccessPolicy
func (tr *HPCCacheAccessPolicy) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this HPCCacheAccessPolicy
func (tr *HPCCacheAccessPolicy) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this HPCCacheAccessPolicy
func (tr *HPCCacheAccessPolicy) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this HPCCacheAccessPolicy
func (tr *HPCCacheAccessPolicy) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this HPCCacheAccessPolicy
func (tr *HPCCacheAccessPolicy) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this HPCCacheAccessPolicy
func (tr *HPCCacheAccessPolicy) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this HPCCacheAccessPolicy using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *HPCCacheAccessPolicy) LateInitialize(attrs []byte) (bool, error) {
	params := &HPCCacheAccessPolicyParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *HPCCacheAccessPolicy) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this HPCCacheBlobNFSTarget
func (mg *HPCCacheBlobNFSTarget) GetTerraformResourceType() string {
	return "azurerm_hpc_cache_blob_nfs_target"
}

// GetConnectionDetailsMapping for this HPCCacheBlobNFSTarget
func (tr *HPCCacheBlobNFSTarget) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this HPCCacheBlobNFSTarget
func (tr *HPCCacheBlobNFSTarget) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this HPCCacheBlobNFSTarget
func (tr *HPCCacheBlobNFSTarget) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this HPCCacheBlobNFSTarget
func (tr *HPCCacheBlobNFSTarget) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this HPCCacheBlobNFSTarget
func (tr *HPCCacheBlobNFSTarget) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this HPCCacheBlobNFSTarget
func (tr *HPCCacheBlobNFSTarget) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this HPCCacheBlobNFSTarget using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *HPCCacheBlobNFSTarget) LateInitialize(attrs []byte) (bool, error) {
	params := &HPCCacheBlobNFSTargetParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *HPCCacheBlobNFSTarget) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this HPCCacheBlobTarget
func (mg *HPCCacheBlobTarget) GetTerraformResourceType() string {
	return "azurerm_hpc_cache_blob_target"
}

// GetConnectionDetailsMapping for this HPCCacheBlobTarget
func (tr *HPCCacheBlobTarget) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this HPCCacheBlobTarget
func (tr *HPCCacheBlobTarget) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this HPCCacheBlobTarget
func (tr *HPCCacheBlobTarget) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this HPCCacheBlobTarget
func (tr *HPCCacheBlobTarget) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this HPCCacheBlobTarget
func (tr *HPCCacheBlobTarget) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this HPCCacheBlobTarget
func (tr *HPCCacheBlobTarget) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this HPCCacheBlobTarget using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *HPCCacheBlobTarget) LateInitialize(attrs []byte) (bool, error) {
	params := &HPCCacheBlobTargetParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *HPCCacheBlobTarget) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this HPCCacheNFSTarget
func (mg *HPCCacheNFSTarget) GetTerraformResourceType() string {
	return "azurerm_hpc_cache_nfs_target"
}

// GetConnectionDetailsMapping for this HPCCacheNFSTarget
func (tr *HPCCacheNFSTarget) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this HPCCacheNFSTarget
func (tr *HPCCacheNFSTarget) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this HPCCacheNFSTarget
func (tr *HPCCacheNFSTarget) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this HPCCacheNFSTarget
func (tr *HPCCacheNFSTarget) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this HPCCacheNFSTarget
func (tr *HPCCacheNFSTarget) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this HPCCacheNFSTarget
func (tr *HPCCacheNFSTarget) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this HPCCacheNFSTarget using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *HPCCacheNFSTarget) LateInitialize(attrs []byte) (bool, error) {
	params := &HPCCacheNFSTargetParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *HPCCacheNFSTarget) GetTerraformSchemaVersion() int {
	return 0
}
