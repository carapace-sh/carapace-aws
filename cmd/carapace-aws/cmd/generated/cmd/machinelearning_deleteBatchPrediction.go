package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var machinelearning_deleteBatchPredictionCmd = &cobra.Command{
	Use:   "delete-batch-prediction",
	Short: "Assigns the DELETED status to a `BatchPrediction`, rendering it unusable.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machinelearning_deleteBatchPredictionCmd).Standalone()

	machinelearning_deleteBatchPredictionCmd.Flags().String("batch-prediction-id", "", "A user-supplied ID that uniquely identifies the `BatchPrediction`.")
	machinelearning_deleteBatchPredictionCmd.MarkFlagRequired("batch-prediction-id")
	machinelearningCmd.AddCommand(machinelearning_deleteBatchPredictionCmd)
}
