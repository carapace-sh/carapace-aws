package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_getBatchPredictionJobsCmd = &cobra.Command{
	Use:   "get-batch-prediction-jobs",
	Short: "Gets all batch prediction jobs or a specific job if you specify a job ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_getBatchPredictionJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(frauddetector_getBatchPredictionJobsCmd).Standalone()

		frauddetector_getBatchPredictionJobsCmd.Flags().String("job-id", "", "The batch prediction job for which to get the details.")
		frauddetector_getBatchPredictionJobsCmd.Flags().String("max-results", "", "The maximum number of objects to return for the request.")
		frauddetector_getBatchPredictionJobsCmd.Flags().String("next-token", "", "The next token from the previous request.")
	})
	frauddetectorCmd.AddCommand(frauddetector_getBatchPredictionJobsCmd)
}
