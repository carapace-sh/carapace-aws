package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_getQueueLimitAssociationCmd = &cobra.Command{
	Use:   "get-queue-limit-association",
	Short: "Gets information about a specific association between a queue and a limit.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_getQueueLimitAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_getQueueLimitAssociationCmd).Standalone()

		deadline_getQueueLimitAssociationCmd.Flags().String("farm-id", "", "The unique identifier of the farm that contains the associated queue and limit.")
		deadline_getQueueLimitAssociationCmd.Flags().String("limit-id", "", "The unique identifier of the limit associated with the queue.")
		deadline_getQueueLimitAssociationCmd.Flags().String("queue-id", "", "The unique identifier of the queue associated with the limit.")
		deadline_getQueueLimitAssociationCmd.MarkFlagRequired("farm-id")
		deadline_getQueueLimitAssociationCmd.MarkFlagRequired("limit-id")
		deadline_getQueueLimitAssociationCmd.MarkFlagRequired("queue-id")
	})
	deadlineCmd.AddCommand(deadline_getQueueLimitAssociationCmd)
}
