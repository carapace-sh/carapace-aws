package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_listManagedAccountsCmd = &cobra.Command{
	Use:   "list-managed-accounts",
	Short: "List Managed Accounts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_listManagedAccountsCmd).Standalone()

	mgn_listManagedAccountsCmd.Flags().String("max-results", "", "List managed accounts request max results.")
	mgn_listManagedAccountsCmd.Flags().String("next-token", "", "List managed accounts request next token.")
	mgnCmd.AddCommand(mgn_listManagedAccountsCmd)
}
