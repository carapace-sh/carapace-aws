package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_updateProvisioningArtifactCmd = &cobra.Command{
	Use:   "update-provisioning-artifact",
	Short: "Updates the specified provisioning artifact (also known as a version) for the specified product.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_updateProvisioningArtifactCmd).Standalone()

	servicecatalog_updateProvisioningArtifactCmd.Flags().String("accept-language", "", "The language code.")
	servicecatalog_updateProvisioningArtifactCmd.Flags().String("active", "", "Indicates whether the product version is active.")
	servicecatalog_updateProvisioningArtifactCmd.Flags().String("description", "", "The updated description of the provisioning artifact.")
	servicecatalog_updateProvisioningArtifactCmd.Flags().String("guidance", "", "Information set by the administrator to provide guidance to end users about which provisioning artifacts to use.")
	servicecatalog_updateProvisioningArtifactCmd.Flags().String("name", "", "The updated name of the provisioning artifact.")
	servicecatalog_updateProvisioningArtifactCmd.Flags().String("product-id", "", "The product identifier.")
	servicecatalog_updateProvisioningArtifactCmd.Flags().String("provisioning-artifact-id", "", "The identifier of the provisioning artifact.")
	servicecatalog_updateProvisioningArtifactCmd.MarkFlagRequired("product-id")
	servicecatalog_updateProvisioningArtifactCmd.MarkFlagRequired("provisioning-artifact-id")
	servicecatalogCmd.AddCommand(servicecatalog_updateProvisioningArtifactCmd)
}
