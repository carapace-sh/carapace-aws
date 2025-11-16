package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_describeWebAppCmd = &cobra.Command{
	Use:   "describe-web-app",
	Short: "Describes the web app that's identified by `WebAppId`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_describeWebAppCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(transfer_describeWebAppCmd).Standalone()

		transfer_describeWebAppCmd.Flags().String("web-app-id", "", "Provide the unique identifier for the web app.")
		transfer_describeWebAppCmd.MarkFlagRequired("web-app-id")
	})
	transferCmd.AddCommand(transfer_describeWebAppCmd)
}
