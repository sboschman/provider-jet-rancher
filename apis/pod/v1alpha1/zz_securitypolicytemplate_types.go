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

type AllowedCsiDriverObservation struct {
}

type AllowedCsiDriverParameters struct {

	// Name is the registered name of the CSI driver
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type AllowedFlexVolumeObservation struct {
}

type AllowedFlexVolumeParameters struct {

	// driver is the name of the Flexvolume driver.
	// +kubebuilder:validation:Required
	Driver *string `json:"driver" tf:"driver,omitempty"`
}

type AllowedHostPathObservation struct {
}

type AllowedHostPathParameters struct {

	// pathPrefix is the path prefix that the host volume must match. It does not support `*`. Trailing slashes are trimmed when validating the path prefix with a host path.
	// +kubebuilder:validation:Required
	PathPrefix *string `json:"pathPrefix" tf:"path_prefix,omitempty"`

	// when set to true, will allow host volumes matching the pathPrefix only if all volume mounts are readOnly.
	// +kubebuilder:validation:Optional
	ReadOnly *bool `json:"readOnly,omitempty" tf:"read_only,omitempty"`
}

type FsGroupObservation struct {
}

type FsGroupParameters struct {

	// ranges are the allowed ranges of fs groups.  If you would like to force a single fs group then supply a single range with the same start and end. Required for MustRunAs.
	// +kubebuilder:validation:Optional
	Range []RangeParameters `json:"range,omitempty" tf:"range,omitempty"`

	// rule is the strategy that will dictate what FSGroup is used in the SecurityContext.
	// +kubebuilder:validation:Optional
	Rule *string `json:"rule,omitempty" tf:"rule,omitempty"`
}

type HostPortObservation struct {
}

type HostPortParameters struct {

	// max is the end of the range, inclusive.
	// +kubebuilder:validation:Required
	Max *float64 `json:"max" tf:"max,omitempty"`

	// min is the start of the range, inclusive.
	// +kubebuilder:validation:Required
	Min *float64 `json:"min" tf:"min,omitempty"`
}

type RangeObservation struct {
}

type RangeParameters struct {

	// max is the end of the range, inclusive.
	// +kubebuilder:validation:Required
	Max *float64 `json:"max" tf:"max,omitempty"`

	// min is the start of the range, inclusive.
	// +kubebuilder:validation:Required
	Min *float64 `json:"min" tf:"min,omitempty"`
}

type RunAsGroupObservation struct {
}

type RunAsGroupParameters struct {

	// ranges are the allowed ranges of gids that may be used. If you would like to force a single gid then supply a single range with the same start and end. Required for MustRunAs.
	// +kubebuilder:validation:Optional
	Range []RunAsGroupRangeParameters `json:"range,omitempty" tf:"range,omitempty"`

	// rule is the strategy that will dictate the allowable RunAsGroup values that may be set.
	// +kubebuilder:validation:Required
	Rule *string `json:"rule" tf:"rule,omitempty"`
}

type RunAsGroupRangeObservation struct {
}

type RunAsGroupRangeParameters struct {

	// max is the end of the range, inclusive.
	// +kubebuilder:validation:Required
	Max *float64 `json:"max" tf:"max,omitempty"`

	// min is the start of the range, inclusive.
	// +kubebuilder:validation:Required
	Min *float64 `json:"min" tf:"min,omitempty"`
}

type RunAsUserObservation struct {
}

type RunAsUserParameters struct {

	// ranges are the allowed ranges of uids that may be used. If you would like to force a single uid then supply a single range with the same start and end. Required for MustRunAs.
	// +kubebuilder:validation:Optional
	Range []RunAsUserRangeParameters `json:"range,omitempty" tf:"range,omitempty"`

	// rule is the strategy that will dictate the allowable RunAsUser values that may be set.
	// +kubebuilder:validation:Required
	Rule *string `json:"rule" tf:"rule,omitempty"`
}

type RunAsUserRangeObservation struct {
}

type RunAsUserRangeParameters struct {

	// max is the end of the range, inclusive.
	// +kubebuilder:validation:Required
	Max *float64 `json:"max" tf:"max,omitempty"`

	// min is the start of the range, inclusive.
	// +kubebuilder:validation:Required
	Min *float64 `json:"min" tf:"min,omitempty"`
}

type RuntimeClassObservation struct {
}

type RuntimeClassParameters struct {

	// allowedRuntimeClassNames is a whitelist of RuntimeClass names that may be specified on a pod. A value of "*" means that any RuntimeClass name is allowed, and must be the only item in the list. An empty list requires the RuntimeClassName field to be unset.
	// +kubebuilder:validation:Required
	AllowedRuntimeClassNames []*string `json:"allowedRuntimeClassNames" tf:"allowed_runtime_class_names,omitempty"`

	// defaultRuntimeClassName is the default RuntimeClassName to set on the pod. The default MUST be allowed by the allowedRuntimeClassNames list. A value of nil does not mutate the Pod.
	// +kubebuilder:validation:Optional
	DefaultRuntimeClassName *string `json:"defaultRuntimeClassName,omitempty" tf:"default_runtime_class_name,omitempty"`
}

type SeLinuxObservation struct {
}

type SeLinuxOptionObservation struct {
}

type SeLinuxOptionParameters struct {

	// Level is SELinux level label that applies to the container.
	// +kubebuilder:validation:Optional
	Level *string `json:"level,omitempty" tf:"level,omitempty"`

	// Role is a SELinux role label that applies to the container.
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// Type is a SELinux type label that applies to the container.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// User is a SELinux user label that applies to the container.
	// +kubebuilder:validation:Optional
	User *string `json:"user,omitempty" tf:"user,omitempty"`
}

type SeLinuxParameters struct {

	// rule is the strategy that will dictate the allowable labels that may be set.
	// +kubebuilder:validation:Required
	Rule *string `json:"rule" tf:"rule,omitempty"`

	// seLinuxOptions required to run as; required for MustRunAs. More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
	// +kubebuilder:validation:Optional
	SeLinuxOption []SeLinuxOptionParameters `json:"seLinuxOption,omitempty" tf:"se_linux_option,omitempty"`
}

type SecurityPolicyTemplateObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SecurityPolicyTemplateParameters struct {

	// allowPrivilegeEscalation determines if a pod can request to allow privilege escalation. If unspecified, defaults to true.
	// +kubebuilder:validation:Optional
	AllowPrivilegeEscalation *bool `json:"allowPrivilegeEscalation,omitempty" tf:"allow_privilege_escalation,omitempty"`

	// allowedCapabilities is a list of capabilities that can be requested to add to the container. Capabilities in this field may be added at the pod author's discretion. You must not list a capability in both allowedCapabilities and requiredDropCapabilities.
	// +kubebuilder:validation:Optional
	AllowedCapabilities []*string `json:"allowedCapabilities,omitempty" tf:"allowed_capabilities,omitempty"`

	// AllowedCSIDrivers is a whitelist of inline CSI drivers that must be explicitly set to be embedded within a pod spec. An empty value indicates that any CSI driver can be used for inline ephemeral volumes. This is an alpha field, and is only honored if the API server enables the CSIInlineVolume feature gate.
	// +kubebuilder:validation:Optional
	AllowedCsiDriver []AllowedCsiDriverParameters `json:"allowedCsiDriver,omitempty" tf:"allowed_csi_driver,omitempty"`

	// allowedFlexVolumes is a whitelist of allowed Flexvolumes.  Empty or nil indicates that all Flexvolumes may be used.  This parameter is effective only when the usage of the Flexvolumes is allowed in the "volumes" field.
	// +kubebuilder:validation:Optional
	AllowedFlexVolume []AllowedFlexVolumeParameters `json:"allowedFlexVolume,omitempty" tf:"allowed_flex_volume,omitempty"`

	// allowedHostPaths is a white list of allowed host paths. Empty indicates that all host paths may be used.
	// +kubebuilder:validation:Optional
	AllowedHostPath []AllowedHostPathParameters `json:"allowedHostPath,omitempty" tf:"allowed_host_path,omitempty"`

	// AllowedProcMountTypes is a whitelist of allowed ProcMountTypes. Empty or nil indicates that only the DefaultProcMountType may be used. This requires the ProcMountType feature flag to be enabled.
	// +kubebuilder:validation:Optional
	AllowedProcMountTypes []*string `json:"allowedProcMountTypes,omitempty" tf:"allowed_proc_mount_types,omitempty"`

	// allowedUnsafeSysctls is a list of explicitly allowed unsafe sysctls, defaults to none. Each entry is either a plain sysctl name or ends in "*" in which case it is considered as a prefix of allowed sysctls. Single * means all unsafe sysctls are allowed. Kubelet has to whitelist all allowed unsafe sysctls explicitly to avoid rejection.
	// +kubebuilder:validation:Optional
	AllowedUnsafeSysctls []*string `json:"allowedUnsafeSysctls,omitempty" tf:"allowed_unsafe_sysctls,omitempty"`

	// Annotations of the resource
	// +kubebuilder:validation:Optional
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// defaultAddCapabilities is the default set of capabilities that will be added to the container unless the pod spec specifically drops the capability.  You may not list a capability in both defaultAddCapabilities and requiredDropCapabilities. Capabilities added here are implicitly allowed, and need not be included in the allowedCapabilities list.
	// +kubebuilder:validation:Optional
	DefaultAddCapabilities []*string `json:"defaultAddCapabilities,omitempty" tf:"default_add_capabilities,omitempty"`

	// defaultAllowPrivilegeEscalation controls the default setting for whether a process can gain more privileges than its parent process.
	// +kubebuilder:validation:Optional
	DefaultAllowPrivilegeEscalation *bool `json:"defaultAllowPrivilegeEscalation,omitempty" tf:"default_allow_privilege_escalation,omitempty"`

	// Pod Security Policy template policy description
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// forbiddenSysctls is a list of explicitly forbidden sysctls, defaults to none. Each entry is either a plain sysctl name or ends in "*" in which case it is considered as a prefix of forbidden sysctls. Single * means all sysctls are forbidden.
	// +kubebuilder:validation:Optional
	ForbiddenSysctls []*string `json:"forbiddenSysctls,omitempty" tf:"forbidden_sysctls,omitempty"`

	// fsGroup is the strategy that will dictate what fs group is used by the SecurityContext.
	// +kubebuilder:validation:Optional
	FsGroup []FsGroupParameters `json:"fsGroup,omitempty" tf:"fs_group,omitempty"`

	// hostIPC determines if the policy allows the use of HostIPC in the pod spec.
	// +kubebuilder:validation:Optional
	HostIpc *bool `json:"hostIpc,omitempty" tf:"host_ipc,omitempty"`

	// hostNetwork determines if the policy allows the use of HostNetwork in the pod spec.
	// +kubebuilder:validation:Optional
	HostNetwork *bool `json:"hostNetwork,omitempty" tf:"host_network,omitempty"`

	// hostPID determines if the policy allows the use of HostPID in the pod spec.
	// +kubebuilder:validation:Optional
	HostPid *bool `json:"hostPid,omitempty" tf:"host_pid,omitempty"`

	// hostPorts determines which host port ranges are allowed to be exposed.
	// +kubebuilder:validation:Optional
	HostPort []HostPortParameters `json:"hostPort,omitempty" tf:"host_port,omitempty"`

	// Labels of the resource
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// privileged determines if a pod can request to be run as privileged.
	// +kubebuilder:validation:Optional
	Privileged *bool `json:"privileged,omitempty" tf:"privileged,omitempty"`

	// readOnlyRootFilesystem when set to true will force containers to run with a read only root file system.  If the container specifically requests to run with a non-read only root file system the PSP should deny the pod. If set to false the container may run with a read only root file system if it wishes but it will not be forced to.
	// +kubebuilder:validation:Optional
	ReadOnlyRootFilesystem *bool `json:"readOnlyRootFilesystem,omitempty" tf:"read_only_root_filesystem,omitempty"`

	// requiredDropCapabilities are the capabilities that will be dropped from the container.  These are required to be dropped and cannot be added.
	// +kubebuilder:validation:Optional
	RequiredDropCapabilities []*string `json:"requiredDropCapabilities,omitempty" tf:"required_drop_capabilities,omitempty"`

	// RunAsGroup is the strategy that will dictate the allowable RunAsGroup values that may be set. If this field is omitted, the pod's RunAsGroup can take any value. This field requires the RunAsGroup feature gate to be enabled.
	// +kubebuilder:validation:Optional
	RunAsGroup []RunAsGroupParameters `json:"runAsGroup,omitempty" tf:"run_as_group,omitempty"`

	// runAsUser is the strategy that will dictate the allowable RunAsUser values that may be set.
	// +kubebuilder:validation:Optional
	RunAsUser []RunAsUserParameters `json:"runAsUser,omitempty" tf:"run_as_user,omitempty"`

	// runtimeClass is the strategy that will dictate the allowable RuntimeClasses for a pod. If this field is omitted, the pod's runtimeClassName field is unrestricted. Enforcement of this field depends on the RuntimeClass feature gate being enabled.
	// +kubebuilder:validation:Optional
	RuntimeClass []RuntimeClassParameters `json:"runtimeClass,omitempty" tf:"runtime_class,omitempty"`

	// seLinux is the strategy that will dictate the allowable labels that may be set.
	// +kubebuilder:validation:Optional
	SeLinux []SeLinuxParameters `json:"seLinux,omitempty" tf:"se_linux,omitempty"`

	// supplementalGroups is the strategy that will dictate what supplemental groups are used by the SecurityContext.
	// +kubebuilder:validation:Optional
	SupplementalGroup []SupplementalGroupParameters `json:"supplementalGroup,omitempty" tf:"supplemental_group,omitempty"`

	// volumes is a white list of allowed volume plugins. Empty indicates that no volumes may be used. To allow all volumes you may use '*'
	// +kubebuilder:validation:Optional
	Volumes []*string `json:"volumes,omitempty" tf:"volumes,omitempty"`
}

type SupplementalGroupObservation struct {
}

type SupplementalGroupParameters struct {

	// ranges are the allowed ranges of supplemental groups.  If you would like to force a single supplemental group then supply a single range with the same start and end. Required for MustRunAs.
	// +kubebuilder:validation:Optional
	Range []SupplementalGroupRangeParameters `json:"range,omitempty" tf:"range,omitempty"`

	// rule is the strategy that will dictate what supplemental groups is used in the SecurityContext.
	// +kubebuilder:validation:Optional
	Rule *string `json:"rule,omitempty" tf:"rule,omitempty"`
}

type SupplementalGroupRangeObservation struct {
}

type SupplementalGroupRangeParameters struct {

	// max is the end of the range, inclusive.
	// +kubebuilder:validation:Required
	Max *float64 `json:"max" tf:"max,omitempty"`

	// min is the start of the range, inclusive.
	// +kubebuilder:validation:Required
	Min *float64 `json:"min" tf:"min,omitempty"`
}

// SecurityPolicyTemplateSpec defines the desired state of SecurityPolicyTemplate
type SecurityPolicyTemplateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecurityPolicyTemplateParameters `json:"forProvider"`
}

// SecurityPolicyTemplateStatus defines the observed state of SecurityPolicyTemplate.
type SecurityPolicyTemplateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecurityPolicyTemplateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityPolicyTemplate is the Schema for the SecurityPolicyTemplates API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,rancherjet}
type SecurityPolicyTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecurityPolicyTemplateSpec   `json:"spec"`
	Status            SecurityPolicyTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityPolicyTemplateList contains a list of SecurityPolicyTemplates
type SecurityPolicyTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityPolicyTemplate `json:"items"`
}

// Repository type metadata.
var (
	SecurityPolicyTemplate_Kind             = "SecurityPolicyTemplate"
	SecurityPolicyTemplate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecurityPolicyTemplate_Kind}.String()
	SecurityPolicyTemplate_KindAPIVersion   = SecurityPolicyTemplate_Kind + "." + CRDGroupVersion.String()
	SecurityPolicyTemplate_GroupVersionKind = CRDGroupVersion.WithKind(SecurityPolicyTemplate_Kind)
)

func init() {
	SchemeBuilder.Register(&SecurityPolicyTemplate{}, &SecurityPolicyTemplateList{})
}
