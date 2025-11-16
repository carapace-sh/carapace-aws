package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_listOrganizationPortfolioAccessCmd = &cobra.Command{
	Use:   "list-organization-portfolio-access",
	Short: "Lists the organization nodes that have access to the specified portfolio.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_listOrganizationPortfolioAccessCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_listOrganizationPortfolioAccessCmd).Standalone()

		servicecatalog_listOrganizationPortfolioAccessCmd.Flags().String("accept-language", "", "The language code.")
		servicecatalog_listOrganizationPortfolioAccessCmd.Flags().String("organization-node-type", "", "The organization node type that will be returned in the output.")
		servicecatalog_listOrganizationPortfolioAccessCmd.Flags().String("page-size", "", "The maximum number of items to return with this call.")
		servicecatalog_listOrganizationPortfolioAccessCmd.Flags().String("page-token", "", "The page token for the next set of results.")
		servicecatalog_listOrganizationPortfolioAccessCmd.Flags().String("portfolio-id", "", "The portfolio identifier.")
		servicecatalog_listOrganizationPortfolioAccessCmd.MarkFlagRequired("organization-node-type")
		servicecatalog_listOrganizationPortfolioAccessCmd.MarkFlagRequired("portfolio-id")
	})
	servicecatalogCmd.AddCommand(servicecatalog_listOrganizationPortfolioAccessCmd)
}
