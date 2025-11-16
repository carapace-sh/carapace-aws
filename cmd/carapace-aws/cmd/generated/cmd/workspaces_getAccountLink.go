package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_getAccountLinkCmd = &cobra.Command{
	Use:   "get-account-link",
	Short: "Retrieves account link information.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_getAccountLinkCmd).Standalone()

	workspaces_getAccountLinkCmd.Flags().String("link-id", "", "The identifier of the account to link.")
	workspaces_getAccountLinkCmd.Flags().String("linked-account-id", "", "The identifier of the account link")
	workspacesCmd.AddCommand(workspaces_getAccountLinkCmd)
}
