package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_updatePortfolioShareCmd = &cobra.Command{
	Use:   "update-portfolio-share",
	Short: "Updates the specified portfolio share.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_updatePortfolioShareCmd).Standalone()

	servicecatalog_updatePortfolioShareCmd.Flags().String("accept-language", "", "The language code.")
	servicecatalog_updatePortfolioShareCmd.Flags().String("account-id", "", "The Amazon Web Services account Id of the recipient account.")
	servicecatalog_updatePortfolioShareCmd.Flags().String("organization-node", "", "")
	servicecatalog_updatePortfolioShareCmd.Flags().String("portfolio-id", "", "The unique identifier of the portfolio for which the share will be updated.")
	servicecatalog_updatePortfolioShareCmd.Flags().String("share-principals", "", "A flag to enables or disables `Principals` sharing in the portfolio.")
	servicecatalog_updatePortfolioShareCmd.Flags().String("share-tag-options", "", "Enables or disables `TagOptions` sharing for the portfolio share.")
	servicecatalog_updatePortfolioShareCmd.MarkFlagRequired("portfolio-id")
	servicecatalogCmd.AddCommand(servicecatalog_updatePortfolioShareCmd)
}
