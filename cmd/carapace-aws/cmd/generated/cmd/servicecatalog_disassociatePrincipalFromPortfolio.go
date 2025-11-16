package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_disassociatePrincipalFromPortfolioCmd = &cobra.Command{
	Use:   "disassociate-principal-from-portfolio",
	Short: "Disassociates a previously associated principal ARN from a specified portfolio.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_disassociatePrincipalFromPortfolioCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_disassociatePrincipalFromPortfolioCmd).Standalone()

		servicecatalog_disassociatePrincipalFromPortfolioCmd.Flags().String("accept-language", "", "The language code.")
		servicecatalog_disassociatePrincipalFromPortfolioCmd.Flags().String("portfolio-id", "", "The portfolio identifier.")
		servicecatalog_disassociatePrincipalFromPortfolioCmd.Flags().String("principal-arn", "", "The ARN of the principal (user, role, or group).")
		servicecatalog_disassociatePrincipalFromPortfolioCmd.Flags().String("principal-type", "", "The supported value is `IAM` if you use a fully defined ARN, or `IAM_PATTERN` if you specify an `IAM` ARN with no AccountId, with or without wildcard characters.")
		servicecatalog_disassociatePrincipalFromPortfolioCmd.MarkFlagRequired("portfolio-id")
		servicecatalog_disassociatePrincipalFromPortfolioCmd.MarkFlagRequired("principal-arn")
	})
	servicecatalogCmd.AddCommand(servicecatalog_disassociatePrincipalFromPortfolioCmd)
}
