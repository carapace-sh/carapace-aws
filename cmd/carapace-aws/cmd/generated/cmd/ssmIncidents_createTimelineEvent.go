package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmIncidents_createTimelineEventCmd = &cobra.Command{
	Use:   "create-timeline-event",
	Short: "Creates a custom timeline event on the incident details page of an incident record.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmIncidents_createTimelineEventCmd).Standalone()

	ssmIncidents_createTimelineEventCmd.Flags().String("client-token", "", "A token that ensures that a client calls the action only once with the specified details.")
	ssmIncidents_createTimelineEventCmd.Flags().String("event-data", "", "A short description of the event.")
	ssmIncidents_createTimelineEventCmd.Flags().String("event-references", "", "Adds one or more references to the `TimelineEvent`.")
	ssmIncidents_createTimelineEventCmd.Flags().String("event-time", "", "The timestamp for when the event occurred.")
	ssmIncidents_createTimelineEventCmd.Flags().String("event-type", "", "The type of event.")
	ssmIncidents_createTimelineEventCmd.Flags().String("incident-record-arn", "", "The Amazon Resource Name (ARN) of the incident record that the action adds the incident to.")
	ssmIncidents_createTimelineEventCmd.MarkFlagRequired("event-data")
	ssmIncidents_createTimelineEventCmd.MarkFlagRequired("event-time")
	ssmIncidents_createTimelineEventCmd.MarkFlagRequired("event-type")
	ssmIncidents_createTimelineEventCmd.MarkFlagRequired("incident-record-arn")
	ssmIncidentsCmd.AddCommand(ssmIncidents_createTimelineEventCmd)
}
