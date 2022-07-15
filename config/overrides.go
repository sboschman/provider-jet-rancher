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

package config

import (
	"strings"

	tjconfig "github.com/crossplane/terrajet/pkg/config"
	"github.com/crossplane/terrajet/pkg/types/name"

	"github.com/crossplane-contrib/provider-jet-rancher/config/common"
)

// GroupKindCalculator returns the correct group and kind name for given TF
// resource.
type GroupKindCalculator func(resource string) (string, string)

// GroupMap contains all overrides we'd like to make to the default group search.
var GroupMap = map[string]GroupKindCalculator{
	"rancher2_cloud_credential": ReplaceGroupWords("", 0),

	"rancher2_cluster":    ReplaceGroupWords("cluster", 0),
	"rancher2_cluster_v2": ReplaceGroupWords("clusterv2", 0),

	"rancher2_project": ReplaceGroupWords("project", 0),

	"rancher2_namespace": ReplaceGroupWords("namespace", 0),

	"rancher2_auth_config_activedirectory": ReplaceGroupWords("authconfig", 2),
	"rancher2_auth_config_adfs":            ReplaceGroupWords("authconfig", 2),
	"rancher2_auth_config_azuread":         ReplaceGroupWords("authconfig", 2),
	"rancher2_auth_config_freeipa":         ReplaceGroupWords("authconfig", 2),
	"rancher2_auth_config_github":          ReplaceGroupWords("authconfig", 2),
	"rancher2_auth_config_keycloak":        ReplaceGroupWords("authconfig", 2),
	"rancher2_auth_config_okta":            ReplaceGroupWords("authconfig", 2),
	"rancher2_auth_config_openldap":        ReplaceGroupWords("authconfig", 2),
	"rancher2_auth_config_ping":            ReplaceGroupWords("authconfig", 2),
}

// ReplaceGroupWords uses given group as the group of the resource and removes
// a number of words in resource name before calculating the kind of the resource.
func ReplaceGroupWords(group string, count int) GroupKindCalculator {
	return func(resource string) (string, string) {
		words := strings.Split(strings.TrimPrefix(resource, "rancher2_"), "_")
		snakeKind := strings.Join(words[count:], "_")
		return group, name.NewFromSnake(snakeKind).Camel
	}
}

func VersionOverride() tjconfig.ResourceOption {
	return func(r *tjconfig.Resource) {
		r.Version = common.VersionV1alpha2
	}
}

// KindMap contains kind string overrides.
var KindMap = map[string]string{
	"rancher2_cloud_credential": "CloudCredential",
	"rancher2_namespace":        "RancherNamespace",
	"rancher2_cluster_v2":       "Cluster",
}

// GroupKindOverrides overrides the group and kind of the resource if it matches
// any entry in the GroupMap.
func GroupKindOverrides() tjconfig.ResourceOption {
	return func(r *tjconfig.Resource) {
		if f, ok := GroupMap[r.Name]; ok {
			r.ShortGroup, r.Kind = f(r.Name)
		}
	}
}

// KindOverrides overrides the kind of the resources given in KindMap.
func KindOverrides() tjconfig.ResourceOption {
	return func(r *tjconfig.Resource) {
		if k, ok := KindMap[r.Name]; ok {
			r.Kind = k
		}
	}
}

// KnownReferencers adds referencers for fields that are known and common among
// more than a few resources.
func KnownReferencers() tjconfig.ResourceOption { //nolint:gocyclo
	return func(r *tjconfig.Resource) {
		for k, s := range r.TerraformResource.Schema {
			// We shouldn't add referencers for status fields and sensitive fields
			// since they already have secret referencer.
			if (s.Computed && !s.Optional) || s.Sensitive {
				continue
			}
			switch k {
			case "cluster_id":
				r.References["cluster_id"] = tjconfig.Reference{
					Type: "github.com/crossplane-contrib/provider-jet-rancher/apis/cluster/v1alpha2.Cluster",
				}
			case "project_id":
				r.References["project_id"] = tjconfig.Reference{
					Type: "github.com/crossplane-contrib/provider-jet-rancher/apis/project/v1alpha2.Project",
				}
			}
		}
	}
}
