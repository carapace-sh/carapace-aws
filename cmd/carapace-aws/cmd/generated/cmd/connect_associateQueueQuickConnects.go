package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_associateQueueQuickConnectsCmd = &cobra.Command{
	Use:   "associate-queue-quick-connects",
	Short: "Associates a set of quick connects with a queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_associateQueueQuickConnectsCmd).Standalone()

	connect_associateQueueQuickConnectsCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_associateQueueQuickConnectsCmd.Flags().String("queue-id", "", "The identifier for the queue.")
	connect_associateQueueQuickConnectsCmd.Flags().String("quick-connect-ids", "", "The quick connects to associate with this queue.")
	connect_associateQueueQuickConnectsCmd.MarkFlagRequired("instance-id")
	connect_associateQueueQuickConnectsCmd.MarkFlagRequired("queue-id")
	connect_associateQueueQuickConnectsCmd.MarkFlagRequired("quick-connect-ids")
	connectCmd.AddCommand(connect_associateQueueQuickConnectsCmd)
}
