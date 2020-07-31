package cmd

import (
	"fmt"

	"github.com/golangspell/golangspell-redis/config"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "golangspell-redis-version",
	Short: "golangspell-redis version number",
	Long:  `Shows the golangspell-redis current installed version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("golangspell-redis v%s -- HEAD\n", config.Version)
	},
}
