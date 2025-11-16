package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_assumeQueueRoleForUserCmd = &cobra.Command{
	Use:   "assume-queue-role-for-user",
	Short: "Allows a user to assume a role for a queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_assumeQueueRoleForUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_assumeQueueRoleForUserCmd).Standalone()

		deadline_assumeQueueRoleForUserCmd.Flags().String("farm-id", "", "The farm ID of the queue that the user assumes the role for.")
		deadline_assumeQueueRoleForUserCmd.Flags().String("queue-id", "", "The queue ID of the queue that the user assumes the role for.")
		deadline_assumeQueueRoleForUserCmd.MarkFlagRequired("farm-id")
		deadline_assumeQueueRoleForUserCmd.MarkFlagRequired("queue-id")
	})
	deadlineCmd.AddCommand(deadline_assumeQueueRoleForUserCmd)
}
