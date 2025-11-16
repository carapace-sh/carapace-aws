package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listPipelineVersionsCmd = &cobra.Command{
	Use:   "list-pipeline-versions",
	Short: "Gets a list of all versions of the pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listPipelineVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listPipelineVersionsCmd).Standalone()

		sagemaker_listPipelineVersionsCmd.Flags().String("created-after", "", "A filter that returns the pipeline versions that were created after a specified time.")
		sagemaker_listPipelineVersionsCmd.Flags().String("created-before", "", "A filter that returns the pipeline versions that were created before a specified time.")
		sagemaker_listPipelineVersionsCmd.Flags().String("max-results", "", "The maximum number of pipeline versions to return in the response.")
		sagemaker_listPipelineVersionsCmd.Flags().String("next-token", "", "If the result of the previous `ListPipelineVersions` request was truncated, the response includes a `NextToken`.")
		sagemaker_listPipelineVersionsCmd.Flags().String("pipeline-name", "", "The Amazon Resource Name (ARN) of the pipeline.")
		sagemaker_listPipelineVersionsCmd.Flags().String("sort-order", "", "The sort order for the results.")
		sagemaker_listPipelineVersionsCmd.MarkFlagRequired("pipeline-name")
	})
	sagemakerCmd.AddCommand(sagemaker_listPipelineVersionsCmd)
}
