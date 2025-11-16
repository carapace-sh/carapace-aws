package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_batchAssociateServiceActionWithProvisioningArtifactCmd = &cobra.Command{
	Use:   "batch-associate-service-action-with-provisioning-artifact",
	Short: "Associates multiple self-service actions with provisioning artifacts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_batchAssociateServiceActionWithProvisioningArtifactCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_batchAssociateServiceActionWithProvisioningArtifactCmd).Standalone()

		servicecatalog_batchAssociateServiceActionWithProvisioningArtifactCmd.Flags().String("accept-language", "", "The language code.")
		servicecatalog_batchAssociateServiceActionWithProvisioningArtifactCmd.Flags().String("service-action-associations", "", "One or more associations, each consisting of the Action ID, the Product ID, and the Provisioning Artifact ID.")
		servicecatalog_batchAssociateServiceActionWithProvisioningArtifactCmd.MarkFlagRequired("service-action-associations")
	})
	servicecatalogCmd.AddCommand(servicecatalog_batchAssociateServiceActionWithProvisioningArtifactCmd)
}
