// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package params

import (
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

// CurrentConfigParam is input parameters for the sacloud API
type CurrentConfigParam struct {
	ParamTemplate     string `json:"param-template"`
	ParamTemplateFile string `json:"param-template-file"`
	GenerateSkeleton  bool   `json:"generate-skeleton"`
}

// NewCurrentConfigParam return new CurrentConfigParam
func NewCurrentConfigParam() *CurrentConfigParam {
	return &CurrentConfigParam{}
}

// FillValueToSkeleton fill values to empty fields
func (p *CurrentConfigParam) FillValueToSkeleton() {
	if isEmpty(p.ParamTemplate) {
		p.ParamTemplate = ""
	}
	if isEmpty(p.ParamTemplateFile) {
		p.ParamTemplateFile = ""
	}
	if isEmpty(p.GenerateSkeleton) {
		p.GenerateSkeleton = false
	}

}

// Validate checks current values in model
func (p *CurrentConfigParam) Validate() []error {
	errors := []error{}

	return errors
}

func (p *CurrentConfigParam) GetResourceDef() *schema.Resource {
	return define.Resources["Config"]
}

func (p *CurrentConfigParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["current"]
}

func (p *CurrentConfigParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *CurrentConfigParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *CurrentConfigParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *CurrentConfigParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *CurrentConfigParam) GetOutputFormat() string {
	return "table"
}

func (p *CurrentConfigParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *CurrentConfigParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *CurrentConfigParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *CurrentConfigParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *CurrentConfigParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *CurrentConfigParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}

// DeleteConfigParam is input parameters for the sacloud API
type DeleteConfigParam struct {
	Assumeyes         bool   `json:"assumeyes"`
	ParamTemplate     string `json:"param-template"`
	ParamTemplateFile string `json:"param-template-file"`
	GenerateSkeleton  bool   `json:"generate-skeleton"`
}

// NewDeleteConfigParam return new DeleteConfigParam
func NewDeleteConfigParam() *DeleteConfigParam {
	return &DeleteConfigParam{}
}

// FillValueToSkeleton fill values to empty fields
func (p *DeleteConfigParam) FillValueToSkeleton() {
	if isEmpty(p.Assumeyes) {
		p.Assumeyes = false
	}
	if isEmpty(p.ParamTemplate) {
		p.ParamTemplate = ""
	}
	if isEmpty(p.ParamTemplateFile) {
		p.ParamTemplateFile = ""
	}
	if isEmpty(p.GenerateSkeleton) {
		p.GenerateSkeleton = false
	}

}

// Validate checks current values in model
func (p *DeleteConfigParam) Validate() []error {
	errors := []error{}

	return errors
}

func (p *DeleteConfigParam) GetResourceDef() *schema.Resource {
	return define.Resources["Config"]
}

func (p *DeleteConfigParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["delete"]
}

func (p *DeleteConfigParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *DeleteConfigParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *DeleteConfigParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *DeleteConfigParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *DeleteConfigParam) GetOutputFormat() string {
	return "table"
}

func (p *DeleteConfigParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *DeleteConfigParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *DeleteConfigParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *DeleteConfigParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *DeleteConfigParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *DeleteConfigParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *DeleteConfigParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *DeleteConfigParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}

// EditConfigParam is input parameters for the sacloud API
type EditConfigParam struct {
	Token             string `json:"token"`
	Secret            string `json:"secret"`
	Zone              string `json:"zone"`
	ParamTemplate     string `json:"param-template"`
	ParamTemplateFile string `json:"param-template-file"`
	GenerateSkeleton  bool   `json:"generate-skeleton"`
}

// NewEditConfigParam return new EditConfigParam
func NewEditConfigParam() *EditConfigParam {
	return &EditConfigParam{}
}

// FillValueToSkeleton fill values to empty fields
func (p *EditConfigParam) FillValueToSkeleton() {
	if isEmpty(p.Token) {
		p.Token = ""
	}
	if isEmpty(p.Secret) {
		p.Secret = ""
	}
	if isEmpty(p.Zone) {
		p.Zone = ""
	}
	if isEmpty(p.ParamTemplate) {
		p.ParamTemplate = ""
	}
	if isEmpty(p.ParamTemplateFile) {
		p.ParamTemplateFile = ""
	}
	if isEmpty(p.GenerateSkeleton) {
		p.GenerateSkeleton = false
	}

}

// Validate checks current values in model
func (p *EditConfigParam) Validate() []error {
	errors := []error{}
	{
		validator := define.Resources["Config"].Commands["edit"].Params["zone"].ValidateFunc
		errs := validator("--zone", p.Zone)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *EditConfigParam) GetResourceDef() *schema.Resource {
	return define.Resources["Config"]
}

func (p *EditConfigParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["edit"]
}

func (p *EditConfigParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *EditConfigParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *EditConfigParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *EditConfigParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *EditConfigParam) GetOutputFormat() string {
	return "table"
}

func (p *EditConfigParam) SetToken(v string) {
	p.Token = v
}

func (p *EditConfigParam) GetToken() string {
	return p.Token
}
func (p *EditConfigParam) SetSecret(v string) {
	p.Secret = v
}

func (p *EditConfigParam) GetSecret() string {
	return p.Secret
}
func (p *EditConfigParam) SetZone(v string) {
	p.Zone = v
}

func (p *EditConfigParam) GetZone() string {
	return p.Zone
}
func (p *EditConfigParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *EditConfigParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *EditConfigParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *EditConfigParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *EditConfigParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *EditConfigParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}

// ListConfigParam is input parameters for the sacloud API
type ListConfigParam struct {
	ParamTemplate     string `json:"param-template"`
	ParamTemplateFile string `json:"param-template-file"`
	GenerateSkeleton  bool   `json:"generate-skeleton"`
}

// NewListConfigParam return new ListConfigParam
func NewListConfigParam() *ListConfigParam {
	return &ListConfigParam{}
}

// FillValueToSkeleton fill values to empty fields
func (p *ListConfigParam) FillValueToSkeleton() {
	if isEmpty(p.ParamTemplate) {
		p.ParamTemplate = ""
	}
	if isEmpty(p.ParamTemplateFile) {
		p.ParamTemplateFile = ""
	}
	if isEmpty(p.GenerateSkeleton) {
		p.GenerateSkeleton = false
	}

}

// Validate checks current values in model
func (p *ListConfigParam) Validate() []error {
	errors := []error{}

	return errors
}

func (p *ListConfigParam) GetResourceDef() *schema.Resource {
	return define.Resources["Config"]
}

func (p *ListConfigParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["list"]
}

func (p *ListConfigParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *ListConfigParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *ListConfigParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *ListConfigParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *ListConfigParam) GetOutputFormat() string {
	return "table"
}

func (p *ListConfigParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ListConfigParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ListConfigParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ListConfigParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ListConfigParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *ListConfigParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}

// MigrateConfigParam is input parameters for the sacloud API
type MigrateConfigParam struct {
	ParamTemplate     string `json:"param-template"`
	ParamTemplateFile string `json:"param-template-file"`
	GenerateSkeleton  bool   `json:"generate-skeleton"`
}

// NewMigrateConfigParam return new MigrateConfigParam
func NewMigrateConfigParam() *MigrateConfigParam {
	return &MigrateConfigParam{}
}

// FillValueToSkeleton fill values to empty fields
func (p *MigrateConfigParam) FillValueToSkeleton() {
	if isEmpty(p.ParamTemplate) {
		p.ParamTemplate = ""
	}
	if isEmpty(p.ParamTemplateFile) {
		p.ParamTemplateFile = ""
	}
	if isEmpty(p.GenerateSkeleton) {
		p.GenerateSkeleton = false
	}

}

// Validate checks current values in model
func (p *MigrateConfigParam) Validate() []error {
	errors := []error{}

	return errors
}

func (p *MigrateConfigParam) GetResourceDef() *schema.Resource {
	return define.Resources["Config"]
}

func (p *MigrateConfigParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["migrate"]
}

func (p *MigrateConfigParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *MigrateConfigParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *MigrateConfigParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *MigrateConfigParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *MigrateConfigParam) GetOutputFormat() string {
	return "table"
}

func (p *MigrateConfigParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *MigrateConfigParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *MigrateConfigParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *MigrateConfigParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *MigrateConfigParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *MigrateConfigParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}

// ShowConfigParam is input parameters for the sacloud API
type ShowConfigParam struct {
	ParamTemplate     string `json:"param-template"`
	ParamTemplateFile string `json:"param-template-file"`
	GenerateSkeleton  bool   `json:"generate-skeleton"`
}

// NewShowConfigParam return new ShowConfigParam
func NewShowConfigParam() *ShowConfigParam {
	return &ShowConfigParam{}
}

// FillValueToSkeleton fill values to empty fields
func (p *ShowConfigParam) FillValueToSkeleton() {
	if isEmpty(p.ParamTemplate) {
		p.ParamTemplate = ""
	}
	if isEmpty(p.ParamTemplateFile) {
		p.ParamTemplateFile = ""
	}
	if isEmpty(p.GenerateSkeleton) {
		p.GenerateSkeleton = false
	}

}

// Validate checks current values in model
func (p *ShowConfigParam) Validate() []error {
	errors := []error{}

	return errors
}

func (p *ShowConfigParam) GetResourceDef() *schema.Resource {
	return define.Resources["Config"]
}

func (p *ShowConfigParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["show"]
}

func (p *ShowConfigParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *ShowConfigParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *ShowConfigParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *ShowConfigParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *ShowConfigParam) GetOutputFormat() string {
	return "table"
}

func (p *ShowConfigParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ShowConfigParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ShowConfigParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ShowConfigParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ShowConfigParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *ShowConfigParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}

// UseConfigParam is input parameters for the sacloud API
type UseConfigParam struct {
	ParamTemplate     string `json:"param-template"`
	ParamTemplateFile string `json:"param-template-file"`
	GenerateSkeleton  bool   `json:"generate-skeleton"`
}

// NewUseConfigParam return new UseConfigParam
func NewUseConfigParam() *UseConfigParam {
	return &UseConfigParam{}
}

// FillValueToSkeleton fill values to empty fields
func (p *UseConfigParam) FillValueToSkeleton() {
	if isEmpty(p.ParamTemplate) {
		p.ParamTemplate = ""
	}
	if isEmpty(p.ParamTemplateFile) {
		p.ParamTemplateFile = ""
	}
	if isEmpty(p.GenerateSkeleton) {
		p.GenerateSkeleton = false
	}

}

// Validate checks current values in model
func (p *UseConfigParam) Validate() []error {
	errors := []error{}

	return errors
}

func (p *UseConfigParam) GetResourceDef() *schema.Resource {
	return define.Resources["Config"]
}

func (p *UseConfigParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["use"]
}

func (p *UseConfigParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *UseConfigParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *UseConfigParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *UseConfigParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *UseConfigParam) GetOutputFormat() string {
	return "table"
}

func (p *UseConfigParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *UseConfigParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *UseConfigParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *UseConfigParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *UseConfigParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *UseConfigParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}