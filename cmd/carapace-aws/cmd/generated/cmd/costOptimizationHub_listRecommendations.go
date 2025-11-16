package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var costOptimizationHub_listRecommendationsCmd = &cobra.Command{
	Use:   "list-recommendations",
	Short: "Returns a list of recommendations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(costOptimizationHub_listRecommendationsCmd).Standalone()

	costOptimizationHub_listRecommendationsCmd.Flags().String("filter", "", "The constraints that you want all returned recommendations to match.")
	costOptimizationHub_listRecommendationsCmd.Flags().String("include-all-recommendations", "", "List of all recommendations for a resource, or a single recommendation if de-duped by `resourceId`.")
	costOptimizationHub_listRecommendationsCmd.Flags().String("max-results", "", "The maximum number of recommendations that are returned for the request.")
	costOptimizationHub_listRecommendationsCmd.Flags().String("next-token", "", "The token to retrieve the next set of results.")
	costOptimizationHub_listRecommendationsCmd.Flags().String("order-by", "", "The ordering of recommendations by a dimension.")
	costOptimizationHubCmd.AddCommand(costOptimizationHub_listRecommendationsCmd)
}
