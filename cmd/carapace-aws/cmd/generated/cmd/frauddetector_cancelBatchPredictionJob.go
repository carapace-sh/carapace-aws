package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_cancelBatchPredictionJobCmd = &cobra.Command{
	Use:   "cancel-batch-prediction-job",
	Short: "Cancels the specified batch prediction job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_cancelBatchPredictionJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(frauddetector_cancelBatchPredictionJobCmd).Standalone()

		frauddetector_cancelBatchPredictionJobCmd.Flags().String("job-id", "", "The ID of the batch prediction job to cancel.")
		frauddetector_cancelBatchPredictionJobCmd.MarkFlagRequired("job-id")
	})
	frauddetectorCmd.AddCommand(frauddetector_cancelBatchPredictionJobCmd)
}
