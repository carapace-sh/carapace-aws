package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_batchUpdateAutomatedDiscoveryAccountsCmd = &cobra.Command{
	Use:   "batch-update-automated-discovery-accounts",
	Short: "Changes the status of automated sensitive data discovery for one or more accounts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_batchUpdateAutomatedDiscoveryAccountsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(macie2_batchUpdateAutomatedDiscoveryAccountsCmd).Standalone()

		macie2_batchUpdateAutomatedDiscoveryAccountsCmd.Flags().String("accounts", "", "An array of objects, one for each account to change the status of automated sensitive data discovery for.")
	})
	macie2Cmd.AddCommand(macie2_batchUpdateAutomatedDiscoveryAccountsCmd)
}
