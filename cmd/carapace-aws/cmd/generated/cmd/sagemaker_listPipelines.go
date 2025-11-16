package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listPipelinesCmd = &cobra.Command{
	Use:   "list-pipelines",
	Short: "Gets a list of pipelines.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listPipelinesCmd).Standalone()

	sagemaker_listPipelinesCmd.Flags().String("created-after", "", "A filter that returns the pipelines that were created after a specified time.")
	sagemaker_listPipelinesCmd.Flags().String("created-before", "", "A filter that returns the pipelines that were created before a specified time.")
	sagemaker_listPipelinesCmd.Flags().String("max-results", "", "The maximum number of pipelines to return in the response.")
	sagemaker_listPipelinesCmd.Flags().String("next-token", "", "If the result of the previous `ListPipelines` request was truncated, the response includes a `NextToken`.")
	sagemaker_listPipelinesCmd.Flags().String("pipeline-name-prefix", "", "The prefix of the pipeline name.")
	sagemaker_listPipelinesCmd.Flags().String("sort-by", "", "The field by which to sort results.")
	sagemaker_listPipelinesCmd.Flags().String("sort-order", "", "The sort order for results.")
	sagemakerCmd.AddCommand(sagemaker_listPipelinesCmd)
}
