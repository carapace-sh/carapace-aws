package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalizeRuntime_getRecommendationsCmd = &cobra.Command{
	Use:   "get-recommendations",
	Short: "Returns a list of recommended items.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalizeRuntime_getRecommendationsCmd).Standalone()

	personalizeRuntime_getRecommendationsCmd.Flags().String("campaign-arn", "", "The Amazon Resource Name (ARN) of the campaign to use for getting recommendations.")
	personalizeRuntime_getRecommendationsCmd.Flags().String("context", "", "The contextual metadata to use when getting recommendations.")
	personalizeRuntime_getRecommendationsCmd.Flags().String("filter-arn", "", "The ARN of the filter to apply to the returned recommendations.")
	personalizeRuntime_getRecommendationsCmd.Flags().String("filter-values", "", "The values to use when filtering recommendations.")
	personalizeRuntime_getRecommendationsCmd.Flags().String("item-id", "", "The item ID to provide recommendations for.")
	personalizeRuntime_getRecommendationsCmd.Flags().String("metadata-columns", "", "If you enabled metadata in recommendations when you created or updated the campaign or recommender, specify the metadata columns from your Items dataset to include in item recommendations.")
	personalizeRuntime_getRecommendationsCmd.Flags().String("num-results", "", "The number of results to return.")
	personalizeRuntime_getRecommendationsCmd.Flags().String("promotions", "", "The promotions to apply to the recommendation request.")
	personalizeRuntime_getRecommendationsCmd.Flags().String("recommender-arn", "", "The Amazon Resource Name (ARN) of the recommender to use to get recommendations.")
	personalizeRuntime_getRecommendationsCmd.Flags().String("user-id", "", "The user ID to provide recommendations for.")
	personalizeRuntimeCmd.AddCommand(personalizeRuntime_getRecommendationsCmd)
}
