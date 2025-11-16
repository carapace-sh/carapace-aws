package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_deleteQueueCmd = &cobra.Command{
	Use:   "delete-queue",
	Short: "Deletes a queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_deleteQueueCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_deleteQueueCmd).Standalone()

		connect_deleteQueueCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_deleteQueueCmd.Flags().String("queue-id", "", "The identifier for the queue.")
		connect_deleteQueueCmd.MarkFlagRequired("instance-id")
		connect_deleteQueueCmd.MarkFlagRequired("queue-id")
	})
	connectCmd.AddCommand(connect_deleteQueueCmd)
}
