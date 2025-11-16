package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_sendPipelineExecutionStepSuccessCmd = &cobra.Command{
	Use:   "send-pipeline-execution-step-success",
	Short: "Notifies the pipeline that the execution of a callback step succeeded and provides a list of the step's output parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_sendPipelineExecutionStepSuccessCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_sendPipelineExecutionStepSuccessCmd).Standalone()

		sagemaker_sendPipelineExecutionStepSuccessCmd.Flags().String("callback-token", "", "The pipeline generated token from the Amazon SQS queue.")
		sagemaker_sendPipelineExecutionStepSuccessCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the operation.")
		sagemaker_sendPipelineExecutionStepSuccessCmd.Flags().String("output-parameters", "", "A list of the output parameters of the callback step.")
		sagemaker_sendPipelineExecutionStepSuccessCmd.MarkFlagRequired("callback-token")
	})
	sagemakerCmd.AddCommand(sagemaker_sendPipelineExecutionStepSuccessCmd)
}
