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

	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/funcs"
	"github.com/sacloud/usacloud/pkg/params"
	"github.com/sacloud/usacloud/pkg/util"
	"github.com/spf13/cobra"
)

// webAccelCmd represents the command to manage SAKURA Cloud WebAccel
func webAccelCmd() *cobra.Command {
	return &cobra.Command{
		Use: "web-accel",

		Short: "A manage commands of WebAccel",
		Long:  `A manage commands of WebAccel`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.HelpFunc()(cmd, args)
			return nil
		},
	}
}

func webAccelListCmd() *cobra.Command {
	webAccelListParam := params.NewListWebAccelParam()
	cmd := &cobra.Command{
		Use:          "list",
		Aliases:      []string{"ls", "find", "select"},
		Short:        "List WebAccel",
		Long:         `List WebAccel`,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := cli.NewCLIContext(globalFlags(), args, webAccelListParam)
			if err != nil {
				return err
			}
			if err := webAccelListParam.Initialize(newParamsAdapter(cmd.Flags()), args, ctx.Option()); err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if webAccelListParam.GenerateSkeleton {
				return generateSkeleton(ctx, webAccelListParam)
			}

			return funcs.WebAccelList(ctx, webAccelListParam)

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&webAccelListParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&webAccelListParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&webAccelListParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&webAccelListParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&webAccelListParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&webAccelListParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&webAccelListParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&webAccelListParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&webAccelListParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&webAccelListParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&webAccelListParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&webAccelListParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.SetNormalizeFunc(webAccelListNormalizeFlagNames)
	buildFlagsUsage(cmd, webAccelListFlagOrder(cmd))
	return cmd
}

func webAccelReadCmd() *cobra.Command {
	webAccelReadParam := params.NewReadWebAccelParam()
	cmd := &cobra.Command{
		Use: "read",

		Short:        "Read WebAccel",
		Long:         `Read WebAccel`,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := cli.NewCLIContext(globalFlags(), args, webAccelReadParam)
			if err != nil {
				return err
			}
			if err := webAccelReadParam.Initialize(newParamsAdapter(cmd.Flags()), args, ctx.Option()); err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if webAccelReadParam.GenerateSkeleton {
				return generateSkeleton(ctx, webAccelReadParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findWebAccelReadTargets(ctx, webAccelReadParam)
			if err != nil {
				return err
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				webAccelReadParam.SetId(id)
				go func(p *params.ReadWebAccelParam) {
					err := funcs.WebAccelRead(ctx, p)
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(webAccelReadParam)
			}
			wg.Wait()
			return cli.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&webAccelReadParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.StringVarP(&webAccelReadParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&webAccelReadParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&webAccelReadParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&webAccelReadParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&webAccelReadParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&webAccelReadParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&webAccelReadParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&webAccelReadParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&webAccelReadParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&webAccelReadParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&webAccelReadParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&webAccelReadParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &webAccelReadParam.Id), "id", "", "Set target ID")
	fs.SetNormalizeFunc(webAccelReadNormalizeFlagNames)
	buildFlagsUsage(cmd, webAccelReadFlagOrder(cmd))
	return cmd
}

func webAccelCertificateInfoCmd() *cobra.Command {
	webAccelCertificateInfoParam := params.NewCertificateInfoWebAccelParam()
	cmd := &cobra.Command{
		Use:          "certificate-info",
		Aliases:      []string{"cert-info"},
		Short:        "CertificateInfo WebAccel",
		Long:         `CertificateInfo WebAccel`,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := cli.NewCLIContext(globalFlags(), args, webAccelCertificateInfoParam)
			if err != nil {
				return err
			}
			if err := webAccelCertificateInfoParam.Initialize(newParamsAdapter(cmd.Flags()), args, ctx.Option()); err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if webAccelCertificateInfoParam.GenerateSkeleton {
				return generateSkeleton(ctx, webAccelCertificateInfoParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findWebAccelCertificateInfoTargets(ctx, webAccelCertificateInfoParam)
			if err != nil {
				return err
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				webAccelCertificateInfoParam.SetId(id)
				go func(p *params.CertificateInfoWebAccelParam) {
					err := funcs.WebAccelCertificateInfo(ctx, p)
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(webAccelCertificateInfoParam)
			}
			wg.Wait()
			return cli.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&webAccelCertificateInfoParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.StringVarP(&webAccelCertificateInfoParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&webAccelCertificateInfoParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&webAccelCertificateInfoParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&webAccelCertificateInfoParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&webAccelCertificateInfoParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&webAccelCertificateInfoParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&webAccelCertificateInfoParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&webAccelCertificateInfoParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&webAccelCertificateInfoParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&webAccelCertificateInfoParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&webAccelCertificateInfoParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&webAccelCertificateInfoParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &webAccelCertificateInfoParam.Id), "id", "", "Set target ID")
	fs.SetNormalizeFunc(webAccelCertificateInfoNormalizeFlagNames)
	buildFlagsUsage(cmd, webAccelCertificateInfoFlagOrder(cmd))
	return cmd
}

func webAccelCertificateNewCmd() *cobra.Command {
	webAccelCertificateNewParam := params.NewCertificateNewWebAccelParam()
	cmd := &cobra.Command{
		Use:          "certificate-new",
		Aliases:      []string{"cert-new", "cert-create", "certificate-create"},
		Short:        "CertificateNew WebAccel",
		Long:         `CertificateNew WebAccel`,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := cli.NewCLIContext(globalFlags(), args, webAccelCertificateNewParam)
			if err != nil {
				return err
			}
			if err := webAccelCertificateNewParam.Initialize(newParamsAdapter(cmd.Flags()), args, ctx.Option()); err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if webAccelCertificateNewParam.GenerateSkeleton {
				return generateSkeleton(ctx, webAccelCertificateNewParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findWebAccelCertificateNewTargets(ctx, webAccelCertificateNewParam)
			if err != nil {
				return err
			}

			// confirm
			if !webAccelCertificateNewParam.Assumeyes {
				if !util.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := util.ConfirmContinue("certificate-new", ctx.IO().In(), ctx.IO().Out(), ids...)
				if err != nil || !result {
					return err
				}
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				webAccelCertificateNewParam.SetId(id)
				go func(p *params.CertificateNewWebAccelParam) {
					err := funcs.WebAccelCertificateNew(ctx, p)
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(webAccelCertificateNewParam)
			}
			wg.Wait()
			return cli.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&webAccelCertificateNewParam.Cert, "cert", "", "", "set certificate(from file)")
	fs.StringVarP(&webAccelCertificateNewParam.Key, "key", "", "", "set private key(from file)")
	fs.StringVarP(&webAccelCertificateNewParam.CertContent, "cert-content", "", "", "set certificate(from text)")
	fs.StringVarP(&webAccelCertificateNewParam.KeyContent, "key-content", "", "", "set private key(from text)")
	fs.StringSliceVarP(&webAccelCertificateNewParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&webAccelCertificateNewParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&webAccelCertificateNewParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&webAccelCertificateNewParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&webAccelCertificateNewParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&webAccelCertificateNewParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&webAccelCertificateNewParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&webAccelCertificateNewParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&webAccelCertificateNewParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&webAccelCertificateNewParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&webAccelCertificateNewParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&webAccelCertificateNewParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&webAccelCertificateNewParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&webAccelCertificateNewParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &webAccelCertificateNewParam.Id), "id", "", "Set target ID")
	fs.SetNormalizeFunc(webAccelCertificateNewNormalizeFlagNames)
	buildFlagsUsage(cmd, webAccelCertificateNewFlagOrder(cmd))
	return cmd
}

func webAccelCertificateUpdateCmd() *cobra.Command {
	webAccelCertificateUpdateParam := params.NewCertificateUpdateWebAccelParam()
	cmd := &cobra.Command{
		Use:          "certificate-update",
		Aliases:      []string{"cert-update"},
		Short:        "CertificateUpdate WebAccel",
		Long:         `CertificateUpdate WebAccel`,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := cli.NewCLIContext(globalFlags(), args, webAccelCertificateUpdateParam)
			if err != nil {
				return err
			}
			if err := webAccelCertificateUpdateParam.Initialize(newParamsAdapter(cmd.Flags()), args, ctx.Option()); err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if webAccelCertificateUpdateParam.GenerateSkeleton {
				return generateSkeleton(ctx, webAccelCertificateUpdateParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findWebAccelCertificateUpdateTargets(ctx, webAccelCertificateUpdateParam)
			if err != nil {
				return err
			}

			// confirm
			if !webAccelCertificateUpdateParam.Assumeyes {
				if !util.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := util.ConfirmContinue("certificate-update", ctx.IO().In(), ctx.IO().Out(), ids...)
				if err != nil || !result {
					return err
				}
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				webAccelCertificateUpdateParam.SetId(id)
				go func(p *params.CertificateUpdateWebAccelParam) {
					err := funcs.WebAccelCertificateUpdate(ctx, p)
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(webAccelCertificateUpdateParam)
			}
			wg.Wait()
			return cli.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&webAccelCertificateUpdateParam.Cert, "cert", "", "", "set certificate(from file)")
	fs.StringVarP(&webAccelCertificateUpdateParam.Key, "key", "", "", "set private key(from file)")
	fs.StringVarP(&webAccelCertificateUpdateParam.CertContent, "cert-content", "", "", "set certificate(from text)")
	fs.StringVarP(&webAccelCertificateUpdateParam.KeyContent, "key-content", "", "", "set private key(from text)")
	fs.StringSliceVarP(&webAccelCertificateUpdateParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&webAccelCertificateUpdateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&webAccelCertificateUpdateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&webAccelCertificateUpdateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&webAccelCertificateUpdateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&webAccelCertificateUpdateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&webAccelCertificateUpdateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&webAccelCertificateUpdateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&webAccelCertificateUpdateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&webAccelCertificateUpdateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&webAccelCertificateUpdateParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&webAccelCertificateUpdateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&webAccelCertificateUpdateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&webAccelCertificateUpdateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &webAccelCertificateUpdateParam.Id), "id", "", "Set target ID")
	fs.SetNormalizeFunc(webAccelCertificateUpdateNormalizeFlagNames)
	buildFlagsUsage(cmd, webAccelCertificateUpdateFlagOrder(cmd))
	return cmd
}

func webAccelDeleteCacheCmd() *cobra.Command {
	webAccelDeleteCacheParam := params.NewDeleteCacheWebAccelParam()
	cmd := &cobra.Command{
		Use:          "delete-cache",
		Aliases:      []string{"purge"},
		Short:        "DeleteCache WebAccel",
		Long:         `DeleteCache WebAccel`,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := cli.NewCLIContext(globalFlags(), args, webAccelDeleteCacheParam)
			if err != nil {
				return err
			}
			if err := webAccelDeleteCacheParam.Initialize(newParamsAdapter(cmd.Flags()), args, ctx.Option()); err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if webAccelDeleteCacheParam.GenerateSkeleton {
				return generateSkeleton(ctx, webAccelDeleteCacheParam)
			}

			// confirm
			if !webAccelDeleteCacheParam.Assumeyes {
				if !util.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := util.ConfirmContinue("delete-cache", ctx.IO().In(), ctx.IO().Out())
				if err != nil || !result {
					return err
				}
			}

			return funcs.WebAccelDeleteCache(ctx, webAccelDeleteCacheParam)

		},
	}

	fs := cmd.Flags()
	fs.BoolVarP(&webAccelDeleteCacheParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&webAccelDeleteCacheParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&webAccelDeleteCacheParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&webAccelDeleteCacheParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&webAccelDeleteCacheParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&webAccelDeleteCacheParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&webAccelDeleteCacheParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&webAccelDeleteCacheParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&webAccelDeleteCacheParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&webAccelDeleteCacheParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&webAccelDeleteCacheParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&webAccelDeleteCacheParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&webAccelDeleteCacheParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.SetNormalizeFunc(webAccelDeleteCacheNormalizeFlagNames)
	buildFlagsUsage(cmd, webAccelDeleteCacheFlagOrder(cmd))
	return cmd
}
