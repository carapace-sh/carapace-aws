package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_getSessionCmd = &cobra.Command{
	Use:   "get-session",
	Short: "Gets a session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_getSessionCmd).Standalone()

	deadline_getSessionCmd.Flags().String("farm-id", "", "The farm ID for the session.")
	deadline_getSessionCmd.Flags().String("job-id", "", "The job ID for the session.")
	deadline_getSessionCmd.Flags().String("queue-id", "", "The queue ID for the session.")
	deadline_getSessionCmd.Flags().String("session-id", "", "The session ID.")
	deadline_getSessionCmd.MarkFlagRequired("farm-id")
	deadline_getSessionCmd.MarkFlagRequired("job-id")
	deadline_getSessionCmd.MarkFlagRequired("queue-id")
	deadline_getSessionCmd.MarkFlagRequired("session-id")
	deadlineCmd.AddCommand(deadline_getSessionCmd)
}
