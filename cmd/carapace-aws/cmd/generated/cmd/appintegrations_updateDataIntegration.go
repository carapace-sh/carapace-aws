package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appintegrations_updateDataIntegrationCmd = &cobra.Command{
	Use:   "update-data-integration",
	Short: "Updates the description of a DataIntegration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appintegrations_updateDataIntegrationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appintegrations_updateDataIntegrationCmd).Standalone()

		appintegrations_updateDataIntegrationCmd.Flags().String("description", "", "A description of the DataIntegration.")
		appintegrations_updateDataIntegrationCmd.Flags().String("identifier", "", "A unique identifier for the DataIntegration.")
		appintegrations_updateDataIntegrationCmd.Flags().String("name", "", "The name of the DataIntegration.")
		appintegrations_updateDataIntegrationCmd.MarkFlagRequired("identifier")
	})
	appintegrationsCmd.AddCommand(appintegrations_updateDataIntegrationCmd)
}
