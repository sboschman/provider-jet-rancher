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

type RegistriesObservation struct {
}

type RegistriesParameters struct {

	// +kubebuilder:validation:Required
	Address *string `json:"address" tf:"address,omitempty"`

	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type RegistryObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RegistryParameters struct {

	// Annotations of the resource
	// +kubebuilder:validation:Optional
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// Description of the docker registry
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Labels of the resource
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Namespace ID to add docker registry
	// +kubebuilder:validation:Optional
	NamespaceID *string `json:"namespaceId,omitempty" tf:"namespace_id,omitempty"`

	// Project ID to add docker registry
	// +kubebuilder:validation:Required
	ProjectID *string `json:"projectId" tf:"project_id,omitempty"`

	// +kubebuilder:validation:Required
	Registries []RegistriesParameters `json:"registries" tf:"registries,omitempty"`
}

// RegistrySpec defines the desired state of Registry
type RegistrySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RegistryParameters `json:"forProvider"`
}

// RegistryStatus defines the observed state of Registry.
type RegistryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RegistryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Registry is the Schema for the Registrys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,rancherjet}
type Registry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RegistrySpec   `json:"spec"`
	Status            RegistryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RegistryList contains a list of Registrys
type RegistryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Registry `json:"items"`
}

// Repository type metadata.
var (
	Registry_Kind             = "Registry"
	Registry_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Registry_Kind}.String()
	Registry_KindAPIVersion   = Registry_Kind + "." + CRDGroupVersion.String()
	Registry_GroupVersionKind = CRDGroupVersion.WithKind(Registry_Kind)
)

func init() {
	SchemeBuilder.Register(&Registry{}, &RegistryList{})
}
