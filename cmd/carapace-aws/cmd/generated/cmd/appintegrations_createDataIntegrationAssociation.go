package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appintegrations_createDataIntegrationAssociationCmd = &cobra.Command{
	Use:   "create-data-integration-association",
	Short: "Creates and persists a DataIntegrationAssociation resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appintegrations_createDataIntegrationAssociationCmd).Standalone()

	appintegrations_createDataIntegrationAssociationCmd.Flags().String("client-association-metadata", "", "The mapping of metadata to be extracted from the data.")
	appintegrations_createDataIntegrationAssociationCmd.Flags().String("client-id", "", "The identifier for the client that is associated with the DataIntegration association.")
	appintegrations_createDataIntegrationAssociationCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	appintegrations_createDataIntegrationAssociationCmd.Flags().String("data-integration-identifier", "", "A unique identifier for the DataIntegration.")
	appintegrations_createDataIntegrationAssociationCmd.Flags().String("destination-uri", "", "The URI of the data destination.")
	appintegrations_createDataIntegrationAssociationCmd.Flags().String("execution-configuration", "", "The configuration for how the files should be pulled from the source.")
	appintegrations_createDataIntegrationAssociationCmd.Flags().String("object-configuration", "", "")
	appintegrations_createDataIntegrationAssociationCmd.MarkFlagRequired("data-integration-identifier")
	appintegrationsCmd.AddCommand(appintegrations_createDataIntegrationAssociationCmd)
}
