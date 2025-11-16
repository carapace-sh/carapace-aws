package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codepipeline_listDeployActionExecutionTargetsCmd = &cobra.Command{
	Use:   "list-deploy-action-execution-targets",
	Short: "Lists the targets for the deploy action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codepipeline_listDeployActionExecutionTargetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codepipeline_listDeployActionExecutionTargetsCmd).Standalone()

		codepipeline_listDeployActionExecutionTargetsCmd.Flags().String("action-execution-id", "", "The execution ID for the deploy action.")
		codepipeline_listDeployActionExecutionTargetsCmd.Flags().String("filters", "", "Filters the targets for a specified deploy action.")
		codepipeline_listDeployActionExecutionTargetsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
		codepipeline_listDeployActionExecutionTargetsCmd.Flags().String("next-token", "", "An identifier that was returned from the previous list action types call, which can be used to return the next set of action types in the list.")
		codepipeline_listDeployActionExecutionTargetsCmd.Flags().String("pipeline-name", "", "The name of the pipeline with the deploy action.")
		codepipeline_listDeployActionExecutionTargetsCmd.MarkFlagRequired("action-execution-id")
	})
	codepipelineCmd.AddCommand(codepipeline_listDeployActionExecutionTargetsCmd)
}
