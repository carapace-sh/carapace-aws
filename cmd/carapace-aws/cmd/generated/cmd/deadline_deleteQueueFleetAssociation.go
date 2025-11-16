package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_deleteQueueFleetAssociationCmd = &cobra.Command{
	Use:   "delete-queue-fleet-association",
	Short: "Deletes a queue-fleet association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_deleteQueueFleetAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_deleteQueueFleetAssociationCmd).Standalone()

		deadline_deleteQueueFleetAssociationCmd.Flags().String("farm-id", "", "The farm ID of the farm that holds the queue-fleet association.")
		deadline_deleteQueueFleetAssociationCmd.Flags().String("fleet-id", "", "The fleet ID of the queue-fleet association.")
		deadline_deleteQueueFleetAssociationCmd.Flags().String("queue-id", "", "The queue ID of the queue-fleet association.")
		deadline_deleteQueueFleetAssociationCmd.MarkFlagRequired("farm-id")
		deadline_deleteQueueFleetAssociationCmd.MarkFlagRequired("fleet-id")
		deadline_deleteQueueFleetAssociationCmd.MarkFlagRequired("queue-id")
	})
	deadlineCmd.AddCommand(deadline_deleteQueueFleetAssociationCmd)
}
