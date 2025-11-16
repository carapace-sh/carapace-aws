package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codepipeline_listActionExecutionsCmd = &cobra.Command{
	Use:   "list-action-executions",
	Short: "Lists the action executions that have occurred in a pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codepipeline_listActionExecutionsCmd).Standalone()

	codepipeline_listActionExecutionsCmd.Flags().String("filter", "", "Input information used to filter action execution history.")
	codepipeline_listActionExecutionsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
	codepipeline_listActionExecutionsCmd.Flags().String("next-token", "", "The token that was returned from the previous `ListActionExecutions` call, which can be used to return the next set of action executions in the list.")
	codepipeline_listActionExecutionsCmd.Flags().String("pipeline-name", "", "The name of the pipeline for which you want to list action execution history.")
	codepipeline_listActionExecutionsCmd.MarkFlagRequired("pipeline-name")
	codepipelineCmd.AddCommand(codepipeline_listActionExecutionsCmd)
}
