package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_deleteProvisioningArtifactCmd = &cobra.Command{
	Use:   "delete-provisioning-artifact",
	Short: "Deletes the specified provisioning artifact (also known as a version) for the specified product.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_deleteProvisioningArtifactCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_deleteProvisioningArtifactCmd).Standalone()

		servicecatalog_deleteProvisioningArtifactCmd.Flags().String("accept-language", "", "The language code.")
		servicecatalog_deleteProvisioningArtifactCmd.Flags().String("product-id", "", "The product identifier.")
		servicecatalog_deleteProvisioningArtifactCmd.Flags().String("provisioning-artifact-id", "", "The identifier of the provisioning artifact.")
		servicecatalog_deleteProvisioningArtifactCmd.MarkFlagRequired("product-id")
		servicecatalog_deleteProvisioningArtifactCmd.MarkFlagRequired("provisioning-artifact-id")
	})
	servicecatalogCmd.AddCommand(servicecatalog_deleteProvisioningArtifactCmd)
}
