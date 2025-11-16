package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_disassociateServiceActionFromProvisioningArtifactCmd = &cobra.Command{
	Use:   "disassociate-service-action-from-provisioning-artifact",
	Short: "Disassociates the specified self-service action association from the specified provisioning artifact.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_disassociateServiceActionFromProvisioningArtifactCmd).Standalone()

	servicecatalog_disassociateServiceActionFromProvisioningArtifactCmd.Flags().String("accept-language", "", "The language code.")
	servicecatalog_disassociateServiceActionFromProvisioningArtifactCmd.Flags().String("idempotency-token", "", "A unique identifier that you provide to ensure idempotency.")
	servicecatalog_disassociateServiceActionFromProvisioningArtifactCmd.Flags().String("product-id", "", "The product identifier.")
	servicecatalog_disassociateServiceActionFromProvisioningArtifactCmd.Flags().String("provisioning-artifact-id", "", "The identifier of the provisioning artifact.")
	servicecatalog_disassociateServiceActionFromProvisioningArtifactCmd.Flags().String("service-action-id", "", "The self-service action identifier.")
	servicecatalog_disassociateServiceActionFromProvisioningArtifactCmd.MarkFlagRequired("product-id")
	servicecatalog_disassociateServiceActionFromProvisioningArtifactCmd.MarkFlagRequired("provisioning-artifact-id")
	servicecatalog_disassociateServiceActionFromProvisioningArtifactCmd.MarkFlagRequired("service-action-id")
	servicecatalogCmd.AddCommand(servicecatalog_disassociateServiceActionFromProvisioningArtifactCmd)
}
