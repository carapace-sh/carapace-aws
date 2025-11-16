package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var computeOptimizer_getLambdaFunctionRecommendationsCmd = &cobra.Command{
	Use:   "get-lambda-function-recommendations",
	Short: "Returns Lambda function recommendations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(computeOptimizer_getLambdaFunctionRecommendationsCmd).Standalone()

	computeOptimizer_getLambdaFunctionRecommendationsCmd.Flags().String("account-ids", "", "The ID of the Amazon Web Services account for which to return function recommendations.")
	computeOptimizer_getLambdaFunctionRecommendationsCmd.Flags().String("filters", "", "An array of objects to specify a filter that returns a more specific list of function recommendations.")
	computeOptimizer_getLambdaFunctionRecommendationsCmd.Flags().String("function-arns", "", "The Amazon Resource Name (ARN) of the functions for which to return recommendations.")
	computeOptimizer_getLambdaFunctionRecommendationsCmd.Flags().String("max-results", "", "The maximum number of function recommendations to return with a single request.")
	computeOptimizer_getLambdaFunctionRecommendationsCmd.Flags().String("next-token", "", "The token to advance to the next page of function recommendations.")
	computeOptimizerCmd.AddCommand(computeOptimizer_getLambdaFunctionRecommendationsCmd)
}
