package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amp_putAlertManagerDefinitionCmd = &cobra.Command{
	Use:   "put-alert-manager-definition",
	Short: "Updates an existing alert manager definition in a workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amp_putAlertManagerDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amp_putAlertManagerDefinitionCmd).Standalone()

		amp_putAlertManagerDefinitionCmd.Flags().String("client-token", "", "A unique identifier that you can provide to ensure the idempotency of the request.")
		amp_putAlertManagerDefinitionCmd.Flags().String("data", "", "The alert manager definition to use.")
		amp_putAlertManagerDefinitionCmd.Flags().String("workspace-id", "", "The ID of the workspace to update the alert manager definition in.")
		amp_putAlertManagerDefinitionCmd.MarkFlagRequired("data")
		amp_putAlertManagerDefinitionCmd.MarkFlagRequired("workspace-id")
	})
	ampCmd.AddCommand(amp_putAlertManagerDefinitionCmd)
}
