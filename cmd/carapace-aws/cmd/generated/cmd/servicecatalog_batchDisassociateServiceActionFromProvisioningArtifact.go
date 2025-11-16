package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_batchDisassociateServiceActionFromProvisioningArtifactCmd = &cobra.Command{
	Use:   "batch-disassociate-service-action-from-provisioning-artifact",
	Short: "Disassociates a batch of self-service actions from the specified provisioning artifact.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_batchDisassociateServiceActionFromProvisioningArtifactCmd).Standalone()

	servicecatalog_batchDisassociateServiceActionFromProvisioningArtifactCmd.Flags().String("accept-language", "", "The language code.")
	servicecatalog_batchDisassociateServiceActionFromProvisioningArtifactCmd.Flags().String("service-action-associations", "", "One or more associations, each consisting of the Action ID, the Product ID, and the Provisioning Artifact ID.")
	servicecatalog_batchDisassociateServiceActionFromProvisioningArtifactCmd.MarkFlagRequired("service-action-associations")
	servicecatalogCmd.AddCommand(servicecatalog_batchDisassociateServiceActionFromProvisioningArtifactCmd)
}
