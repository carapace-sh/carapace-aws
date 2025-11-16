package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listPipelineExecutionStepsCmd = &cobra.Command{
	Use:   "list-pipeline-execution-steps",
	Short: "Gets a list of `PipeLineExecutionStep` objects.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listPipelineExecutionStepsCmd).Standalone()

	sagemaker_listPipelineExecutionStepsCmd.Flags().String("max-results", "", "The maximum number of pipeline execution steps to return in the response.")
	sagemaker_listPipelineExecutionStepsCmd.Flags().String("next-token", "", "If the result of the previous `ListPipelineExecutionSteps` request was truncated, the response includes a `NextToken`.")
	sagemaker_listPipelineExecutionStepsCmd.Flags().String("pipeline-execution-arn", "", "The Amazon Resource Name (ARN) of the pipeline execution.")
	sagemaker_listPipelineExecutionStepsCmd.Flags().String("sort-order", "", "The field by which to sort results.")
	sagemakerCmd.AddCommand(sagemaker_listPipelineExecutionStepsCmd)
}
