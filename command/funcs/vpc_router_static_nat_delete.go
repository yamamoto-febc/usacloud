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

package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func VPCRouterStaticNatDelete(ctx command.Context, params *params.StaticNatDeleteVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterStaticNatDelete is failed: %s", e)
	}
	if p.IsStandardPlan() {
		return fmt.Errorf("Static NAT is not supported on standard plan")
	}
	if !p.HasStaticNAT() {
		return fmt.Errorf("VPCRouter[%d] don't have any static NAT settings", params.Id)
	}

	// validate

	// index
	if params.Index <= 0 || params.Index-1 >= len(p.Settings.Router.StaticNAT.Config) {
		return fmt.Errorf("index(%d) is out of range", params.Index)
	}

	natConfig := p.Settings.Router.StaticNAT.Config[params.Index-1]

	p.Settings.Router.RemoveStaticNAT(natConfig.GlobalAddress, natConfig.PrivateAddress)

	_, err := api.UpdateSetting(params.Id, p)
	if err != nil {
		return fmt.Errorf("VPCRouterStaticNatDelete is failed: %s", err)
	}
	_, err = api.Config(params.Id)
	if err != nil {
		return fmt.Errorf("VPCRouterStaticNatDelete is failed: %s", err)
	}

	return nil

}