package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	_ = viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	_ = viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("license", "Apache")
}

func initConfig() {
	viper.AutomaticEnv()
}
