package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var trustedadvisor_listOrganizationRecommendationResourcesCmd = &cobra.Command{
	Use:   "list-organization-recommendation-resources",
	Short: "List Resources of a Recommendation within an Organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(trustedadvisor_listOrganizationRecommendationResourcesCmd).Standalone()

	trustedadvisor_listOrganizationRecommendationResourcesCmd.Flags().String("affected-account-id", "", "An account affected by this organization recommendation")
	trustedadvisor_listOrganizationRecommendationResourcesCmd.Flags().String("exclusion-status", "", "The exclusion status of the resource")
	trustedadvisor_listOrganizationRecommendationResourcesCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	trustedadvisor_listOrganizationRecommendationResourcesCmd.Flags().String("next-token", "", "The token for the next set of results.")
	trustedadvisor_listOrganizationRecommendationResourcesCmd.Flags().String("organization-recommendation-identifier", "", "The AWS Organization organization's Recommendation identifier")
	trustedadvisor_listOrganizationRecommendationResourcesCmd.Flags().String("region-code", "", "The AWS Region code of the resource")
	trustedadvisor_listOrganizationRecommendationResourcesCmd.Flags().String("status", "", "The status of the resource")
	trustedadvisor_listOrganizationRecommendationResourcesCmd.MarkFlagRequired("organization-recommendation-identifier")
	trustedadvisorCmd.AddCommand(trustedadvisor_listOrganizationRecommendationResourcesCmd)
}
