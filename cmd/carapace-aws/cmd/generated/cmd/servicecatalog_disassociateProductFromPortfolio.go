package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_disassociateProductFromPortfolioCmd = &cobra.Command{
	Use:   "disassociate-product-from-portfolio",
	Short: "Disassociates the specified product from the specified portfolio.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_disassociateProductFromPortfolioCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_disassociateProductFromPortfolioCmd).Standalone()

		servicecatalog_disassociateProductFromPortfolioCmd.Flags().String("accept-language", "", "The language code.")
		servicecatalog_disassociateProductFromPortfolioCmd.Flags().String("portfolio-id", "", "The portfolio identifier.")
		servicecatalog_disassociateProductFromPortfolioCmd.Flags().String("product-id", "", "The product identifier.")
		servicecatalog_disassociateProductFromPortfolioCmd.MarkFlagRequired("portfolio-id")
		servicecatalog_disassociateProductFromPortfolioCmd.MarkFlagRequired("product-id")
	})
	servicecatalogCmd.AddCommand(servicecatalog_disassociateProductFromPortfolioCmd)
}
