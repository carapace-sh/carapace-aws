package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_associateMemberToQueueCmd = &cobra.Command{
	Use:   "associate-member-to-queue",
	Short: "Assigns a queue membership level to a member",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_associateMemberToQueueCmd).Standalone()

	deadline_associateMemberToQueueCmd.Flags().String("farm-id", "", "The farm ID of the queue to associate with the member.")
	deadline_associateMemberToQueueCmd.Flags().String("identity-store-id", "", "The member's identity store ID to associate with the queue.")
	deadline_associateMemberToQueueCmd.Flags().String("membership-level", "", "The principal's membership level for the associated queue.")
	deadline_associateMemberToQueueCmd.Flags().String("principal-id", "", "The member's principal ID to associate with the queue.")
	deadline_associateMemberToQueueCmd.Flags().String("principal-type", "", "The member's principal type to associate with the queue.")
	deadline_associateMemberToQueueCmd.Flags().String("queue-id", "", "The ID of the queue to associate to the member.")
	deadline_associateMemberToQueueCmd.MarkFlagRequired("farm-id")
	deadline_associateMemberToQueueCmd.MarkFlagRequired("identity-store-id")
	deadline_associateMemberToQueueCmd.MarkFlagRequired("membership-level")
	deadline_associateMemberToQueueCmd.MarkFlagRequired("principal-id")
	deadline_associateMemberToQueueCmd.MarkFlagRequired("principal-type")
	deadline_associateMemberToQueueCmd.MarkFlagRequired("queue-id")
	deadlineCmd.AddCommand(deadline_associateMemberToQueueCmd)
}
