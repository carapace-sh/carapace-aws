package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeModelCardExportJobCmd = &cobra.Command{
	Use:   "describe-model-card-export-job",
	Short: "Describes an Amazon SageMaker Model Card export job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeModelCardExportJobCmd).Standalone()

	sagemaker_describeModelCardExportJobCmd.Flags().String("model-card-export-job-arn", "", "The Amazon Resource Name (ARN) of the model card export job to describe.")
	sagemaker_describeModelCardExportJobCmd.MarkFlagRequired("model-card-export-job-arn")
	sagemakerCmd.AddCommand(sagemaker_describeModelCardExportJobCmd)
}
