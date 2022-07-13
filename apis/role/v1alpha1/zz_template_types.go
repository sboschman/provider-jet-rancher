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

type RulesObservation struct {
}

type RulesParameters struct {

	// Policy rule api groups
	// +kubebuilder:validation:Optional
	APIGroups []*string `json:"apiGroups,omitempty" tf:"api_groups,omitempty"`

	// Policy rule non resource urls
	// +kubebuilder:validation:Optional
	NonResourceUrls []*string `json:"nonResourceUrls,omitempty" tf:"non_resource_urls,omitempty"`

	// Policy rule resource names
	// +kubebuilder:validation:Optional
	ResourceNames []*string `json:"resourceNames,omitempty" tf:"resource_names,omitempty"`

	// Policy rule resources
	// +kubebuilder:validation:Optional
	Resources []*string `json:"resources,omitempty" tf:"resources,omitempty"`

	// Policy rule verbs
	// +kubebuilder:validation:Optional
	Verbs []*string `json:"verbs,omitempty" tf:"verbs,omitempty"`
}

type TemplateObservation struct {
	Builtin *bool `json:"builtin,omitempty" tf:"builtin,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TemplateParameters struct {

	// Administrative role template
	// +kubebuilder:validation:Optional
	Administrative *bool `json:"administrative,omitempty" tf:"administrative,omitempty"`

	// Annotations of the resource
	// +kubebuilder:validation:Optional
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// Context role template
	// +kubebuilder:validation:Optional
	Context *string `json:"context,omitempty" tf:"context,omitempty"`

	// Default role template for new created cluster or project
	// +kubebuilder:validation:Optional
	DefaultRole *bool `json:"defaultRole,omitempty" tf:"default_role,omitempty"`

	// Role template policy description
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// External role template
	// +kubebuilder:validation:Optional
	External *bool `json:"external,omitempty" tf:"external,omitempty"`

	// Hidden role template
	// +kubebuilder:validation:Optional
	Hidden *bool `json:"hidden,omitempty" tf:"hidden,omitempty"`

	// Labels of the resource
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Locked role template
	// +kubebuilder:validation:Optional
	Locked *bool `json:"locked,omitempty" tf:"locked,omitempty"`

	// Inherit role template IDs
	// +kubebuilder:validation:Optional
	RoleTemplateIds []*string `json:"roleTemplateIds,omitempty" tf:"role_template_ids,omitempty"`

	// Role template policy rules
	// +kubebuilder:validation:Optional
	Rules []RulesParameters `json:"rules,omitempty" tf:"rules,omitempty"`
}

// TemplateSpec defines the desired state of Template
type TemplateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TemplateParameters `json:"forProvider"`
}

// TemplateStatus defines the observed state of Template.
type TemplateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TemplateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Template is the Schema for the Templates API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,rancherjet}
type Template struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TemplateSpec   `json:"spec"`
	Status            TemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TemplateList contains a list of Templates
type TemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Template `json:"items"`
}

// Repository type metadata.
var (
	Template_Kind             = "Template"
	Template_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Template_Kind}.String()
	Template_KindAPIVersion   = Template_Kind + "." + CRDGroupVersion.String()
	Template_GroupVersionKind = CRDGroupVersion.WithKind(Template_Kind)
)

func init() {
	SchemeBuilder.Register(&Template{}, &TemplateList{})
}
