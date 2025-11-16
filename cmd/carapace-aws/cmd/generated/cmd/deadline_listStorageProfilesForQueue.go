package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_listStorageProfilesForQueueCmd = &cobra.Command{
	Use:   "list-storage-profiles-for-queue",
	Short: "Lists storage profiles for a queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_listStorageProfilesForQueueCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_listStorageProfilesForQueueCmd).Standalone()

		deadline_listStorageProfilesForQueueCmd.Flags().String("farm-id", "", "The farm ID of the queue's storage profile.")
		deadline_listStorageProfilesForQueueCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		deadline_listStorageProfilesForQueueCmd.Flags().String("next-token", "", "The token for the next set of results, or `null` to start from the beginning.")
		deadline_listStorageProfilesForQueueCmd.Flags().String("queue-id", "", "The queue ID for the storage profile.")
		deadline_listStorageProfilesForQueueCmd.MarkFlagRequired("farm-id")
		deadline_listStorageProfilesForQueueCmd.MarkFlagRequired("queue-id")
	})
	deadlineCmd.AddCommand(deadline_listStorageProfilesForQueueCmd)
}
