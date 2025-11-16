package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var trustedadvisor_listRecommendationResourcesCmd = &cobra.Command{
	Use:   "list-recommendation-resources",
	Short: "List Resources of a Recommendation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(trustedadvisor_listRecommendationResourcesCmd).Standalone()

	trustedadvisor_listRecommendationResourcesCmd.Flags().String("exclusion-status", "", "The exclusion status of the resource")
	trustedadvisor_listRecommendationResourcesCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	trustedadvisor_listRecommendationResourcesCmd.Flags().String("next-token", "", "The token for the next set of results.")
	trustedadvisor_listRecommendationResourcesCmd.Flags().String("recommendation-identifier", "", "The Recommendation identifier")
	trustedadvisor_listRecommendationResourcesCmd.Flags().String("region-code", "", "The AWS Region code of the resource")
	trustedadvisor_listRecommendationResourcesCmd.Flags().String("status", "", "The status of the resource")
	trustedadvisor_listRecommendationResourcesCmd.MarkFlagRequired("recommendation-identifier")
	trustedadvisorCmd.AddCommand(trustedadvisor_listRecommendationResourcesCmd)
}
