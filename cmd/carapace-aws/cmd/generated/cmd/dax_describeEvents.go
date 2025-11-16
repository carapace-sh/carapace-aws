package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dax_describeEventsCmd = &cobra.Command{
	Use:   "describe-events",
	Short: "Returns events related to DAX clusters and parameter groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dax_describeEventsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dax_describeEventsCmd).Standalone()

		dax_describeEventsCmd.Flags().String("duration", "", "The number of minutes' worth of events to retrieve.")
		dax_describeEventsCmd.Flags().String("end-time", "", "The end of the time interval for which to retrieve events, specified in ISO 8601 format.")
		dax_describeEventsCmd.Flags().String("max-results", "", "The maximum number of results to include in the response.")
		dax_describeEventsCmd.Flags().String("next-token", "", "An optional token returned from a prior request.")
		dax_describeEventsCmd.Flags().String("source-name", "", "The identifier of the event source for which events will be returned.")
		dax_describeEventsCmd.Flags().String("source-type", "", "The event source to retrieve events for.")
		dax_describeEventsCmd.Flags().String("start-time", "", "The beginning of the time interval to retrieve events for, specified in ISO 8601 format.")
	})
	daxCmd.AddCommand(dax_describeEventsCmd)
}
