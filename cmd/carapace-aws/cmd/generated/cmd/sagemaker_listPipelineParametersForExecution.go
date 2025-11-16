package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listPipelineParametersForExecutionCmd = &cobra.Command{
	Use:   "list-pipeline-parameters-for-execution",
	Short: "Gets a list of parameters for a pipeline execution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listPipelineParametersForExecutionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listPipelineParametersForExecutionCmd).Standalone()

		sagemaker_listPipelineParametersForExecutionCmd.Flags().String("max-results", "", "The maximum number of parameters to return in the response.")
		sagemaker_listPipelineParametersForExecutionCmd.Flags().String("next-token", "", "If the result of the previous `ListPipelineParametersForExecution` request was truncated, the response includes a `NextToken`.")
		sagemaker_listPipelineParametersForExecutionCmd.Flags().String("pipeline-execution-arn", "", "The Amazon Resource Name (ARN) of the pipeline execution.")
		sagemaker_listPipelineParametersForExecutionCmd.MarkFlagRequired("pipeline-execution-arn")
	})
	sagemakerCmd.AddCommand(sagemaker_listPipelineParametersForExecutionCmd)
}
