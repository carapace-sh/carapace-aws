package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_updateProvisionedProductPropertiesCmd = &cobra.Command{
	Use:   "update-provisioned-product-properties",
	Short: "Requests updates to the properties of the specified provisioned product.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_updateProvisionedProductPropertiesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_updateProvisionedProductPropertiesCmd).Standalone()

		servicecatalog_updateProvisionedProductPropertiesCmd.Flags().String("accept-language", "", "The language code.")
		servicecatalog_updateProvisionedProductPropertiesCmd.Flags().String("idempotency-token", "", "The idempotency token that uniquely identifies the provisioning product update request.")
		servicecatalog_updateProvisionedProductPropertiesCmd.Flags().String("provisioned-product-id", "", "The identifier of the provisioned product.")
		servicecatalog_updateProvisionedProductPropertiesCmd.Flags().String("provisioned-product-properties", "", "A map that contains the provisioned product properties to be updated.")
		servicecatalog_updateProvisionedProductPropertiesCmd.MarkFlagRequired("idempotency-token")
		servicecatalog_updateProvisionedProductPropertiesCmd.MarkFlagRequired("provisioned-product-id")
		servicecatalog_updateProvisionedProductPropertiesCmd.MarkFlagRequired("provisioned-product-properties")
	})
	servicecatalogCmd.AddCommand(servicecatalog_updateProvisionedProductPropertiesCmd)
}
