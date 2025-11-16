package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_importAsProvisionedProductCmd = &cobra.Command{
	Use:   "import-as-provisioned-product",
	Short: "Requests the import of a resource as an Service Catalog provisioned product that is associated to an Service Catalog product and provisioning artifact.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_importAsProvisionedProductCmd).Standalone()

	servicecatalog_importAsProvisionedProductCmd.Flags().String("accept-language", "", "The language code.")
	servicecatalog_importAsProvisionedProductCmd.Flags().String("idempotency-token", "", "A unique identifier that you provide to ensure idempotency.")
	servicecatalog_importAsProvisionedProductCmd.Flags().String("physical-id", "", "The unique identifier of the resource to be imported.")
	servicecatalog_importAsProvisionedProductCmd.Flags().String("product-id", "", "The product identifier.")
	servicecatalog_importAsProvisionedProductCmd.Flags().String("provisioned-product-name", "", "The user-friendly name of the provisioned product.")
	servicecatalog_importAsProvisionedProductCmd.Flags().String("provisioning-artifact-id", "", "The identifier of the provisioning artifact.")
	servicecatalog_importAsProvisionedProductCmd.MarkFlagRequired("idempotency-token")
	servicecatalog_importAsProvisionedProductCmd.MarkFlagRequired("physical-id")
	servicecatalog_importAsProvisionedProductCmd.MarkFlagRequired("product-id")
	servicecatalog_importAsProvisionedProductCmd.MarkFlagRequired("provisioned-product-name")
	servicecatalog_importAsProvisionedProductCmd.MarkFlagRequired("provisioning-artifact-id")
	servicecatalogCmd.AddCommand(servicecatalog_importAsProvisionedProductCmd)
}
