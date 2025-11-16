package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteTrainingJobCmd = &cobra.Command{
	Use:   "delete-training-job",
	Short: "Deletes a training job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteTrainingJobCmd).Standalone()

	sagemaker_deleteTrainingJobCmd.Flags().String("training-job-name", "", "The name of the training job to delete.")
	sagemaker_deleteTrainingJobCmd.MarkFlagRequired("training-job-name")
	sagemakerCmd.AddCommand(sagemaker_deleteTrainingJobCmd)
}
