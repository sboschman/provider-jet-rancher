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

type Amazonec2CredentialConfigObservation struct {
}

type Amazonec2CredentialConfigParameters struct {

	// AWS Access Key
	// +kubebuilder:validation:Required
	AccessKeySecretRef v1.SecretKeySelector `json:"accessKeySecretRef" tf:"-"`

	// AWS default region
	// +kubebuilder:validation:Optional
	DefaultRegion *string `json:"defaultRegion,omitempty" tf:"default_region,omitempty"`

	// AWS Secret Key
	// +kubebuilder:validation:Required
	SecretKeySecretRef v1.SecretKeySelector `json:"secretKeySecretRef" tf:"-"`
}

type AzureCredentialConfigObservation struct {
}

type AzureCredentialConfigParameters struct {

	// Azure Service Principal Account ID
	// +kubebuilder:validation:Required
	ClientIDSecretRef v1.SecretKeySelector `json:"clientIdSecretRef" tf:"-"`

	// Azure Service Principal Account password
	// +kubebuilder:validation:Required
	ClientSecretSecretRef v1.SecretKeySelector `json:"clientSecretSecretRef" tf:"-"`

	// Azure environment (e.g. AzurePublicCloud, AzureChinaCloud)
	// +kubebuilder:validation:Optional
	Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

	// Azure Subscription ID
	// +kubebuilder:validation:Required
	SubscriptionIDSecretRef v1.SecretKeySelector `json:"subscriptionIdSecretRef" tf:"-"`

	// Azure Tenant ID
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type CredentialObservation struct {
	Driver *string `json:"driver,omitempty" tf:"driver,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type CredentialParameters struct {

	// +kubebuilder:validation:Optional
	Amazonec2CredentialConfig []Amazonec2CredentialConfigParameters `json:"amazonec2CredentialConfig,omitempty" tf:"amazonec2_credential_config,omitempty"`

	// Annotations of the resource
	// +kubebuilder:validation:Optional
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// +kubebuilder:validation:Optional
	AzureCredentialConfig []AzureCredentialConfigParameters `json:"azureCredentialConfig,omitempty" tf:"azure_credential_config,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	DigitaloceanCredentialConfig []DigitaloceanCredentialConfigParameters `json:"digitaloceanCredentialConfig,omitempty" tf:"digitalocean_credential_config,omitempty"`

	// +kubebuilder:validation:Optional
	GoogleCredentialConfig []GoogleCredentialConfigParameters `json:"googleCredentialConfig,omitempty" tf:"google_credential_config,omitempty"`

	// +kubebuilder:validation:Optional
	HarvesterCredentialConfig []HarvesterCredentialConfigParameters `json:"harvesterCredentialConfig,omitempty" tf:"harvester_credential_config,omitempty"`

	// Labels of the resource
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// +kubebuilder:validation:Optional
	LinodeCredentialConfig []LinodeCredentialConfigParameters `json:"linodeCredentialConfig,omitempty" tf:"linode_credential_config,omitempty"`

	// +kubebuilder:validation:Optional
	OpenstackCredentialConfig []OpenstackCredentialConfigParameters `json:"openstackCredentialConfig,omitempty" tf:"openstack_credential_config,omitempty"`

	// +kubebuilder:validation:Optional
	S3CredentialConfig []S3CredentialConfigParameters `json:"s3CredentialConfig,omitempty" tf:"s3_credential_config,omitempty"`

	// +kubebuilder:validation:Optional
	VsphereCredentialConfig []VsphereCredentialConfigParameters `json:"vsphereCredentialConfig,omitempty" tf:"vsphere_credential_config,omitempty"`
}

type DigitaloceanCredentialConfigObservation struct {
}

type DigitaloceanCredentialConfigParameters struct {

	// Digital Ocean access token
	// +kubebuilder:validation:Required
	AccessTokenSecretRef v1.SecretKeySelector `json:"accessTokenSecretRef" tf:"-"`
}

type GoogleCredentialConfigObservation struct {
}

type GoogleCredentialConfigParameters struct {

	// Google auth encoded json
	// +kubebuilder:validation:Required
	AuthEncodedJSONSecretRef v1.SecretKeySelector `json:"authEncodedJsonSecretRef" tf:"-"`
}

type HarvesterCredentialConfigObservation struct {
}

type HarvesterCredentialConfigParameters struct {

	// The cluster id of imported Harvester cluster
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Harvester cluster type. must be imported or external
	// +kubebuilder:validation:Required
	ClusterType *string `json:"clusterType" tf:"cluster_type,omitempty"`

	// Harvester cluster kubeconfig content
	// +kubebuilder:validation:Required
	KubeconfigContentSecretRef v1.SecretKeySelector `json:"kubeconfigContentSecretRef" tf:"-"`
}

type LinodeCredentialConfigObservation struct {
}

type LinodeCredentialConfigParameters struct {

	// Linode API token
	// +kubebuilder:validation:Required
	TokenSecretRef v1.SecretKeySelector `json:"tokenSecretRef" tf:"-"`
}

type OpenstackCredentialConfigObservation struct {
}

type OpenstackCredentialConfigParameters struct {

	// OpenStack password
	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`
}

type S3CredentialConfigObservation struct {
}

type S3CredentialConfigParameters struct {

	// AWS Access Key
	// +kubebuilder:validation:Required
	AccessKeySecretRef v1.SecretKeySelector `json:"accessKeySecretRef" tf:"-"`

	// AWS default bucket
	// +kubebuilder:validation:Optional
	DefaultBucket *string `json:"defaultBucket,omitempty" tf:"default_bucket,omitempty"`

	// AWS default endpoint
	// +kubebuilder:validation:Optional
	DefaultEndpoint *string `json:"defaultEndpoint,omitempty" tf:"default_endpoint,omitempty"`

	// AWS default endpoint CA
	// +kubebuilder:validation:Optional
	DefaultEndpointCASecretRef *v1.SecretKeySelector `json:"defaultEndpointCaSecretRef,omitempty" tf:"-"`

	// AWS default folder
	// +kubebuilder:validation:Optional
	DefaultFolder *string `json:"defaultFolder,omitempty" tf:"default_folder,omitempty"`

	// AWS default region
	// +kubebuilder:validation:Optional
	DefaultRegion *string `json:"defaultRegion,omitempty" tf:"default_region,omitempty"`

	// AWS default skip ssl verify
	// +kubebuilder:validation:Optional
	DefaultSkipSSLVerify *bool `json:"defaultSkipSslVerify,omitempty" tf:"default_skip_ssl_verify,omitempty"`

	// AWS Secret Key
	// +kubebuilder:validation:Required
	SecretKeySecretRef v1.SecretKeySelector `json:"secretKeySecretRef" tf:"-"`
}

type VsphereCredentialConfigObservation struct {
}

type VsphereCredentialConfigParameters struct {

	// vSphere password
	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// vSphere username
	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`

	// vSphere IP/hostname for vCenter
	// +kubebuilder:validation:Required
	Vcenter *string `json:"vcenter" tf:"vcenter,omitempty"`

	// vSphere Port for vCenter
	// +kubebuilder:validation:Optional
	VcenterPort *string `json:"vcenterPort,omitempty" tf:"vcenter_port,omitempty"`
}

// CredentialSpec defines the desired state of Credential
type CredentialSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CredentialParameters `json:"forProvider"`
}

// CredentialStatus defines the observed state of Credential.
type CredentialStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CredentialObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Credential is the Schema for the Credentials API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,rancherjet}
type Credential struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CredentialSpec   `json:"spec"`
	Status            CredentialStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CredentialList contains a list of Credentials
type CredentialList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Credential `json:"items"`
}

// Repository type metadata.
var (
	Credential_Kind             = "Credential"
	Credential_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Credential_Kind}.String()
	Credential_KindAPIVersion   = Credential_Kind + "." + CRDGroupVersion.String()
	Credential_GroupVersionKind = CRDGroupVersion.WithKind(Credential_Kind)
)

func init() {
	SchemeBuilder.Register(&Credential{}, &CredentialList{})
}
