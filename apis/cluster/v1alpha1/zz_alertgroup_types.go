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

type AlertGroupObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AlertGroupParameters struct {

	// Annotations of the resource
	// +kubebuilder:validation:Optional
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// Alert group Cluster ID
	// +kubebuilder:validation:Required
	ClusterID *string `json:"clusterId" tf:"cluster_id,omitempty"`

	// Alert group description
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Alert group interval seconds
	// +kubebuilder:validation:Optional
	GroupIntervalSeconds *float64 `json:"groupIntervalSeconds,omitempty" tf:"group_interval_seconds,omitempty"`

	// Alert group wait seconds
	// +kubebuilder:validation:Optional
	GroupWaitSeconds *float64 `json:"groupWaitSeconds,omitempty" tf:"group_wait_seconds,omitempty"`

	// Labels of the resource
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Alert group recipients
	// +kubebuilder:validation:Optional
	Recipients []RecipientsParameters `json:"recipients,omitempty" tf:"recipients,omitempty"`

	// Alert group repeat interval seconds
	// +kubebuilder:validation:Optional
	RepeatIntervalSeconds *float64 `json:"repeatIntervalSeconds,omitempty" tf:"repeat_interval_seconds,omitempty"`
}

type RecipientsObservation struct {
	NotifierType *string `json:"notifierType,omitempty" tf:"notifier_type,omitempty"`
}

type RecipientsParameters struct {

	// Use notifier default recipient
	// +kubebuilder:validation:Optional
	DefaultRecipient *bool `json:"defaultRecipient,omitempty" tf:"default_recipient,omitempty"`

	// Recipient notifier ID
	// +kubebuilder:validation:Required
	NotifierID *string `json:"notifierId" tf:"notifier_id,omitempty"`

	// Recipient
	// +kubebuilder:validation:Optional
	Recipient *string `json:"recipient,omitempty" tf:"recipient,omitempty"`
}

// AlertGroupSpec defines the desired state of AlertGroup
type AlertGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AlertGroupParameters `json:"forProvider"`
}

// AlertGroupStatus defines the observed state of AlertGroup.
type AlertGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AlertGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AlertGroup is the Schema for the AlertGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,rancherjet}
type AlertGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AlertGroupSpec   `json:"spec"`
	Status            AlertGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AlertGroupList contains a list of AlertGroups
type AlertGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AlertGroup `json:"items"`
}

// Repository type metadata.
var (
	AlertGroup_Kind             = "AlertGroup"
	AlertGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AlertGroup_Kind}.String()
	AlertGroup_KindAPIVersion   = AlertGroup_Kind + "." + CRDGroupVersion.String()
	AlertGroup_GroupVersionKind = CRDGroupVersion.WithKind(AlertGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&AlertGroup{}, &AlertGroupList{})
}
