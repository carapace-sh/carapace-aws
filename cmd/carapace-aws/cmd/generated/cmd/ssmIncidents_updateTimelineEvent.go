package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmIncidents_updateTimelineEventCmd = &cobra.Command{
	Use:   "update-timeline-event",
	Short: "Updates a timeline event.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmIncidents_updateTimelineEventCmd).Standalone()

	ssmIncidents_updateTimelineEventCmd.Flags().String("client-token", "", "A token that ensures that a client calls the operation only once with the specified details.")
	ssmIncidents_updateTimelineEventCmd.Flags().String("event-data", "", "A short description of the event.")
	ssmIncidents_updateTimelineEventCmd.Flags().String("event-id", "", "The ID of the event to update.")
	ssmIncidents_updateTimelineEventCmd.Flags().String("event-references", "", "Updates all existing references in a `TimelineEvent`.")
	ssmIncidents_updateTimelineEventCmd.Flags().String("event-time", "", "The timestamp for when the event occurred.")
	ssmIncidents_updateTimelineEventCmd.Flags().String("event-type", "", "The type of event.")
	ssmIncidents_updateTimelineEventCmd.Flags().String("incident-record-arn", "", "The Amazon Resource Name (ARN) of the incident that includes the timeline event.")
	ssmIncidents_updateTimelineEventCmd.MarkFlagRequired("event-id")
	ssmIncidents_updateTimelineEventCmd.MarkFlagRequired("incident-record-arn")
	ssmIncidentsCmd.AddCommand(ssmIncidents_updateTimelineEventCmd)
}
