package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_searchProvisionedProductsCmd = &cobra.Command{
	Use:   "search-provisioned-products",
	Short: "Gets information about the provisioned products that meet the specified criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_searchProvisionedProductsCmd).Standalone()

	servicecatalog_searchProvisionedProductsCmd.Flags().String("accept-language", "", "The language code.")
	servicecatalog_searchProvisionedProductsCmd.Flags().String("access-level-filter", "", "The access level to use to obtain results.")
	servicecatalog_searchProvisionedProductsCmd.Flags().String("filters", "", "The search filters.")
	servicecatalog_searchProvisionedProductsCmd.Flags().String("page-size", "", "The maximum number of items to return with this call.")
	servicecatalog_searchProvisionedProductsCmd.Flags().String("page-token", "", "The page token for the next set of results.")
	servicecatalog_searchProvisionedProductsCmd.Flags().String("sort-by", "", "The sort field.")
	servicecatalog_searchProvisionedProductsCmd.Flags().String("sort-order", "", "The sort order.")
	servicecatalogCmd.AddCommand(servicecatalog_searchProvisionedProductsCmd)
}
