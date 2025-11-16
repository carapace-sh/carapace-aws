package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_listPortfolioAccessCmd = &cobra.Command{
	Use:   "list-portfolio-access",
	Short: "Lists the account IDs that have access to the specified portfolio.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_listPortfolioAccessCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_listPortfolioAccessCmd).Standalone()

		servicecatalog_listPortfolioAccessCmd.Flags().String("accept-language", "", "The language code.")
		servicecatalog_listPortfolioAccessCmd.Flags().String("organization-parent-id", "", "The ID of an organization node the portfolio is shared with.")
		servicecatalog_listPortfolioAccessCmd.Flags().String("page-size", "", "The maximum number of items to return with this call.")
		servicecatalog_listPortfolioAccessCmd.Flags().String("page-token", "", "The page token for the next set of results.")
		servicecatalog_listPortfolioAccessCmd.Flags().String("portfolio-id", "", "The portfolio identifier.")
		servicecatalog_listPortfolioAccessCmd.MarkFlagRequired("portfolio-id")
	})
	servicecatalogCmd.AddCommand(servicecatalog_listPortfolioAccessCmd)
}
