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

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/terrajet/pkg/controller"

	activedirectory "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/authconfig/activedirectory"
	adfs "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/authconfig/adfs"
	azuread "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/authconfig/azuread"
	freeipa "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/authconfig/freeipa"
	github "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/authconfig/github"
	keycloak "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/authconfig/keycloak"
	okta "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/authconfig/okta"
	openldap "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/authconfig/openldap"
	ping "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/authconfig/ping"
	alertgroup "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/cluster/alertgroup"
	alertrule "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/cluster/alertrule"
	cluster "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/cluster/cluster"
	driver "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/cluster/driver"
	logging "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/cluster/logging"
	roletemplatebinding "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/cluster/roletemplatebinding"
	sync "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/cluster/sync"
	template "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/cluster/template"
	clusterclusterv2 "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/clusterv2/cluster"
	dns "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/global/dns"
	dnsprovider "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/global/dnsprovider"
	role "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/global/role"
	rolebinding "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/global/rolebinding"
	ranchernamespace "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/namespace/ranchernamespace"
	alertgroupproject "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/project/alertgroup"
	alertruleproject "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/project/alertrule"
	loggingproject "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/project/logging"
	project "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/project/project"
	roletemplatebindingproject "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/project/roletemplatebinding"
	providerconfig "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/providerconfig"
	cloudcredential "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/rancher/cloudcredential"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		activedirectory.Setup,
		adfs.Setup,
		azuread.Setup,
		freeipa.Setup,
		github.Setup,
		keycloak.Setup,
		okta.Setup,
		openldap.Setup,
		ping.Setup,
		alertgroup.Setup,
		alertrule.Setup,
		cluster.Setup,
		driver.Setup,
		logging.Setup,
		roletemplatebinding.Setup,
		sync.Setup,
		template.Setup,
		clusterclusterv2.Setup,
		dns.Setup,
		dnsprovider.Setup,
		role.Setup,
		rolebinding.Setup,
		ranchernamespace.Setup,
		alertgroupproject.Setup,
		alertruleproject.Setup,
		loggingproject.Setup,
		project.Setup,
		roletemplatebindingproject.Setup,
		providerconfig.Setup,
		cloudcredential.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
