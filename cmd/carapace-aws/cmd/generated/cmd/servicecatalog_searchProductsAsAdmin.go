package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_searchProductsAsAdminCmd = &cobra.Command{
	Use:   "search-products-as-admin",
	Short: "Gets information about the products for the specified portfolio or all products.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_searchProductsAsAdminCmd).Standalone()

	servicecatalog_searchProductsAsAdminCmd.Flags().String("accept-language", "", "The language code.")
	servicecatalog_searchProductsAsAdminCmd.Flags().String("filters", "", "The search filters.")
	servicecatalog_searchProductsAsAdminCmd.Flags().String("page-size", "", "The maximum number of items to return with this call.")
	servicecatalog_searchProductsAsAdminCmd.Flags().String("page-token", "", "The page token for the next set of results.")
	servicecatalog_searchProductsAsAdminCmd.Flags().String("portfolio-id", "", "The portfolio identifier.")
	servicecatalog_searchProductsAsAdminCmd.Flags().String("product-source", "", "Access level of the source of the product.")
	servicecatalog_searchProductsAsAdminCmd.Flags().String("sort-by", "", "The sort field.")
	servicecatalog_searchProductsAsAdminCmd.Flags().String("sort-order", "", "The sort order.")
	servicecatalogCmd.AddCommand(servicecatalog_searchProductsAsAdminCmd)
}
