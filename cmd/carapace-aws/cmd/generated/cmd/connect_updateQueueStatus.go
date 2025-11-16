package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updateQueueStatusCmd = &cobra.Command{
	Use:   "update-queue-status",
	Short: "Updates the status of the queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updateQueueStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_updateQueueStatusCmd).Standalone()

		connect_updateQueueStatusCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_updateQueueStatusCmd.Flags().String("queue-id", "", "The identifier for the queue.")
		connect_updateQueueStatusCmd.Flags().String("status", "", "The status of the queue.")
		connect_updateQueueStatusCmd.MarkFlagRequired("instance-id")
		connect_updateQueueStatusCmd.MarkFlagRequired("queue-id")
		connect_updateQueueStatusCmd.MarkFlagRequired("status")
	})
	connectCmd.AddCommand(connect_updateQueueStatusCmd)
}
