package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_getStorageProfileForQueueCmd = &cobra.Command{
	Use:   "get-storage-profile-for-queue",
	Short: "Gets a storage profile for a queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_getStorageProfileForQueueCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_getStorageProfileForQueueCmd).Standalone()

		deadline_getStorageProfileForQueueCmd.Flags().String("farm-id", "", "The farm ID for the queue in storage profile.")
		deadline_getStorageProfileForQueueCmd.Flags().String("queue-id", "", "The queue ID the queue in the storage profile.")
		deadline_getStorageProfileForQueueCmd.Flags().String("storage-profile-id", "", "The storage profile ID for the storage profile in the queue.")
		deadline_getStorageProfileForQueueCmd.MarkFlagRequired("farm-id")
		deadline_getStorageProfileForQueueCmd.MarkFlagRequired("queue-id")
		deadline_getStorageProfileForQueueCmd.MarkFlagRequired("storage-profile-id")
	})
	deadlineCmd.AddCommand(deadline_getStorageProfileForQueueCmd)
}
