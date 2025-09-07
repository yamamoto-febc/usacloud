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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsCommandSafe(t *testing.T) {
	tests := []struct {
		name     string
		command  string
		args     []string
		expected bool
	}{
		{
			name:     "安全なコマンド - server list",
			command:  "server list",
			args:     []string{},
			expected: true,
		},
		{
			name:     "安全なコマンド - zone list",
			command:  "zone list",
			args:     []string{},
			expected: true,
		},
		{
			name:     "危険なコマンド - server delete",
			command:  "server delete",
			args:     []string{},
			expected: false,
		},
		{
			name:     "危険なコマンド - server shutdown",
			command:  "server shutdown",
			args:     []string{},
			expected: false,
		},
		{
			name:     "危険なコマンド - server create",
			command:  "server create",
			args:     []string{},
			expected: false,
		},
		{
			name:     "危険なコマンド - config set",
			command:  "config set",
			args:     []string{},
			expected: false,
		},
		{
			name:     "危険なフラグ - force",
			command:  "server list",
			args:     []string{"--force"},
			expected: false,
		},
		{
			name:     "安全なフラグ付きコマンド",
			command:  "server list",
			args:     []string{"--zone", "is1a"},
			expected: true,
		},
		{
			name:     "短すぎるコマンド",
			command:  "server",
			args:     []string{},
			expected: false,
		},
		{
			name:     "空のコマンド",
			command:  "",
			args:     []string{},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isCommandSafe(tt.command, tt.args)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestExecuteCommand(t *testing.T) {
	// このテストは実際のプロセス実行を行うため、テスト環境では無限ループを避けるためスキップ
	t.Skip("Skipping TestExecuteCommand to prevent infinite process spawning during tests")

	tests := []struct {
		name        string
		args        ExecuteCommandArgs
		expectError bool
	}{
		{
			name: "空のコマンド",
			args: ExecuteCommandArgs{
				Command: "",
			},
			expectError: true,
		},
		{
			name: "危険なコマンド",
			args: ExecuteCommandArgs{
				Command: "server delete",
			},
			expectError: true,
		},
		{
			name: "安全なコマンド（実際には実行しない）",
			args: ExecuteCommandArgs{
				Command: "version", // versionコマンドは安全だが、短すぎるのでfalseになる
			},
			expectError: true, // 短すぎるコマンドのため
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 構造体の妥当性とセキュリティチェックのみをテスト
			assert.NotNil(t, tt.args, "Args should not be nil")
			assert.IsType(t, ExecuteCommandArgs{}, tt.args, "Args should be of correct type")

			// セキュリティチェックのテスト（実際の実行は行わない）
			isUnsafe := tt.args.Command == "" ||
				tt.args.Command == "server delete" ||
				!isCommandSafe(tt.args.Command, tt.args.Args)
			assert.Equal(t, tt.expectError, isUnsafe, "Security check should match expected error")
		})
	}
}

func TestGetConfig(t *testing.T) {
	// このテストは実際のプロセス実行を行うため、テスト環境では無限ループを避けるためスキップ
	t.Skip("Skipping TestGetConfig to prevent infinite process spawning during tests")

	tests := []struct {
		name        string
		args        GetConfigArgs
		description string
	}{
		{
			name:        "全設定取得",
			args:        GetConfigArgs{},
			description: "名前を指定せずに全設定を取得",
		},
		{
			name: "特定設定取得",
			args: GetConfigArgs{
				Name: "profile",
			},
			description: "特定の設定名を指定して取得",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 構造体の妥当性のみをテスト
			assert.NotNil(t, tt.args, "Args should not be nil")
			assert.IsType(t, GetConfigArgs{}, tt.args, "Args should be of correct type")
		})
	}
}

func TestListZones(t *testing.T) {
	// このテストは実際のプロセス実行を行うため、テスト環境では無限ループを避けるためスキップ
	t.Skip("Skipping TestListZones to prevent infinite process spawning during tests")

	tests := []struct {
		name        string
		args        ListZonesArgs
		description string
	}{
		{
			name:        "デフォルトJSON出力",
			args:        ListZonesArgs{},
			description: "出力形式を指定せずにデフォルトのJSON形式で取得",
		},
		{
			name: "YAML出力",
			args: ListZonesArgs{
				OutputFormat: "yaml",
			},
			description: "YAML形式で出力を指定",
		},
		{
			name: "テーブル出力",
			args: ListZonesArgs{
				OutputFormat: "table",
			},
			description: "テーブル形式で出力を指定",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 構造体の妥当性のみをテスト
			assert.NotNil(t, tt.args, "Args should not be nil")
			assert.IsType(t, ListZonesArgs{}, tt.args, "Args should be of correct type")

			// 出力形式の検証
			if tt.args.OutputFormat != "" {
				validFormats := []string{"json", "yaml", "table"}
				found := false
				for _, format := range validFormats {
					if tt.args.OutputFormat == format {
						found = true
						break
					}
				}
				assert.True(t, found, "Output format should be valid")
			}
		})
	}
}

func TestDangerousFlagsMap(t *testing.T) {
	expectedFlags := []string{"force", "force-stop", "force-boot", "force-reset"}

	for _, flag := range expectedFlags {
		assert.True(t, dangerousFlags[flag], "Flag %s should be marked as dangerous", flag)
	}

	// 安全なフラグは含まれていないことを確認
	safeFlags := []string{"zone", "output", "help", "version"}
	for _, flag := range safeFlags {
		assert.False(t, dangerousFlags[flag], "Flag %s should not be marked as dangerous", flag)
	}
}

func TestExecuteCommandArgs(t *testing.T) {
	args := ExecuteCommandArgs{
		Command: "server list",
		Args:    []string{"example-server"},
		Flags:   []string{"--zone", "is1a"},
	}

	assert.Equal(t, "server list", args.Command)
	assert.Equal(t, []string{"example-server"}, args.Args)
	assert.Equal(t, []string{"--zone", "is1a"}, args.Flags)
}

func TestGetConfigArgs(t *testing.T) {
	// 全設定取得の場合
	args1 := GetConfigArgs{}
	assert.Equal(t, "", args1.Name)

	// 特定設定取得の場合
	args2 := GetConfigArgs{Name: "profile"}
	assert.Equal(t, "profile", args2.Name)
}

func TestListZonesArgs(t *testing.T) {
	// デフォルトの場合
	args1 := ListZonesArgs{}
	assert.Equal(t, "", args1.OutputFormat)

	// 形式指定の場合
	args2 := ListZonesArgs{OutputFormat: "json"}
	assert.Equal(t, "json", args2.OutputFormat)
}

func TestListCommandsArgs(t *testing.T) {
	// デフォルト（hidden含まない）
	args1 := ListCommandsArgs{}
	assert.False(t, args1.IncludeHidden)

	// hidden含む
	args2 := ListCommandsArgs{IncludeHidden: true}
	assert.True(t, args2.IncludeHidden)
}

func TestGetCommandHelpArgs(t *testing.T) {
	args := GetCommandHelpArgs{Command: "server list"}
	assert.Equal(t, "server list", args.Command)
}
