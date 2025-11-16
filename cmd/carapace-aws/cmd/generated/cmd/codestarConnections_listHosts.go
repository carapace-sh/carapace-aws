package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codestarConnections_listHostsCmd = &cobra.Command{
	Use:   "list-hosts",
	Short: "Lists the hosts associated with your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codestarConnections_listHostsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codestarConnections_listHostsCmd).Standalone()

		codestarConnections_listHostsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
		codestarConnections_listHostsCmd.Flags().String("next-token", "", "The token that was returned from the previous `ListHosts` call, which can be used to return the next set of hosts in the list.")
	})
	codestarConnectionsCmd.AddCommand(codestarConnections_listHostsCmd)
}
