package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_acceptPortfolioShareCmd = &cobra.Command{
	Use:   "accept-portfolio-share",
	Short: "Accepts an offer to share the specified portfolio.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_acceptPortfolioShareCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_acceptPortfolioShareCmd).Standalone()

		servicecatalog_acceptPortfolioShareCmd.Flags().String("accept-language", "", "The language code.")
		servicecatalog_acceptPortfolioShareCmd.Flags().String("portfolio-id", "", "The portfolio identifier.")
		servicecatalog_acceptPortfolioShareCmd.Flags().String("portfolio-share-type", "", "The type of shared portfolios to accept.")
		servicecatalog_acceptPortfolioShareCmd.MarkFlagRequired("portfolio-id")
	})
	servicecatalogCmd.AddCommand(servicecatalog_acceptPortfolioShareCmd)
}
