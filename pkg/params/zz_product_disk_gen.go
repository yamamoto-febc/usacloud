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

// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-params'; DO NOT EDIT

package params

import (
	"fmt"
	"io"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/pkg/define"
	"github.com/sacloud/usacloud/pkg/flags"
	"github.com/sacloud/usacloud/pkg/output"
	"github.com/sacloud/usacloud/pkg/schema"
	"github.com/sacloud/usacloud/pkg/util"
	"github.com/sacloud/usacloud/pkg/validation"
)

// ListProductDiskParam is input parameters for the sacloud API
type ListProductDiskParam struct {
	Name              []string
	Id                []sacloud.ID
	From              int
	Max               int
	Sort              []string
	ParamTemplate     string
	Parameters        string
	ParamTemplateFile string
	ParameterFile     string
	GenerateSkeleton  bool
	OutputType        string
	Column            []string
	Quiet             bool
	Format            string
	FormatFile        string
	Query             string
	QueryFile         string

	options *flags.Flags
	input   Input
}

// NewListProductDiskParam return new ListProductDiskParam
func NewListProductDiskParam() *ListProductDiskParam {
	return &ListProductDiskParam{}
}

// Initialize init ListProductDiskParam
func (p *ListProductDiskParam) Initialize(in Input, args []string, options *flags.Flags) error {
	p.input = in
	p.options = options
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *ListProductDiskParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

// FillValueToSkeleton fills empty value to the parameter
func (p *ListProductDiskParam) FillValueToSkeleton() {
	if util.IsEmpty(p.Name) {
		p.Name = []string{""}
	}
	if util.IsEmpty(p.Id) {
		p.Id = []sacloud.ID{}
	}
	if util.IsEmpty(p.From) {
		p.From = 0
	}
	if util.IsEmpty(p.Max) {
		p.Max = 0
	}
	if util.IsEmpty(p.Sort) {
		p.Sort = []string{""}
	}
	if util.IsEmpty(p.ParamTemplate) {
		p.ParamTemplate = ""
	}
	if util.IsEmpty(p.Parameters) {
		p.Parameters = ""
	}
	if util.IsEmpty(p.ParamTemplateFile) {
		p.ParamTemplateFile = ""
	}
	if util.IsEmpty(p.ParameterFile) {
		p.ParameterFile = ""
	}
	if util.IsEmpty(p.GenerateSkeleton) {
		p.GenerateSkeleton = false
	}
	if util.IsEmpty(p.OutputType) {
		p.OutputType = ""
	}
	if util.IsEmpty(p.Column) {
		p.Column = []string{""}
	}
	if util.IsEmpty(p.Quiet) {
		p.Quiet = false
	}
	if util.IsEmpty(p.Format) {
		p.Format = ""
	}
	if util.IsEmpty(p.FormatFile) {
		p.FormatFile = ""
	}
	if util.IsEmpty(p.Query) {
		p.Query = ""
	}
	if util.IsEmpty(p.QueryFile) {
		p.QueryFile = ""
	}

}

func (p *ListProductDiskParam) validate() error {
	var errors []error

	{
		errs := validation.ConflictsWith("--name", p.Name, map[string]interface{}{

			"--id": p.Id,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["ProductDisk"].Commands["list"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validation.ConflictsWith("--id", p.Id, map[string]interface{}{

			"--name": p.Name,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		errs := validateParameterOptions(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := schema.ValidateInStrValues(define.AllowOutputTypes...)
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateOutputOption(p, p.options.DefaultOutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	return util.FlattenErrors(errors)
}

func (p *ListProductDiskParam) ResourceDef() *schema.Resource {
	return define.Resources["ProductDisk"]
}

func (p *ListProductDiskParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["list"]
}

func (p *ListProductDiskParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ListProductDiskParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ListProductDiskParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ListProductDiskParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

/*
 * v0系との互換性維持のための実装
 */
func (p *ListProductDiskParam) GetResourceDef() *schema.Resource {
	return define.Resources["ProductDisk"]
}

func (p *ListProductDiskParam) GetCommandDef() *schema.Command {
	return p.ResourceDef().Commands["list"]
}

func (p *ListProductDiskParam) GetIncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ListProductDiskParam) GetExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ListProductDiskParam) GetTableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ListProductDiskParam) GetColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *ListProductDiskParam) SetName(v []string) {
	p.Name = v
}

func (p *ListProductDiskParam) GetName() []string {
	return p.Name
}
func (p *ListProductDiskParam) SetId(v []sacloud.ID) {
	p.Id = v
}

func (p *ListProductDiskParam) GetId() []sacloud.ID {
	return p.Id
}
func (p *ListProductDiskParam) SetFrom(v int) {
	p.From = v
}

func (p *ListProductDiskParam) GetFrom() int {
	return p.From
}
func (p *ListProductDiskParam) SetMax(v int) {
	p.Max = v
}

func (p *ListProductDiskParam) GetMax() int {
	return p.Max
}
func (p *ListProductDiskParam) SetSort(v []string) {
	p.Sort = v
}

func (p *ListProductDiskParam) GetSort() []string {
	return p.Sort
}
func (p *ListProductDiskParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ListProductDiskParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ListProductDiskParam) SetParameters(v string) {
	p.Parameters = v
}

func (p *ListProductDiskParam) GetParameters() string {
	return p.Parameters
}
func (p *ListProductDiskParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ListProductDiskParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ListProductDiskParam) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *ListProductDiskParam) GetParameterFile() string {
	return p.ParameterFile
}
func (p *ListProductDiskParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *ListProductDiskParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *ListProductDiskParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ListProductDiskParam) GetOutputType() string {
	return p.OutputType
}
func (p *ListProductDiskParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ListProductDiskParam) GetColumn() []string {
	return p.Column
}
func (p *ListProductDiskParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ListProductDiskParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ListProductDiskParam) SetFormat(v string) {
	p.Format = v
}

func (p *ListProductDiskParam) GetFormat() string {
	return p.Format
}
func (p *ListProductDiskParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ListProductDiskParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *ListProductDiskParam) SetQuery(v string) {
	p.Query = v
}

func (p *ListProductDiskParam) GetQuery() string {
	return p.Query
}
func (p *ListProductDiskParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *ListProductDiskParam) GetQueryFile() string {
	return p.QueryFile
}

// Changed usacloud v0系との互換性維持のための実装
func (p *ListProductDiskParam) Changed(name string) bool {
	return p.input.Changed(name)
}

// ReadProductDiskParam is input parameters for the sacloud API
type ReadProductDiskParam struct {
	Assumeyes         bool
	ParamTemplate     string
	Parameters        string
	ParamTemplateFile string
	ParameterFile     string
	GenerateSkeleton  bool
	OutputType        string
	Column            []string
	Quiet             bool
	Format            string
	FormatFile        string
	Query             string
	QueryFile         string
	Id                sacloud.ID

	options *flags.Flags
	input   Input
}

// NewReadProductDiskParam return new ReadProductDiskParam
func NewReadProductDiskParam() *ReadProductDiskParam {
	return &ReadProductDiskParam{}
}

// Initialize init ReadProductDiskParam
func (p *ReadProductDiskParam) Initialize(in Input, args []string, options *flags.Flags) error {
	p.input = in
	p.options = options

	if len(args) == 0 {
		return fmt.Errorf("argument <ID> is required")
	}
	p.Id = sacloud.StringID(args[0])
	if p.Id.IsEmpty() {
		return fmt.Errorf("argument <ID> is required")
	}
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *ReadProductDiskParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

// FillValueToSkeleton fills empty value to the parameter
func (p *ReadProductDiskParam) FillValueToSkeleton() {
	if util.IsEmpty(p.Assumeyes) {
		p.Assumeyes = false
	}
	if util.IsEmpty(p.ParamTemplate) {
		p.ParamTemplate = ""
	}
	if util.IsEmpty(p.Parameters) {
		p.Parameters = ""
	}
	if util.IsEmpty(p.ParamTemplateFile) {
		p.ParamTemplateFile = ""
	}
	if util.IsEmpty(p.ParameterFile) {
		p.ParameterFile = ""
	}
	if util.IsEmpty(p.GenerateSkeleton) {
		p.GenerateSkeleton = false
	}
	if util.IsEmpty(p.OutputType) {
		p.OutputType = ""
	}
	if util.IsEmpty(p.Column) {
		p.Column = []string{""}
	}
	if util.IsEmpty(p.Quiet) {
		p.Quiet = false
	}
	if util.IsEmpty(p.Format) {
		p.Format = ""
	}
	if util.IsEmpty(p.FormatFile) {
		p.FormatFile = ""
	}
	if util.IsEmpty(p.Query) {
		p.Query = ""
	}
	if util.IsEmpty(p.QueryFile) {
		p.QueryFile = ""
	}
	if util.IsEmpty(p.Id) {
		p.Id = sacloud.ID(0)
	}

}

func (p *ReadProductDiskParam) validate() error {
	var errors []error

	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["ProductDisk"].Commands["read"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		errs := validateParameterOptions(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := schema.ValidateInStrValues(define.AllowOutputTypes...)
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateOutputOption(p, p.options.DefaultOutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	return util.FlattenErrors(errors)
}

func (p *ReadProductDiskParam) ResourceDef() *schema.Resource {
	return define.Resources["ProductDisk"]
}

func (p *ReadProductDiskParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["read"]
}

func (p *ReadProductDiskParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ReadProductDiskParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ReadProductDiskParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ReadProductDiskParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

/*
 * v0系との互換性維持のための実装
 */
func (p *ReadProductDiskParam) GetResourceDef() *schema.Resource {
	return define.Resources["ProductDisk"]
}

func (p *ReadProductDiskParam) GetCommandDef() *schema.Command {
	return p.ResourceDef().Commands["read"]
}

func (p *ReadProductDiskParam) GetIncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ReadProductDiskParam) GetExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ReadProductDiskParam) GetTableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ReadProductDiskParam) GetColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *ReadProductDiskParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *ReadProductDiskParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *ReadProductDiskParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ReadProductDiskParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ReadProductDiskParam) SetParameters(v string) {
	p.Parameters = v
}

func (p *ReadProductDiskParam) GetParameters() string {
	return p.Parameters
}
func (p *ReadProductDiskParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ReadProductDiskParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ReadProductDiskParam) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *ReadProductDiskParam) GetParameterFile() string {
	return p.ParameterFile
}
func (p *ReadProductDiskParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *ReadProductDiskParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *ReadProductDiskParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ReadProductDiskParam) GetOutputType() string {
	return p.OutputType
}
func (p *ReadProductDiskParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ReadProductDiskParam) GetColumn() []string {
	return p.Column
}
func (p *ReadProductDiskParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ReadProductDiskParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ReadProductDiskParam) SetFormat(v string) {
	p.Format = v
}

func (p *ReadProductDiskParam) GetFormat() string {
	return p.Format
}
func (p *ReadProductDiskParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ReadProductDiskParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *ReadProductDiskParam) SetQuery(v string) {
	p.Query = v
}

func (p *ReadProductDiskParam) GetQuery() string {
	return p.Query
}
func (p *ReadProductDiskParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *ReadProductDiskParam) GetQueryFile() string {
	return p.QueryFile
}
func (p *ReadProductDiskParam) SetId(v sacloud.ID) {
	p.Id = v
}

func (p *ReadProductDiskParam) GetId() sacloud.ID {
	return p.Id
}

// Changed usacloud v0系との互換性維持のための実装
func (p *ReadProductDiskParam) Changed(name string) bool {
	return p.input.Changed(name)
}
