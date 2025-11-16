package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var computeOptimizer_getEcsserviceRecommendationsCmd = &cobra.Command{
	Use:   "get-ecsservice-recommendations",
	Short: "Returns Amazon ECS service recommendations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(computeOptimizer_getEcsserviceRecommendationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(computeOptimizer_getEcsserviceRecommendationsCmd).Standalone()

		computeOptimizer_getEcsserviceRecommendationsCmd.Flags().String("account-ids", "", "Return the Amazon ECS service recommendations to the specified Amazon Web Services account IDs.")
		computeOptimizer_getEcsserviceRecommendationsCmd.Flags().String("filters", "", "An array of objects to specify a filter that returns a more specific list of Amazon ECS service recommendations.")
		computeOptimizer_getEcsserviceRecommendationsCmd.Flags().String("max-results", "", "The maximum number of Amazon ECS service recommendations to return with a single request.")
		computeOptimizer_getEcsserviceRecommendationsCmd.Flags().String("next-token", "", "The token to advance to the next page of Amazon ECS service recommendations.")
		computeOptimizer_getEcsserviceRecommendationsCmd.Flags().String("service-arns", "", "The ARN that identifies the Amazon ECS service.")
	})
	computeOptimizerCmd.AddCommand(computeOptimizer_getEcsserviceRecommendationsCmd)
}
