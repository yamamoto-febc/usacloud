// Copyright 2017-2025 The sacloud/usacloud Authors
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

package mcp

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var Command = &cobra.Command{
	Use:   "mcp",
	Short: "Model Context Protocol サーバを起動",
	Long: `Model Context Protocol (MCP) サーバを起動します。

MCPは、言語モデルがアプリケーション固有のツール、データソース、
プロンプトテンプレートにアクセスできるようにするための統一プロトコルです。

このコマンドは標準入出力を使用してJSON-RPC形式で通信を行います。`,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		if err := runMCPServer(); err != nil {
			fmt.Fprintf(os.Stderr, "MCPサーバでエラーが発生しました: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	// MCPコマンドは単体で動作するため、不要なグローバルフラグを非表示にする
	// 継承されたPersistentフラグをすべて隠す
	Command.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		// 継承されたPersistentFlagsを一時的に隠す
		cmd.InheritedFlags().VisitAll(func(flag *pflag.Flag) {
			cmd.InheritedFlags().MarkHidden(flag.Name) //nolint:errcheck
		})
		// デフォルトのヘルプを表示
		cmd.Parent().HelpFunc()(cmd, args)
	})
}

func runMCPServer() error {
	// MCPサーバを作成
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "usacloud",
		Version: "1.0.0",
	}, nil)

	// この時点では何もツールを提供しない（起動のみ）
	// 将来的にはusacloudの各機能をツールとして公開する

	// コンテキストを作成
	ctx := context.Background()

	// stdio transportを使用してサーバを起動
	if err := server.Run(ctx, &mcp.StdioTransport{}); err != nil {
		log.Printf("MCPサーバの実行中にエラーが発生しました: %v", err)
		return err
	}

	return nil
}
