package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var swf_countPendingActivityTasksCmd = &cobra.Command{
	Use:   "count-pending-activity-tasks",
	Short: "Returns the estimated number of activity tasks in the specified task list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swf_countPendingActivityTasksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(swf_countPendingActivityTasksCmd).Standalone()

		swf_countPendingActivityTasksCmd.Flags().String("domain", "", "The name of the domain that contains the task list.")
		swf_countPendingActivityTasksCmd.Flags().String("task-list", "", "The name of the task list.")
		swf_countPendingActivityTasksCmd.MarkFlagRequired("domain")
		swf_countPendingActivityTasksCmd.MarkFlagRequired("task-list")
	})
	swfCmd.AddCommand(swf_countPendingActivityTasksCmd)
}
