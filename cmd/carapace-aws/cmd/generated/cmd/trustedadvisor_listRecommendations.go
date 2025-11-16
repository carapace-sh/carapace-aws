package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var trustedadvisor_listRecommendationsCmd = &cobra.Command{
	Use:   "list-recommendations",
	Short: "List a filterable set of Recommendations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(trustedadvisor_listRecommendationsCmd).Standalone()

	trustedadvisor_listRecommendationsCmd.Flags().String("after-last-updated-at", "", "After the last update of the Recommendation")
	trustedadvisor_listRecommendationsCmd.Flags().String("aws-service", "", "The aws service associated with the Recommendation")
	trustedadvisor_listRecommendationsCmd.Flags().String("before-last-updated-at", "", "Before the last update of the Recommendation")
	trustedadvisor_listRecommendationsCmd.Flags().String("check-identifier", "", "The check identifier of the Recommendation")
	trustedadvisor_listRecommendationsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	trustedadvisor_listRecommendationsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	trustedadvisor_listRecommendationsCmd.Flags().String("pillar", "", "The pillar of the Recommendation")
	trustedadvisor_listRecommendationsCmd.Flags().String("source", "", "The source of the Recommendation")
	trustedadvisor_listRecommendationsCmd.Flags().String("status", "", "The status of the Recommendation")
	trustedadvisor_listRecommendationsCmd.Flags().String("type", "", "The type of the Recommendation")
	trustedadvisorCmd.AddCommand(trustedadvisor_listRecommendationsCmd)
}
