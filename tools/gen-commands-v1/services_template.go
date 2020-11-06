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

package main

var serviceCommandTemplate = `// Code generated by 'github.com/sacloud/usacloud/tools/gen-commands-v1'; DO NOT EDIT

package {{ .PackageDirName }}

import (
	service "github.com/sacloud/libsacloud/v2/helper/service/{{ .PackageDirName }}"
	"github.com/sacloud/libsacloud/v2/sacloud/pointer"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/cmd/base"
	"github.com/sacloud/usacloud/pkg/cmd/conv"
	"github.com/spf13/pflag"
)

{{ range .Commands }}{{ if .Parameters }}
func {{ .ServiceCommandFuncName }}(ctx cli.Context, parameter *{{.CLICommandParameterTypeName}}) error { 
	svc := service.New(ctx.Client())

	req, err := parameter.ServiceRequest()
	if err != nil {
		return err
	}
	if err := req.Validate(); err != nil {
		return err
	}

	var result interface{}
	err = ctx.ExecWithProgress(func() error {
		res, err := svc.{{ .ServiceFuncName }}(ctx, req)
		if err != nil {
			return err
		}
		result = res
		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to run libsacloud service: %s", err)
	}
	return ctx.Output().Print(result)
}
{{ end }}{{ end }}
`