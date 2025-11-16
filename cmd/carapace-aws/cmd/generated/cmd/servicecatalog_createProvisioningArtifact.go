package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_createProvisioningArtifactCmd = &cobra.Command{
	Use:   "create-provisioning-artifact",
	Short: "Creates a provisioning artifact (also known as a version) for the specified product.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_createProvisioningArtifactCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_createProvisioningArtifactCmd).Standalone()

		servicecatalog_createProvisioningArtifactCmd.Flags().String("accept-language", "", "The language code.")
		servicecatalog_createProvisioningArtifactCmd.Flags().String("idempotency-token", "", "A unique identifier that you provide to ensure idempotency.")
		servicecatalog_createProvisioningArtifactCmd.Flags().String("parameters", "", "The configuration for the provisioning artifact.")
		servicecatalog_createProvisioningArtifactCmd.Flags().String("product-id", "", "The product identifier.")
		servicecatalog_createProvisioningArtifactCmd.MarkFlagRequired("idempotency-token")
		servicecatalog_createProvisioningArtifactCmd.MarkFlagRequired("parameters")
		servicecatalog_createProvisioningArtifactCmd.MarkFlagRequired("product-id")
	})
	servicecatalogCmd.AddCommand(servicecatalog_createProvisioningArtifactCmd)
}
