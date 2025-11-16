package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var computeOptimizer_getAutoScalingGroupRecommendationsCmd = &cobra.Command{
	Use:   "get-auto-scaling-group-recommendations",
	Short: "Returns Auto Scaling group recommendations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(computeOptimizer_getAutoScalingGroupRecommendationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(computeOptimizer_getAutoScalingGroupRecommendationsCmd).Standalone()

		computeOptimizer_getAutoScalingGroupRecommendationsCmd.Flags().String("account-ids", "", "The ID of the Amazon Web Services account for which to return Auto Scaling group recommendations.")
		computeOptimizer_getAutoScalingGroupRecommendationsCmd.Flags().String("auto-scaling-group-arns", "", "The Amazon Resource Name (ARN) of the Auto Scaling groups for which to return recommendations.")
		computeOptimizer_getAutoScalingGroupRecommendationsCmd.Flags().String("filters", "", "An array of objects to specify a filter that returns a more specific list of Auto Scaling group recommendations.")
		computeOptimizer_getAutoScalingGroupRecommendationsCmd.Flags().String("max-results", "", "The maximum number of Auto Scaling group recommendations to return with a single request.")
		computeOptimizer_getAutoScalingGroupRecommendationsCmd.Flags().String("next-token", "", "The token to advance to the next page of Auto Scaling group recommendations.")
		computeOptimizer_getAutoScalingGroupRecommendationsCmd.Flags().String("recommendation-preferences", "", "An object to specify the preferences for the Auto Scaling group recommendations to return in the response.")
	})
	computeOptimizerCmd.AddCommand(computeOptimizer_getAutoScalingGroupRecommendationsCmd)
}
