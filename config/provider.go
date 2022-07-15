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
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	tjconfig "github.com/crossplane/terrajet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/crossplane-contrib/provider-jet-rancher/config/cloudcredentials"
	"github.com/crossplane-contrib/provider-jet-rancher/config/project"
)

const (
	resourcePrefix = "rancher2"
	modulePath     = "github.com/crossplane-contrib/provider-jet-rancher"
)

//go:embed schema.json
var providerSchema string

var skipList = []string{
	"rancher2_secret",
}

var includeList = []string{
	"rancher2_cloud_credential",
	"rancher2_cluster",
	"rancher2_project",
	"rancher2_auth_config_",
	"rancher2_global_",
	"rancher2_namespace",
}

// GetProvider returns provider configuration
func GetProvider() *tjconfig.Provider {
	pc := tjconfig.NewProviderWithSchema([]byte(providerSchema), resourcePrefix, modulePath,
		tjconfig.WithShortName("rancherjet"),
		tjconfig.WithRootGroup("rancher.jet.crossplane.io"),
		tjconfig.WithIncludeList(includeList),
		tjconfig.WithSkipList(skipList),
		tjconfig.WithDefaultResourceFn(DefaultResource(
			VersionOverride(),
			GroupKindOverrides(),
			KindOverrides(),
			KnownReferencers(),
		)),
	)

	for _, configure := range []func(provider *tjconfig.Provider){
		// add custom config functions
		cloudcredentials.Configure,
		project.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// DefaultResource returns a DefaultResoruceFn that makes sure the original
// DefaultResource call is made with given options here.
func DefaultResource(opts ...tjconfig.ResourceOption) tjconfig.DefaultResourceFn {
	return func(name string, terraformResource *schema.Resource, orgOpts ...tjconfig.ResourceOption) *tjconfig.Resource {
		return tjconfig.DefaultResource(name, terraformResource, append(orgOpts, opts...)...)
	}
}
