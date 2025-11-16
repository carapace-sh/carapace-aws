package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_assumeQueueRoleForReadCmd = &cobra.Command{
	Use:   "assume-queue-role-for-read",
	Short: "Gets Amazon Web Services credentials from the queue role.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_assumeQueueRoleForReadCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_assumeQueueRoleForReadCmd).Standalone()

		deadline_assumeQueueRoleForReadCmd.Flags().String("farm-id", "", "The farm ID of the farm containing the queue.")
		deadline_assumeQueueRoleForReadCmd.Flags().String("queue-id", "", "The queue ID.")
		deadline_assumeQueueRoleForReadCmd.MarkFlagRequired("farm-id")
		deadline_assumeQueueRoleForReadCmd.MarkFlagRequired("queue-id")
	})
	deadlineCmd.AddCommand(deadline_assumeQueueRoleForReadCmd)
}
