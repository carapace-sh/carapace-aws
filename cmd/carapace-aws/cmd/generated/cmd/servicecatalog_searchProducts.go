package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_searchProductsCmd = &cobra.Command{
	Use:   "search-products",
	Short: "Gets information about the products to which the caller has access.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_searchProductsCmd).Standalone()

	servicecatalog_searchProductsCmd.Flags().String("accept-language", "", "The language code.")
	servicecatalog_searchProductsCmd.Flags().String("filters", "", "The search filters.")
	servicecatalog_searchProductsCmd.Flags().String("page-size", "", "The maximum number of items to return with this call.")
	servicecatalog_searchProductsCmd.Flags().String("page-token", "", "The page token for the next set of results.")
	servicecatalog_searchProductsCmd.Flags().String("sort-by", "", "The sort field.")
	servicecatalog_searchProductsCmd.Flags().String("sort-order", "", "The sort order.")
	servicecatalogCmd.AddCommand(servicecatalog_searchProductsCmd)
}
