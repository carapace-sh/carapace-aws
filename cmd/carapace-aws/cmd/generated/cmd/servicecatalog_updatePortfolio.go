package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_updatePortfolioCmd = &cobra.Command{
	Use:   "update-portfolio",
	Short: "Updates the specified portfolio.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_updatePortfolioCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_updatePortfolioCmd).Standalone()

		servicecatalog_updatePortfolioCmd.Flags().String("accept-language", "", "The language code.")
		servicecatalog_updatePortfolioCmd.Flags().String("add-tags", "", "The tags to add.")
		servicecatalog_updatePortfolioCmd.Flags().String("description", "", "The updated description of the portfolio.")
		servicecatalog_updatePortfolioCmd.Flags().String("display-name", "", "The name to use for display purposes.")
		servicecatalog_updatePortfolioCmd.Flags().String("id", "", "The portfolio identifier.")
		servicecatalog_updatePortfolioCmd.Flags().String("provider-name", "", "The updated name of the portfolio provider.")
		servicecatalog_updatePortfolioCmd.Flags().String("remove-tags", "", "The tags to remove.")
		servicecatalog_updatePortfolioCmd.MarkFlagRequired("id")
	})
	servicecatalogCmd.AddCommand(servicecatalog_updatePortfolioCmd)
}
