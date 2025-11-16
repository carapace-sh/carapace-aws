package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codepipeline_listRuleExecutionsCmd = &cobra.Command{
	Use:   "list-rule-executions",
	Short: "Lists the rule executions that have occurred in a pipeline configured for conditions with rules.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codepipeline_listRuleExecutionsCmd).Standalone()

	codepipeline_listRuleExecutionsCmd.Flags().String("filter", "", "Input information used to filter rule execution history.")
	codepipeline_listRuleExecutionsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
	codepipeline_listRuleExecutionsCmd.Flags().String("next-token", "", "The token that was returned from the previous `ListRuleExecutions` call, which can be used to return the next set of rule executions in the list.")
	codepipeline_listRuleExecutionsCmd.Flags().String("pipeline-name", "", "The name of the pipeline for which you want to get execution summary information.")
	codepipeline_listRuleExecutionsCmd.MarkFlagRequired("pipeline-name")
	codepipelineCmd.AddCommand(codepipeline_listRuleExecutionsCmd)
}
