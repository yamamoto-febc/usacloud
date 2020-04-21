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
	"github.com/sacloud/usacloud/command/internal"
	"github.com/sacloud/usacloud/command/params"
)

func InternetSubnetUpdate(ctx command.Context, params *params.SubnetUpdateInternetParam) error {

	client := ctx.GetAPIClient()
	api := client.GetInternetAPI()
	_, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("InternetSubnetUpdate is failed: %s", e)
	}

	var sn *subnet
	err := internal.ExecWithProgress(
		fmt.Sprintf("Still updating[ID:%d]...", params.SubnetId),
		fmt.Sprintf("Update subnet[ID:%d]", params.SubnetId),
		command.GlobalOption.Progress,
		func(compChan chan bool, errChan chan error) {
			subnet, err := api.UpdateSubnet(params.Id, params.SubnetId, params.NextHop)
			if err != nil {
				errChan <- err
				return
			}
			sn, err = getSubnetByID(ctx, subnet.ID)
			if err != nil {
				errChan <- err
				return
			}
			compChan <- true
		},
	)
	if err != nil {
		return fmt.Errorf("InternetSubnetUpdate is failed: %s", err)
	}
	return ctx.GetOutput().Print(sn)
}