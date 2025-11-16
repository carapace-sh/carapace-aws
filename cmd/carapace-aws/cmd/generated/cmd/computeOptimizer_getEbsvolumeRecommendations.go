package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var computeOptimizer_getEbsvolumeRecommendationsCmd = &cobra.Command{
	Use:   "get-ebsvolume-recommendations",
	Short: "Returns Amazon Elastic Block Store (Amazon EBS) volume recommendations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(computeOptimizer_getEbsvolumeRecommendationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(computeOptimizer_getEbsvolumeRecommendationsCmd).Standalone()

		computeOptimizer_getEbsvolumeRecommendationsCmd.Flags().String("account-ids", "", "The ID of the Amazon Web Services account for which to return volume recommendations.")
		computeOptimizer_getEbsvolumeRecommendationsCmd.Flags().String("filters", "", "An array of objects to specify a filter that returns a more specific list of volume recommendations.")
		computeOptimizer_getEbsvolumeRecommendationsCmd.Flags().String("max-results", "", "The maximum number of volume recommendations to return with a single request.")
		computeOptimizer_getEbsvolumeRecommendationsCmd.Flags().String("next-token", "", "The token to advance to the next page of volume recommendations.")
		computeOptimizer_getEbsvolumeRecommendationsCmd.Flags().String("volume-arns", "", "The Amazon Resource Name (ARN) of the volumes for which to return recommendations.")
	})
	computeOptimizerCmd.AddCommand(computeOptimizer_getEbsvolumeRecommendationsCmd)
}
