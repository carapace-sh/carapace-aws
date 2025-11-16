package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_listEndpointsCmd = &cobra.Command{
	Use:   "list-endpoints",
	Short: "List the global endpoints associated with this account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_listEndpointsCmd).Standalone()

	events_listEndpointsCmd.Flags().String("home-region", "", "The primary Region of the endpoints associated with this account.")
	events_listEndpointsCmd.Flags().String("max-results", "", "The maximum number of results returned by the call.")
	events_listEndpointsCmd.Flags().String("name-prefix", "", "A value that will return a subset of the endpoints associated with this account.")
	events_listEndpointsCmd.Flags().String("next-token", "", "The token returned by a previous call, which you can use to retrieve the next set of results.")
	eventsCmd.AddCommand(events_listEndpointsCmd)
}
