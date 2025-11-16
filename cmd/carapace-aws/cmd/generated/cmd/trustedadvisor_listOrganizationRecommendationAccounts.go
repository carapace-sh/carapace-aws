package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var trustedadvisor_listOrganizationRecommendationAccountsCmd = &cobra.Command{
	Use:   "list-organization-recommendation-accounts",
	Short: "Lists the accounts that own the resources for an organization aggregate recommendation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(trustedadvisor_listOrganizationRecommendationAccountsCmd).Standalone()

	trustedadvisor_listOrganizationRecommendationAccountsCmd.Flags().String("affected-account-id", "", "An account affected by this organization recommendation")
	trustedadvisor_listOrganizationRecommendationAccountsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	trustedadvisor_listOrganizationRecommendationAccountsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	trustedadvisor_listOrganizationRecommendationAccountsCmd.Flags().String("organization-recommendation-identifier", "", "The Recommendation identifier")
	trustedadvisor_listOrganizationRecommendationAccountsCmd.MarkFlagRequired("organization-recommendation-identifier")
	trustedadvisorCmd.AddCommand(trustedadvisor_listOrganizationRecommendationAccountsCmd)
}
