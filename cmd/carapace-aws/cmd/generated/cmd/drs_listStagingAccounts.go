package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_listStagingAccountsCmd = &cobra.Command{
	Use:   "list-staging-accounts",
	Short: "Returns an array of staging accounts for existing extended source servers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_listStagingAccountsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(drs_listStagingAccountsCmd).Standalone()

		drs_listStagingAccountsCmd.Flags().String("max-results", "", "The maximum number of staging Accounts to retrieve.")
		drs_listStagingAccountsCmd.Flags().String("next-token", "", "The token of the next staging Account to retrieve.")
	})
	drsCmd.AddCommand(drs_listStagingAccountsCmd)
}
