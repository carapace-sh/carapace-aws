package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_updateQueueLimitAssociationCmd = &cobra.Command{
	Use:   "update-queue-limit-association",
	Short: "Updates the status of the queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_updateQueueLimitAssociationCmd).Standalone()

	deadline_updateQueueLimitAssociationCmd.Flags().String("farm-id", "", "The unique identifier of the farm that contains the associated queues and limits.")
	deadline_updateQueueLimitAssociationCmd.Flags().String("limit-id", "", "The unique identifier of the limit associated to the queue.")
	deadline_updateQueueLimitAssociationCmd.Flags().String("queue-id", "", "The unique identifier of the queue associated to the limit.")
	deadline_updateQueueLimitAssociationCmd.Flags().String("status", "", "Sets the status of the limit.")
	deadline_updateQueueLimitAssociationCmd.MarkFlagRequired("farm-id")
	deadline_updateQueueLimitAssociationCmd.MarkFlagRequired("limit-id")
	deadline_updateQueueLimitAssociationCmd.MarkFlagRequired("queue-id")
	deadline_updateQueueLimitAssociationCmd.MarkFlagRequired("status")
	deadlineCmd.AddCommand(deadline_updateQueueLimitAssociationCmd)
}
