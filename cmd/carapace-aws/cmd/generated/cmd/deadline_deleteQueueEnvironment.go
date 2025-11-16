package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_deleteQueueEnvironmentCmd = &cobra.Command{
	Use:   "delete-queue-environment",
	Short: "Deletes a queue environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_deleteQueueEnvironmentCmd).Standalone()

	deadline_deleteQueueEnvironmentCmd.Flags().String("farm-id", "", "The farm ID of the farm from which to remove the queue environment.")
	deadline_deleteQueueEnvironmentCmd.Flags().String("queue-environment-id", "", "The queue environment ID of the queue environment to delete.")
	deadline_deleteQueueEnvironmentCmd.Flags().String("queue-id", "", "The queue ID of the queue environment to delete.")
	deadline_deleteQueueEnvironmentCmd.MarkFlagRequired("farm-id")
	deadline_deleteQueueEnvironmentCmd.MarkFlagRequired("queue-environment-id")
	deadline_deleteQueueEnvironmentCmd.MarkFlagRequired("queue-id")
	deadlineCmd.AddCommand(deadline_deleteQueueEnvironmentCmd)
}
