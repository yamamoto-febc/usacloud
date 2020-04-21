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

func PacketFilterListCompleteFlags(ctx command.Context, params *params.ListPacketFilterParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "name":
		param := define.Resources["PacketFilter"].Commands["list"].BuildedParams().Get("name")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["PacketFilter"].Commands["list"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "from", "offset":
		param := define.Resources["PacketFilter"].Commands["list"].BuildedParams().Get("from")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "max", "limit":
		param := define.Resources["PacketFilter"].Commands["list"].BuildedParams().Get("max")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "sort":
		param := define.Resources["PacketFilter"].Commands["list"].BuildedParams().Get("sort")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func PacketFilterCreateCompleteFlags(ctx command.Context, params *params.CreatePacketFilterParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "name":
		param := define.Resources["PacketFilter"].Commands["create"].BuildedParams().Get("name")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "description", "desc":
		param := define.Resources["PacketFilter"].Commands["create"].BuildedParams().Get("description")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func PacketFilterReadCompleteFlags(ctx command.Context, params *params.ReadPacketFilterParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "id":
		param := define.Resources["PacketFilter"].Commands["read"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func PacketFilterUpdateCompleteFlags(ctx command.Context, params *params.UpdatePacketFilterParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "name":
		param := define.Resources["PacketFilter"].Commands["update"].BuildedParams().Get("name")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "description", "desc":
		param := define.Resources["PacketFilter"].Commands["update"].BuildedParams().Get("description")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["PacketFilter"].Commands["update"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func PacketFilterDeleteCompleteFlags(ctx command.Context, params *params.DeletePacketFilterParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "id":
		param := define.Resources["PacketFilter"].Commands["delete"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func PacketFilterRuleInfoCompleteFlags(ctx command.Context, params *params.RuleInfoPacketFilterParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "id":
		param := define.Resources["PacketFilter"].Commands["rule-info"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func PacketFilterRuleAddCompleteFlags(ctx command.Context, params *params.RuleAddPacketFilterParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "index":
		param := define.Resources["PacketFilter"].Commands["rule-add"].BuildedParams().Get("index")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "protocol":
		param := define.Resources["PacketFilter"].Commands["rule-add"].BuildedParams().Get("protocol")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "source-network":
		param := define.Resources["PacketFilter"].Commands["rule-add"].BuildedParams().Get("source-network")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "source-port":
		param := define.Resources["PacketFilter"].Commands["rule-add"].BuildedParams().Get("source-port")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "destination-port", "dest-port":
		param := define.Resources["PacketFilter"].Commands["rule-add"].BuildedParams().Get("destination-port")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "action":
		param := define.Resources["PacketFilter"].Commands["rule-add"].BuildedParams().Get("action")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "description", "desc":
		param := define.Resources["PacketFilter"].Commands["rule-add"].BuildedParams().Get("description")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["PacketFilter"].Commands["rule-add"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func PacketFilterRuleUpdateCompleteFlags(ctx command.Context, params *params.RuleUpdatePacketFilterParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "index":
		param := define.Resources["PacketFilter"].Commands["rule-update"].BuildedParams().Get("index")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "protocol":
		param := define.Resources["PacketFilter"].Commands["rule-update"].BuildedParams().Get("protocol")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "source-network":
		param := define.Resources["PacketFilter"].Commands["rule-update"].BuildedParams().Get("source-network")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "source-port":
		param := define.Resources["PacketFilter"].Commands["rule-update"].BuildedParams().Get("source-port")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "destination-port", "dest-port":
		param := define.Resources["PacketFilter"].Commands["rule-update"].BuildedParams().Get("destination-port")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "action":
		param := define.Resources["PacketFilter"].Commands["rule-update"].BuildedParams().Get("action")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "description", "desc":
		param := define.Resources["PacketFilter"].Commands["rule-update"].BuildedParams().Get("description")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["PacketFilter"].Commands["rule-update"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func PacketFilterRuleDeleteCompleteFlags(ctx command.Context, params *params.RuleDeletePacketFilterParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "index":
		param := define.Resources["PacketFilter"].Commands["rule-delete"].BuildedParams().Get("index")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["PacketFilter"].Commands["rule-delete"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func PacketFilterInterfaceConnectCompleteFlags(ctx command.Context, params *params.InterfaceConnectPacketFilterParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "interface-id":
		param := define.Resources["PacketFilter"].Commands["interface-connect"].BuildedParams().Get("interface-id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["PacketFilter"].Commands["interface-connect"].BuildedParams().Get("id")
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

func PacketFilterInterfaceDisconnectCompleteFlags(ctx command.Context, params *params.InterfaceDisconnectPacketFilterParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "interface-id":
		param := define.Resources["PacketFilter"].Commands["interface-disconnect"].BuildedParams().Get("interface-id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["PacketFilter"].Commands["interface-disconnect"].BuildedParams().Get("id")
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