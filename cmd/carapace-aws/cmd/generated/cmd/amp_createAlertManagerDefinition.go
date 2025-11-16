package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amp_createAlertManagerDefinitionCmd = &cobra.Command{
	Use:   "create-alert-manager-definition",
	Short: "The `CreateAlertManagerDefinition` operation creates the alert manager definition in a workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amp_createAlertManagerDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amp_createAlertManagerDefinitionCmd).Standalone()

		amp_createAlertManagerDefinitionCmd.Flags().String("client-token", "", "A unique identifier that you can provide to ensure the idempotency of the request.")
		amp_createAlertManagerDefinitionCmd.Flags().String("data", "", "The alert manager definition to add.")
		amp_createAlertManagerDefinitionCmd.Flags().String("workspace-id", "", "The ID of the workspace to add the alert manager definition to.")
		amp_createAlertManagerDefinitionCmd.MarkFlagRequired("data")
		amp_createAlertManagerDefinitionCmd.MarkFlagRequired("workspace-id")
	})
	ampCmd.AddCommand(amp_createAlertManagerDefinitionCmd)
}
