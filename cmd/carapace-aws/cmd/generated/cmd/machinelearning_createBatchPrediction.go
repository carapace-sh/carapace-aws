package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var machinelearning_createBatchPredictionCmd = &cobra.Command{
	Use:   "create-batch-prediction",
	Short: "Generates predictions for a group of observations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machinelearning_createBatchPredictionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(machinelearning_createBatchPredictionCmd).Standalone()

		machinelearning_createBatchPredictionCmd.Flags().String("batch-prediction-data-source-id", "", "The ID of the `DataSource` that points to the group of observations to predict.")
		machinelearning_createBatchPredictionCmd.Flags().String("batch-prediction-id", "", "A user-supplied ID that uniquely identifies the `BatchPrediction`.")
		machinelearning_createBatchPredictionCmd.Flags().String("batch-prediction-name", "", "A user-supplied name or description of the `BatchPrediction`.")
		machinelearning_createBatchPredictionCmd.Flags().String("mlmodel-id", "", "The ID of the `MLModel` that will generate predictions for the group of observations.")
		machinelearning_createBatchPredictionCmd.Flags().String("output-uri", "", "The location of an Amazon Simple Storage Service (Amazon S3) bucket or directory to store the batch prediction results.")
		machinelearning_createBatchPredictionCmd.MarkFlagRequired("batch-prediction-data-source-id")
		machinelearning_createBatchPredictionCmd.MarkFlagRequired("batch-prediction-id")
		machinelearning_createBatchPredictionCmd.MarkFlagRequired("mlmodel-id")
		machinelearning_createBatchPredictionCmd.MarkFlagRequired("output-uri")
	})
	machinelearningCmd.AddCommand(machinelearning_createBatchPredictionCmd)
}
