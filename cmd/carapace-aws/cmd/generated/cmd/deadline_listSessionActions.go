package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_listSessionActionsCmd = &cobra.Command{
	Use:   "list-session-actions",
	Short: "Lists session actions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_listSessionActionsCmd).Standalone()

	deadline_listSessionActionsCmd.Flags().String("farm-id", "", "The farm ID for the session actions list.")
	deadline_listSessionActionsCmd.Flags().String("job-id", "", "The job ID for the session actions list.")
	deadline_listSessionActionsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	deadline_listSessionActionsCmd.Flags().String("next-token", "", "The token for the next set of results, or `null` to start from the beginning.")
	deadline_listSessionActionsCmd.Flags().String("queue-id", "", "The queue ID for the session actions list.")
	deadline_listSessionActionsCmd.Flags().String("session-id", "", "The session ID to include on the sessions action list.")
	deadline_listSessionActionsCmd.Flags().String("task-id", "", "The task ID for the session actions list.")
	deadline_listSessionActionsCmd.MarkFlagRequired("farm-id")
	deadline_listSessionActionsCmd.MarkFlagRequired("job-id")
	deadline_listSessionActionsCmd.MarkFlagRequired("queue-id")
	deadlineCmd.AddCommand(deadline_listSessionActionsCmd)
}
