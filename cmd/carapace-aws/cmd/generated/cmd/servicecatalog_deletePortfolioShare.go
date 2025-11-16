package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_deletePortfolioShareCmd = &cobra.Command{
	Use:   "delete-portfolio-share",
	Short: "Stops sharing the specified portfolio with the specified account or organization node.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_deletePortfolioShareCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_deletePortfolioShareCmd).Standalone()

		servicecatalog_deletePortfolioShareCmd.Flags().String("accept-language", "", "The language code.")
		servicecatalog_deletePortfolioShareCmd.Flags().String("account-id", "", "The Amazon Web Services account ID.")
		servicecatalog_deletePortfolioShareCmd.Flags().String("organization-node", "", "The organization node to whom you are going to stop sharing.")
		servicecatalog_deletePortfolioShareCmd.Flags().String("portfolio-id", "", "The portfolio identifier.")
		servicecatalog_deletePortfolioShareCmd.MarkFlagRequired("portfolio-id")
	})
	servicecatalogCmd.AddCommand(servicecatalog_deletePortfolioShareCmd)
}
