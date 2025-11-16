package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_stopPipelineExecutionCmd = &cobra.Command{
	Use:   "stop-pipeline-execution",
	Short: "Stops a pipeline execution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_stopPipelineExecutionCmd).Standalone()

	sagemaker_stopPipelineExecutionCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the operation.")
	sagemaker_stopPipelineExecutionCmd.Flags().String("pipeline-execution-arn", "", "The Amazon Resource Name (ARN) of the pipeline execution.")
	sagemaker_stopPipelineExecutionCmd.MarkFlagRequired("client-request-token")
	sagemaker_stopPipelineExecutionCmd.MarkFlagRequired("pipeline-execution-arn")
	sagemakerCmd.AddCommand(sagemaker_stopPipelineExecutionCmd)
}
