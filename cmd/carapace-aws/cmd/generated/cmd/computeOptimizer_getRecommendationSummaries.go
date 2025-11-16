package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var computeOptimizer_getRecommendationSummariesCmd = &cobra.Command{
	Use:   "get-recommendation-summaries",
	Short: "Returns the optimization findings for an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(computeOptimizer_getRecommendationSummariesCmd).Standalone()

	computeOptimizer_getRecommendationSummariesCmd.Flags().String("account-ids", "", "The ID of the Amazon Web Services account for which to return recommendation summaries.")
	computeOptimizer_getRecommendationSummariesCmd.Flags().String("max-results", "", "The maximum number of recommendation summaries to return with a single request.")
	computeOptimizer_getRecommendationSummariesCmd.Flags().String("next-token", "", "The token to advance to the next page of recommendation summaries.")
	computeOptimizerCmd.AddCommand(computeOptimizer_getRecommendationSummariesCmd)
}
