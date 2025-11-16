package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_updateSessionCmd = &cobra.Command{
	Use:   "update-session",
	Short: "Updates a session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_updateSessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_updateSessionCmd).Standalone()

		deadline_updateSessionCmd.Flags().String("client-token", "", "The unique token which the server uses to recognize retries of the same request.")
		deadline_updateSessionCmd.Flags().String("farm-id", "", "The farm ID to update in the session.")
		deadline_updateSessionCmd.Flags().String("job-id", "", "The job ID to update in the session.")
		deadline_updateSessionCmd.Flags().String("queue-id", "", "The queue ID to update in the session.")
		deadline_updateSessionCmd.Flags().String("session-id", "", "The session ID to update.")
		deadline_updateSessionCmd.Flags().String("target-lifecycle-status", "", "The life cycle status to update in the session.")
		deadline_updateSessionCmd.MarkFlagRequired("farm-id")
		deadline_updateSessionCmd.MarkFlagRequired("job-id")
		deadline_updateSessionCmd.MarkFlagRequired("queue-id")
		deadline_updateSessionCmd.MarkFlagRequired("session-id")
		deadline_updateSessionCmd.MarkFlagRequired("target-lifecycle-status")
	})
	deadlineCmd.AddCommand(deadline_updateSessionCmd)
}
