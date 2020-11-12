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

var serviceCommandTemplate = `// Code generated by 'github.com/sacloud/usacloud/tools/gen-commands'; DO NOT EDIT

package services

import (
	service "github.com/sacloud/libsacloud/v2/helper/service/{{ .PackageDirName }}"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/cmd/conv"
)

{{ if .ServiceType }}
func init() { {{ range .Commands }}{{ if .ParameterInitializer }}
	setDefaultServiceFunc("{{ .Resource.Name }}", "{{.Name}}", 
		func (ctx cli.Context, parameter interface{}) ([]interface{}, error) { 
			svc := service.New(ctx.Client())
			{{ if .ServiceFuncReturnValueType.HasRequestValue }}
			req := &service.{{.ServiceRequestTypeName}}{}
			if err := conv.ConvertTo(parameter, req); err != nil {
				return nil, err
			}
			if err := req.Validate(); err != nil {
				return nil, err
			}
			{{ end }}
		
			{{ if .ServiceFuncReturnValueType.HasReturnValue }}res, {{ end }}err := svc.{{ .ServiceFuncName }}(ctx{{ if .ServiceFuncReturnValueType.HasRequestValue }}, req{{ end }})
			if err != nil {
				return nil, err
			}

			{{ if .ServiceFuncReturnValueType.HasReturnValue }}
			{{ if .ServiceFuncReturnValueType.IsReturnValueSlice }}
			var results []interface{}
			for _ , v := range res {
				results = append(results, v)
			}
			return results, nil
			{{ else }}
			return []interface{}{res}, nil
			{{ end -}}
			{{ else }}
			return nil, nil
			{{ end }}
		},
	)
	{{ if .Resource.ServiceMeta.HasFindMethod }}setDefaultListAllFunc("{{ .Resource.Name }}", "{{.Name}}", 
		func (ctx cli.Context, parameter interface{}) ([]interface{}, error) { 
			svc := service.New(ctx.Client())
			res, err := svc.FindWithContext(ctx, &service.FindRequest{ {{ if not .Resource.IsGlobalResource}} Zone: (parameter.(cflag.ZoneParameterValueHandler)).ZoneFlagValue(){{ end }} })
			if err != nil {
				return nil, err
			}
		
			var results []interface{}
			for _, v := range res {
				results = append(results, v)
			}
			return results, nil
		},
	){{ end -}}
{{ end }}{{ end -}}
}
{{ end }}
`