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

type AgentEnvVarsObservation struct {
}

type AgentEnvVarsParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type ClusterRegistrationTokenObservation struct {
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	Command *string `json:"command,omitempty" tf:"command,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	InsecureCommand *string `json:"insecureCommand,omitempty" tf:"insecure_command,omitempty"`

	InsecureNodeCommand *string `json:"insecureNodeCommand,omitempty" tf:"insecure_node_command,omitempty"`

	InsecureWindowsNodeCommand *string `json:"insecureWindowsNodeCommand,omitempty" tf:"insecure_windows_node_command,omitempty"`

	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	ManifestURL *string `json:"manifestUrl,omitempty" tf:"manifest_url,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	NodeCommand *string `json:"nodeCommand,omitempty" tf:"node_command,omitempty"`

	Token *string `json:"token,omitempty" tf:"token,omitempty"`

	WindowsNodeCommand *string `json:"windowsNodeCommand,omitempty" tf:"windows_node_command,omitempty"`
}

type ClusterRegistrationTokenParameters struct {
}

type ConfigsObservation struct {
}

type ConfigsParameters struct {

	// Registry auth config secret name
	// +kubebuilder:validation:Optional
	AuthConfigSecretName *string `json:"authConfigSecretName,omitempty" tf:"auth_config_secret_name,omitempty"`

	// Registry CA bundle
	// +kubebuilder:validation:Optional
	CABundle *string `json:"caBundle,omitempty" tf:"ca_bundle,omitempty"`

	// Registry hostname
	// +kubebuilder:validation:Required
	Hostname *string `json:"hostname" tf:"hostname,omitempty"`

	// Registry insecure connectivity
	// +kubebuilder:validation:Optional
	Insecure *bool `json:"insecure,omitempty" tf:"insecure,omitempty"`

	// Registry TLS secret name. TLS is a pair of Cert/Key
	// +kubebuilder:validation:Optional
	TLSSecretName *string `json:"tlsSecretName,omitempty" tf:"tls_secret_name,omitempty"`
}

type ControlPlaneDrainOptionsObservation struct {
}

type ControlPlaneDrainOptionsParameters struct {

	// Drain options delete empty dir data
	// +kubebuilder:validation:Optional
	DeleteEmptyDirData *bool `json:"deleteEmptyDirData,omitempty" tf:"delete_empty_dir_data,omitempty"`

	// Drain options disable eviction
	// +kubebuilder:validation:Optional
	DisableEviction *bool `json:"disableEviction,omitempty" tf:"disable_eviction,omitempty"`

	// Drain options enabled?
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Drain options force
	// +kubebuilder:validation:Optional
	Force *bool `json:"force,omitempty" tf:"force,omitempty"`

	// Drain options grace period
	// +kubebuilder:validation:Optional
	GracePeriod *float64 `json:"gracePeriod,omitempty" tf:"grace_period,omitempty"`

	// Drain options ignore daemon sets
	// +kubebuilder:validation:Optional
	IgnoreDaemonSets *bool `json:"ignoreDaemonSets,omitempty" tf:"ignore_daemon_sets,omitempty"`

	// Drain options ignore errors
	// +kubebuilder:validation:Optional
	IgnoreErrors *bool `json:"ignoreErrors,omitempty" tf:"ignore_errors,omitempty"`

	// Drain options skip wait for delete timeout seconds
	// +kubebuilder:validation:Optional
	SkipWaitForDeleteTimeoutSeconds *float64 `json:"skipWaitForDeleteTimeoutSeconds,omitempty" tf:"skip_wait_for_delete_timeout_seconds,omitempty"`

	// Drain options timeout
	// +kubebuilder:validation:Optional
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`
}

type EtcdSnapshotCreateObservation struct {
}

type EtcdSnapshotCreateParameters struct {

	// ETCD generation to initiate a snapshot
	// +kubebuilder:validation:Required
	Generation *float64 `json:"generation" tf:"generation,omitempty"`
}

type EtcdSnapshotRestoreObservation struct {
}

type EtcdSnapshotRestoreParameters struct {

	// ETCD snapshot desired generation
	// +kubebuilder:validation:Required
	Generation *float64 `json:"generation" tf:"generation,omitempty"`

	// ETCD snapshot name to restore
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// ETCD restore RKE config (set to none, all, or kubernetesVersion)
	// +kubebuilder:validation:Optional
	RestoreRkeConfig *string `json:"restoreRkeConfig,omitempty" tf:"restore_rke_config,omitempty"`
}

type LocalAuthEndpointObservation struct {
}

type LocalAuthEndpointParameters struct {

	// +kubebuilder:validation:Optional
	CACerts *string `json:"caCerts,omitempty" tf:"ca_certs,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`
}

type MachineConfigObservation struct {
}

type MachineConfigParameters struct {

	// Machine config kind
	// +kubebuilder:validation:Required
	Kind *string `json:"kind" tf:"kind,omitempty"`

	// Machine config name
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type MachineLabelSelectorObservation struct {
}

type MachineLabelSelectorParameters struct {

	// Label selector match expressions
	// +kubebuilder:validation:Optional
	MatchExpressions []MatchExpressionsParameters `json:"matchExpressions,omitempty" tf:"match_expressions,omitempty"`

	// Label selector match labels
	// +kubebuilder:validation:Optional
	MatchLabels map[string]*string `json:"matchLabels,omitempty" tf:"match_labels,omitempty"`
}

type MachinePoolsObservation struct {
}

type MachinePoolsParameters struct {

	// Annotations of the resource
	// +kubebuilder:validation:Optional
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// Machine pool cloud credential secret name
	// +kubebuilder:validation:Required
	CloudCredentialSecretName *string `json:"cloudCredentialSecretName" tf:"cloud_credential_secret_name,omitempty"`

	// Machine pool control plane role
	// +kubebuilder:validation:Optional
	ControlPlaneRole *bool `json:"controlPlaneRole,omitempty" tf:"control_plane_role,omitempty"`

	// Machine pool drain before delete
	// +kubebuilder:validation:Optional
	DrainBeforeDelete *bool `json:"drainBeforeDelete,omitempty" tf:"drain_before_delete,omitempty"`

	// Machine pool etcd role
	// +kubebuilder:validation:Optional
	EtcdRole *bool `json:"etcdRole,omitempty" tf:"etcd_role,omitempty"`

	// Labels of the resource
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Machine config data
	// +kubebuilder:validation:Required
	MachineConfig []MachineConfigParameters `json:"machineConfig" tf:"machine_config,omitempty"`

	// max unhealthy nodes for automated replacement to be allowed
	// +kubebuilder:validation:Optional
	MaxUnhealthy *string `json:"maxUnhealthy,omitempty" tf:"max_unhealthy,omitempty"`

	// Machine pool name
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// seconds to wait for machine pool drain to complete before machine deletion
	// +kubebuilder:validation:Optional
	NodeDrainTimeout *float64 `json:"nodeDrainTimeout,omitempty" tf:"node_drain_timeout,omitempty"`

	// seconds a new node has to become active before it is replaced
	// +kubebuilder:validation:Optional
	NodeStartupTimeoutSeconds *float64 `json:"nodeStartupTimeoutSeconds,omitempty" tf:"node_startup_timeout_seconds,omitempty"`

	// Machine pool paused
	// +kubebuilder:validation:Optional
	Paused *bool `json:"paused,omitempty" tf:"paused,omitempty"`

	// Machine pool quantity
	// +kubebuilder:validation:Optional
	Quantity *float64 `json:"quantity,omitempty" tf:"quantity,omitempty"`

	// Machine pool rolling update
	// +kubebuilder:validation:Optional
	RollingUpdate []MachinePoolsRollingUpdateParameters `json:"rollingUpdate,omitempty" tf:"rolling_update,omitempty"`

	// Machine pool taints
	// +kubebuilder:validation:Optional
	Taints []TaintsParameters `json:"taints,omitempty" tf:"taints,omitempty"`

	// seconds an unhealthy node has to become active before it is replaced
	// +kubebuilder:validation:Optional
	UnhealthyNodeTimeoutSeconds *float64 `json:"unhealthyNodeTimeoutSeconds,omitempty" tf:"unhealthy_node_timeout_seconds,omitempty"`

	// range of unhealthy nodes for automated replacement to be allowed
	// +kubebuilder:validation:Optional
	UnhealthyRange *string `json:"unhealthyRange,omitempty" tf:"unhealthy_range,omitempty"`

	// Machine pool worker role
	// +kubebuilder:validation:Optional
	WorkerRole *bool `json:"workerRole,omitempty" tf:"worker_role,omitempty"`
}

type MachinePoolsRollingUpdateObservation struct {
}

type MachinePoolsRollingUpdateParameters struct {

	// Rolling update max surge
	// +kubebuilder:validation:Optional
	MaxSurge *string `json:"maxSurge,omitempty" tf:"max_surge,omitempty"`

	// Rolling update max unavailable
	// +kubebuilder:validation:Optional
	MaxUnavailable *string `json:"maxUnavailable,omitempty" tf:"max_unavailable,omitempty"`
}

type MachineSelectorConfigObservation struct {
}

type MachineSelectorConfigParameters struct {

	// Machine selector config
	// +kubebuilder:validation:Optional
	Config map[string]*string `json:"config,omitempty" tf:"config,omitempty"`

	// Machine label selector
	// +kubebuilder:validation:Optional
	MachineLabelSelector []MachineLabelSelectorParameters `json:"machineLabelSelector,omitempty" tf:"machine_label_selector,omitempty"`
}

type MatchExpressionsObservation struct {
}

type MatchExpressionsParameters struct {

	// Label selector requirement key
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Label selector operator
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Label selector requirement values
	// +kubebuilder:validation:Optional
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type MirrorsObservation struct {
}

type MirrorsParameters struct {

	// Registry mirror endpoints
	// +kubebuilder:validation:Optional
	Endpoints []*string `json:"endpoints,omitempty" tf:"endpoints,omitempty"`

	// Registry hostname
	// +kubebuilder:validation:Required
	Hostname *string `json:"hostname" tf:"hostname,omitempty"`

	// Registry mirror rewrites
	// +kubebuilder:validation:Optional
	Rewrites map[string]*string `json:"rewrites,omitempty" tf:"rewrites,omitempty"`
}

type RegistriesObservation struct {
}

type RegistriesParameters struct {

	// Registry config
	// +kubebuilder:validation:Optional
	Configs []ConfigsParameters `json:"configs,omitempty" tf:"configs,omitempty"`

	// Registry mirrors
	// +kubebuilder:validation:Optional
	Mirrors []MirrorsParameters `json:"mirrors,omitempty" tf:"mirrors,omitempty"`
}

type RkeConfigEtcdObservation struct {
}

type RkeConfigEtcdParameters struct {

	// Disable ETCD snapshots
	// +kubebuilder:validation:Optional
	DisableSnapshots *bool `json:"disableSnapshots,omitempty" tf:"disable_snapshots,omitempty"`

	// ETCD snapshot S3 config
	// +kubebuilder:validation:Optional
	S3Config []S3ConfigParameters `json:"s3Config,omitempty" tf:"s3_config,omitempty"`

	// ETCD snapshot retention
	// +kubebuilder:validation:Optional
	SnapshotRetention *float64 `json:"snapshotRetention,omitempty" tf:"snapshot_retention,omitempty"`

	// ETCD snapshot schedule cron (e.g `"0 */5 * * *"`)
	// +kubebuilder:validation:Optional
	SnapshotScheduleCron *string `json:"snapshotScheduleCron,omitempty" tf:"snapshot_schedule_cron,omitempty"`
}

type RkeConfigLocalAuthEndpointObservation struct {
}

type RkeConfigLocalAuthEndpointParameters struct {

	// +kubebuilder:validation:Optional
	CACerts *string `json:"caCerts,omitempty" tf:"ca_certs,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`
}

type RkeConfigUpgradeStrategyObservation struct {
}

type RkeConfigUpgradeStrategyParameters struct {

	// How many controlplane nodes should be upgrade at time, 0 is infinite. Percentages are also accepted
	// +kubebuilder:validation:Optional
	ControlPlaneConcurrency *string `json:"controlPlaneConcurrency,omitempty" tf:"control_plane_concurrency,omitempty"`

	// Controlplane nodes drain options
	// +kubebuilder:validation:Optional
	ControlPlaneDrainOptions []ControlPlaneDrainOptionsParameters `json:"controlPlaneDrainOptions,omitempty" tf:"control_plane_drain_options,omitempty"`

	// How many worker nodes should be upgrade at time
	// +kubebuilder:validation:Optional
	WorkerConcurrency *string `json:"workerConcurrency,omitempty" tf:"worker_concurrency,omitempty"`

	// Worker nodes drain options
	// +kubebuilder:validation:Optional
	WorkerDrainOptions []WorkerDrainOptionsParameters `json:"workerDrainOptions,omitempty" tf:"worker_drain_options,omitempty"`
}

type RotateCertificatesObservation struct {
}

type RotateCertificatesParameters struct {

	// Desired certificate rotation generation.
	// +kubebuilder:validation:Required
	Generation *float64 `json:"generation" tf:"generation,omitempty"`

	// Service certificates to rotate with this generation.
	// +kubebuilder:validation:Optional
	Services []*string `json:"services,omitempty" tf:"services,omitempty"`
}

type S3ConfigObservation struct {
}

type S3ConfigParameters struct {

	// ETCD snapshot S3 bucket
	// +kubebuilder:validation:Required
	Bucket *string `json:"bucket" tf:"bucket,omitempty"`

	// ETCD snapshot S3 cloud credential name
	// +kubebuilder:validation:Optional
	CloudCredentialName *string `json:"cloudCredentialName,omitempty" tf:"cloud_credential_name,omitempty"`

	// ETCD snapshot S3 endpoint
	// +kubebuilder:validation:Required
	Endpoint *string `json:"endpoint" tf:"endpoint,omitempty"`

	// ETCD snapshot S3 endpoint CA
	// +kubebuilder:validation:Optional
	EndpointCA *string `json:"endpointCa,omitempty" tf:"endpoint_ca,omitempty"`

	// ETCD snapshot S3 folder
	// +kubebuilder:validation:Optional
	Folder *string `json:"folder,omitempty" tf:"folder,omitempty"`

	// ETCD snapshot S3 region
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Disable ETCD skip ssl verify
	// +kubebuilder:validation:Optional
	SkipSSLVerify *bool `json:"skipSslVerify,omitempty" tf:"skip_ssl_verify,omitempty"`
}

type TaintsObservation struct {
}

type TaintsParameters struct {

	// +kubebuilder:validation:Optional
	Effect *string `json:"effect,omitempty" tf:"effect,omitempty"`

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type V2Observation struct {
	ClusterV1ID *string `json:"clusterV1Id,omitempty" tf:"cluster_v1_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	ResourceVersion *string `json:"resourceVersion,omitempty" tf:"resource_version,omitempty"`
}

type V2Parameters struct {

	// Cluster V2 default agent env vars
	// +kubebuilder:validation:Optional
	AgentEnvVars []AgentEnvVarsParameters `json:"agentEnvVars,omitempty" tf:"agent_env_vars,omitempty"`

	// Annotations of the resource
	// +kubebuilder:validation:Optional
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// Cluster V2 cloud credential secret name
	// +kubebuilder:validation:Optional
	CloudCredentialSecretName *string `json:"cloudCredentialSecretName,omitempty" tf:"cloud_credential_secret_name,omitempty"`

	// Cluster V2 default cluster role for project members
	// +kubebuilder:validation:Optional
	DefaultClusterRoleForProjectMembers *string `json:"defaultClusterRoleForProjectMembers,omitempty" tf:"default_cluster_role_for_project_members,omitempty"`

	// Cluster V2 default pod security policy template name
	// +kubebuilder:validation:Optional
	DefaultPodSecurityPolicyTemplateName *string `json:"defaultPodSecurityPolicyTemplateName,omitempty" tf:"default_pod_security_policy_template_name,omitempty"`

	// Enable k8s network policy
	// +kubebuilder:validation:Optional
	EnableNetworkPolicy *bool `json:"enableNetworkPolicy,omitempty" tf:"enable_network_policy,omitempty"`

	// +kubebuilder:validation:Optional
	FleetNamespace *string `json:"fleetNamespace,omitempty" tf:"fleet_namespace,omitempty"`

	// Cluster V2 kubernetes version
	// +kubebuilder:validation:Required
	KubernetesVersion *string `json:"kubernetesVersion" tf:"kubernetes_version,omitempty"`

	// Labels of the resource
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Cluster V2 local auth endpoint
	// +kubebuilder:validation:Optional
	LocalAuthEndpoint []LocalAuthEndpointParameters `json:"localAuthEndpoint,omitempty" tf:"local_auth_endpoint,omitempty"`

	// Cluster V2 rke config
	// +kubebuilder:validation:Optional
	RkeConfig []V2RkeConfigParameters `json:"rkeConfig,omitempty" tf:"rke_config,omitempty"`
}

type V2RkeConfigObservation struct {
}

type V2RkeConfigParameters struct {

	// Cluster V2 additional manifest
	// +kubebuilder:validation:Optional
	AdditionalManifest *string `json:"additionalManifest,omitempty" tf:"additional_manifest,omitempty"`

	// Cluster V2 chart values. It should be in YAML format
	// +kubebuilder:validation:Optional
	ChartValues *string `json:"chartValues,omitempty" tf:"chart_values,omitempty"`

	// Cluster V2 etcd
	// +kubebuilder:validation:Optional
	Etcd []RkeConfigEtcdParameters `json:"etcd,omitempty" tf:"etcd,omitempty"`

	// Cluster V2 etcd snapshot create
	// +kubebuilder:validation:Optional
	EtcdSnapshotCreate []EtcdSnapshotCreateParameters `json:"etcdSnapshotCreate,omitempty" tf:"etcd_snapshot_create,omitempty"`

	// Cluster V2 etcd snapshot restore
	// +kubebuilder:validation:Optional
	EtcdSnapshotRestore []EtcdSnapshotRestoreParameters `json:"etcdSnapshotRestore,omitempty" tf:"etcd_snapshot_restore,omitempty"`

	// Cluster V2 local auth endpoint
	// +kubebuilder:validation:Optional
	LocalAuthEndpoint []RkeConfigLocalAuthEndpointParameters `json:"localAuthEndpoint,omitempty" tf:"local_auth_endpoint,omitempty"`

	// Cluster V2 machine global config
	// +kubebuilder:validation:Optional
	MachineGlobalConfig *string `json:"machineGlobalConfig,omitempty" tf:"machine_global_config,omitempty"`

	// Cluster V2 machine pools
	// +kubebuilder:validation:Optional
	MachinePools []MachinePoolsParameters `json:"machinePools,omitempty" tf:"machine_pools,omitempty"`

	// Cluster V2 machine selector config
	// +kubebuilder:validation:Optional
	MachineSelectorConfig []MachineSelectorConfigParameters `json:"machineSelectorConfig,omitempty" tf:"machine_selector_config,omitempty"`

	// Cluster V2 registries
	// +kubebuilder:validation:Optional
	Registries []RegistriesParameters `json:"registries,omitempty" tf:"registries,omitempty"`

	// Cluster V2 certificate rotation
	// +kubebuilder:validation:Optional
	RotateCertificates []RotateCertificatesParameters `json:"rotateCertificates,omitempty" tf:"rotate_certificates,omitempty"`

	// Cluster V2 upgrade strategy
	// +kubebuilder:validation:Optional
	UpgradeStrategy []RkeConfigUpgradeStrategyParameters `json:"upgradeStrategy,omitempty" tf:"upgrade_strategy,omitempty"`
}

type WorkerDrainOptionsObservation struct {
}

type WorkerDrainOptionsParameters struct {

	// Drain options delete empty dir data
	// +kubebuilder:validation:Optional
	DeleteEmptyDirData *bool `json:"deleteEmptyDirData,omitempty" tf:"delete_empty_dir_data,omitempty"`

	// Drain options disable eviction
	// +kubebuilder:validation:Optional
	DisableEviction *bool `json:"disableEviction,omitempty" tf:"disable_eviction,omitempty"`

	// Drain options enabled?
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Drain options force
	// +kubebuilder:validation:Optional
	Force *bool `json:"force,omitempty" tf:"force,omitempty"`

	// Drain options grace period
	// +kubebuilder:validation:Optional
	GracePeriod *float64 `json:"gracePeriod,omitempty" tf:"grace_period,omitempty"`

	// Drain options ignore daemon sets
	// +kubebuilder:validation:Optional
	IgnoreDaemonSets *bool `json:"ignoreDaemonSets,omitempty" tf:"ignore_daemon_sets,omitempty"`

	// Drain options ignore errors
	// +kubebuilder:validation:Optional
	IgnoreErrors *bool `json:"ignoreErrors,omitempty" tf:"ignore_errors,omitempty"`

	// Drain options skip wait for delete timeout seconds
	// +kubebuilder:validation:Optional
	SkipWaitForDeleteTimeoutSeconds *float64 `json:"skipWaitForDeleteTimeoutSeconds,omitempty" tf:"skip_wait_for_delete_timeout_seconds,omitempty"`

	// Drain options timeout
	// +kubebuilder:validation:Optional
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`
}

// V2Spec defines the desired state of V2
type V2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     V2Parameters `json:"forProvider"`
}

// V2Status defines the observed state of V2.
type V2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        V2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// V2 is the Schema for the V2s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,rancherjet}
type V2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              V2Spec   `json:"spec"`
	Status            V2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// V2List contains a list of V2s
type V2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []V2 `json:"items"`
}

// Repository type metadata.
var (
	V2_Kind             = "V2"
	V2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: V2_Kind}.String()
	V2_KindAPIVersion   = V2_Kind + "." + CRDGroupVersion.String()
	V2_GroupVersionKind = CRDGroupVersion.WithKind(V2_Kind)
)

func init() {
	SchemeBuilder.Register(&V2{}, &V2List{})
}
