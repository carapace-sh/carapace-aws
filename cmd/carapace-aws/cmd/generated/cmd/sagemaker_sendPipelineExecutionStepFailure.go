package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_sendPipelineExecutionStepFailureCmd = &cobra.Command{
	Use:   "send-pipeline-execution-step-failure",
	Short: "Notifies the pipeline that the execution of a callback step failed, along with a message describing why.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_sendPipelineExecutionStepFailureCmd).Standalone()

	sagemaker_sendPipelineExecutionStepFailureCmd.Flags().String("callback-token", "", "The pipeline generated token from the Amazon SQS queue.")
	sagemaker_sendPipelineExecutionStepFailureCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the operation.")
	sagemaker_sendPipelineExecutionStepFailureCmd.Flags().String("failure-reason", "", "A message describing why the step failed.")
	sagemaker_sendPipelineExecutionStepFailureCmd.MarkFlagRequired("callback-token")
	sagemakerCmd.AddCommand(sagemaker_sendPipelineExecutionStepFailureCmd)
}
