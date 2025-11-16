package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeconnections_listHostsCmd = &cobra.Command{
	Use:   "list-hosts",
	Short: "Lists the hosts associated with your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeconnections_listHostsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeconnections_listHostsCmd).Standalone()

		codeconnections_listHostsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
		codeconnections_listHostsCmd.Flags().String("next-token", "", "The token that was returned from the previous `ListHosts` call, which can be used to return the next set of hosts in the list.")
	})
	codeconnectionsCmd.AddCommand(codeconnections_listHostsCmd)
}
