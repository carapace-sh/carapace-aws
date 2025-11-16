package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_listPortfoliosCmd = &cobra.Command{
	Use:   "list-portfolios",
	Short: "Lists all portfolios in the catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_listPortfoliosCmd).Standalone()

	servicecatalog_listPortfoliosCmd.Flags().String("accept-language", "", "The language code.")
	servicecatalog_listPortfoliosCmd.Flags().String("page-size", "", "The maximum number of items to return with this call.")
	servicecatalog_listPortfoliosCmd.Flags().String("page-token", "", "The page token for the next set of results.")
	servicecatalogCmd.AddCommand(servicecatalog_listPortfoliosCmd)
}
