package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_listTasksCmd = &cobra.Command{
	Use:   "list-tasks",
	Short: "Lists tasks for a job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_listTasksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_listTasksCmd).Standalone()

		deadline_listTasksCmd.Flags().String("farm-id", "", "The farm ID connected to the tasks.")
		deadline_listTasksCmd.Flags().String("job-id", "", "The job ID for the tasks.")
		deadline_listTasksCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		deadline_listTasksCmd.Flags().String("next-token", "", "The token for the next set of results, or `null` to start from the beginning.")
		deadline_listTasksCmd.Flags().String("queue-id", "", "The queue ID connected to the tasks.")
		deadline_listTasksCmd.Flags().String("step-id", "", "The step ID for the tasks.")
		deadline_listTasksCmd.MarkFlagRequired("farm-id")
		deadline_listTasksCmd.MarkFlagRequired("job-id")
		deadline_listTasksCmd.MarkFlagRequired("queue-id")
		deadline_listTasksCmd.MarkFlagRequired("step-id")
	})
	deadlineCmd.AddCommand(deadline_listTasksCmd)
}
