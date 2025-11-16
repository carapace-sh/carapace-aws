package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_describePortfolioSharesCmd = &cobra.Command{
	Use:   "describe-portfolio-shares",
	Short: "Returns a summary of each of the portfolio shares that were created for the specified portfolio.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_describePortfolioSharesCmd).Standalone()

	servicecatalog_describePortfolioSharesCmd.Flags().String("page-size", "", "The maximum number of items to return with this call.")
	servicecatalog_describePortfolioSharesCmd.Flags().String("page-token", "", "The page token for the next set of results.")
	servicecatalog_describePortfolioSharesCmd.Flags().String("portfolio-id", "", "The unique identifier of the portfolio for which shares will be retrieved.")
	servicecatalog_describePortfolioSharesCmd.Flags().String("type", "", "The type of portfolio share to summarize.")
	servicecatalog_describePortfolioSharesCmd.MarkFlagRequired("portfolio-id")
	servicecatalog_describePortfolioSharesCmd.MarkFlagRequired("type")
	servicecatalogCmd.AddCommand(servicecatalog_describePortfolioSharesCmd)
}
