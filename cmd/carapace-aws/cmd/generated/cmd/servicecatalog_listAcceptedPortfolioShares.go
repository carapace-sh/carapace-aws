package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_listAcceptedPortfolioSharesCmd = &cobra.Command{
	Use:   "list-accepted-portfolio-shares",
	Short: "Lists all imported portfolios for which account-to-account shares were accepted by this account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_listAcceptedPortfolioSharesCmd).Standalone()

	servicecatalog_listAcceptedPortfolioSharesCmd.Flags().String("accept-language", "", "The language code.")
	servicecatalog_listAcceptedPortfolioSharesCmd.Flags().String("page-size", "", "The maximum number of items to return with this call.")
	servicecatalog_listAcceptedPortfolioSharesCmd.Flags().String("page-token", "", "The page token for the next set of results.")
	servicecatalog_listAcceptedPortfolioSharesCmd.Flags().String("portfolio-share-type", "", "The type of shared portfolios to list.")
	servicecatalogCmd.AddCommand(servicecatalog_listAcceptedPortfolioSharesCmd)
}
