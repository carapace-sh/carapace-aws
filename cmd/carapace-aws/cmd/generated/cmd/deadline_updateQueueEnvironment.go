package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_updateQueueEnvironmentCmd = &cobra.Command{
	Use:   "update-queue-environment",
	Short: "Updates the queue environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_updateQueueEnvironmentCmd).Standalone()

	deadline_updateQueueEnvironmentCmd.Flags().String("client-token", "", "The unique token which the server uses to recognize retries of the same request.")
	deadline_updateQueueEnvironmentCmd.Flags().String("farm-id", "", "The farm ID of the queue environment to update.")
	deadline_updateQueueEnvironmentCmd.Flags().String("priority", "", "The priority to update.")
	deadline_updateQueueEnvironmentCmd.Flags().String("queue-environment-id", "", "The queue environment ID to update.")
	deadline_updateQueueEnvironmentCmd.Flags().String("queue-id", "", "The queue ID of the queue environment to update.")
	deadline_updateQueueEnvironmentCmd.Flags().String("template", "", "The template to update.")
	deadline_updateQueueEnvironmentCmd.Flags().String("template-type", "", "The template type to update.")
	deadline_updateQueueEnvironmentCmd.MarkFlagRequired("farm-id")
	deadline_updateQueueEnvironmentCmd.MarkFlagRequired("queue-environment-id")
	deadline_updateQueueEnvironmentCmd.MarkFlagRequired("queue-id")
	deadlineCmd.AddCommand(deadline_updateQueueEnvironmentCmd)
}
