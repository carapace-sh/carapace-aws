package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_listEventBusesCmd = &cobra.Command{
	Use:   "list-event-buses",
	Short: "Lists all the event buses in your account, including the default event bus, custom event buses, and partner event buses.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_listEventBusesCmd).Standalone()

	events_listEventBusesCmd.Flags().String("limit", "", "Specifying this limits the number of results returned by this operation.")
	events_listEventBusesCmd.Flags().String("name-prefix", "", "Specifying this limits the results to only those event buses with names that start with the specified prefix.")
	events_listEventBusesCmd.Flags().String("next-token", "", "The token returned by a previous call, which you can use to retrieve the next set of results.")
	eventsCmd.AddCommand(events_listEventBusesCmd)
}
