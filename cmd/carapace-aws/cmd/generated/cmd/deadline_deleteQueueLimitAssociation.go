package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_deleteQueueLimitAssociationCmd = &cobra.Command{
	Use:   "delete-queue-limit-association",
	Short: "Removes the association between a queue and a limit.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_deleteQueueLimitAssociationCmd).Standalone()

	deadline_deleteQueueLimitAssociationCmd.Flags().String("farm-id", "", "The unique identifier of the farm that contains the queue and limit to disassociate.")
	deadline_deleteQueueLimitAssociationCmd.Flags().String("limit-id", "", "The unique identifier of the limit to disassociate.")
	deadline_deleteQueueLimitAssociationCmd.Flags().String("queue-id", "", "The unique identifier of the queue to disassociate.")
	deadline_deleteQueueLimitAssociationCmd.MarkFlagRequired("farm-id")
	deadline_deleteQueueLimitAssociationCmd.MarkFlagRequired("limit-id")
	deadline_deleteQueueLimitAssociationCmd.MarkFlagRequired("queue-id")
	deadlineCmd.AddCommand(deadline_deleteQueueLimitAssociationCmd)
}
