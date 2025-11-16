package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmIncidents_deleteTimelineEventCmd = &cobra.Command{
	Use:   "delete-timeline-event",
	Short: "Deletes a timeline event from an incident.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmIncidents_deleteTimelineEventCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmIncidents_deleteTimelineEventCmd).Standalone()

		ssmIncidents_deleteTimelineEventCmd.Flags().String("event-id", "", "The ID of the event to update.")
		ssmIncidents_deleteTimelineEventCmd.Flags().String("incident-record-arn", "", "The Amazon Resource Name (ARN) of the incident that includes the timeline event.")
		ssmIncidents_deleteTimelineEventCmd.MarkFlagRequired("event-id")
		ssmIncidents_deleteTimelineEventCmd.MarkFlagRequired("incident-record-arn")
	})
	ssmIncidentsCmd.AddCommand(ssmIncidents_deleteTimelineEventCmd)
}
