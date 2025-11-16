package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_getQueueCmd = &cobra.Command{
	Use:   "get-queue",
	Short: "Gets a queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_getQueueCmd).Standalone()

	deadline_getQueueCmd.Flags().String("farm-id", "", "The farm ID of the farm in the queue.")
	deadline_getQueueCmd.Flags().String("queue-id", "", "The queue ID for the queue to retrieve.")
	deadline_getQueueCmd.MarkFlagRequired("farm-id")
	deadline_getQueueCmd.MarkFlagRequired("queue-id")
	deadlineCmd.AddCommand(deadline_getQueueCmd)
}
