package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appintegrations_deleteDataIntegrationCmd = &cobra.Command{
	Use:   "delete-data-integration",
	Short: "Deletes the DataIntegration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appintegrations_deleteDataIntegrationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appintegrations_deleteDataIntegrationCmd).Standalone()

		appintegrations_deleteDataIntegrationCmd.Flags().String("data-integration-identifier", "", "A unique identifier for the DataIntegration.")
		appintegrations_deleteDataIntegrationCmd.MarkFlagRequired("data-integration-identifier")
	})
	appintegrationsCmd.AddCommand(appintegrations_deleteDataIntegrationCmd)
}
