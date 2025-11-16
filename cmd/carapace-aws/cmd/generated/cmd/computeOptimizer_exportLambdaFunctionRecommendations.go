package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var computeOptimizer_exportLambdaFunctionRecommendationsCmd = &cobra.Command{
	Use:   "export-lambda-function-recommendations",
	Short: "Exports optimization recommendations for Lambda functions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(computeOptimizer_exportLambdaFunctionRecommendationsCmd).Standalone()

	computeOptimizer_exportLambdaFunctionRecommendationsCmd.Flags().String("account-ids", "", "The IDs of the Amazon Web Services accounts for which to export Lambda function recommendations.")
	computeOptimizer_exportLambdaFunctionRecommendationsCmd.Flags().String("fields-to-export", "", "The recommendations data to include in the export file.")
	computeOptimizer_exportLambdaFunctionRecommendationsCmd.Flags().String("file-format", "", "The format of the export file.")
	computeOptimizer_exportLambdaFunctionRecommendationsCmd.Flags().String("filters", "", "An array of objects to specify a filter that exports a more specific set of Lambda function recommendations.")
	computeOptimizer_exportLambdaFunctionRecommendationsCmd.Flags().String("include-member-accounts", "", "Indicates whether to include recommendations for resources in all member accounts of the organization if your account is the management account of an organization.")
	computeOptimizer_exportLambdaFunctionRecommendationsCmd.Flags().String("s3-destination-config", "", "")
	computeOptimizer_exportLambdaFunctionRecommendationsCmd.MarkFlagRequired("s3-destination-config")
	computeOptimizerCmd.AddCommand(computeOptimizer_exportLambdaFunctionRecommendationsCmd)
}
