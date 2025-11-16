package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_associateServiceActionWithProvisioningArtifactCmd = &cobra.Command{
	Use:   "associate-service-action-with-provisioning-artifact",
	Short: "Associates a self-service action with a provisioning artifact.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_associateServiceActionWithProvisioningArtifactCmd).Standalone()

	servicecatalog_associateServiceActionWithProvisioningArtifactCmd.Flags().String("accept-language", "", "The language code.")
	servicecatalog_associateServiceActionWithProvisioningArtifactCmd.Flags().String("idempotency-token", "", "A unique identifier that you provide to ensure idempotency.")
	servicecatalog_associateServiceActionWithProvisioningArtifactCmd.Flags().String("product-id", "", "The product identifier.")
	servicecatalog_associateServiceActionWithProvisioningArtifactCmd.Flags().String("provisioning-artifact-id", "", "The identifier of the provisioning artifact.")
	servicecatalog_associateServiceActionWithProvisioningArtifactCmd.Flags().String("service-action-id", "", "The self-service action identifier.")
	servicecatalog_associateServiceActionWithProvisioningArtifactCmd.MarkFlagRequired("product-id")
	servicecatalog_associateServiceActionWithProvisioningArtifactCmd.MarkFlagRequired("provisioning-artifact-id")
	servicecatalog_associateServiceActionWithProvisioningArtifactCmd.MarkFlagRequired("service-action-id")
	servicecatalogCmd.AddCommand(servicecatalog_associateServiceActionWithProvisioningArtifactCmd)
}
