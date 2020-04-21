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

// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package completion

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/schema"
)

func ConfigCurrentCompleteFlags(ctx command.Context, params *params.CurrentConfigParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func ConfigDeleteCompleteFlags(ctx command.Context, params *params.DeleteConfigParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func ConfigEditCompleteFlags(ctx command.Context, params *params.EditConfigParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "token":
		param := define.Resources["Config"].Commands["edit"].BuildedParams().Get("token")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "secret":
		param := define.Resources["Config"].Commands["edit"].BuildedParams().Get("secret")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "zone":
		param := define.Resources["Config"].Commands["edit"].BuildedParams().Get("zone")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "default-output-type":
		param := define.Resources["Config"].Commands["edit"].BuildedParams().Get("default-output-type")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func ConfigListCompleteFlags(ctx command.Context, params *params.ListConfigParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func ConfigMigrateCompleteFlags(ctx command.Context, params *params.MigrateConfigParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func ConfigShowCompleteFlags(ctx command.Context, params *params.ShowConfigParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func ConfigUseCompleteFlags(ctx command.Context, params *params.UseConfigParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}