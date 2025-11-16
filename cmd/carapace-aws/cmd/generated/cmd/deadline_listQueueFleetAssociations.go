package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_listQueueFleetAssociationsCmd = &cobra.Command{
	Use:   "list-queue-fleet-associations",
	Short: "Lists queue-fleet associations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_listQueueFleetAssociationsCmd).Standalone()

	deadline_listQueueFleetAssociationsCmd.Flags().String("farm-id", "", "The farm ID for the queue-fleet association list.")
	deadline_listQueueFleetAssociationsCmd.Flags().String("fleet-id", "", "The fleet ID for the queue-fleet association list.")
	deadline_listQueueFleetAssociationsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	deadline_listQueueFleetAssociationsCmd.Flags().String("next-token", "", "The token for the next set of results, or `null` to start from the beginning.")
	deadline_listQueueFleetAssociationsCmd.Flags().String("queue-id", "", "The queue ID for the queue-fleet association list.")
	deadline_listQueueFleetAssociationsCmd.MarkFlagRequired("farm-id")
	deadlineCmd.AddCommand(deadline_listQueueFleetAssociationsCmd)
}
