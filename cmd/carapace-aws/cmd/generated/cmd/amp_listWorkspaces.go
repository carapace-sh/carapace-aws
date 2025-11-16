package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amp_listWorkspacesCmd = &cobra.Command{
	Use:   "list-workspaces",
	Short: "Lists all of the Amazon Managed Service for Prometheus workspaces in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amp_listWorkspacesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amp_listWorkspacesCmd).Standalone()

		amp_listWorkspacesCmd.Flags().String("alias", "", "If this is included, it filters the results to only the workspaces with names that start with the value that you specify here.")
		amp_listWorkspacesCmd.Flags().String("max-results", "", "The maximum number of workspaces to return per request.")
		amp_listWorkspacesCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	})
	ampCmd.AddCommand(amp_listWorkspacesCmd)
}
