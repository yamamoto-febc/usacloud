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

//go:generate go run github.com/sacloud/usacloud/tools/gen-commands/
package cmd

import (
	"github.com/sacloud/usacloud/pkg/cmd/commands/archive"
	"github.com/sacloud/usacloud/pkg/cmd/commands/authstatus"
	"github.com/sacloud/usacloud/pkg/cmd/commands/autobackup"
	"github.com/sacloud/usacloud/pkg/cmd/commands/bill"
	"github.com/sacloud/usacloud/pkg/cmd/commands/bridge"
	"github.com/sacloud/usacloud/pkg/cmd/commands/cdrom"
	"github.com/sacloud/usacloud/pkg/cmd/commands/containerregistry"
	"github.com/sacloud/usacloud/pkg/cmd/commands/coupon"
	"github.com/sacloud/usacloud/pkg/cmd/commands/database"
	"github.com/sacloud/usacloud/pkg/cmd/commands/disk"
	"github.com/sacloud/usacloud/pkg/cmd/commands/diskplan"
	"github.com/sacloud/usacloud/pkg/cmd/commands/dns"
	"github.com/sacloud/usacloud/pkg/cmd/commands/esme"
	"github.com/sacloud/usacloud/pkg/cmd/commands/gslb"
	"github.com/sacloud/usacloud/pkg/cmd/commands/icon"
	"github.com/sacloud/usacloud/pkg/cmd/commands/iface"
	"github.com/sacloud/usacloud/pkg/cmd/commands/internet"
	"github.com/sacloud/usacloud/pkg/cmd/commands/internetplan"
	"github.com/sacloud/usacloud/pkg/cmd/commands/ipaddress"
	"github.com/sacloud/usacloud/pkg/cmd/commands/ipv6addr"
	"github.com/sacloud/usacloud/pkg/cmd/commands/ipv6net"
	"github.com/sacloud/usacloud/pkg/cmd/commands/license"
	"github.com/sacloud/usacloud/pkg/cmd/commands/licenseinfo"
	"github.com/sacloud/usacloud/pkg/cmd/commands/loadbalancer"
	"github.com/sacloud/usacloud/pkg/cmd/commands/note"
	"github.com/sacloud/usacloud/pkg/cmd/commands/privatehostplan"
	"github.com/sacloud/usacloud/pkg/cmd/commands/serverplan"
	"github.com/sacloud/usacloud/pkg/cmd/commands/serviceclass"
	"github.com/sacloud/usacloud/pkg/cmd/commands/swytch"
	"github.com/sacloud/usacloud/pkg/cmd/commands/zone"
	"github.com/sacloud/usacloud/pkg/cmd/core"
	"github.com/sacloud/usacloud/pkg/cmd/root"
)

var Resources = core.Resources{
	archive.Resource,
	authstatus.Resource,
	autobackup.Resource,
	bill.Resource,
	bridge.Resource,
	cdrom.Resource,
	containerregistry.Resource,
	coupon.Resource,
	database.Resource,
	disk.Resource,
	diskplan.Resource,
	dns.Resource,
	esme.Resource,
	gslb.Resource,
	icon.Resource,
	iface.Resource,
	internet.Resource,
	internetplan.Resource,
	ipaddress.Resource,
	ipv6addr.Resource,
	ipv6net.Resource,
	license.Resource,
	licenseinfo.Resource,
	loadbalancer.Resource,
	note.Resource,
	privatehostplan.Resource,
	serverplan.Resource,
	serviceclass.Resource,
	swytch.Resource,
	zone.Resource,
}

func initCommands() {
	for _, r := range Resources {
		root.Command.AddCommand(r.CLICommand())
	}
	core.BuildRootCommandsUsage(root.Command, Resources.CategorizedResources())
}
