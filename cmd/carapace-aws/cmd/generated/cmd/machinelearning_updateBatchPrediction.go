package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var machinelearning_updateBatchPredictionCmd = &cobra.Command{
	Use:   "update-batch-prediction",
	Short: "Updates the `BatchPredictionName` of a `BatchPrediction`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machinelearning_updateBatchPredictionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(machinelearning_updateBatchPredictionCmd).Standalone()

		machinelearning_updateBatchPredictionCmd.Flags().String("batch-prediction-id", "", "The ID assigned to the `BatchPrediction` during creation.")
		machinelearning_updateBatchPredictionCmd.Flags().String("batch-prediction-name", "", "A new user-supplied name or description of the `BatchPrediction`.")
		machinelearning_updateBatchPredictionCmd.MarkFlagRequired("batch-prediction-id")
		machinelearning_updateBatchPredictionCmd.MarkFlagRequired("batch-prediction-name")
	})
	machinelearningCmd.AddCommand(machinelearning_updateBatchPredictionCmd)
}
