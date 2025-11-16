package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmIncidents_getTimelineEventCmd = &cobra.Command{
	Use:   "get-timeline-event",
	Short: "Retrieves a timeline event based on its ID and incident record.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmIncidents_getTimelineEventCmd).Standalone()

	ssmIncidents_getTimelineEventCmd.Flags().String("event-id", "", "The ID of the event.")
	ssmIncidents_getTimelineEventCmd.Flags().String("incident-record-arn", "", "The Amazon Resource Name (ARN) of the incident that includes the timeline event.")
	ssmIncidents_getTimelineEventCmd.MarkFlagRequired("event-id")
	ssmIncidents_getTimelineEventCmd.MarkFlagRequired("incident-record-arn")
	ssmIncidentsCmd.AddCommand(ssmIncidents_getTimelineEventCmd)
}
