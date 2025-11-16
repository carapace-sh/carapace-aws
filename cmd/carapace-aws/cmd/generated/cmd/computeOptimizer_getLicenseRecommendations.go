package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var computeOptimizer_getLicenseRecommendationsCmd = &cobra.Command{
	Use:   "get-license-recommendations",
	Short: "Returns license recommendations for Amazon EC2 instances that run on a specific license.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(computeOptimizer_getLicenseRecommendationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(computeOptimizer_getLicenseRecommendationsCmd).Standalone()

		computeOptimizer_getLicenseRecommendationsCmd.Flags().String("account-ids", "", "The ID of the Amazon Web Services account for which to return license recommendations.")
		computeOptimizer_getLicenseRecommendationsCmd.Flags().String("filters", "", "An array of objects to specify a filter that returns a more specific list of license recommendations.")
		computeOptimizer_getLicenseRecommendationsCmd.Flags().String("max-results", "", "The maximum number of license recommendations to return with a single request.")
		computeOptimizer_getLicenseRecommendationsCmd.Flags().String("next-token", "", "The token to advance to the next page of license recommendations.")
		computeOptimizer_getLicenseRecommendationsCmd.Flags().String("resource-arns", "", "The ARN that identifies the Amazon EC2 instance.")
	})
	computeOptimizerCmd.AddCommand(computeOptimizer_getLicenseRecommendationsCmd)
}
