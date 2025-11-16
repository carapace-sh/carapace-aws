package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_deletePortfolioCmd = &cobra.Command{
	Use:   "delete-portfolio",
	Short: "Deletes the specified portfolio.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_deletePortfolioCmd).Standalone()

	servicecatalog_deletePortfolioCmd.Flags().String("accept-language", "", "The language code.")
	servicecatalog_deletePortfolioCmd.Flags().String("id", "", "The portfolio identifier.")
	servicecatalog_deletePortfolioCmd.MarkFlagRequired("id")
	servicecatalogCmd.AddCommand(servicecatalog_deletePortfolioCmd)
}
