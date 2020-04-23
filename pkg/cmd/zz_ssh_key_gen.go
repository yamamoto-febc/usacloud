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

// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-commands'; DO NOT EDIT

package cmd

import (
	"errors"
	"sync"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/funcs"
	"github.com/sacloud/usacloud/pkg/params"
	"github.com/sacloud/usacloud/pkg/util"
	"github.com/spf13/cobra"
)

// sshKeyCmd represents the command to manage SAKURA Cloud SSHKey
func sshKeyCmd() *cobra.Command {
	return &cobra.Command{
		Use: "ssh-key",

		Short: "A manage commands of SSHKey",
		Long:  `A manage commands of SSHKey`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.HelpFunc()(cmd, args)
			return nil
		},
	}
}

func sshKeyListCmd() *cobra.Command {
	sshKeyListParam := params.NewListSSHKeyParam()
	cmd := &cobra.Command{
		Use:          "list",
		Aliases:      []string{"ls", "find"},
		Short:        "List SSHKey",
		Long:         `List SSHKey`,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := cli.NewCLIContext(globalFlags(), args, sshKeyListParam)
			if err != nil {
				return err
			}
			if err := sshKeyListParam.Initialize(newParamsAdapter(cmd.Flags()), args, ctx.Option()); err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if sshKeyListParam.GenerateSkeleton {
				return generateSkeleton(ctx, sshKeyListParam)
			}

			return funcs.SSHKeyList(ctx, sshKeyListParam)

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&sshKeyListParam.Name, "name", "", []string{}, "set filter by name(s)")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &sshKeyListParam.Id), "id", "", "set filter by id(s)")
	fs.IntVarP(&sshKeyListParam.From, "from", "", 0, "set offset (aliases: offset)")
	fs.IntVarP(&sshKeyListParam.Max, "max", "", 0, "set limit (aliases: limit)")
	fs.StringSliceVarP(&sshKeyListParam.Sort, "sort", "", []string{}, "set field(s) for sort")
	fs.StringVarP(&sshKeyListParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&sshKeyListParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&sshKeyListParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&sshKeyListParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&sshKeyListParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&sshKeyListParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&sshKeyListParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&sshKeyListParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&sshKeyListParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&sshKeyListParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&sshKeyListParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&sshKeyListParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.SetNormalizeFunc(sshKeyListNormalizeFlagNames)
	buildFlagsUsage(cmd, sshKeyListFlagOrder(cmd))
	return cmd
}

func sshKeyCreateCmd() *cobra.Command {
	sshKeyCreateParam := params.NewCreateSSHKeyParam()
	cmd := &cobra.Command{
		Use: "create",

		Short:        "Create SSHKey",
		Long:         `Create SSHKey`,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := cli.NewCLIContext(globalFlags(), args, sshKeyCreateParam)
			if err != nil {
				return err
			}
			if err := sshKeyCreateParam.Initialize(newParamsAdapter(cmd.Flags()), args, ctx.Option()); err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if sshKeyCreateParam.GenerateSkeleton {
				return generateSkeleton(ctx, sshKeyCreateParam)
			}

			// confirm
			if !sshKeyCreateParam.Assumeyes {
				if !util.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := util.ConfirmContinue("create", ctx.IO().In(), ctx.IO().Out())
				if err != nil || !result {
					return err
				}
			}

			return funcs.SSHKeyCreate(ctx, sshKeyCreateParam)

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&sshKeyCreateParam.PublicKey, "public-key", "", "", "set public-key from file")
	fs.StringVarP(&sshKeyCreateParam.Name, "name", "", "", "set resource display name")
	fs.StringVarP(&sshKeyCreateParam.Description, "description", "", "", "set resource description (aliases: desc)")
	fs.BoolVarP(&sshKeyCreateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&sshKeyCreateParam.PublicKeyContent, "public-key-content", "", "", "set public-key")
	fs.StringVarP(&sshKeyCreateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&sshKeyCreateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&sshKeyCreateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&sshKeyCreateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&sshKeyCreateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&sshKeyCreateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&sshKeyCreateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&sshKeyCreateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&sshKeyCreateParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&sshKeyCreateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&sshKeyCreateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&sshKeyCreateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.SetNormalizeFunc(sshKeyCreateNormalizeFlagNames)
	buildFlagsUsage(cmd, sshKeyCreateFlagOrder(cmd))
	return cmd
}

func sshKeyReadCmd() *cobra.Command {
	sshKeyReadParam := params.NewReadSSHKeyParam()
	cmd := &cobra.Command{
		Use: "read",

		Short:        "Read SSHKey",
		Long:         `Read SSHKey`,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := cli.NewCLIContext(globalFlags(), args, sshKeyReadParam)
			if err != nil {
				return err
			}
			if err := sshKeyReadParam.Initialize(newParamsAdapter(cmd.Flags()), args, ctx.Option()); err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if sshKeyReadParam.GenerateSkeleton {
				return generateSkeleton(ctx, sshKeyReadParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findSSHKeyReadTargets(ctx, sshKeyReadParam)
			if err != nil {
				return err
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				sshKeyReadParam.SetId(id)
				go func(p *params.ReadSSHKeyParam) {
					err := funcs.SSHKeyRead(ctx, p)
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(sshKeyReadParam)
			}
			wg.Wait()
			return cli.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&sshKeyReadParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&sshKeyReadParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&sshKeyReadParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&sshKeyReadParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&sshKeyReadParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&sshKeyReadParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&sshKeyReadParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&sshKeyReadParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&sshKeyReadParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&sshKeyReadParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&sshKeyReadParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&sshKeyReadParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &sshKeyReadParam.Id), "id", "", "Set target ID")
	fs.SetNormalizeFunc(sshKeyReadNormalizeFlagNames)
	buildFlagsUsage(cmd, sshKeyReadFlagOrder(cmd))
	return cmd
}

func sshKeyUpdateCmd() *cobra.Command {
	sshKeyUpdateParam := params.NewUpdateSSHKeyParam()
	cmd := &cobra.Command{
		Use: "update",

		Short:        "Update SSHKey",
		Long:         `Update SSHKey`,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := cli.NewCLIContext(globalFlags(), args, sshKeyUpdateParam)
			if err != nil {
				return err
			}
			if err := sshKeyUpdateParam.Initialize(newParamsAdapter(cmd.Flags()), args, ctx.Option()); err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if sshKeyUpdateParam.GenerateSkeleton {
				return generateSkeleton(ctx, sshKeyUpdateParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findSSHKeyUpdateTargets(ctx, sshKeyUpdateParam)
			if err != nil {
				return err
			}

			// confirm
			if !sshKeyUpdateParam.Assumeyes {
				if !util.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := util.ConfirmContinue("update", ctx.IO().In(), ctx.IO().Out(), ids...)
				if err != nil || !result {
					return err
				}
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				sshKeyUpdateParam.SetId(id)
				go func(p *params.UpdateSSHKeyParam) {
					err := funcs.SSHKeyUpdate(ctx, p)
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(sshKeyUpdateParam)
			}
			wg.Wait()
			return cli.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&sshKeyUpdateParam.Name, "name", "", "", "set resource display name")
	fs.StringVarP(&sshKeyUpdateParam.Description, "description", "", "", "set resource description (aliases: desc)")
	fs.BoolVarP(&sshKeyUpdateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&sshKeyUpdateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&sshKeyUpdateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&sshKeyUpdateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&sshKeyUpdateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&sshKeyUpdateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&sshKeyUpdateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&sshKeyUpdateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&sshKeyUpdateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&sshKeyUpdateParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&sshKeyUpdateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&sshKeyUpdateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&sshKeyUpdateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &sshKeyUpdateParam.Id), "id", "", "Set target ID")
	fs.SetNormalizeFunc(sshKeyUpdateNormalizeFlagNames)
	buildFlagsUsage(cmd, sshKeyUpdateFlagOrder(cmd))
	return cmd
}

func sshKeyDeleteCmd() *cobra.Command {
	sshKeyDeleteParam := params.NewDeleteSSHKeyParam()
	cmd := &cobra.Command{
		Use:          "delete",
		Aliases:      []string{"rm"},
		Short:        "Delete SSHKey",
		Long:         `Delete SSHKey`,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := cli.NewCLIContext(globalFlags(), args, sshKeyDeleteParam)
			if err != nil {
				return err
			}
			if err := sshKeyDeleteParam.Initialize(newParamsAdapter(cmd.Flags()), args, ctx.Option()); err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if sshKeyDeleteParam.GenerateSkeleton {
				return generateSkeleton(ctx, sshKeyDeleteParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findSSHKeyDeleteTargets(ctx, sshKeyDeleteParam)
			if err != nil {
				return err
			}

			// confirm
			if !sshKeyDeleteParam.Assumeyes {
				if !util.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := util.ConfirmContinue("delete", ctx.IO().In(), ctx.IO().Out(), ids...)
				if err != nil || !result {
					return err
				}
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				sshKeyDeleteParam.SetId(id)
				go func(p *params.DeleteSSHKeyParam) {
					err := funcs.SSHKeyDelete(ctx, p)
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(sshKeyDeleteParam)
			}
			wg.Wait()
			return cli.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.BoolVarP(&sshKeyDeleteParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&sshKeyDeleteParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&sshKeyDeleteParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&sshKeyDeleteParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&sshKeyDeleteParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&sshKeyDeleteParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&sshKeyDeleteParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&sshKeyDeleteParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&sshKeyDeleteParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&sshKeyDeleteParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&sshKeyDeleteParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&sshKeyDeleteParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&sshKeyDeleteParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &sshKeyDeleteParam.Id), "id", "", "Set target ID")
	fs.SetNormalizeFunc(sshKeyDeleteNormalizeFlagNames)
	buildFlagsUsage(cmd, sshKeyDeleteFlagOrder(cmd))
	return cmd
}

func sshKeyGenerateCmd() *cobra.Command {
	sshKeyGenerateParam := params.NewGenerateSSHKeyParam()
	cmd := &cobra.Command{
		Use:          "generate",
		Aliases:      []string{"gen"},
		Short:        "Generate SSHKey",
		Long:         `Generate SSHKey`,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := cli.NewCLIContext(globalFlags(), args, sshKeyGenerateParam)
			if err != nil {
				return err
			}
			if err := sshKeyGenerateParam.Initialize(newParamsAdapter(cmd.Flags()), args, ctx.Option()); err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if sshKeyGenerateParam.GenerateSkeleton {
				return generateSkeleton(ctx, sshKeyGenerateParam)
			}

			// confirm
			if !sshKeyGenerateParam.Assumeyes {
				if !util.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := util.ConfirmContinue("generate", ctx.IO().In(), ctx.IO().Out())
				if err != nil || !result {
					return err
				}
			}

			return funcs.SSHKeyGenerate(ctx, sshKeyGenerateParam)

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&sshKeyGenerateParam.PassPhrase, "pass-phrase", "", "", "set ssh-key pass phrase")
	fs.StringVarP(&sshKeyGenerateParam.Name, "name", "", "", "set resource display name")
	fs.StringVarP(&sshKeyGenerateParam.Description, "description", "", "", "set resource description (aliases: desc)")
	fs.BoolVarP(&sshKeyGenerateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&sshKeyGenerateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&sshKeyGenerateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&sshKeyGenerateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&sshKeyGenerateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&sshKeyGenerateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&sshKeyGenerateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringVarP(&sshKeyGenerateParam.PrivateKeyOutput, "private-key-output", "", "", "set ssh-key privatekey output path (aliases: file)")
	fs.StringSliceVarP(&sshKeyGenerateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&sshKeyGenerateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&sshKeyGenerateParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&sshKeyGenerateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&sshKeyGenerateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&sshKeyGenerateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.SetNormalizeFunc(sshKeyGenerateNormalizeFlagNames)
	buildFlagsUsage(cmd, sshKeyGenerateFlagOrder(cmd))
	return cmd
}
