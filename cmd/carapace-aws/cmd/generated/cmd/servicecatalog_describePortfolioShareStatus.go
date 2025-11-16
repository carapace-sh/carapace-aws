package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_describePortfolioShareStatusCmd = &cobra.Command{
	Use:   "describe-portfolio-share-status",
	Short: "Gets the status of the specified portfolio share operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_describePortfolioShareStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_describePortfolioShareStatusCmd).Standalone()

		servicecatalog_describePortfolioShareStatusCmd.Flags().String("portfolio-share-token", "", "The token for the portfolio share operation.")
		servicecatalog_describePortfolioShareStatusCmd.MarkFlagRequired("portfolio-share-token")
	})
	servicecatalogCmd.AddCommand(servicecatalog_describePortfolioShareStatusCmd)
}
