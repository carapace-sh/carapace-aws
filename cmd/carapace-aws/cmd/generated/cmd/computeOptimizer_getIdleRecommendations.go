package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var computeOptimizer_getIdleRecommendationsCmd = &cobra.Command{
	Use:   "get-idle-recommendations",
	Short: "Returns idle resource recommendations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(computeOptimizer_getIdleRecommendationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(computeOptimizer_getIdleRecommendationsCmd).Standalone()

		computeOptimizer_getIdleRecommendationsCmd.Flags().String("account-ids", "", "Return the idle resource recommendations to the specified Amazon Web Services account IDs.")
		computeOptimizer_getIdleRecommendationsCmd.Flags().String("filters", "", "An array of objects to specify a filter that returns a more specific list of idle resource recommendations.")
		computeOptimizer_getIdleRecommendationsCmd.Flags().String("max-results", "", "The maximum number of idle resource recommendations to return with a single request.")
		computeOptimizer_getIdleRecommendationsCmd.Flags().String("next-token", "", "The token to advance to the next page of idle resource recommendations.")
		computeOptimizer_getIdleRecommendationsCmd.Flags().String("order-by", "", "The order to sort the idle resource recommendations.")
		computeOptimizer_getIdleRecommendationsCmd.Flags().String("resource-arns", "", "The ARN that identifies the idle resource.")
	})
	computeOptimizerCmd.AddCommand(computeOptimizer_getIdleRecommendationsCmd)
}
