package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_createPortfolioCmd = &cobra.Command{
	Use:   "create-portfolio",
	Short: "Creates a portfolio.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_createPortfolioCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_createPortfolioCmd).Standalone()

		servicecatalog_createPortfolioCmd.Flags().String("accept-language", "", "The language code.")
		servicecatalog_createPortfolioCmd.Flags().String("description", "", "The description of the portfolio.")
		servicecatalog_createPortfolioCmd.Flags().String("display-name", "", "The name to use for display purposes.")
		servicecatalog_createPortfolioCmd.Flags().String("idempotency-token", "", "A unique identifier that you provide to ensure idempotency.")
		servicecatalog_createPortfolioCmd.Flags().String("provider-name", "", "The name of the portfolio provider.")
		servicecatalog_createPortfolioCmd.Flags().String("tags", "", "One or more tags.")
		servicecatalog_createPortfolioCmd.MarkFlagRequired("display-name")
		servicecatalog_createPortfolioCmd.MarkFlagRequired("idempotency-token")
		servicecatalog_createPortfolioCmd.MarkFlagRequired("provider-name")
	})
	servicecatalogCmd.AddCommand(servicecatalog_createPortfolioCmd)
}
