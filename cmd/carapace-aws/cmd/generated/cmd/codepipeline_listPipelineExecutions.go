package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codepipeline_listPipelineExecutionsCmd = &cobra.Command{
	Use:   "list-pipeline-executions",
	Short: "Gets a summary of the most recent executions for a pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codepipeline_listPipelineExecutionsCmd).Standalone()

	codepipeline_listPipelineExecutionsCmd.Flags().String("filter", "", "The pipeline execution to filter on.")
	codepipeline_listPipelineExecutionsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
	codepipeline_listPipelineExecutionsCmd.Flags().String("next-token", "", "The token that was returned from the previous `ListPipelineExecutions` call, which can be used to return the next set of pipeline executions in the list.")
	codepipeline_listPipelineExecutionsCmd.Flags().String("pipeline-name", "", "The name of the pipeline for which you want to get execution summary information.")
	codepipeline_listPipelineExecutionsCmd.MarkFlagRequired("pipeline-name")
	codepipelineCmd.AddCommand(codepipeline_listPipelineExecutionsCmd)
}
