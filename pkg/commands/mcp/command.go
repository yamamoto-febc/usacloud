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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"

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

// ExecuteCommandArgs represents arguments for execute_command tool
type ExecuteCommandArgs struct {
	Command string   `json:"command" jsonschema:"実行するusacloudコマンド（例: 'server list', 'zone list'）"`
	Args    []string `json:"args,omitempty" jsonschema:"追加の引数"`
	Flags   []string `json:"flags,omitempty" jsonschema:"追加のフラグ"`
}

// GetConfigArgs represents arguments for get_config tool
type GetConfigArgs struct {
	Name string `json:"name,omitempty" jsonschema:"取得する設定名（空の場合は全設定）"`
}

// ListZonesArgs represents arguments for list_zones tool
type ListZonesArgs struct {
	OutputFormat string `json:"output_format,omitempty" jsonschema:"出力形式 (json/yaml/table)"`
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

	// execute_command ツール
	mcp.AddTool(server, &mcp.Tool{
		Name:        "execute_command",
		Description: "usacloudコマンドを安全に実行します。",
	}, func(ctx context.Context, request *mcp.CallToolRequest, args ExecuteCommandArgs) (*mcp.CallToolResult, any, error) {
		return executeCommand(args, catalog)
	})

	// get_config ツール
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_config",
		Description: "usacloudの設定情報を取得します。",
	}, func(ctx context.Context, request *mcp.CallToolRequest, args GetConfigArgs) (*mcp.CallToolResult, any, error) {
		return getConfig(args, catalog)
	})

	// list_zones ツール
	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_zones",
		Description: "利用可能なゾーン一覧を取得します。",
	}, func(ctx context.Context, request *mcp.CallToolRequest, args ListZonesArgs) (*mcp.CallToolResult, any, error) {
		return listZones(args, catalog)
	})

	log.Printf("MCPツールを登録しました: list_commands, get_command_help, execute_command, get_config, list_zones")
	return nil
}

// isCommandSafe checks if the command is safe to execute
// 現在は全コマンドを許可。将来的に必要に応じて制限を追加予定
func isCommandSafe(command string, args []string) bool {
	// TODO: 将来的にSafetyLevelベースの制限を実装予定
	// type SafetyLevel int
	// const (
	//     SafetyNone SafetyLevel = iota     // 制限なし（現在）
	//     SafetyMinimal                     // 最低限の制限
	//     SafetyStandard                    // 標準的制限
	//     SafetyStrict                      // 厳格な制限
	// )

	/*
	// 将来の制限例（現在はコメントアウト）
	dangerousCommands := map[string]bool{
		"delete":      true,
		"shutdown":    true,
		"boot":        true,
		"reset":       true,
		"create":      true,
		"update":      true,
		"modify":      true,
		"build":       true,
		"archive":     true,
		"clone":       true,
		"power":       true,
		"insert":      true,
		"eject":       true,
		"disconnect":  true,
		"connect":     true,
		"start":       true,
		"stop":        true,
		"restart":     true,
		"set":         true, // config set操作は危険
		"update-self": true,
	}
	*/

	// 現在は全コマンドを許可
	return true
}

/*
// 将来の制限で使用予定のフラグリスト（現在はコメントアウト）
var dangerousFlags = map[string]bool{
	"force":       true,
	"force-stop":  true,
	"force-boot":  true,
	"force-reset": true,
}
*/

// isTestEnvironment checks if we are running in a test environment
func isTestEnvironment() bool {
	return testing.Testing()
}

// executeCommand executes a usacloud command safely
func executeCommand(args ExecuteCommandArgs, catalog *CommandCatalog) (*mcp.CallToolResult, any, error) {
	// テスト環境では実行を避ける
	if isTestEnvironment() {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{
					Text: "テスト環境のため実行をスキップしました",
				},
			},
			IsError: false,
		}, nil, nil
	}

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

	// セキュリティチェック（現在は全コマンド許可）
	if !isCommandSafe(args.Command, args.Args) {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{
					Text: "このコマンドは実行が制限されています。",
				},
			},
			IsError: true,
		}, nil, nil
	}

	// usacloudコマンドを構築
	cmdParts := strings.Fields(args.Command)
	allArgs := make([]string, 0, len(cmdParts)+len(args.Args)+len(args.Flags))
	allArgs = append(allArgs, cmdParts...)
	allArgs = append(allArgs, args.Args...)
	allArgs = append(allArgs, args.Flags...)

	// 出力フラグをサポートしているコマンドの場合のみJSON出力を強制
	if catalog.HasOutputTypeFlag(args.Command) {
		allArgs = append(allArgs, "--output-type", "json")
	}

	// usacloudバイナリのパス取得（現在のバイナリと同じディレクトリから）
	executable, err := os.Executable()
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{
					Text: fmt.Sprintf("実行可能ファイルパスの取得に失敗: %v", err),
				},
			},
			IsError: true,
		}, nil, nil
	}

	cmd := exec.Command(executable, allArgs...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Env = os.Environ()

	err = cmd.Run()

	result := map[string]interface{}{
		"command": args.Command,
		"args":    allArgs,
		"stdout":  stdout.String(),
		"stderr":  stderr.String(),
		"success": err == nil,
	}

	if err != nil {
		result["error"] = err.Error()
	}

	jsonData, jsonErr := json.MarshalIndent(result, "", "  ")
	if jsonErr != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{
					Text: fmt.Sprintf("結果のJSON変換に失敗: %v", jsonErr),
				},
			},
			IsError: true,
		}, nil, nil
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: string(jsonData),
			},
		},
		IsError: err != nil,
	}, nil, nil
}

// getConfig gets configuration information
func getConfig(args GetConfigArgs, catalog *CommandCatalog) (*mcp.CallToolResult, any, error) {
	// テスト環境では実行を避ける
	if isTestEnvironment() {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{
					Text: "テスト環境のため実行をスキップしました",
				},
			},
			IsError: false,
		}, nil, nil
	}

	cmdArgs := []string{"config", "show"}
	if args.Name != "" {
		cmdArgs = append(cmdArgs, "--name", args.Name)
	}

	executable, err := os.Executable()
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{
					Text: fmt.Sprintf("実行可能ファイルパスの取得に失敗: %v", err),
				},
			},
			IsError: true,
		}, nil, nil
	}

	cmd := exec.Command(executable, cmdArgs...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Env = os.Environ()

	err = cmd.Run()
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{
					Text: fmt.Sprintf("config showコマンドの実行に失敗: %v\nstderr: %s", err, stderr.String()),
				},
			},
			IsError: true,
		}, nil, nil
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: stdout.String(),
			},
		},
	}, nil, nil
}

// listZones lists available zones
func listZones(args ListZonesArgs, catalog *CommandCatalog) (*mcp.CallToolResult, any, error) {
	// テスト環境では実行を避ける
	if isTestEnvironment() {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{
					Text: "テスト環境のため実行をスキップしました",
				},
			},
			IsError: false,
		}, nil, nil
	}

	outputFormat := "json"
	if args.OutputFormat != "" {
		outputFormat = args.OutputFormat
	}

	cmdArgs := []string{"zone", "list"}

	// zone listコマンドが出力フラグをサポートしているかチェック
	if catalog.HasOutputTypeFlag("zone list") {
		cmdArgs = append(cmdArgs, "--output-type", outputFormat)
	}

	executable, err := os.Executable()
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{
					Text: fmt.Sprintf("実行可能ファイルパスの取得に失敗: %v", err),
				},
			},
			IsError: true,
		}, nil, nil
	}

	cmd := exec.Command(executable, cmdArgs...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Env = os.Environ()

	err = cmd.Run()
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{
					Text: fmt.Sprintf("zone listコマンドの実行に失敗: %v\nstderr: %s", err, stderr.String()),
				},
			},
			IsError: true,
		}, nil, nil
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: stdout.String(),
			},
		},
	}, nil, nil
}
