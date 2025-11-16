package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var costOptimizationHub_listRecommendationSummariesCmd = &cobra.Command{
	Use:   "list-recommendation-summaries",
	Short: "Returns a concise representation of savings estimates for resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(costOptimizationHub_listRecommendationSummariesCmd).Standalone()

	costOptimizationHub_listRecommendationSummariesCmd.Flags().String("filter", "", "")
	costOptimizationHub_listRecommendationSummariesCmd.Flags().String("group-by", "", "The grouping of recommendations by a dimension.")
	costOptimizationHub_listRecommendationSummariesCmd.Flags().String("max-results", "", "The maximum number of recommendations to be returned for the request.")
	costOptimizationHub_listRecommendationSummariesCmd.Flags().String("metrics", "", "Additional metrics to be returned for the request.")
	costOptimizationHub_listRecommendationSummariesCmd.Flags().String("next-token", "", "The token to retrieve the next set of results.")
	costOptimizationHub_listRecommendationSummariesCmd.MarkFlagRequired("group-by")
	costOptimizationHubCmd.AddCommand(costOptimizationHub_listRecommendationSummariesCmd)
}
