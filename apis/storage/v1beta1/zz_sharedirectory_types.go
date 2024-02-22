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

type ShareDirectoryInitParameters struct {

	// A mapping of metadata to assign to this Directory.
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The name (or path) of the Directory that should be created within this File Share. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the File Share where this Directory should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Share
	ShareName *string `json:"shareName,omitempty" tf:"share_name,omitempty"`

	// Reference to a Share in storage to populate shareName.
	// +kubebuilder:validation:Optional
	ShareNameRef *v1.Reference `json:"shareNameRef,omitempty" tf:"-"`

	// Selector for a Share in storage to populate shareName.
	// +kubebuilder:validation:Optional
	ShareNameSelector *v1.Selector `json:"shareNameSelector,omitempty" tf:"-"`

	// The name of the Storage Account within which the File Share is located. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Account
	StorageAccountName *string `json:"storageAccountName,omitempty" tf:"storage_account_name,omitempty"`

	// Reference to a Account in storage to populate storageAccountName.
	// +kubebuilder:validation:Optional
	StorageAccountNameRef *v1.Reference `json:"storageAccountNameRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate storageAccountName.
	// +kubebuilder:validation:Optional
	StorageAccountNameSelector *v1.Selector `json:"storageAccountNameSelector,omitempty" tf:"-"`
}

type ShareDirectoryObservation struct {

	// The ID of the Directory within the File Share.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A mapping of metadata to assign to this Directory.
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The name (or path) of the Directory that should be created within this File Share. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the File Share where this Directory should be created. Changing this forces a new resource to be created.
	ShareName *string `json:"shareName,omitempty" tf:"share_name,omitempty"`

	// The name of the Storage Account within which the File Share is located. Changing this forces a new resource to be created.
	StorageAccountName *string `json:"storageAccountName,omitempty" tf:"storage_account_name,omitempty"`
}

type ShareDirectoryParameters struct {

	// A mapping of metadata to assign to this Directory.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The name (or path) of the Directory that should be created within this File Share. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the File Share where this Directory should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Share
	// +kubebuilder:validation:Optional
	ShareName *string `json:"shareName,omitempty" tf:"share_name,omitempty"`

	// Reference to a Share in storage to populate shareName.
	// +kubebuilder:validation:Optional
	ShareNameRef *v1.Reference `json:"shareNameRef,omitempty" tf:"-"`

	// Selector for a Share in storage to populate shareName.
	// +kubebuilder:validation:Optional
	ShareNameSelector *v1.Selector `json:"shareNameSelector,omitempty" tf:"-"`

	// The name of the Storage Account within which the File Share is located. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Account
	// +kubebuilder:validation:Optional
	StorageAccountName *string `json:"storageAccountName,omitempty" tf:"storage_account_name,omitempty"`

	// Reference to a Account in storage to populate storageAccountName.
	// +kubebuilder:validation:Optional
	StorageAccountNameRef *v1.Reference `json:"storageAccountNameRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate storageAccountName.
	// +kubebuilder:validation:Optional
	StorageAccountNameSelector *v1.Selector `json:"storageAccountNameSelector,omitempty" tf:"-"`
}

// ShareDirectorySpec defines the desired state of ShareDirectory
type ShareDirectorySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ShareDirectoryParameters `json:"forProvider"`
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
	InitProvider ShareDirectoryInitParameters `json:"initProvider,omitempty"`
}

// ShareDirectoryStatus defines the observed state of ShareDirectory.
type ShareDirectoryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ShareDirectoryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ShareDirectory is the Schema for the ShareDirectorys API. Manages a Directory within an Azure Storage File Share.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ShareDirectory struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   ShareDirectorySpec   `json:"spec"`
	Status ShareDirectoryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ShareDirectoryList contains a list of ShareDirectorys
type ShareDirectoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ShareDirectory `json:"items"`
}

// Repository type metadata.
var (
	ShareDirectory_Kind             = "ShareDirectory"
	ShareDirectory_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ShareDirectory_Kind}.String()
	ShareDirectory_KindAPIVersion   = ShareDirectory_Kind + "." + CRDGroupVersion.String()
	ShareDirectory_GroupVersionKind = CRDGroupVersion.WithKind(ShareDirectory_Kind)
)

func init() {
	SchemeBuilder.Register(&ShareDirectory{}, &ShareDirectoryList{})
}
