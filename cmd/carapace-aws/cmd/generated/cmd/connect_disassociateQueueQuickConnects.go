package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_disassociateQueueQuickConnectsCmd = &cobra.Command{
	Use:   "disassociate-queue-quick-connects",
	Short: "Disassociates a set of quick connects from a queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_disassociateQueueQuickConnectsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_disassociateQueueQuickConnectsCmd).Standalone()

		connect_disassociateQueueQuickConnectsCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_disassociateQueueQuickConnectsCmd.Flags().String("queue-id", "", "The identifier for the queue.")
		connect_disassociateQueueQuickConnectsCmd.Flags().String("quick-connect-ids", "", "The quick connects to disassociate from the queue.")
		connect_disassociateQueueQuickConnectsCmd.MarkFlagRequired("instance-id")
		connect_disassociateQueueQuickConnectsCmd.MarkFlagRequired("queue-id")
		connect_disassociateQueueQuickConnectsCmd.MarkFlagRequired("quick-connect-ids")
	})
	connectCmd.AddCommand(connect_disassociateQueueQuickConnectsCmd)
}
