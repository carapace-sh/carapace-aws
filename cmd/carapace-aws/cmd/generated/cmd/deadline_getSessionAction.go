package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_getSessionActionCmd = &cobra.Command{
	Use:   "get-session-action",
	Short: "Gets a session action for the job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_getSessionActionCmd).Standalone()

	deadline_getSessionActionCmd.Flags().String("farm-id", "", "The farm ID for the session action.")
	deadline_getSessionActionCmd.Flags().String("job-id", "", "The job ID for the session.")
	deadline_getSessionActionCmd.Flags().String("queue-id", "", "The queue ID for the session action.")
	deadline_getSessionActionCmd.Flags().String("session-action-id", "", "The session action ID for the session.")
	deadline_getSessionActionCmd.MarkFlagRequired("farm-id")
	deadline_getSessionActionCmd.MarkFlagRequired("job-id")
	deadline_getSessionActionCmd.MarkFlagRequired("queue-id")
	deadline_getSessionActionCmd.MarkFlagRequired("session-action-id")
	deadlineCmd.AddCommand(deadline_getSessionActionCmd)
}
