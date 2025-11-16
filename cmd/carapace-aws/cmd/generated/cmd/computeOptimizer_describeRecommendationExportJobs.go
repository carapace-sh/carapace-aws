package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var computeOptimizer_describeRecommendationExportJobsCmd = &cobra.Command{
	Use:   "describe-recommendation-export-jobs",
	Short: "Describes recommendation export jobs created in the last seven days.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(computeOptimizer_describeRecommendationExportJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(computeOptimizer_describeRecommendationExportJobsCmd).Standalone()

		computeOptimizer_describeRecommendationExportJobsCmd.Flags().String("filters", "", "An array of objects to specify a filter that returns a more specific list of export jobs.")
		computeOptimizer_describeRecommendationExportJobsCmd.Flags().String("job-ids", "", "The identification numbers of the export jobs to return.")
		computeOptimizer_describeRecommendationExportJobsCmd.Flags().String("max-results", "", "The maximum number of export jobs to return with a single request.")
		computeOptimizer_describeRecommendationExportJobsCmd.Flags().String("next-token", "", "The token to advance to the next page of export jobs.")
	})
	computeOptimizerCmd.AddCommand(computeOptimizer_describeRecommendationExportJobsCmd)
}
