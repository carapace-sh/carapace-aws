package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_listPortfoliosForProductCmd = &cobra.Command{
	Use:   "list-portfolios-for-product",
	Short: "Lists all portfolios that the specified product is associated with.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_listPortfoliosForProductCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_listPortfoliosForProductCmd).Standalone()

		servicecatalog_listPortfoliosForProductCmd.Flags().String("accept-language", "", "The language code.")
		servicecatalog_listPortfoliosForProductCmd.Flags().String("page-size", "", "The maximum number of items to return with this call.")
		servicecatalog_listPortfoliosForProductCmd.Flags().String("page-token", "", "The page token for the next set of results.")
		servicecatalog_listPortfoliosForProductCmd.Flags().String("product-id", "", "The product identifier.")
		servicecatalog_listPortfoliosForProductCmd.MarkFlagRequired("product-id")
	})
	servicecatalogCmd.AddCommand(servicecatalog_listPortfoliosForProductCmd)
}
