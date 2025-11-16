package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_listCrossAccountResourceAccountsCmd = &cobra.Command{
	Use:   "list-cross-account-resource-accounts",
	Short: "List the accounts that have cross-account resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_listCrossAccountResourceAccountsCmd).Standalone()

	globalacceleratorCmd.AddCommand(globalaccelerator_listCrossAccountResourceAccountsCmd)
}
