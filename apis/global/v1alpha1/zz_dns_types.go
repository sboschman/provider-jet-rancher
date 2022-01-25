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

type DNSObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DNSParameters struct {

	// Annotations of the resource
	// +kubebuilder:validation:Optional
	Annotations map[string]string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// +kubebuilder:validation:Required
	Fqdn *string `json:"fqdn" tf:"fqdn,omitempty"`

	// Labels of the resource
	// +kubebuilder:validation:Optional
	Labels map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`

	// +kubebuilder:validation:Optional
	MultiClusterAppID *string `json:"multiClusterAppId,omitempty" tf:"multi_cluster_app_id,omitempty"`

	// +kubebuilder:validation:Optional
	ProjectIds []*string `json:"projectIds,omitempty" tf:"project_ids,omitempty"`

	// +kubebuilder:validation:Required
	ProviderID *string `json:"providerId" tf:"provider_id,omitempty"`

	// +kubebuilder:validation:Optional
	TTL *int64 `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

// DNSSpec defines the desired state of DNS
type DNSSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DNSParameters `json:"forProvider"`
}

// DNSStatus defines the observed state of DNS.
type DNSStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DNSObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DNS is the Schema for the DNSs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,rancherjet}
type DNS struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DNSSpec   `json:"spec"`
	Status            DNSStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DNSList contains a list of DNSs
type DNSList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DNS `json:"items"`
}

// Repository type metadata.
var (
	DNS_Kind             = "DNS"
	DNS_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DNS_Kind}.String()
	DNS_KindAPIVersion   = DNS_Kind + "." + CRDGroupVersion.String()
	DNS_GroupVersionKind = CRDGroupVersion.WithKind(DNS_Kind)
)

func init() {
	SchemeBuilder.Register(&DNS{}, &DNSList{})
}