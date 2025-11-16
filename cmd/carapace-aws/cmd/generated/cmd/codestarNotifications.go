package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codestarNotificationsCmd = &cobra.Command{
	Use:   "codestar-notifications",
	Short: "This CodeStar Notifications API Reference provides descriptions and usage examples of the operations and data types for the CodeStar Notifications API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codestarNotificationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codestarNotificationsCmd).Standalone()

	})
	rootCmd.AddCommand(codestarNotificationsCmd)
}
