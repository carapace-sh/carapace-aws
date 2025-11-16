package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_updateProvisionedProductCmd = &cobra.Command{
	Use:   "update-provisioned-product",
	Short: "Requests updates to the configuration of the specified provisioned product.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_updateProvisionedProductCmd).Standalone()

	servicecatalog_updateProvisionedProductCmd.Flags().String("accept-language", "", "The language code.")
	servicecatalog_updateProvisionedProductCmd.Flags().String("path-id", "", "The path identifier.")
	servicecatalog_updateProvisionedProductCmd.Flags().String("path-name", "", "The name of the path.")
	servicecatalog_updateProvisionedProductCmd.Flags().String("product-id", "", "The identifier of the product.")
	servicecatalog_updateProvisionedProductCmd.Flags().String("product-name", "", "The name of the product.")
	servicecatalog_updateProvisionedProductCmd.Flags().String("provisioned-product-id", "", "The identifier of the provisioned product.")
	servicecatalog_updateProvisionedProductCmd.Flags().String("provisioned-product-name", "", "The name of the provisioned product.")
	servicecatalog_updateProvisionedProductCmd.Flags().String("provisioning-artifact-id", "", "The identifier of the provisioning artifact.")
	servicecatalog_updateProvisionedProductCmd.Flags().String("provisioning-artifact-name", "", "The name of the provisioning artifact.")
	servicecatalog_updateProvisionedProductCmd.Flags().String("provisioning-parameters", "", "The new parameters.")
	servicecatalog_updateProvisionedProductCmd.Flags().String("provisioning-preferences", "", "An object that contains information about the provisioning preferences for a stack set.")
	servicecatalog_updateProvisionedProductCmd.Flags().String("tags", "", "One or more tags.")
	servicecatalog_updateProvisionedProductCmd.Flags().String("update-token", "", "The idempotency token that uniquely identifies the provisioning update request.")
	servicecatalog_updateProvisionedProductCmd.MarkFlagRequired("update-token")
	servicecatalogCmd.AddCommand(servicecatalog_updateProvisionedProductCmd)
}
