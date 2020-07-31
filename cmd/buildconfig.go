package cmd

import (
	"fmt"

	"encoding/json"

	"github.com/golangspell/golangspell/domain"
	"github.com/spf13/cobra"
)

func init() {
	RunCommandFunctions["build-config"] = runBuildConfigCommand
}

func runBuildConfigCommand(cmd *cobra.Command, args []string) {
	configBytes, err := json.MarshalIndent(buildSpellConfig(), "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(configBytes))
}

func buildSpellConfig() domain.Spell {
	return domain.Spell{
		Name: "golangspell-redis",
		URL:  "github.com/golangspell/golangspell-redis",
		Commands: map[string]*domain.Command{
			"build-config": &domain.Command{
				Name:             "build-config",
				ShortDescription: "Builds the config necessary for adding this plugin to the Golang Spell tool",
				LongDescription: `Builds the config necessary for adding this plugin to the Golang Spell tool.
This command must be available in all Golang Spell plugins to make it possible the plugin addition to the platform.

Syntax: 
golangspell build-config
`,
			},
			"redisinit": &domain.Command{
				Name:             "redisinit",
				ShortDescription: "The redisinit command adds the redis cache and locking features to the project",
				LongDescription: `The redisinit command adds the redis cache and locking features to the project
Args:
No arguments required

Syntax: 
golangspell redisinit
`,
			},
		},
	}
}
