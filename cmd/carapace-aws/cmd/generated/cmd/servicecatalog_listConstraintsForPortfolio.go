package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_listConstraintsForPortfolioCmd = &cobra.Command{
	Use:   "list-constraints-for-portfolio",
	Short: "Lists the constraints for the specified portfolio and product.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_listConstraintsForPortfolioCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_listConstraintsForPortfolioCmd).Standalone()

		servicecatalog_listConstraintsForPortfolioCmd.Flags().String("accept-language", "", "The language code.")
		servicecatalog_listConstraintsForPortfolioCmd.Flags().String("page-size", "", "The maximum number of items to return with this call.")
		servicecatalog_listConstraintsForPortfolioCmd.Flags().String("page-token", "", "The page token for the next set of results.")
		servicecatalog_listConstraintsForPortfolioCmd.Flags().String("portfolio-id", "", "The portfolio identifier.")
		servicecatalog_listConstraintsForPortfolioCmd.Flags().String("product-id", "", "The product identifier.")
		servicecatalog_listConstraintsForPortfolioCmd.MarkFlagRequired("portfolio-id")
	})
	servicecatalogCmd.AddCommand(servicecatalog_listConstraintsForPortfolioCmd)
}
