package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_stopTrainingJobCmd = &cobra.Command{
	Use:   "stop-training-job",
	Short: "Stops a training job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_stopTrainingJobCmd).Standalone()

	sagemaker_stopTrainingJobCmd.Flags().String("training-job-name", "", "The name of the training job to stop.")
	sagemaker_stopTrainingJobCmd.MarkFlagRequired("training-job-name")
	sagemakerCmd.AddCommand(sagemaker_stopTrainingJobCmd)
}
