package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	userLicense string

	rootCmd = &cobra.Command{
		Use:              "golangspell-redis",
		Short:            "golangspell-redis: [Add your Spell's short description here]",
		Long:             `golangspell-redis - [Add your Spell's long description here]`,
		TraverseChildren: true,
	}
)

// Execute executes the root command.
func Execute() error {
	addInnerCommands()
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringP("author", "a", "", "author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "Apache", "name of license for the project")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
}

func initConfig() {
}
