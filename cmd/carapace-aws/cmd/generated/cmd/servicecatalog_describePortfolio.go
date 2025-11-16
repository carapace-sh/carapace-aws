package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_describePortfolioCmd = &cobra.Command{
	Use:   "describe-portfolio",
	Short: "Gets information about the specified portfolio.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_describePortfolioCmd).Standalone()

	servicecatalog_describePortfolioCmd.Flags().String("accept-language", "", "The language code.")
	servicecatalog_describePortfolioCmd.Flags().String("id", "", "The portfolio identifier.")
	servicecatalog_describePortfolioCmd.MarkFlagRequired("id")
	servicecatalogCmd.AddCommand(servicecatalog_describePortfolioCmd)
}
