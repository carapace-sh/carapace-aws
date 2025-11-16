package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_getQueueFleetAssociationCmd = &cobra.Command{
	Use:   "get-queue-fleet-association",
	Short: "Gets a queue-fleet association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_getQueueFleetAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_getQueueFleetAssociationCmd).Standalone()

		deadline_getQueueFleetAssociationCmd.Flags().String("farm-id", "", "The farm ID of the farm that contains the queue-fleet association.")
		deadline_getQueueFleetAssociationCmd.Flags().String("fleet-id", "", "The fleet ID for the queue-fleet association.")
		deadline_getQueueFleetAssociationCmd.Flags().String("queue-id", "", "The queue ID for the queue-fleet association.")
		deadline_getQueueFleetAssociationCmd.MarkFlagRequired("farm-id")
		deadline_getQueueFleetAssociationCmd.MarkFlagRequired("fleet-id")
		deadline_getQueueFleetAssociationCmd.MarkFlagRequired("queue-id")
	})
	deadlineCmd.AddCommand(deadline_getQueueFleetAssociationCmd)
}
