package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createPipelineCmd = &cobra.Command{
	Use:   "create-pipeline",
	Short: "Creates a pipeline using a JSON pipeline definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createPipelineCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_createPipelineCmd).Standalone()

		sagemaker_createPipelineCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the operation.")
		sagemaker_createPipelineCmd.Flags().String("parallelism-configuration", "", "This is the configuration that controls the parallelism of the pipeline.")
		sagemaker_createPipelineCmd.Flags().String("pipeline-definition", "", "The [JSON pipeline definition](https://aws-sagemaker-mlops.github.io/sagemaker-model-building-pipeline-definition-JSON-schema/) of the pipeline.")
		sagemaker_createPipelineCmd.Flags().String("pipeline-definition-s3-location", "", "The location of the pipeline definition stored in Amazon S3.")
		sagemaker_createPipelineCmd.Flags().String("pipeline-description", "", "A description of the pipeline.")
		sagemaker_createPipelineCmd.Flags().String("pipeline-display-name", "", "The display name of the pipeline.")
		sagemaker_createPipelineCmd.Flags().String("pipeline-name", "", "The name of the pipeline.")
		sagemaker_createPipelineCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the role used by the pipeline to access and create resources.")
		sagemaker_createPipelineCmd.Flags().String("tags", "", "A list of tags to apply to the created pipeline.")
		sagemaker_createPipelineCmd.MarkFlagRequired("client-request-token")
		sagemaker_createPipelineCmd.MarkFlagRequired("pipeline-name")
		sagemaker_createPipelineCmd.MarkFlagRequired("role-arn")
	})
	sagemakerCmd.AddCommand(sagemaker_createPipelineCmd)
}
