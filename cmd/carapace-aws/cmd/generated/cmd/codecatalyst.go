package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecatalystCmd = &cobra.Command{
	Use:   "codecatalyst",
	Short: "Welcome to the Amazon CodeCatalyst API reference.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecatalystCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecatalystCmd).Standalone()

	})
	rootCmd.AddCommand(codecatalystCmd)
}
