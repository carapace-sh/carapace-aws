package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appintegrations_updateDataIntegrationAssociationCmd = &cobra.Command{
	Use:   "update-data-integration-association",
	Short: "Updates and persists a DataIntegrationAssociation resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appintegrations_updateDataIntegrationAssociationCmd).Standalone()

	appintegrations_updateDataIntegrationAssociationCmd.Flags().String("data-integration-association-identifier", "", "A unique identifier.")
	appintegrations_updateDataIntegrationAssociationCmd.Flags().String("data-integration-identifier", "", "A unique identifier for the DataIntegration.")
	appintegrations_updateDataIntegrationAssociationCmd.Flags().String("execution-configuration", "", "The configuration for how the files should be pulled from the source.")
	appintegrations_updateDataIntegrationAssociationCmd.MarkFlagRequired("data-integration-association-identifier")
	appintegrations_updateDataIntegrationAssociationCmd.MarkFlagRequired("data-integration-identifier")
	appintegrations_updateDataIntegrationAssociationCmd.MarkFlagRequired("execution-configuration")
	appintegrationsCmd.AddCommand(appintegrations_updateDataIntegrationAssociationCmd)
}
