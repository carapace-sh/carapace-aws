package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_associateProductWithPortfolioCmd = &cobra.Command{
	Use:   "associate-product-with-portfolio",
	Short: "Associates the specified product with the specified portfolio.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_associateProductWithPortfolioCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_associateProductWithPortfolioCmd).Standalone()

		servicecatalog_associateProductWithPortfolioCmd.Flags().String("accept-language", "", "The language code.")
		servicecatalog_associateProductWithPortfolioCmd.Flags().String("portfolio-id", "", "The portfolio identifier.")
		servicecatalog_associateProductWithPortfolioCmd.Flags().String("product-id", "", "The product identifier.")
		servicecatalog_associateProductWithPortfolioCmd.Flags().String("source-portfolio-id", "", "The identifier of the source portfolio.")
		servicecatalog_associateProductWithPortfolioCmd.MarkFlagRequired("portfolio-id")
		servicecatalog_associateProductWithPortfolioCmd.MarkFlagRequired("product-id")
	})
	servicecatalogCmd.AddCommand(servicecatalog_associateProductWithPortfolioCmd)
}
