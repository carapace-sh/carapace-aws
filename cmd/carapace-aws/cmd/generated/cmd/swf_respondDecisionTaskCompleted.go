package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var swf_respondDecisionTaskCompletedCmd = &cobra.Command{
	Use:   "respond-decision-task-completed",
	Short: "Used by deciders to tell the service that the [DecisionTask]() identified by the `taskToken` has successfully completed.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swf_respondDecisionTaskCompletedCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(swf_respondDecisionTaskCompletedCmd).Standalone()

		swf_respondDecisionTaskCompletedCmd.Flags().String("decisions", "", "The list of decisions (possibly empty) made by the decider while processing this decision task.")
		swf_respondDecisionTaskCompletedCmd.Flags().String("execution-context", "", "User defined context to add to workflow execution.")
		swf_respondDecisionTaskCompletedCmd.Flags().String("task-list", "", "The task list to use for the future decision tasks of this workflow execution.")
		swf_respondDecisionTaskCompletedCmd.Flags().String("task-list-schedule-to-start-timeout", "", "Specifies a timeout (in seconds) for the task list override.")
		swf_respondDecisionTaskCompletedCmd.Flags().String("task-token", "", "The `taskToken` from the [DecisionTask]().")
		swf_respondDecisionTaskCompletedCmd.MarkFlagRequired("task-token")
	})
	swfCmd.AddCommand(swf_respondDecisionTaskCompletedCmd)
}
