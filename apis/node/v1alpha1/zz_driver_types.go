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

type DriverObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DriverParameters struct {

	// +kubebuilder:validation:Required
	Active *bool `json:"active" tf:"active,omitempty"`

	// Annotations of the resource
	// +kubebuilder:validation:Optional
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// +kubebuilder:validation:Required
	Builtin *bool `json:"builtin" tf:"builtin,omitempty"`

	// +kubebuilder:validation:Optional
	Checksum *string `json:"checksum,omitempty" tf:"checksum,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	ExternalID *string `json:"externalId,omitempty" tf:"external_id,omitempty"`

	// Labels of the resource
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// +kubebuilder:validation:Optional
	UIURL *string `json:"uiUrl,omitempty" tf:"ui_url,omitempty"`

	// +kubebuilder:validation:Required
	URL *string `json:"url" tf:"url,omitempty"`

	// +kubebuilder:validation:Optional
	WhitelistDomains []*string `json:"whitelistDomains,omitempty" tf:"whitelist_domains,omitempty"`
}

// DriverSpec defines the desired state of Driver
type DriverSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DriverParameters `json:"forProvider"`
}

// DriverStatus defines the observed state of Driver.
type DriverStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DriverObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Driver is the Schema for the Drivers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,rancherjet}
type Driver struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DriverSpec   `json:"spec"`
	Status            DriverStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DriverList contains a list of Drivers
type DriverList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Driver `json:"items"`
}

// Repository type metadata.
var (
	Driver_Kind             = "Driver"
	Driver_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Driver_Kind}.String()
	Driver_KindAPIVersion   = Driver_Kind + "." + CRDGroupVersion.String()
	Driver_GroupVersionKind = CRDGroupVersion.WithKind(Driver_Kind)
)

func init() {
	SchemeBuilder.Register(&Driver{}, &DriverList{})
}
