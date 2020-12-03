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

package disk

import (
	"github.com/sacloud/usacloud/pkg/cmd/core"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func (p *editParameter) CleanupEmptyValue(fs *pflag.FlagSet) {

}

func (p *editParameter) buildFlags(fs *pflag.FlagSet) {

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
	fs.StringVarP(&p.EditDisk.HostName, "host-name", "", p.EditDisk.HostName, "")
	fs.StringVarP(&p.EditDisk.Password, "password", "", p.EditDisk.Password, "")
	fs.BoolVarP(&p.EditDisk.DisablePWAuth, "disable-pw-auth", "", p.EditDisk.DisablePWAuth, "")
	fs.BoolVarP(&p.EditDisk.EnableDHCP, "enable-dhcp", "", p.EditDisk.EnableDHCP, "")
	fs.BoolVarP(&p.EditDisk.ChangePartitionUUID, "change-partition-uuid", "", p.EditDisk.ChangePartitionUUID, "")
	fs.StringVarP(&p.EditDisk.IPAddress, "ip-address", "", p.EditDisk.IPAddress, "")
	fs.IntVarP(&p.EditDisk.NetworkMaskLen, "network-mask-len", "", p.EditDisk.NetworkMaskLen, "")
	fs.StringVarP(&p.EditDisk.DefaultRoute, "default-route", "", p.EditDisk.DefaultRoute, "")
	fs.StringSliceVarP(&p.EditDisk.SSHKeys, "ssh-keys", "", p.EditDisk.SSHKeys, "")
	fs.VarP(core.NewIDSliceFlag(&p.EditDisk.SSHKeyIDs, &p.EditDisk.SSHKeyIDs), "ssh-key-ids", "", "")
	fs.BoolVarP(&p.EditDisk.IsSSHKeysEphemeral, "make-ssh-keys-ephemeral", "", p.EditDisk.IsSSHKeysEphemeral, "")
	fs.VarP(core.NewIDSliceFlag(&p.EditDisk.NoteIDs, &p.EditDisk.NoteIDs), "note-ids", "", "")
	fs.StringVarP(&p.EditDisk.NotesData, "notes", "", p.EditDisk.NotesData, "")
	fs.BoolVarP(&p.EditDisk.IsNotesEphemeral, "make-notes-ephemeral", "", p.EditDisk.IsNotesEphemeral, "")
	fs.BoolVarP(&p.NoWait, "no-wait", "", p.NoWait, "")
	fs.SetNormalizeFunc(p.normalizeFlagName)
}

func (p *editParameter) normalizeFlagName(_ *pflag.FlagSet, name string) pflag.NormalizedName {
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

func (p *editParameter) buildFlagsUsage(cmd *cobra.Command) {
	var sets []*core.FlagSet
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("disk", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("zone"))
		fs.AddFlag(cmd.LocalFlags().Lookup("host-name"))
		fs.AddFlag(cmd.LocalFlags().Lookup("password"))
		fs.AddFlag(cmd.LocalFlags().Lookup("disable-pw-auth"))
		fs.AddFlag(cmd.LocalFlags().Lookup("enable-dhcp"))
		fs.AddFlag(cmd.LocalFlags().Lookup("change-partition-uuid"))
		fs.AddFlag(cmd.LocalFlags().Lookup("ip-address"))
		fs.AddFlag(cmd.LocalFlags().Lookup("network-mask-len"))
		fs.AddFlag(cmd.LocalFlags().Lookup("default-route"))
		fs.AddFlag(cmd.LocalFlags().Lookup("ssh-keys"))
		fs.AddFlag(cmd.LocalFlags().Lookup("ssh-key-ids"))
		fs.AddFlag(cmd.LocalFlags().Lookup("make-ssh-keys-ephemeral"))
		fs.AddFlag(cmd.LocalFlags().Lookup("note-ids"))
		fs.AddFlag(cmd.LocalFlags().Lookup("notes"))
		fs.AddFlag(cmd.LocalFlags().Lookup("make-notes-ephemeral"))
		fs.AddFlag(cmd.LocalFlags().Lookup("no-wait"))
		sets = append(sets, &core.FlagSet{
			Title: "Disk options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("Input", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("parameters"))
		fs.AddFlag(cmd.LocalFlags().Lookup("generate-skeleton"))
		fs.AddFlag(cmd.LocalFlags().Lookup("assumeyes"))
		sets = append(sets, &core.FlagSet{
			Title: "Input options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("output", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("output-type"))
		fs.AddFlag(cmd.LocalFlags().Lookup("quiet"))
		fs.AddFlag(cmd.LocalFlags().Lookup("format"))
		fs.AddFlag(cmd.LocalFlags().Lookup("format-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("query"))
		fs.AddFlag(cmd.LocalFlags().Lookup("query-file"))
		sets = append(sets, &core.FlagSet{
			Title: "Output options",
			Flags: fs,
		})
	}

	core.BuildFlagsUsage(cmd, sets)
}

func (p *editParameter) setCompletionFunc(cmd *cobra.Command) {

}

func (p *editParameter) SetupCobraCommandFlags(cmd *cobra.Command) {
	p.buildFlags(cmd.Flags())
	p.buildFlagsUsage(cmd)
	p.setCompletionFunc(cmd)
}
