package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_getQueueEnvironmentCmd = &cobra.Command{
	Use:   "get-queue-environment",
	Short: "Gets a queue environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_getQueueEnvironmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_getQueueEnvironmentCmd).Standalone()

		deadline_getQueueEnvironmentCmd.Flags().String("farm-id", "", "The farm ID for the queue environment.")
		deadline_getQueueEnvironmentCmd.Flags().String("queue-environment-id", "", "The queue environment ID.")
		deadline_getQueueEnvironmentCmd.Flags().String("queue-id", "", "The queue ID for the queue environment.")
		deadline_getQueueEnvironmentCmd.MarkFlagRequired("farm-id")
		deadline_getQueueEnvironmentCmd.MarkFlagRequired("queue-environment-id")
		deadline_getQueueEnvironmentCmd.MarkFlagRequired("queue-id")
	})
	deadlineCmd.AddCommand(deadline_getQueueEnvironmentCmd)
}
