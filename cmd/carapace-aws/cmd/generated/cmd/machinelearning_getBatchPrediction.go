package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var machinelearning_getBatchPredictionCmd = &cobra.Command{
	Use:   "get-batch-prediction",
	Short: "Returns a `BatchPrediction` that includes detailed metadata, status, and data file information for a `Batch Prediction` request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machinelearning_getBatchPredictionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(machinelearning_getBatchPredictionCmd).Standalone()

		machinelearning_getBatchPredictionCmd.Flags().String("batch-prediction-id", "", "An ID assigned to the `BatchPrediction` at creation.")
		machinelearning_getBatchPredictionCmd.MarkFlagRequired("batch-prediction-id")
	})
	machinelearningCmd.AddCommand(machinelearning_getBatchPredictionCmd)
}
