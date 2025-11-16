package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_retryPipelineExecutionCmd = &cobra.Command{
	Use:   "retry-pipeline-execution",
	Short: "Retry the execution of the pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_retryPipelineExecutionCmd).Standalone()

	sagemaker_retryPipelineExecutionCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the operation.")
	sagemaker_retryPipelineExecutionCmd.Flags().String("parallelism-configuration", "", "This configuration, if specified, overrides the parallelism configuration of the parent pipeline.")
	sagemaker_retryPipelineExecutionCmd.Flags().String("pipeline-execution-arn", "", "The Amazon Resource Name (ARN) of the pipeline execution.")
	sagemaker_retryPipelineExecutionCmd.MarkFlagRequired("client-request-token")
	sagemaker_retryPipelineExecutionCmd.MarkFlagRequired("pipeline-execution-arn")
	sagemakerCmd.AddCommand(sagemaker_retryPipelineExecutionCmd)
}
