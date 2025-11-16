package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_createProductCmd = &cobra.Command{
	Use:   "create-product",
	Short: "Creates a product.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_createProductCmd).Standalone()

	servicecatalog_createProductCmd.Flags().String("accept-language", "", "The language code.")
	servicecatalog_createProductCmd.Flags().String("description", "", "The description of the product.")
	servicecatalog_createProductCmd.Flags().String("distributor", "", "The distributor of the product.")
	servicecatalog_createProductCmd.Flags().String("idempotency-token", "", "A unique identifier that you provide to ensure idempotency.")
	servicecatalog_createProductCmd.Flags().String("name", "", "The name of the product.")
	servicecatalog_createProductCmd.Flags().String("owner", "", "The owner of the product.")
	servicecatalog_createProductCmd.Flags().String("product-type", "", "The type of product.")
	servicecatalog_createProductCmd.Flags().String("provisioning-artifact-parameters", "", "The configuration of the provisioning artifact.")
	servicecatalog_createProductCmd.Flags().String("source-connection", "", "Specifies connection details for the created product and syncs the product to the connection source artifact.")
	servicecatalog_createProductCmd.Flags().String("support-description", "", "The support information about the product.")
	servicecatalog_createProductCmd.Flags().String("support-email", "", "The contact email for product support.")
	servicecatalog_createProductCmd.Flags().String("support-url", "", "The contact URL for product support.")
	servicecatalog_createProductCmd.Flags().String("tags", "", "One or more tags.")
	servicecatalog_createProductCmd.MarkFlagRequired("idempotency-token")
	servicecatalog_createProductCmd.MarkFlagRequired("name")
	servicecatalog_createProductCmd.MarkFlagRequired("owner")
	servicecatalog_createProductCmd.MarkFlagRequired("product-type")
	servicecatalogCmd.AddCommand(servicecatalog_createProductCmd)
}
