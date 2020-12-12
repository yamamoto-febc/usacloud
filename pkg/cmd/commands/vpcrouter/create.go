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
	"github.com/sacloud/libsacloud/v2/helper/service/vpcrouter"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/cmd/core"
	"github.com/sacloud/usacloud/pkg/cmd/examples"
	"github.com/sacloud/usacloud/pkg/util"
)

var createCommand = &core.Command{
	Name:     "create",
	Category: "basic",
	Order:    20,

	ColumnDefs: defaultColumnDefs,

	ParameterInitializer: func() interface{} {
		return newCreateParameter()
	},
}

type createParameter struct {
	cflag.ZoneParameter    `cli:",squash" mapconv:",squash"`
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	cflag.NameParameter   `cli:",squash" mapconv:",squash"`
	cflag.DescParameter   `cli:",squash" mapconv:",squash"`
	cflag.TagsParameter   `cli:",squash" mapconv:",squash"`
	cflag.IconIDParameter `cli:",squash" mapconv:",squash"`

	Plan string `cli:"plan,options=vpc_router_plan_premium,category=plan" mapconv:"PlanID,filters=vpc_router_plan_premium_to_value" validate:"required,vpc_router_plan_premium"`

	PublicNetworkInterface vpcrouter.PremiumNICSetting `cli:",category=network,order=10" mapconv:"NICSetting,omitempty"`

	PrivateNetworkInterfacesData string                                   `cli:"private-network-interfaces,category=network,order=20" mapconv:"-" json:"-"`
	PrivateNetworkInterfaces     []*vpcrouter.AdditionalPremiumNICSetting `cli:"-" mapconv:"AdditionalNICSettings"`

	RouterSetting routerSetting `cli:",squash" mapconv:",recursive"`

	BootAfterCreate       bool
	cflag.NoWaitParameter `cli:",squash" mapconv:",squash"`
}

func newCreateParameter() *createParameter {
	return &createParameter{
		Plan: "premium",
	}
}

func init() {
	Resource.AddCommand(createCommand)
}

// Customize パラメータ変換処理
func (p *createParameter) Customize(ctx cli.Context) error {
	if p.PrivateNetworkInterfacesData != "" {
		var nics []*vpcrouter.AdditionalPremiumNICSetting
		if err := util.MarshalJSONFromPathOrContent(p.PrivateNetworkInterfacesData, &nics); err != nil {
			return err
		}
		p.PrivateNetworkInterfaces = append(p.PrivateNetworkInterfaces, nics...)
	}
	return p.RouterSetting.Customize(ctx)
}

func (p *createParameter) ExampleParameters(ctx cli.Context) interface{} {
	return &createParameter{
		ZoneParameter:   examples.Zones(ctx.Option().Zones),
		NameParameter:   examples.Name,
		DescParameter:   examples.Description,
		TagsParameter:   examples.Tags,
		IconIDParameter: examples.IconID,
		Plan:            examples.OptionsString("vpc_router_plan_premium"),
		PublicNetworkInterface: vpcrouter.PremiumNICSetting{
			SwitchID:         examples.ID,
			IPAddresses:      examples.IPAddresses,
			VirtualIPAddress: examples.VirtualIPAddress,
			IPAliases:        []string{"192.0.2.102"},
		},
		PrivateNetworkInterfaces: []*vpcrouter.AdditionalPremiumNICSetting{
			{
				SwitchID:         examples.ID,
				IPAddresses:      []string{"192.168.0.11", "192.168.0.12"},
				VirtualIPAddress: "192.168.0.1",
				NetworkMaskLen:   24,
				Index:            1,
			},
		},
		RouterSetting: routerSetting{
			VRID:                      1,
			InternetConnectionEnabled: true,
			StaticNAT: []*sacloud.VPCRouterStaticNAT{
				{
					GlobalAddress:  examples.VirtualIPAddress,
					PrivateAddress: "192.168.0.1",
					Description:    "example",
				},
			},
			PortForwarding: []*sacloud.VPCRouterPortForwarding{
				{
					Protocol:       types.EVPCRouterPortForwardingProtocol(examples.OptionsString("vpc_router_port_forwarding_protocol")),
					GlobalPort:     22,
					PrivateAddress: "192.168.0.11",
					PrivatePort:    22,
					Description:    "example",
				},
			},
			Firewall: []*sacloud.VPCRouterFirewall{
				{
					Send: []*sacloud.VPCRouterFirewallRule{
						{
							Protocol:           types.Protocol(examples.OptionsString("vpc_router_firewall_protocol")),
							SourceNetwork:      "192.0.2.1 | 192.0.2.0/24",
							SourcePort:         "1024 | 1024-2048",
							DestinationNetwork: "192.0.2.1 | 192.0.2.0/24",
							DestinationPort:    "1024 | 1024-2048",
							Action:             types.Action(examples.OptionsString("packetfilter_action")),
							Logging:            true,
							Description:        "example",
						},
					},
					Receive: []*sacloud.VPCRouterFirewallRule{
						{
							Protocol:           types.Protocol(examples.OptionsString("vpc_router_firewall_protocol")),
							SourceNetwork:      "192.0.2.1 | 192.0.2.0/24",
							SourcePort:         "1024 | 1024-2048",
							DestinationNetwork: "192.0.2.1 | 192.0.2.0/24",
							DestinationPort:    "1024 | 1024-2048",
							Action:             types.Action(examples.OptionsString("packetfilter_action")),
							Logging:            true,
							Description:        "example",
						},
					},
					Index: 0,
				},
			},
			DHCPServer: []*sacloud.VPCRouterDHCPServer{
				{
					Interface:  "eth1",
					RangeStart: "192.168.0.240",
					RangeStop:  "192.168.0.244",
					DNSServers: []string{"133.242.0.3", "133.242.0.4"},
				},
			},
			DHCPStaticMapping: []*sacloud.VPCRouterDHCPStaticMapping{
				{
					MACAddress: "9C:A3:BA:xx:xx:xx",
					IPAddress:  "192.168.0.245",
				},
			},
			PPTPServer: sacloud.VPCRouterPPTPServer{
				RangeStart: "192.168.0.246",
				RangeStop:  "192.168.0.249",
			},
			L2TPIPsecServer: sacloud.VPCRouterL2TPIPsecServer{
				RangeStart:      "192.168.0.250",
				RangeStop:       "192.168.0.254",
				PreSharedSecret: "presharedsecret",
			},
			RemoteAccessUsers: []*sacloud.VPCRouterRemoteAccessUser{
				{
					UserName: "username",
					Password: "password",
				},
			},
			SiteToSiteIPsecVPN: []*sacloud.VPCRouterSiteToSiteIPsecVPN{
				{
					Peer:            "192.0.2.1",
					PreSharedSecret: "presharedsecret",
					RemoteID:        "192.0.2.1",
					Routes:          []string{"10.0.0.0/8"},
					LocalPrefix:     []string{"192.168.0.0/24"},
				},
			},
			StaticRoute: []*sacloud.VPCRouterStaticRoute{
				{
					Prefix:  "172.16.0.0/16",
					NextHop: "192.168.0.21",
				},
			},
			SyslogHost: "192.168.0.1",
		},
		BootAfterCreate: true,
		NoWaitParameter: cflag.NoWaitParameter{
			NoWait: false,
		},
	}
}
