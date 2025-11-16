package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_listQueueQuickConnectsCmd = &cobra.Command{
	Use:   "list-queue-quick-connects",
	Short: "Lists the quick connects associated with a queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_listQueueQuickConnectsCmd).Standalone()

	connect_listQueueQuickConnectsCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_listQueueQuickConnectsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	connect_listQueueQuickConnectsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	connect_listQueueQuickConnectsCmd.Flags().String("queue-id", "", "The identifier for the queue.")
	connect_listQueueQuickConnectsCmd.MarkFlagRequired("instance-id")
	connect_listQueueQuickConnectsCmd.MarkFlagRequired("queue-id")
	connectCmd.AddCommand(connect_listQueueQuickConnectsCmd)
}
