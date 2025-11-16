package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amp_deleteAlertManagerDefinitionCmd = &cobra.Command{
	Use:   "delete-alert-manager-definition",
	Short: "Deletes the alert manager definition from a workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amp_deleteAlertManagerDefinitionCmd).Standalone()

	amp_deleteAlertManagerDefinitionCmd.Flags().String("client-token", "", "A unique identifier that you can provide to ensure the idempotency of the request.")
	amp_deleteAlertManagerDefinitionCmd.Flags().String("workspace-id", "", "The ID of the workspace to delete the alert manager definition from.")
	amp_deleteAlertManagerDefinitionCmd.MarkFlagRequired("workspace-id")
	ampCmd.AddCommand(amp_deleteAlertManagerDefinitionCmd)
}
