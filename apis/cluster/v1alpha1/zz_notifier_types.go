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

type DingtalkConfigObservation struct {
}

type DingtalkConfigParameters struct {

	// Dingtalk proxy URL
	// +kubebuilder:validation:Optional
	ProxyURL *string `json:"proxyUrl,omitempty" tf:"proxy_url,omitempty"`

	// Required for webhook with sign enabled
	// +kubebuilder:validation:Optional
	SecretSecretRef *v1.SecretKeySelector `json:"secretSecretRef,omitempty" tf:"-"`

	// Webhook URL
	// +kubebuilder:validation:Required
	URL *string `json:"url" tf:"url,omitempty"`
}

type MsteamsConfigObservation struct {
}

type MsteamsConfigParameters struct {

	// MS teams proxy URL
	// +kubebuilder:validation:Optional
	ProxyURL *string `json:"proxyUrl,omitempty" tf:"proxy_url,omitempty"`

	// Webhook URL
	// +kubebuilder:validation:Required
	URL *string `json:"url" tf:"url,omitempty"`
}

type NotifierObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type NotifierParameters struct {

	// Annotations of the resource
	// +kubebuilder:validation:Optional
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// Notifier cluster ID
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-rancher/apis/cluster/v1alpha2.Cluster
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// Notifier description
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	DingtalkConfig []DingtalkConfigParameters `json:"dingtalkConfig,omitempty" tf:"dingtalk_config,omitempty"`

	// Labels of the resource
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// +kubebuilder:validation:Optional
	MsteamsConfig []MsteamsConfigParameters `json:"msteamsConfig,omitempty" tf:"msteams_config,omitempty"`

	// +kubebuilder:validation:Optional
	PagerdutyConfig []PagerdutyConfigParameters `json:"pagerdutyConfig,omitempty" tf:"pagerduty_config,omitempty"`

	// +kubebuilder:validation:Optional
	SMTPConfig []SMTPConfigParameters `json:"smtpConfig,omitempty" tf:"smtp_config,omitempty"`

	// Notifier send resolved
	// +kubebuilder:validation:Optional
	SendResolved *bool `json:"sendResolved,omitempty" tf:"send_resolved,omitempty"`

	// +kubebuilder:validation:Optional
	SlackConfig []SlackConfigParameters `json:"slackConfig,omitempty" tf:"slack_config,omitempty"`

	// +kubebuilder:validation:Optional
	WebhookConfig []WebhookConfigParameters `json:"webhookConfig,omitempty" tf:"webhook_config,omitempty"`

	// +kubebuilder:validation:Optional
	WechatConfig []WechatConfigParameters `json:"wechatConfig,omitempty" tf:"wechat_config,omitempty"`
}

type PagerdutyConfigObservation struct {
}

type PagerdutyConfigParameters struct {

	// Pagerduty proxy URL
	// +kubebuilder:validation:Optional
	ProxyURL *string `json:"proxyUrl,omitempty" tf:"proxy_url,omitempty"`

	// Pagerduty service key
	// +kubebuilder:validation:Required
	ServiceKey *string `json:"serviceKey" tf:"service_key,omitempty"`
}

type SMTPConfigObservation struct {
}

type SMTPConfigParameters struct {

	// SMTP default recipient address
	// +kubebuilder:validation:Required
	DefaultRecipient *string `json:"defaultRecipient" tf:"default_recipient,omitempty"`

	// SMTP host
	// +kubebuilder:validation:Required
	Host *string `json:"host" tf:"host,omitempty"`

	// SMTP password
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// SMTP port
	// +kubebuilder:validation:Required
	Port *float64 `json:"port" tf:"port,omitempty"`

	// SMTP sender
	// +kubebuilder:validation:Required
	Sender *string `json:"sender" tf:"sender,omitempty"`

	// SMTP TLS
	// +kubebuilder:validation:Optional
	TLS *bool `json:"tls,omitempty" tf:"tls,omitempty"`

	// SMTP username
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type SlackConfigObservation struct {
}

type SlackConfigParameters struct {

	// Slack default channel
	// +kubebuilder:validation:Required
	DefaultRecipient *string `json:"defaultRecipient" tf:"default_recipient,omitempty"`

	// Slack proxy URL
	// +kubebuilder:validation:Optional
	ProxyURL *string `json:"proxyUrl,omitempty" tf:"proxy_url,omitempty"`

	// Slack URL
	// +kubebuilder:validation:Required
	URL *string `json:"url" tf:"url,omitempty"`
}

type WebhookConfigObservation struct {
}

type WebhookConfigParameters struct {

	// Webhook proxy URL
	// +kubebuilder:validation:Optional
	ProxyURL *string `json:"proxyUrl,omitempty" tf:"proxy_url,omitempty"`

	// Webhook URL
	// +kubebuilder:validation:Required
	URL *string `json:"url" tf:"url,omitempty"`
}

type WechatConfigObservation struct {
}

type WechatConfigParameters struct {

	// Wechat application agent ID
	// +kubebuilder:validation:Required
	Agent *string `json:"agent" tf:"agent,omitempty"`

	// Wechat corporation ID
	// +kubebuilder:validation:Required
	Corp *string `json:"corp" tf:"corp,omitempty"`

	// Wechat default channel
	// +kubebuilder:validation:Required
	DefaultRecipient *string `json:"defaultRecipient" tf:"default_recipient,omitempty"`

	// Wechat proxy URL
	// +kubebuilder:validation:Optional
	ProxyURL *string `json:"proxyUrl,omitempty" tf:"proxy_url,omitempty"`

	// Wechat recipient type
	// +kubebuilder:validation:Optional
	RecipientType *string `json:"recipientType,omitempty" tf:"recipient_type,omitempty"`

	// Wechat application secret
	// +kubebuilder:validation:Required
	SecretSecretRef v1.SecretKeySelector `json:"secretSecretRef" tf:"-"`
}

// NotifierSpec defines the desired state of Notifier
type NotifierSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NotifierParameters `json:"forProvider"`
}

// NotifierStatus defines the observed state of Notifier.
type NotifierStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NotifierObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Notifier is the Schema for the Notifiers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,rancherjet}
type Notifier struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NotifierSpec   `json:"spec"`
	Status            NotifierStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NotifierList contains a list of Notifiers
type NotifierList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Notifier `json:"items"`
}

// Repository type metadata.
var (
	Notifier_Kind             = "Notifier"
	Notifier_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Notifier_Kind}.String()
	Notifier_KindAPIVersion   = Notifier_Kind + "." + CRDGroupVersion.String()
	Notifier_GroupVersionKind = CRDGroupVersion.WithKind(Notifier_Kind)
)

func init() {
	SchemeBuilder.Register(&Notifier{}, &NotifierList{})
}