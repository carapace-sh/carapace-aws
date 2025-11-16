package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entityresolutionCmd = &cobra.Command{
	Use:   "entityresolution",
	Short: "Welcome to the *Entity Resolution API Reference*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entityresolutionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(entityresolutionCmd).Standalone()

	})
	rootCmd.AddCommand(entityresolutionCmd)
}
