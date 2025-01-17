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

type RoleTemplateBindingObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RoleTemplateBindingParameters struct {

	// Annotations of the resource
	// +kubebuilder:validation:Optional
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// +kubebuilder:validation:Required
	ClusterID *string `json:"clusterId" tf:"cluster_id,omitempty"`

	// +kubebuilder:validation:Optional
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// +kubebuilder:validation:Optional
	GroupPrincipalID *string `json:"groupPrincipalId,omitempty" tf:"group_principal_id,omitempty"`

	// Labels of the resource
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// +kubebuilder:validation:Required
	RoleTemplateID *string `json:"roleTemplateId" tf:"role_template_id,omitempty"`

	// +kubebuilder:validation:Optional
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`

	// +kubebuilder:validation:Optional
	UserPrincipalID *string `json:"userPrincipalId,omitempty" tf:"user_principal_id,omitempty"`
}

// RoleTemplateBindingSpec defines the desired state of RoleTemplateBinding
type RoleTemplateBindingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RoleTemplateBindingParameters `json:"forProvider"`
}

// RoleTemplateBindingStatus defines the observed state of RoleTemplateBinding.
type RoleTemplateBindingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RoleTemplateBindingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RoleTemplateBinding is the Schema for the RoleTemplateBindings API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,rancherjet}
type RoleTemplateBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RoleTemplateBindingSpec   `json:"spec"`
	Status            RoleTemplateBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RoleTemplateBindingList contains a list of RoleTemplateBindings
type RoleTemplateBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RoleTemplateBinding `json:"items"`
}

// Repository type metadata.
var (
	RoleTemplateBinding_Kind             = "RoleTemplateBinding"
	RoleTemplateBinding_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RoleTemplateBinding_Kind}.String()
	RoleTemplateBinding_KindAPIVersion   = RoleTemplateBinding_Kind + "." + CRDGroupVersion.String()
	RoleTemplateBinding_GroupVersionKind = CRDGroupVersion.WithKind(RoleTemplateBinding_Kind)
)

func init() {
	SchemeBuilder.Register(&RoleTemplateBinding{}, &RoleTemplateBindingList{})
}
