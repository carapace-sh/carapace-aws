package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_updateQueueCmd = &cobra.Command{
	Use:   "update-queue",
	Short: "Updates a queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_updateQueueCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_updateQueueCmd).Standalone()

		deadline_updateQueueCmd.Flags().String("allowed-storage-profile-ids-to-add", "", "The storage profile IDs to add.")
		deadline_updateQueueCmd.Flags().String("allowed-storage-profile-ids-to-remove", "", "The storage profile ID to remove.")
		deadline_updateQueueCmd.Flags().String("client-token", "", "The idempotency token to update in the queue.")
		deadline_updateQueueCmd.Flags().String("default-budget-action", "", "The default action to take for a queue update if a budget isn't configured.")
		deadline_updateQueueCmd.Flags().String("description", "", "The description of the queue to update.")
		deadline_updateQueueCmd.Flags().String("display-name", "", "The display name of the queue to update.")
		deadline_updateQueueCmd.Flags().String("farm-id", "", "The farm ID to update in the queue.")
		deadline_updateQueueCmd.Flags().String("job-attachment-settings", "", "The job attachment settings to update for the queue.")
		deadline_updateQueueCmd.Flags().String("job-run-as-user", "", "Update the jobs in the queue to run as a specified POSIX user.")
		deadline_updateQueueCmd.Flags().String("queue-id", "", "The queue ID to update.")
		deadline_updateQueueCmd.Flags().String("required-file-system-location-names-to-add", "", "The required file system location names to add to the queue.")
		deadline_updateQueueCmd.Flags().String("required-file-system-location-names-to-remove", "", "The required file system location names to remove from the queue.")
		deadline_updateQueueCmd.Flags().String("role-arn", "", "The IAM role ARN that's used to run jobs from this queue.")
		deadline_updateQueueCmd.MarkFlagRequired("farm-id")
		deadline_updateQueueCmd.MarkFlagRequired("queue-id")
	})
	deadlineCmd.AddCommand(deadline_updateQueueCmd)
}
