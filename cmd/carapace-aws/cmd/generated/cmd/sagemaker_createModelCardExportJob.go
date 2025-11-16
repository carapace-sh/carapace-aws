package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createModelCardExportJobCmd = &cobra.Command{
	Use:   "create-model-card-export-job",
	Short: "Creates an Amazon SageMaker Model Card export job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createModelCardExportJobCmd).Standalone()

	sagemaker_createModelCardExportJobCmd.Flags().String("model-card-export-job-name", "", "The name of the model card export job.")
	sagemaker_createModelCardExportJobCmd.Flags().String("model-card-name", "", "The name or Amazon Resource Name (ARN) of the model card to export.")
	sagemaker_createModelCardExportJobCmd.Flags().String("model-card-version", "", "The version of the model card to export.")
	sagemaker_createModelCardExportJobCmd.Flags().String("output-config", "", "The model card output configuration that specifies the Amazon S3 path for exporting.")
	sagemaker_createModelCardExportJobCmd.MarkFlagRequired("model-card-export-job-name")
	sagemaker_createModelCardExportJobCmd.MarkFlagRequired("model-card-name")
	sagemaker_createModelCardExportJobCmd.MarkFlagRequired("output-config")
	sagemakerCmd.AddCommand(sagemaker_createModelCardExportJobCmd)
}
