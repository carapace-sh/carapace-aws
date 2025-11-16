package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_updateProductCmd = &cobra.Command{
	Use:   "update-product",
	Short: "Updates the specified product.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_updateProductCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_updateProductCmd).Standalone()

		servicecatalog_updateProductCmd.Flags().String("accept-language", "", "The language code.")
		servicecatalog_updateProductCmd.Flags().String("add-tags", "", "The tags to add to the product.")
		servicecatalog_updateProductCmd.Flags().String("description", "", "The updated description of the product.")
		servicecatalog_updateProductCmd.Flags().String("distributor", "", "The updated distributor of the product.")
		servicecatalog_updateProductCmd.Flags().String("id", "", "The product identifier.")
		servicecatalog_updateProductCmd.Flags().String("name", "", "The updated product name.")
		servicecatalog_updateProductCmd.Flags().String("owner", "", "The updated owner of the product.")
		servicecatalog_updateProductCmd.Flags().String("remove-tags", "", "The tags to remove from the product.")
		servicecatalog_updateProductCmd.Flags().String("source-connection", "", "Specifies connection details for the updated product and syncs the product to the connection source artifact.")
		servicecatalog_updateProductCmd.Flags().String("support-description", "", "The updated support description for the product.")
		servicecatalog_updateProductCmd.Flags().String("support-email", "", "The updated support email for the product.")
		servicecatalog_updateProductCmd.Flags().String("support-url", "", "The updated support URL for the product.")
		servicecatalog_updateProductCmd.MarkFlagRequired("id")
	})
	servicecatalogCmd.AddCommand(servicecatalog_updateProductCmd)
}
