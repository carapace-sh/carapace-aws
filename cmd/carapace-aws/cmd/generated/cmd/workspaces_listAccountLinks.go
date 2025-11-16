package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_listAccountLinksCmd = &cobra.Command{
	Use:   "list-account-links",
	Short: "Lists all account links.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_listAccountLinksCmd).Standalone()

	workspaces_listAccountLinksCmd.Flags().String("link-status-filter", "", "Filters the account based on their link status.")
	workspaces_listAccountLinksCmd.Flags().String("max-results", "", "The maximum number of accounts to return.")
	workspaces_listAccountLinksCmd.Flags().String("next-token", "", "The token to use to retrieve the next page of results.")
	workspacesCmd.AddCommand(workspaces_listAccountLinksCmd)
}
