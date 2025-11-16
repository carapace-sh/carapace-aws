package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appintegrations_createDataIntegrationCmd = &cobra.Command{
	Use:   "create-data-integration",
	Short: "Creates and persists a DataIntegration resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appintegrations_createDataIntegrationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appintegrations_createDataIntegrationCmd).Standalone()

		appintegrations_createDataIntegrationCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		appintegrations_createDataIntegrationCmd.Flags().String("description", "", "A description of the DataIntegration.")
		appintegrations_createDataIntegrationCmd.Flags().String("file-configuration", "", "The configuration for what files should be pulled from the source.")
		appintegrations_createDataIntegrationCmd.Flags().String("kms-key", "", "The KMS key ARN for the DataIntegration.")
		appintegrations_createDataIntegrationCmd.Flags().String("name", "", "The name of the DataIntegration.")
		appintegrations_createDataIntegrationCmd.Flags().String("object-configuration", "", "The configuration for what data should be pulled from the source.")
		appintegrations_createDataIntegrationCmd.Flags().String("schedule-config", "", "The name of the data and how often it should be pulled from the source.")
		appintegrations_createDataIntegrationCmd.Flags().String("source-uri", "", "The URI of the data source.")
		appintegrations_createDataIntegrationCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
		appintegrations_createDataIntegrationCmd.MarkFlagRequired("kms-key")
		appintegrations_createDataIntegrationCmd.MarkFlagRequired("name")
	})
	appintegrationsCmd.AddCommand(appintegrations_createDataIntegrationCmd)
}
