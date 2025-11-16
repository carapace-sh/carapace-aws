package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeTrainingJobCmd = &cobra.Command{
	Use:   "describe-training-job",
	Short: "Returns information about a training job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeTrainingJobCmd).Standalone()

	sagemaker_describeTrainingJobCmd.Flags().String("training-job-name", "", "The name of the training job.")
	sagemaker_describeTrainingJobCmd.MarkFlagRequired("training-job-name")
	sagemakerCmd.AddCommand(sagemaker_describeTrainingJobCmd)
}
