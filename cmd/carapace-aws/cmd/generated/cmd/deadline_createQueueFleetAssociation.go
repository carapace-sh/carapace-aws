package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_createQueueFleetAssociationCmd = &cobra.Command{
	Use:   "create-queue-fleet-association",
	Short: "Creates an association between a queue and a fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_createQueueFleetAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_createQueueFleetAssociationCmd).Standalone()

		deadline_createQueueFleetAssociationCmd.Flags().String("farm-id", "", "The ID of the farm that the queue and fleet belong to.")
		deadline_createQueueFleetAssociationCmd.Flags().String("fleet-id", "", "The fleet ID.")
		deadline_createQueueFleetAssociationCmd.Flags().String("queue-id", "", "The queue ID.")
		deadline_createQueueFleetAssociationCmd.MarkFlagRequired("farm-id")
		deadline_createQueueFleetAssociationCmd.MarkFlagRequired("fleet-id")
		deadline_createQueueFleetAssociationCmd.MarkFlagRequired("queue-id")
	})
	deadlineCmd.AddCommand(deadline_createQueueFleetAssociationCmd)
}
