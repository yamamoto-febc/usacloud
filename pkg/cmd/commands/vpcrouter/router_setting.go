// Copyright 2017-2020 The Usacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vpcrouter

import (
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/util"
)

type routerSetting struct {
	VRID                      int
	InternetConnectionEnabled bool

	StaticNATData string                        `cli:"static-nat" mapconv:"-"`
	StaticNAT     []*sacloud.VPCRouterStaticNAT `cli:"-"`

	PortForwardingData string                             `cli:"port-forwarding" mapconv:"-"`
	PortForwarding     []*sacloud.VPCRouterPortForwarding `cli:"-"`

	FirewallData string                       `cli:"firewall" mapconv:"-"`
	Firewall     []*sacloud.VPCRouterFirewall `cli:"-"`

	DHCPServerData string                         `cli:"dhcp-server" mapconv:"-"`
	DHCPServer     []*sacloud.VPCRouterDHCPServer `cli:"-"`

	DHCPStaticMappingData string                                `cli:"dhcp-static-mapping" mapconv:"-"`
	DHCPStaticMapping     []*sacloud.VPCRouterDHCPStaticMapping `cli:"-"`

	PPTPServer sacloud.VPCRouterPPTPServer `cli:"pptp" mapconv:",omitempty"`

	L2TPIPsecServer sacloud.VPCRouterL2TPIPsecServer `cli:"l2tp" mapconv:",omitempty"`

	RemoteAccessUsersData string                               `cli:"users" mapconv:"-"`
	RemoteAccessUsers     []*sacloud.VPCRouterRemoteAccessUser `cli:"-"`

	SiteToSiteIPsecVPNData string                                 `cli:"site-to-site-vpn" mapconv:"-"`
	SiteToSiteIPsecVPN     []*sacloud.VPCRouterSiteToSiteIPsecVPN `cli:"-"`

	StaticRouteData string                          `cli:"static-route" mapconv:"-"`
	StaticRoute     []*sacloud.VPCRouterStaticRoute `cli:"-"`

	SyslogHost string
}

func (r *routerSetting) Customize(_ cli.Context) error {
	if r.StaticNATData != "" {
		var staticNat []*sacloud.VPCRouterStaticNAT
		if err := util.MarshalJSONFromPathOrContent(r.StaticNATData, &staticNat); err != nil {
			return err
		}
		r.StaticNAT = append(r.StaticNAT, staticNat...)
	}

	if r.PortForwardingData != "" {
		var portForwarding []*sacloud.VPCRouterPortForwarding
		if err := util.MarshalJSONFromPathOrContent(r.PortForwardingData, &portForwarding); err != nil {
			return err
		}
		r.PortForwarding = append(r.PortForwarding, portForwarding...)
	}

	if r.FirewallData != "" {
		var firewall []*sacloud.VPCRouterFirewall
		if err := util.MarshalJSONFromPathOrContent(r.FirewallData, &firewall); err != nil {
			return err
		}
		r.Firewall = append(r.Firewall, firewall...)
	}

	if r.DHCPServerData != "" {
		var dhcpServer []*sacloud.VPCRouterDHCPServer
		if err := util.MarshalJSONFromPathOrContent(r.DHCPServerData, &dhcpServer); err != nil {
			return err
		}
		r.DHCPServer = append(r.DHCPServer, dhcpServer...)
	}

	if r.DHCPStaticMappingData != "" {
		var dhcpStaticMapping []*sacloud.VPCRouterDHCPStaticMapping
		if err := util.MarshalJSONFromPathOrContent(r.DHCPStaticMappingData, &dhcpStaticMapping); err != nil {
			return err
		}
		r.DHCPStaticMapping = append(r.DHCPStaticMapping, dhcpStaticMapping...)
	}

	if r.RemoteAccessUsersData != "" {
		var users []*sacloud.VPCRouterRemoteAccessUser
		if err := util.MarshalJSONFromPathOrContent(r.RemoteAccessUsersData, &users); err != nil {
			return err
		}
		r.RemoteAccessUsers = append(r.RemoteAccessUsers, users...)
	}

	if r.SiteToSiteIPsecVPNData != "" {
		var s2s []*sacloud.VPCRouterSiteToSiteIPsecVPN
		if err := util.MarshalJSONFromPathOrContent(r.SiteToSiteIPsecVPNData, &s2s); err != nil {
			return err
		}
		r.SiteToSiteIPsecVPN = append(r.SiteToSiteIPsecVPN, s2s...)
	}

	if r.StaticRouteData != "" {
		var staticRoutes []*sacloud.VPCRouterStaticRoute
		if err := util.MarshalJSONFromPathOrContent(r.StaticRouteData, &staticRoutes); err != nil {
			return err
		}
		r.StaticRoute = append(r.StaticRoute, staticRoutes...)
	}

	return nil
}
