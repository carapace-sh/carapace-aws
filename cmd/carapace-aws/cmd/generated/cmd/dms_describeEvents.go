package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_describeEventsCmd = &cobra.Command{
	Use:   "describe-events",
	Short: "Lists events for a given source identifier and source type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_describeEventsCmd).Standalone()

	dms_describeEventsCmd.Flags().String("duration", "", "The duration of the events to be listed.")
	dms_describeEventsCmd.Flags().String("end-time", "", "The end time for the events to be listed.")
	dms_describeEventsCmd.Flags().String("event-categories", "", "A list of event categories for the source type that you've chosen.")
	dms_describeEventsCmd.Flags().String("filters", "", "Filters applied to events.")
	dms_describeEventsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous request.")
	dms_describeEventsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	dms_describeEventsCmd.Flags().String("source-identifier", "", "The identifier of an event source.")
	dms_describeEventsCmd.Flags().String("source-type", "", "The type of DMS resource that generates events.")
	dms_describeEventsCmd.Flags().String("start-time", "", "The start time for the events to be listed.")
	dmsCmd.AddCommand(dms_describeEventsCmd)
}
