package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codestarConnections_listConnectionsCmd = &cobra.Command{
	Use:   "list-connections",
	Short: "Lists the connections associated with your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codestarConnections_listConnectionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codestarConnections_listConnectionsCmd).Standalone()

		codestarConnections_listConnectionsCmd.Flags().String("host-arn-filter", "", "Filters the list of connections to those associated with a specified host.")
		codestarConnections_listConnectionsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
		codestarConnections_listConnectionsCmd.Flags().String("next-token", "", "The token that was returned from the previous `ListConnections` call, which can be used to return the next set of connections in the list.")
		codestarConnections_listConnectionsCmd.Flags().String("provider-type-filter", "", "Filters the list of connections to those associated with a specified provider, such as Bitbucket.")
	})
	codestarConnectionsCmd.AddCommand(codestarConnections_listConnectionsCmd)
}
