package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_deleteBatchPredictionJobCmd = &cobra.Command{
	Use:   "delete-batch-prediction-job",
	Short: "Deletes a batch prediction job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_deleteBatchPredictionJobCmd).Standalone()

	frauddetector_deleteBatchPredictionJobCmd.Flags().String("job-id", "", "The ID of the batch prediction job to delete.")
	frauddetector_deleteBatchPredictionJobCmd.MarkFlagRequired("job-id")
	frauddetectorCmd.AddCommand(frauddetector_deleteBatchPredictionJobCmd)
}
