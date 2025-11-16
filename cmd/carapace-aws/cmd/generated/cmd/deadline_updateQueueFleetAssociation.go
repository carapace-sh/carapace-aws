package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_updateQueueFleetAssociationCmd = &cobra.Command{
	Use:   "update-queue-fleet-association",
	Short: "Updates a queue-fleet association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_updateQueueFleetAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_updateQueueFleetAssociationCmd).Standalone()

		deadline_updateQueueFleetAssociationCmd.Flags().String("farm-id", "", "The farm ID to update.")
		deadline_updateQueueFleetAssociationCmd.Flags().String("fleet-id", "", "The fleet ID to update.")
		deadline_updateQueueFleetAssociationCmd.Flags().String("queue-id", "", "The queue ID to update.")
		deadline_updateQueueFleetAssociationCmd.Flags().String("status", "", "The status to update.")
		deadline_updateQueueFleetAssociationCmd.MarkFlagRequired("farm-id")
		deadline_updateQueueFleetAssociationCmd.MarkFlagRequired("fleet-id")
		deadline_updateQueueFleetAssociationCmd.MarkFlagRequired("queue-id")
		deadline_updateQueueFleetAssociationCmd.MarkFlagRequired("status")
	})
	deadlineCmd.AddCommand(deadline_updateQueueFleetAssociationCmd)
}
