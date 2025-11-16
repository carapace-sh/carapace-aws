package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmIncidents_listTimelineEventsCmd = &cobra.Command{
	Use:   "list-timeline-events",
	Short: "Lists timeline events for the specified incident record.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmIncidents_listTimelineEventsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmIncidents_listTimelineEventsCmd).Standalone()

		ssmIncidents_listTimelineEventsCmd.Flags().String("filters", "", "Filters the timeline events based on the provided conditional values.")
		ssmIncidents_listTimelineEventsCmd.Flags().String("incident-record-arn", "", "The Amazon Resource Name (ARN) of the incident that includes the timeline event.")
		ssmIncidents_listTimelineEventsCmd.Flags().String("max-results", "", "The maximum number of results per page.")
		ssmIncidents_listTimelineEventsCmd.Flags().String("next-token", "", "The pagination token for the next set of items to return.")
		ssmIncidents_listTimelineEventsCmd.Flags().String("sort-by", "", "Sort timeline events by the specified key value pair.")
		ssmIncidents_listTimelineEventsCmd.Flags().String("sort-order", "", "Sorts the order of timeline events by the value specified in the `sortBy` field.")
		ssmIncidents_listTimelineEventsCmd.MarkFlagRequired("incident-record-arn")
	})
	ssmIncidentsCmd.AddCommand(ssmIncidents_listTimelineEventsCmd)
}
