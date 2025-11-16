package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeconnections_listConnectionsCmd = &cobra.Command{
	Use:   "list-connections",
	Short: "Lists the connections associated with your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeconnections_listConnectionsCmd).Standalone()

	codeconnections_listConnectionsCmd.Flags().String("host-arn-filter", "", "Filters the list of connections to those associated with a specified host.")
	codeconnections_listConnectionsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
	codeconnections_listConnectionsCmd.Flags().String("next-token", "", "The token that was returned from the previous `ListConnections` call, which can be used to return the next set of connections in the list.")
	codeconnections_listConnectionsCmd.Flags().String("provider-type-filter", "", "Filters the list of connections to those associated with a specified provider, such as Bitbucket.")
	codeconnectionsCmd.AddCommand(codeconnections_listConnectionsCmd)
}
