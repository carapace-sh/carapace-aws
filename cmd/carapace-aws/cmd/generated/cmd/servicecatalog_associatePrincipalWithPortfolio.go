package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_associatePrincipalWithPortfolioCmd = &cobra.Command{
	Use:   "associate-principal-with-portfolio",
	Short: "Associates the specified principal ARN with the specified portfolio.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_associatePrincipalWithPortfolioCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_associatePrincipalWithPortfolioCmd).Standalone()

		servicecatalog_associatePrincipalWithPortfolioCmd.Flags().String("accept-language", "", "The language code.")
		servicecatalog_associatePrincipalWithPortfolioCmd.Flags().String("portfolio-id", "", "The portfolio identifier.")
		servicecatalog_associatePrincipalWithPortfolioCmd.Flags().String("principal-arn", "", "The ARN of the principal (user, role, or group).")
		servicecatalog_associatePrincipalWithPortfolioCmd.Flags().String("principal-type", "", "The principal type.")
		servicecatalog_associatePrincipalWithPortfolioCmd.MarkFlagRequired("portfolio-id")
		servicecatalog_associatePrincipalWithPortfolioCmd.MarkFlagRequired("principal-arn")
		servicecatalog_associatePrincipalWithPortfolioCmd.MarkFlagRequired("principal-type")
	})
	servicecatalogCmd.AddCommand(servicecatalog_associatePrincipalWithPortfolioCmd)
}
