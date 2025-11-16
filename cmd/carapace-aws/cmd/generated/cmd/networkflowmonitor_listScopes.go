package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkflowmonitor_listScopesCmd = &cobra.Command{
	Use:   "list-scopes",
	Short: "List all the scopes for an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkflowmonitor_listScopesCmd).Standalone()

	networkflowmonitor_listScopesCmd.Flags().String("max-results", "", "The number of query results that you want to return with this call.")
	networkflowmonitor_listScopesCmd.Flags().String("next-token", "", "The token for the next set of results.")
	networkflowmonitorCmd.AddCommand(networkflowmonitor_listScopesCmd)
}
