package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listPipelineExecutionsCmd = &cobra.Command{
	Use:   "list-pipeline-executions",
	Short: "Gets a list of the pipeline executions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listPipelineExecutionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listPipelineExecutionsCmd).Standalone()

		sagemaker_listPipelineExecutionsCmd.Flags().String("created-after", "", "A filter that returns the pipeline executions that were created after a specified time.")
		sagemaker_listPipelineExecutionsCmd.Flags().String("created-before", "", "A filter that returns the pipeline executions that were created before a specified time.")
		sagemaker_listPipelineExecutionsCmd.Flags().String("max-results", "", "The maximum number of pipeline executions to return in the response.")
		sagemaker_listPipelineExecutionsCmd.Flags().String("next-token", "", "If the result of the previous `ListPipelineExecutions` request was truncated, the response includes a `NextToken`.")
		sagemaker_listPipelineExecutionsCmd.Flags().String("pipeline-name", "", "The name or Amazon Resource Name (ARN) of the pipeline.")
		sagemaker_listPipelineExecutionsCmd.Flags().String("sort-by", "", "The field by which to sort results.")
		sagemaker_listPipelineExecutionsCmd.Flags().String("sort-order", "", "The sort order for results.")
		sagemaker_listPipelineExecutionsCmd.MarkFlagRequired("pipeline-name")
	})
	sagemakerCmd.AddCommand(sagemaker_listPipelineExecutionsCmd)
}
