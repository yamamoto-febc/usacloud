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

func GSLBServerDelete(ctx command.Context, params *params.ServerDeleteGSLBParam) error {

	client := ctx.GetAPIClient()
	api := client.GetGSLBAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("GSLBServerDelete is failed: %s", e)
	}

	if len(p.Settings.GSLB.Servers) == 0 {
		return fmt.Errorf("GSLB don't have any servers")
	}

	// validate index
	if params.Index <= 0 || params.Index-1 >= len(p.Settings.GSLB.Servers) {
		return fmt.Errorf("index(%d) is out of range", params.Index)
	}

	// delete by ipaddress
	p.Settings.GSLB.DeleteServer(p.Settings.GSLB.Servers[params.Index-1].IPAddress)

	p, e = api.Update(params.Id, p)
	if e != nil {
		return fmt.Errorf("GSLBServerDelete is failed: %s", e)
	}

	list := []interface{}{}
	for i := range p.Settings.GSLB.Servers {
		list = append(list, p.Settings.GSLB.Servers[i])
	}

	return ctx.GetOutput().Print(list...)

}