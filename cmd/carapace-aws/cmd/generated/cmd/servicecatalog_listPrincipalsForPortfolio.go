package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_listPrincipalsForPortfolioCmd = &cobra.Command{
	Use:   "list-principals-for-portfolio",
	Short: "Lists all `PrincipalARN`s and corresponding `PrincipalType`s associated with the specified portfolio.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_listPrincipalsForPortfolioCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_listPrincipalsForPortfolioCmd).Standalone()

		servicecatalog_listPrincipalsForPortfolioCmd.Flags().String("accept-language", "", "The language code.")
		servicecatalog_listPrincipalsForPortfolioCmd.Flags().String("page-size", "", "The maximum number of items to return with this call.")
		servicecatalog_listPrincipalsForPortfolioCmd.Flags().String("page-token", "", "The page token for the next set of results.")
		servicecatalog_listPrincipalsForPortfolioCmd.Flags().String("portfolio-id", "", "The portfolio identifier.")
		servicecatalog_listPrincipalsForPortfolioCmd.MarkFlagRequired("portfolio-id")
	})
	servicecatalogCmd.AddCommand(servicecatalog_listPrincipalsForPortfolioCmd)
}
