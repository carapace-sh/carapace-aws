package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_listDelegatedAdminAccountsCmd = &cobra.Command{
	Use:   "list-delegated-admin-accounts",
	Short: "Lists information about the Amazon Inspector delegated administrator of your organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_listDelegatedAdminAccountsCmd).Standalone()

	inspector2_listDelegatedAdminAccountsCmd.Flags().String("max-results", "", "The maximum number of results the response can return.")
	inspector2_listDelegatedAdminAccountsCmd.Flags().String("next-token", "", "A token to use for paginating results that are returned in the response.")
	inspector2Cmd.AddCommand(inspector2_listDelegatedAdminAccountsCmd)
}
