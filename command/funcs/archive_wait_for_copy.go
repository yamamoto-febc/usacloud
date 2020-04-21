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

func ArchiveWaitForCopy(ctx command.Context, params *params.WaitForCopyArchiveParam) error {

	client := ctx.GetAPIClient()
	api := client.GetArchiveAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ArchiveWaitForCopy is failed: %s", e)
	}

	// wait for copy with progress
	err := internal.ExecWithProgress(
		fmt.Sprintf("Still copying[ID:%d]...", params.Id),
		fmt.Sprintf("Copy archive[ID:%d]", params.Id),
		command.GlobalOption.Progress,
		func(compChan chan bool, errChan chan error) {
			err := api.SleepWhileCopying(p.ID, client.DefaultTimeoutDuration)
			if err != nil {
				errChan <- err
			}
			compChan <- true
		},
	)
	if err != nil {
		return fmt.Errorf("ArchiveWaitForCopy is failed: %s", err)
	}

	return nil
}