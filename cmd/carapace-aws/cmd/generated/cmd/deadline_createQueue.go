package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_createQueueCmd = &cobra.Command{
	Use:   "create-queue",
	Short: "Creates a queue to coordinate the order in which jobs run on a farm.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_createQueueCmd).Standalone()

	deadline_createQueueCmd.Flags().String("allowed-storage-profile-ids", "", "The storage profile IDs to include in the queue.")
	deadline_createQueueCmd.Flags().String("client-token", "", "The unique token which the server uses to recognize retries of the same request.")
	deadline_createQueueCmd.Flags().String("default-budget-action", "", "The default action to take on a queue if a budget isn't configured.")
	deadline_createQueueCmd.Flags().String("description", "", "The description of the queue.")
	deadline_createQueueCmd.Flags().String("display-name", "", "The display name of the queue.")
	deadline_createQueueCmd.Flags().String("farm-id", "", "The farm ID of the farm to connect to the queue.")
	deadline_createQueueCmd.Flags().String("job-attachment-settings", "", "The job attachment settings for the queue.")
	deadline_createQueueCmd.Flags().String("job-run-as-user", "", "The jobs in the queue run as the specified POSIX user.")
	deadline_createQueueCmd.Flags().String("required-file-system-location-names", "", "The file system location name to include in the queue.")
	deadline_createQueueCmd.Flags().String("role-arn", "", "The IAM role ARN that workers will use while running jobs for this queue.")
	deadline_createQueueCmd.Flags().String("tags", "", "Each tag consists of a tag key and a tag value.")
	deadline_createQueueCmd.MarkFlagRequired("display-name")
	deadline_createQueueCmd.MarkFlagRequired("farm-id")
	deadlineCmd.AddCommand(deadline_createQueueCmd)
}
