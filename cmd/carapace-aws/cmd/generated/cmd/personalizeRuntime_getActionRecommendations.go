package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalizeRuntime_getActionRecommendationsCmd = &cobra.Command{
	Use:   "get-action-recommendations",
	Short: "Returns a list of recommended actions in sorted in descending order by prediction score.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalizeRuntime_getActionRecommendationsCmd).Standalone()

	personalizeRuntime_getActionRecommendationsCmd.Flags().String("campaign-arn", "", "The Amazon Resource Name (ARN) of the campaign to use for getting action recommendations.")
	personalizeRuntime_getActionRecommendationsCmd.Flags().String("filter-arn", "", "The ARN of the filter to apply to the returned recommendations.")
	personalizeRuntime_getActionRecommendationsCmd.Flags().String("filter-values", "", "The values to use when filtering recommendations.")
	personalizeRuntime_getActionRecommendationsCmd.Flags().String("num-results", "", "The number of results to return.")
	personalizeRuntime_getActionRecommendationsCmd.Flags().String("user-id", "", "The user ID of the user to provide action recommendations for.")
	personalizeRuntimeCmd.AddCommand(personalizeRuntime_getActionRecommendationsCmd)
}
