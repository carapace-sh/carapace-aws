package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_disassociateMemberFromQueueCmd = &cobra.Command{
	Use:   "disassociate-member-from-queue",
	Short: "Disassociates a member from a queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_disassociateMemberFromQueueCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_disassociateMemberFromQueueCmd).Standalone()

		deadline_disassociateMemberFromQueueCmd.Flags().String("farm-id", "", "The farm ID for the queue to disassociate from a member.")
		deadline_disassociateMemberFromQueueCmd.Flags().String("principal-id", "", "A member's principal ID to disassociate from a queue.")
		deadline_disassociateMemberFromQueueCmd.Flags().String("queue-id", "", "The queue ID of the queue in which you're disassociating from a member.")
		deadline_disassociateMemberFromQueueCmd.MarkFlagRequired("farm-id")
		deadline_disassociateMemberFromQueueCmd.MarkFlagRequired("principal-id")
		deadline_disassociateMemberFromQueueCmd.MarkFlagRequired("queue-id")
	})
	deadlineCmd.AddCommand(deadline_disassociateMemberFromQueueCmd)
}
