package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var trustedadvisor_listOrganizationRecommendationsCmd = &cobra.Command{
	Use:   "list-organization-recommendations",
	Short: "List a filterable set of Recommendations within an Organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(trustedadvisor_listOrganizationRecommendationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(trustedadvisor_listOrganizationRecommendationsCmd).Standalone()

		trustedadvisor_listOrganizationRecommendationsCmd.Flags().String("after-last-updated-at", "", "After the last update of the Recommendation")
		trustedadvisor_listOrganizationRecommendationsCmd.Flags().String("aws-service", "", "The aws service associated with the Recommendation")
		trustedadvisor_listOrganizationRecommendationsCmd.Flags().String("before-last-updated-at", "", "Before the last update of the Recommendation")
		trustedadvisor_listOrganizationRecommendationsCmd.Flags().String("check-identifier", "", "The check identifier of the Recommendation")
		trustedadvisor_listOrganizationRecommendationsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		trustedadvisor_listOrganizationRecommendationsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		trustedadvisor_listOrganizationRecommendationsCmd.Flags().String("pillar", "", "The pillar of the Recommendation")
		trustedadvisor_listOrganizationRecommendationsCmd.Flags().String("source", "", "The source of the Recommendation")
		trustedadvisor_listOrganizationRecommendationsCmd.Flags().String("status", "", "The status of the Recommendation")
		trustedadvisor_listOrganizationRecommendationsCmd.Flags().String("type", "", "The type of the Recommendation")
	})
	trustedadvisorCmd.AddCommand(trustedadvisor_listOrganizationRecommendationsCmd)
}
