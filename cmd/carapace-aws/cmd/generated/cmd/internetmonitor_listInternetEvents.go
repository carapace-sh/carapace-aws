package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var internetmonitor_listInternetEventsCmd = &cobra.Command{
	Use:   "list-internet-events",
	Short: "Lists internet events that cause performance or availability issues for client locations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(internetmonitor_listInternetEventsCmd).Standalone()

	internetmonitor_listInternetEventsCmd.Flags().String("end-time", "", "The end time of the time window that you want to get a list of internet events for.")
	internetmonitor_listInternetEventsCmd.Flags().String("event-status", "", "The status of an internet event.")
	internetmonitor_listInternetEventsCmd.Flags().String("event-type", "", "The type of network impairment.")
	internetmonitor_listInternetEventsCmd.Flags().String("max-results", "", "The number of query results that you want to return with this call.")
	internetmonitor_listInternetEventsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	internetmonitor_listInternetEventsCmd.Flags().String("start-time", "", "The start time of the time window that you want to get a list of internet events for.")
	internetmonitorCmd.AddCommand(internetmonitor_listInternetEventsCmd)
}
