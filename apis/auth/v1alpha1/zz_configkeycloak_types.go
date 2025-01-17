/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ConfigKeycloakObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ConfigKeycloakParameters struct {

	// +kubebuilder:validation:Optional
	AccessMode *string `json:"accessMode,omitempty" tf:"access_mode,omitempty"`

	// +kubebuilder:validation:Optional
	AllowedPrincipalIds []*string `json:"allowedPrincipalIds,omitempty" tf:"allowed_principal_ids,omitempty"`

	// Annotations of the resource
	// +kubebuilder:validation:Optional
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// +kubebuilder:validation:Required
	DisplayNameField *string `json:"displayNameField" tf:"display_name_field,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	EntityID *string `json:"entityId,omitempty" tf:"entity_id,omitempty"`

	// +kubebuilder:validation:Required
	GroupsField *string `json:"groupsField" tf:"groups_field,omitempty"`

	// +kubebuilder:validation:Required
	IdpMetadataContentSecretRef v1.SecretKeySelector `json:"idpMetadataContentSecretRef" tf:"-"`

	// Labels of the resource
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// +kubebuilder:validation:Required
	RancherAPIHost *string `json:"rancherApiHost" tf:"rancher_api_host,omitempty"`

	// +kubebuilder:validation:Required
	SpCertSecretRef v1.SecretKeySelector `json:"spCertSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	SpKeySecretRef v1.SecretKeySelector `json:"spKeySecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	UIDField *string `json:"uidField" tf:"uid_field,omitempty"`

	// +kubebuilder:validation:Required
	UserNameField *string `json:"userNameField" tf:"user_name_field,omitempty"`
}

// ConfigKeycloakSpec defines the desired state of ConfigKeycloak
type ConfigKeycloakSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConfigKeycloakParameters `json:"forProvider"`
}

// ConfigKeycloakStatus defines the observed state of ConfigKeycloak.
type ConfigKeycloakStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConfigKeycloakObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigKeycloak is the Schema for the ConfigKeycloaks API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,rancherjet}
type ConfigKeycloak struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConfigKeycloakSpec   `json:"spec"`
	Status            ConfigKeycloakStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigKeycloakList contains a list of ConfigKeycloaks
type ConfigKeycloakList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigKeycloak `json:"items"`
}

// Repository type metadata.
var (
	ConfigKeycloak_Kind             = "ConfigKeycloak"
	ConfigKeycloak_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ConfigKeycloak_Kind}.String()
	ConfigKeycloak_KindAPIVersion   = ConfigKeycloak_Kind + "." + CRDGroupVersion.String()
	ConfigKeycloak_GroupVersionKind = CRDGroupVersion.WithKind(ConfigKeycloak_Kind)
)

func init() {
	SchemeBuilder.Register(&ConfigKeycloak{}, &ConfigKeycloakList{})
}
