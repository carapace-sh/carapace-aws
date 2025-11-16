package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_listEventSourcesCmd = &cobra.Command{
	Use:   "list-event-sources",
	Short: "You can use this to see all the partner event sources that have been shared with your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_listEventSourcesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(events_listEventSourcesCmd).Standalone()

		events_listEventSourcesCmd.Flags().String("limit", "", "Specifying this limits the number of results returned by this operation.")
		events_listEventSourcesCmd.Flags().String("name-prefix", "", "Specifying this limits the results to only those partner event sources with names that start with the specified prefix.")
		events_listEventSourcesCmd.Flags().String("next-token", "", "The token returned by a previous call, which you can use to retrieve the next set of results.")
	})
	eventsCmd.AddCommand(events_listEventSourcesCmd)
}
