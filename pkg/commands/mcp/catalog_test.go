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

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewCommandCatalog(t *testing.T) {
	rootCmd := &cobra.Command{
		Use:   "test",
		Short: "test command",
	}

	catalog := NewCommandCatalog(rootCmd)
	assert.NotNil(t, catalog)
	assert.Equal(t, rootCmd, catalog.rootCmd)
}

func TestCommandCatalog_ListAllCommands(t *testing.T) {
	rootCmd := &cobra.Command{
		Use:   "test",
		Short: "test root command",
	}

	subCmd1 := &cobra.Command{
		Use:   "sub1",
		Short: "sub command 1",
	}

	subCmd2 := &cobra.Command{
		Use:    "sub2",
		Short:  "sub command 2",
		Hidden: true, // 隠されたコマンド
	}

	rootCmd.AddCommand(subCmd1, subCmd2)

	catalog := NewCommandCatalog(rootCmd)
	commands, err := catalog.ListAllCommands()

	require.NoError(t, err)
	assert.Len(t, commands, 1) // 隠されたコマンドは除外される
	assert.Equal(t, "sub1", commands[0].Name)
	assert.Equal(t, "sub command 1", commands[0].Usage)
	assert.Equal(t, "usacloud sub1", commands[0].FullCommand)
}

func TestCommandCatalog_GetCommandHelp(t *testing.T) {
	rootCmd := &cobra.Command{
		Use:   "test",
		Short: "test root command",
	}

	subCmd := &cobra.Command{
		Use:   "sub",
		Short: "sub command",
		Long:  "detailed description of sub command",
	}

	rootCmd.AddCommand(subCmd)

	catalog := NewCommandCatalog(rootCmd)

	// 存在するコマンド
	info, err := catalog.GetCommandHelp("sub")
	require.NoError(t, err)
	assert.Equal(t, "sub", info.Name)
	assert.Equal(t, "sub command", info.Usage)
	assert.Equal(t, "detailed description of sub command", info.Description)
	assert.Equal(t, "usacloud sub", info.FullCommand)

	// 存在しないコマンド
	_, err = catalog.GetCommandHelp("nonexistent")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "command not found")
}

func TestCommandCatalog_findCommand(t *testing.T) {
	rootCmd := &cobra.Command{
		Use: "test",
	}

	level1Cmd := &cobra.Command{
		Use:     "level1",
		Aliases: []string{"l1"},
	}

	level2Cmd := &cobra.Command{
		Use: "level2",
	}

	level1Cmd.AddCommand(level2Cmd)
	rootCmd.AddCommand(level1Cmd)

	catalog := NewCommandCatalog(rootCmd)

	// 直接のサブコマンド
	cmd, err := catalog.findCommand([]string{"level1"})
	require.NoError(t, err)
	assert.Equal(t, "level1", cmd.Name())

	// エイリアス
	cmd, err = catalog.findCommand([]string{"l1"})
	require.NoError(t, err)
	assert.Equal(t, "level1", cmd.Name())

	// ネストしたコマンド
	cmd, err = catalog.findCommand([]string{"level1", "level2"})
	require.NoError(t, err)
	assert.Equal(t, "level2", cmd.Name())

	// 存在しないコマンド
	_, err = catalog.findCommand([]string{"nonexistent"})
	assert.Error(t, err)
}

func TestCommandCatalog_extractFlagInfo(t *testing.T) {
	cmd := &cobra.Command{
		Use: "test",
	}

	// フラグを追加
	cmd.Flags().String("name", "default", "name flag")
	cmd.Flags().StringP("output", "o", "table", "output format")
	cmd.Flags().Bool("verbose", false, "verbose output")

	catalog := NewCommandCatalog(&cobra.Command{Use: "root"})
	flags := catalog.extractFlagInfo(cmd)

	assert.Len(t, flags, 3)

	// フラグは名前順でソートされている
	assert.Equal(t, "name", flags[0].Name)
	assert.Equal(t, "output", flags[1].Name)
	assert.Equal(t, "verbose", flags[2].Name)

	// outputフラグの詳細確認
	outputFlag := flags[1]
	assert.Equal(t, "o", outputFlag.Shorthand)
	assert.Equal(t, "output format", outputFlag.Usage)
	assert.Equal(t, "table", outputFlag.DefaultValue)
	assert.Equal(t, "string", outputFlag.Type)
}

func TestCommandCatalog_buildFullCommand(t *testing.T) {
	cmd := &cobra.Command{
		Use: "test",
	}

	catalog := NewCommandCatalog(&cobra.Command{Use: "root"})

	// 親コマンドなし
	result := catalog.buildFullCommand(cmd, []string{})
	assert.Equal(t, "usacloud test", result)

	// 親コマンドあり
	result = catalog.buildFullCommand(cmd, []string{"parent1", "parent2"})
	assert.Equal(t, "usacloud parent1 parent2 test", result)
}

func TestCommandCatalog_ToJSON(t *testing.T) {
	catalog := NewCommandCatalog(&cobra.Command{})

	testData := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}

	jsonStr, err := catalog.ToJSON(testData)
	require.NoError(t, err)
	assert.Contains(t, jsonStr, "key1")
	assert.Contains(t, jsonStr, "value1")
	assert.Contains(t, jsonStr, "key2")
	assert.Contains(t, jsonStr, "value2")
}

func TestCommandCatalog_containsAlias(t *testing.T) {
	catalog := NewCommandCatalog(&cobra.Command{})

	aliases := []string{"alias1", "alias2", "alias3"}

	assert.True(t, catalog.containsAlias(aliases, "alias1"))
	assert.True(t, catalog.containsAlias(aliases, "alias2"))
	assert.True(t, catalog.containsAlias(aliases, "alias3"))
	assert.False(t, catalog.containsAlias(aliases, "nonexistent"))
	assert.False(t, catalog.containsAlias([]string{}, "any"))
}

func TestCommandCatalog_generateExamples(t *testing.T) {
	catalog := NewCommandCatalog(&cobra.Command{})

	flags := []FlagInfo{
		{Name: "zone", Hidden: false},
		{Name: "output", Hidden: false},
		{Name: "hidden-flag", Hidden: true},
	}

	examples := catalog.generateExamples("usacloud server list", flags)

	assert.Contains(t, examples, "usacloud server list")
	assert.Len(t, examples, 3) // base + zone + output examples

	// 隠されたフラグは例に含まれない
	for _, example := range examples {
		assert.NotContains(t, example, "hidden-flag")
	}
}
