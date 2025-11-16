package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_rejectPortfolioShareCmd = &cobra.Command{
	Use:   "reject-portfolio-share",
	Short: "Rejects an offer to share the specified portfolio.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_rejectPortfolioShareCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_rejectPortfolioShareCmd).Standalone()

		servicecatalog_rejectPortfolioShareCmd.Flags().String("accept-language", "", "The language code.")
		servicecatalog_rejectPortfolioShareCmd.Flags().String("portfolio-id", "", "The portfolio identifier.")
		servicecatalog_rejectPortfolioShareCmd.Flags().String("portfolio-share-type", "", "The type of shared portfolios to reject.")
		servicecatalog_rejectPortfolioShareCmd.MarkFlagRequired("portfolio-id")
	})
	servicecatalogCmd.AddCommand(servicecatalog_rejectPortfolioShareCmd)
}
