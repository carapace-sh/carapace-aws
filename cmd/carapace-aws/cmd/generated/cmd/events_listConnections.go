package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_listConnectionsCmd = &cobra.Command{
	Use:   "list-connections",
	Short: "Retrieves a list of connections from the account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_listConnectionsCmd).Standalone()

	events_listConnectionsCmd.Flags().String("connection-state", "", "The state of the connection.")
	events_listConnectionsCmd.Flags().String("limit", "", "The maximum number of connections to return.")
	events_listConnectionsCmd.Flags().String("name-prefix", "", "A name prefix to filter results returned.")
	events_listConnectionsCmd.Flags().String("next-token", "", "The token returned by a previous call, which you can use to retrieve the next set of results.")
	eventsCmd.AddCommand(events_listConnectionsCmd)
}
