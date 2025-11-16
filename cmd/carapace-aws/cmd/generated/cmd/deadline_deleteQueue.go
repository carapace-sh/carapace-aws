package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_deleteQueueCmd = &cobra.Command{
	Use:   "delete-queue",
	Short: "Deletes a queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_deleteQueueCmd).Standalone()

	deadline_deleteQueueCmd.Flags().String("farm-id", "", "The ID of the farm from which to remove the queue.")
	deadline_deleteQueueCmd.Flags().String("queue-id", "", "The queue ID of the queue to delete.")
	deadline_deleteQueueCmd.MarkFlagRequired("farm-id")
	deadline_deleteQueueCmd.MarkFlagRequired("queue-id")
	deadlineCmd.AddCommand(deadline_deleteQueueCmd)
}
