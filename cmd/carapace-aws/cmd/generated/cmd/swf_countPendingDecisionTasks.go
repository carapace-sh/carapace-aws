package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var swf_countPendingDecisionTasksCmd = &cobra.Command{
	Use:   "count-pending-decision-tasks",
	Short: "Returns the estimated number of decision tasks in the specified task list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swf_countPendingDecisionTasksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(swf_countPendingDecisionTasksCmd).Standalone()

		swf_countPendingDecisionTasksCmd.Flags().String("domain", "", "The name of the domain that contains the task list.")
		swf_countPendingDecisionTasksCmd.Flags().String("task-list", "", "The name of the task list.")
		swf_countPendingDecisionTasksCmd.MarkFlagRequired("domain")
		swf_countPendingDecisionTasksCmd.MarkFlagRequired("task-list")
	})
	swfCmd.AddCommand(swf_countPendingDecisionTasksCmd)
}
