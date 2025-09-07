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
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// CommandInfo represents information about a usacloud command
type CommandInfo struct {
	Name        string        `json:"name"`
	Aliases     []string      `json:"aliases,omitempty"`
	Usage       string        `json:"usage"`
	Description string        `json:"description"`
	Args        string        `json:"args,omitempty"`
	Flags       []FlagInfo    `json:"flags,omitempty"`
	Subcommands []CommandInfo `json:"subcommands,omitempty"`
	Parents     []string      `json:"parents,omitempty"`
	FullCommand string        `json:"full_command"`
	Examples    []string      `json:"examples,omitempty"`
}

// FlagInfo represents information about a command flag
type FlagInfo struct {
	Name         string `json:"name"`
	Shorthand    string `json:"shorthand,omitempty"`
	Usage        string `json:"usage"`
	Type         string `json:"type"`
	DefaultValue string `json:"default_value,omitempty"`
	Required     bool   `json:"required"`
	Hidden       bool   `json:"hidden"`
}

// CommandCatalog manages the command information catalog
type CommandCatalog struct {
	rootCmd *cobra.Command
}

// NewCommandCatalog creates a new command catalog
func NewCommandCatalog(rootCmd *cobra.Command) *CommandCatalog {
	return &CommandCatalog{
		rootCmd: rootCmd,
	}
}

// ListAllCommands returns all commands in the catalog
func (c *CommandCatalog) ListAllCommands() ([]CommandInfo, error) {
	var commands []CommandInfo

	// Root commandから再帰的にコマンドを収集
	for _, cmd := range c.rootCmd.Commands() {
		if cmd.Hidden {
			continue
		}
		cmdInfo := c.extractCommandInfo(cmd, []string{})
		commands = append(commands, cmdInfo)
	}

	return commands, nil
}

// GetCommandHelp returns detailed help information for a specific command
func (c *CommandCatalog) GetCommandHelp(commandPath string) (*CommandInfo, error) {
	parts := strings.Fields(commandPath)
	if len(parts) == 0 {
		return nil, fmt.Errorf("empty command path")
	}

	cmd, err := c.findCommand(parts)
	if err != nil {
		return nil, err
	}

	parents := make([]string, len(parts)-1)
	copy(parents, parts[:len(parts)-1])

	cmdInfo := c.extractCommandInfo(cmd, parents)
	return &cmdInfo, nil
}

// findCommand finds a command by path
func (c *CommandCatalog) findCommand(path []string) (*cobra.Command, error) {
	current := c.rootCmd

	for i, part := range path {
		found := false
		for _, cmd := range current.Commands() {
			if cmd.Name() == part || c.containsAlias(cmd.Aliases, part) {
				current = cmd
				found = true
				break
			}
		}
		if !found {
			return nil, fmt.Errorf("command not found: %s (at position %d in path %v)",
				part, i+1, path)
		}
	}

	return current, nil
}

// containsAlias checks if aliases contain the given name
func (c *CommandCatalog) containsAlias(aliases []string, name string) bool {
	for _, alias := range aliases {
		if alias == name {
			return true
		}
	}
	return false
}

// extractCommandInfo extracts command information from a cobra command
func (c *CommandCatalog) extractCommandInfo(cmd *cobra.Command, parents []string) CommandInfo {
	info := CommandInfo{
		Name:        cmd.Name(),
		Aliases:     cmd.Aliases,
		Usage:       cmd.Short,
		Description: cmd.Long,
		Args:        "",
		Parents:     parents,
		FullCommand: c.buildFullCommand(cmd, parents),
	}

	// Argsの情報を文字列として設定
	if use := cmd.Use; use != "" {
		if spaceIdx := strings.Index(use, " "); spaceIdx != -1 {
			info.Args = strings.TrimSpace(use[spaceIdx+1:])
		}
	}

	// フラグ情報を収集
	info.Flags = c.extractFlagInfo(cmd)

	// サブコマンド情報を収集
	for _, subCmd := range cmd.Commands() {
		if subCmd.Hidden {
			continue
		}
		parents = append(parents, cmd.Name())
		subInfo := c.extractCommandInfo(subCmd, parents)
		info.Subcommands = append(info.Subcommands, subInfo)
	}

	// 使用例を生成
	info.Examples = c.generateExamples(info.FullCommand, info.Flags)

	return info
}

// extractFlagInfo extracts flag information from a cobra command
func (c *CommandCatalog) extractFlagInfo(cmd *cobra.Command) []FlagInfo {
	var flags []FlagInfo

	// Local flags
	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		flagInfo := FlagInfo{
			Name:         flag.Name,
			Shorthand:    flag.Shorthand,
			Usage:        flag.Usage,
			Type:         flag.Value.Type(),
			DefaultValue: flag.DefValue,
			Required:     c.isFlagRequired(flag),
			Hidden:       flag.Hidden,
		}
		flags = append(flags, flagInfo)
	})

	// Inherited flags (global flags)
	cmd.InheritedFlags().VisitAll(func(flag *pflag.Flag) {
		// MCPコマンドの場合は継承フラグをスキップ
		if cmd.Name() == "mcp" {
			return
		}
		flagInfo := FlagInfo{
			Name:         flag.Name,
			Shorthand:    flag.Shorthand,
			Usage:        flag.Usage,
			Type:         flag.Value.Type(),
			DefaultValue: flag.DefValue,
			Required:     c.isFlagRequired(flag),
			Hidden:       flag.Hidden,
		}
		flags = append(flags, flagInfo)
	})

	// フラグを名前順でソート
	sort.Slice(flags, func(i, j int) bool {
		return flags[i].Name < flags[j].Name
	})

	return flags
}

// isFlagRequired checks if a flag is required (simplified heuristic)
func (c *CommandCatalog) isFlagRequired(flag *pflag.Flag) bool {
	// pflagからrequired情報を取得する方法が限定的なので、
	// ここでは簡易的な判定を行う
	return false
}

// buildFullCommand builds the full command string
func (c *CommandCatalog) buildFullCommand(cmd *cobra.Command, parents []string) string {
	parts := make([]string, len(parents)+2) // usacloud + parents + current command
	parts[0] = "usacloud"
	copy(parts[1:], parents)
	parts[len(parts)-1] = cmd.Name()
	return strings.Join(parts, " ")
}

// generateExamples generates example commands
func (c *CommandCatalog) generateExamples(baseCommand string, flags []FlagInfo) []string {
	var examples []string

	// 基本的な使用例
	examples = append(examples, baseCommand)

	// 主要なフラグを含む例
	var commonFlags []string
	for _, flag := range flags {
		if flag.Hidden {
			continue
		}
		// よく使われるフラグの例を生成
		switch flag.Name {
		case "zone":
			commonFlags = append(commonFlags, "--zone is1a")
		case "output", "format":
			commonFlags = append(commonFlags, "--output json")
		}
	}

	if len(commonFlags) > 0 {
		for _, flagExample := range commonFlags[:min(2, len(commonFlags))] {
			examples = append(examples, baseCommand+" "+flagExample)
		}
	}

	return examples
}

// min returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// HasOutputTypeFlag checks if the command supports --output-type flag
func (c *CommandCatalog) HasOutputTypeFlag(commandPath string) bool {
	cmdInfo, err := c.GetCommandHelp(commandPath)
	if err != nil {
		return false
	}

	// フラグ情報から output-type, out, o のいずれかが存在するかチェック
	for _, flag := range cmdInfo.Flags {
		if flag.Name == "output-type" || flag.Name == "out" || flag.Name == "o" {
			return true
		}
	}

	return false
}

// ToJSON converts command info to JSON string
func (c *CommandCatalog) ToJSON(commands interface{}) (string, error) {
	data, err := json.MarshalIndent(commands, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal to JSON: %v", err)
	}
	return string(data), nil
}
