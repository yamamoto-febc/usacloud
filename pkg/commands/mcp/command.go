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
	"github.com/sacloud/usacloud/pkg/commands/root"
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
	// コマンドカタログを作成
	catalog := NewCommandCatalog(root.Command)

	// MCPサーバを作成
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "usacloud",
		Version: "1.0.0",
	}, nil)

	// ツールを登録
	if err := registerTools(server, catalog); err != nil {
		return fmt.Errorf("ツールの登録に失敗しました: %v", err)
	}

	// コンテキストを作成
	ctx := context.Background()

	log.Printf("usacloud MCPサーバを起動します...")
	log.Printf("stdio transport経由でMCP接続を待機中...")

	// stdio transportを使用してサーバを起動
	if err := server.Run(ctx, &mcp.StdioTransport{}); err != nil {
		log.Printf("MCPサーバの実行中にエラーが発生しました: %v", err)
		return err
	}

	log.Printf("MCPサーバが正常に終了しました")

	return nil
}

// ListCommandsArgs represents arguments for list_commands tool
type ListCommandsArgs struct {
	IncludeHidden bool `json:"include_hidden" jsonschema:"隠されたコマンドも含めるかどうか"`
}

// GetCommandHelpArgs represents arguments for get_command_help tool
type GetCommandHelpArgs struct {
	Command string `json:"command" jsonschema:"ヘルプを取得したいコマンド（例: 'server list', 'disk create'）"`
}

// registerTools registers MCP tools
func registerTools(server *mcp.Server, catalog *CommandCatalog) error {
	// list_commands ツール
	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_commands",
		Description: "usacloudの利用可能な全コマンドの一覧を取得します。各コマンドの基本情報（名前、説明、使用法）を含みます。",
	}, func(ctx context.Context, request *mcp.CallToolRequest, args ListCommandsArgs) (*mcp.CallToolResult, any, error) {
		commands, err := catalog.ListAllCommands()
		if err != nil {
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					&mcp.TextContent{
						Text: fmt.Sprintf("コマンド一覧の取得に失敗しました: %v", err),
					},
				},
				IsError: true,
			}, nil, nil
		}

		jsonData, err := catalog.ToJSON(commands)
		if err != nil {
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					&mcp.TextContent{
						Text: fmt.Sprintf("JSON変換に失敗しました: %v", err),
					},
				},
				IsError: true,
			}, nil, nil
		}

		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{
					Text: jsonData,
				},
			},
		}, nil, nil
	})

	// get_command_help ツール
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_command_help",
		Description: "特定のコマンドの詳細なヘルプ情報を取得します。フラグ、引数、使用例などの詳細情報を含みます。",
	}, func(ctx context.Context, request *mcp.CallToolRequest, args GetCommandHelpArgs) (*mcp.CallToolResult, any, error) {
		if args.Command == "" {
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					&mcp.TextContent{
						Text: "commandパラメータが必要です",
					},
				},
				IsError: true,
			}, nil, nil
		}

		commandInfo, err := catalog.GetCommandHelp(args.Command)
		if err != nil {
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					&mcp.TextContent{
						Text: fmt.Sprintf("コマンドヘルプの取得に失敗しました: %v", err),
					},
				},
				IsError: true,
			}, nil, nil
		}

		jsonData, err := catalog.ToJSON(commandInfo)
		if err != nil {
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					&mcp.TextContent{
						Text: fmt.Sprintf("JSON変換に失敗しました: %v", err),
					},
				},
				IsError: true,
			}, nil, nil
		}

		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{
					Text: jsonData,
				},
			},
		}, nil, nil
	})

	log.Printf("MCPツールを登録しました: list_commands, get_command_help")
	return nil
}
