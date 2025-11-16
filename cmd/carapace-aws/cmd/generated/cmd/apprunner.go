package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apprunnerCmd = &cobra.Command{
	Use:   "apprunner",
	Short: "App Runner\n\nApp Runner is an application service that provides a fast, simple, and cost-effective way to go directly from an existing container image or source code to a running service in the Amazon Web Services Cloud in seconds.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apprunnerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apprunnerCmd).Standalone()

	})
	rootCmd.AddCommand(apprunnerCmd)
}
