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

// Code generated by 'github.com/sacloud/usacloud/tools/gen-commands'; DO NOT EDIT

package mobilegateway

import (
	"github.com/sacloud/libsacloud/v2/sacloud/pointer"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/cmd/core"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func (p *updateParameter) CleanupEmptyValue(fs *pflag.FlagSet) {
	if !fs.Changed("name") {
		p.Name = nil
	}
	if !fs.Changed("description") {
		p.Description = nil
	}
	if !fs.Changed("tags") {
		p.Tags = nil
	}
	if !fs.Changed("icon-id") {
		p.IconID = nil
	}
	if !fs.Changed("private-interface-switch-id") {
		p.PrivateInterface.SwitchID = nil
	}
	if !fs.Changed("private-interface-ip-address") {
		p.PrivateInterface.IPAddress = nil
	}
	if !fs.Changed("private-interface-network-mask-len") {
		p.PrivateInterface.NetworkMaskLen = nil
	}
	if !fs.Changed("internet-connection-enabled") {
		p.InternetConnectionEnabled = nil
	}
	if !fs.Changed("inter-device-communication-enabled") {
		p.InterDeviceCommunicationEnabled = nil
	}
	if !fs.Changed("sims") {
		p.SIMsData = nil
	}
	if !fs.Changed("sim-routes") {
		p.SIMRoutesData = nil
	}
	if !fs.Changed("static-routes") {
		p.StaticRoutesData = nil
	}
	if !fs.Changed("dns1") {
		p.DNS.DNS1 = nil
	}
	if !fs.Changed("dns2") {
		p.DNS.DNS2 = nil
	}
	if !fs.Changed("traffic-config-traffic-quota-in-mb") {
		p.TrafficConfig.TrafficQuotaInMB = nil
	}
	if !fs.Changed("traffic-config-band-width-limit-in-kbps") {
		p.TrafficConfig.BandWidthLimitInKbps = nil
	}
	if !fs.Changed("traffic-config-email-notify-enabled") {
		p.TrafficConfig.EmailNotifyEnabled = nil
	}
	if !fs.Changed("traffic-config-slack-notify-enabled") {
		p.TrafficConfig.SlackNotifyEnabled = nil
	}
	if !fs.Changed("traffic-config-slack-notify-webhooks-url") {
		p.TrafficConfig.SlackNotifyWebhooksURL = nil
	}
	if !fs.Changed("traffic-config-auto-traffic-shaping") {
		p.TrafficConfig.AutoTrafficShaping = nil
	}
}

func (p *updateParameter) buildFlags(fs *pflag.FlagSet) {
	if p.Name == nil {
		p.Name = pointer.NewString("")
	}
	if p.Description == nil {
		p.Description = pointer.NewString("")
	}
	if p.Tags == nil {
		p.Tags = pointer.NewStringSlice([]string{})
	}
	if p.IconID == nil {
		p.IconID = pointer.NewID(types.ID(0))
	}
	if p.PrivateInterface.SwitchID == nil {
		p.PrivateInterface.SwitchID = pointer.NewID(types.ID(0))
	}
	if p.PrivateInterface.IPAddress == nil {
		p.PrivateInterface.IPAddress = pointer.NewString("")
	}
	if p.PrivateInterface.NetworkMaskLen == nil {
		p.PrivateInterface.NetworkMaskLen = pointer.NewInt(0)
	}
	if p.InternetConnectionEnabled == nil {
		p.InternetConnectionEnabled = pointer.NewBool(false)
	}
	if p.InterDeviceCommunicationEnabled == nil {
		p.InterDeviceCommunicationEnabled = pointer.NewBool(false)
	}
	if p.SIMsData == nil {
		p.SIMsData = pointer.NewString("")
	}
	if p.SIMRoutesData == nil {
		p.SIMRoutesData = pointer.NewString("")
	}
	if p.StaticRoutesData == nil {
		p.StaticRoutesData = pointer.NewString("")
	}
	if p.DNS.DNS1 == nil {
		p.DNS.DNS1 = pointer.NewString("")
	}
	if p.DNS.DNS2 == nil {
		p.DNS.DNS2 = pointer.NewString("")
	}
	if p.TrafficConfig.TrafficQuotaInMB == nil {
		p.TrafficConfig.TrafficQuotaInMB = pointer.NewInt(0)
	}
	if p.TrafficConfig.BandWidthLimitInKbps == nil {
		p.TrafficConfig.BandWidthLimitInKbps = pointer.NewInt(0)
	}
	if p.TrafficConfig.EmailNotifyEnabled == nil {
		p.TrafficConfig.EmailNotifyEnabled = pointer.NewBool(false)
	}
	if p.TrafficConfig.SlackNotifyEnabled == nil {
		p.TrafficConfig.SlackNotifyEnabled = pointer.NewBool(false)
	}
	if p.TrafficConfig.SlackNotifyWebhooksURL == nil {
		p.TrafficConfig.SlackNotifyWebhooksURL = pointer.NewString("")
	}
	if p.TrafficConfig.AutoTrafficShaping == nil {
		p.TrafficConfig.AutoTrafficShaping = pointer.NewBool(false)
	}
	fs.StringVarP(&p.Zone, "zone", "", p.Zone, "")
	fs.StringVarP(&p.Parameters, "parameters", "", p.Parameters, "Input parameters in JSON format")
	fs.BoolVarP(&p.GenerateSkeleton, "generate-skeleton", "", p.GenerateSkeleton, "Output skeleton of parameters with JSON format (aliases: --skeleton)")
	fs.BoolVarP(&p.AssumeYes, "assumeyes", "y", p.AssumeYes, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&p.OutputType, "output-type", "o", p.OutputType, "Output format: one of the following [table/json/yaml] (aliases: --out)")
	fs.BoolVarP(&p.Quiet, "quiet", "q", p.Quiet, "Output IDs only")
	fs.StringVarP(&p.Format, "format", "", p.Format, "Output format in Go templates (aliases: --fmt)")
	fs.StringVarP(&p.FormatFile, "format-file", "", p.FormatFile, "Output format in Go templates(from file)")
	fs.StringVarP(&p.Query, "query", "", p.Query, "JMESPath query")
	fs.StringVarP(&p.QueryFile, "query-file", "", p.QueryFile, "JMESPath query(from file)")
	fs.StringVarP(p.Name, "name", "", "", "")
	fs.StringVarP(p.Description, "description", "", "", "")
	fs.StringSliceVarP(p.Tags, "tags", "", nil, "")
	fs.VarP(core.NewIDFlag(p.IconID, p.IconID), "icon-id", "", "")
	fs.VarP(core.NewIDFlag(p.PrivateInterface.SwitchID, p.PrivateInterface.SwitchID), "private-interface-switch-id", "", "")
	fs.StringVarP(p.PrivateInterface.IPAddress, "private-interface-ip-address", "", "", "")
	fs.IntVarP(p.PrivateInterface.NetworkMaskLen, "private-interface-network-mask-len", "", 0, "")
	fs.BoolVarP(p.InternetConnectionEnabled, "internet-connection-enabled", "", false, "")
	fs.BoolVarP(p.InterDeviceCommunicationEnabled, "inter-device-communication-enabled", "", false, "")
	fs.StringVarP(p.SIMsData, "sims", "", "", "")
	fs.StringVarP(p.SIMRoutesData, "sim-routes", "", "", "")
	fs.StringVarP(p.StaticRoutesData, "static-routes", "", "", "")
	fs.StringVarP(p.DNS.DNS1, "dns1", "", "", "")
	fs.StringVarP(p.DNS.DNS2, "dns2", "", "", "")
	fs.IntVarP(p.TrafficConfig.TrafficQuotaInMB, "traffic-config-traffic-quota-in-mb", "", 0, "")
	fs.IntVarP(p.TrafficConfig.BandWidthLimitInKbps, "traffic-config-band-width-limit-in-kbps", "", 0, "")
	fs.BoolVarP(p.TrafficConfig.EmailNotifyEnabled, "traffic-config-email-notify-enabled", "", false, "")
	fs.BoolVarP(p.TrafficConfig.SlackNotifyEnabled, "traffic-config-slack-notify-enabled", "", false, "")
	fs.StringVarP(p.TrafficConfig.SlackNotifyWebhooksURL, "traffic-config-slack-notify-webhooks-url", "", "", "")
	fs.BoolVarP(p.TrafficConfig.AutoTrafficShaping, "traffic-config-auto-traffic-shaping", "", false, "")
	fs.BoolVarP(&p.NoWait, "no-wait", "", p.NoWait, "")
	fs.SetNormalizeFunc(p.normalizeFlagName)
}

func (p *updateParameter) normalizeFlagName(_ *pflag.FlagSet, name string) pflag.NormalizedName {
	switch name {
	case "skeleton":
		name = "generate-skeleton"
	case "out":
		name = "output-type"
	case "fmt":
		name = "format"
	}
	return pflag.NormalizedName(name)
}

func (p *updateParameter) buildFlagsUsage(cmd *cobra.Command) {
	var sets []*core.FlagSet
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("common", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("name"))
		fs.AddFlag(cmd.LocalFlags().Lookup("description"))
		fs.AddFlag(cmd.LocalFlags().Lookup("tags"))
		fs.AddFlag(cmd.LocalFlags().Lookup("icon-id"))
		sets = append(sets, &core.FlagSet{
			Title: "Common options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("mobile-gateway", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("dns1"))
		fs.AddFlag(cmd.LocalFlags().Lookup("dns2"))
		fs.AddFlag(cmd.LocalFlags().Lookup("inter-device-communication-enabled"))
		fs.AddFlag(cmd.LocalFlags().Lookup("internet-connection-enabled"))
		fs.AddFlag(cmd.LocalFlags().Lookup("private-interface-ip-address"))
		fs.AddFlag(cmd.LocalFlags().Lookup("private-interface-network-mask-len"))
		fs.AddFlag(cmd.LocalFlags().Lookup("private-interface-switch-id"))
		fs.AddFlag(cmd.LocalFlags().Lookup("sim-routes"))
		fs.AddFlag(cmd.LocalFlags().Lookup("sims"))
		fs.AddFlag(cmd.LocalFlags().Lookup("static-routes"))
		fs.AddFlag(cmd.LocalFlags().Lookup("traffic-config-auto-traffic-shaping"))
		fs.AddFlag(cmd.LocalFlags().Lookup("traffic-config-band-width-limit-in-kbps"))
		fs.AddFlag(cmd.LocalFlags().Lookup("traffic-config-email-notify-enabled"))
		fs.AddFlag(cmd.LocalFlags().Lookup("traffic-config-slack-notify-enabled"))
		fs.AddFlag(cmd.LocalFlags().Lookup("traffic-config-slack-notify-webhooks-url"))
		fs.AddFlag(cmd.LocalFlags().Lookup("traffic-config-traffic-quota-in-mb"))
		sets = append(sets, &core.FlagSet{
			Title: "Mobile-Gateway-specific options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("zone", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("zone"))
		sets = append(sets, &core.FlagSet{
			Title: "Zone options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("wait", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("no-wait"))
		sets = append(sets, &core.FlagSet{
			Title: "Wait options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("input", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("assumeyes"))
		fs.AddFlag(cmd.LocalFlags().Lookup("generate-skeleton"))
		fs.AddFlag(cmd.LocalFlags().Lookup("parameters"))
		sets = append(sets, &core.FlagSet{
			Title: "Input options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("output", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("format"))
		fs.AddFlag(cmd.LocalFlags().Lookup("format-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("output-type"))
		fs.AddFlag(cmd.LocalFlags().Lookup("query"))
		fs.AddFlag(cmd.LocalFlags().Lookup("query-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("quiet"))
		sets = append(sets, &core.FlagSet{
			Title: "Output options",
			Flags: fs,
		})
	}

	core.BuildFlagsUsage(cmd, sets)
}

func (p *updateParameter) setCompletionFunc(cmd *cobra.Command) {

}

func (p *updateParameter) SetupCobraCommandFlags(cmd *cobra.Command) {
	p.buildFlags(cmd.Flags())
	p.buildFlagsUsage(cmd)
	p.setCompletionFunc(cmd)
}
