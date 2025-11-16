package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalizeRuntime_getPersonalizedRankingCmd = &cobra.Command{
	Use:   "get-personalized-ranking",
	Short: "Re-ranks a list of recommended items for the given user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalizeRuntime_getPersonalizedRankingCmd).Standalone()

	personalizeRuntime_getPersonalizedRankingCmd.Flags().String("campaign-arn", "", "The Amazon Resource Name (ARN) of the campaign to use for generating the personalized ranking.")
	personalizeRuntime_getPersonalizedRankingCmd.Flags().String("context", "", "The contextual metadata to use when getting recommendations.")
	personalizeRuntime_getPersonalizedRankingCmd.Flags().String("filter-arn", "", "The Amazon Resource Name (ARN) of a filter you created to include items or exclude items from recommendations for a given user.")
	personalizeRuntime_getPersonalizedRankingCmd.Flags().String("filter-values", "", "The values to use when filtering recommendations.")
	personalizeRuntime_getPersonalizedRankingCmd.Flags().String("input-list", "", "A list of items (by `itemId`) to rank.")
	personalizeRuntime_getPersonalizedRankingCmd.Flags().String("metadata-columns", "", "If you enabled metadata in recommendations when you created or updated the campaign, specify metadata columns from your Items dataset to include in the personalized ranking.")
	personalizeRuntime_getPersonalizedRankingCmd.Flags().String("user-id", "", "The user for which you want the campaign to provide a personalized ranking.")
	personalizeRuntime_getPersonalizedRankingCmd.MarkFlagRequired("campaign-arn")
	personalizeRuntime_getPersonalizedRankingCmd.MarkFlagRequired("input-list")
	personalizeRuntime_getPersonalizedRankingCmd.MarkFlagRequired("user-id")
	personalizeRuntimeCmd.AddCommand(personalizeRuntime_getPersonalizedRankingCmd)
}
