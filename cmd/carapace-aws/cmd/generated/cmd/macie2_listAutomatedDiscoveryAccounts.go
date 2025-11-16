package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_listAutomatedDiscoveryAccountsCmd = &cobra.Command{
	Use:   "list-automated-discovery-accounts",
	Short: "Retrieves the status of automated sensitive data discovery for one or more accounts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_listAutomatedDiscoveryAccountsCmd).Standalone()

	macie2_listAutomatedDiscoveryAccountsCmd.Flags().String("account-ids", "", "The Amazon Web Services account ID for each account, for as many as 50 accounts.")
	macie2_listAutomatedDiscoveryAccountsCmd.Flags().String("max-results", "", "The maximum number of items to include in each page of a paginated response.")
	macie2_listAutomatedDiscoveryAccountsCmd.Flags().String("next-token", "", "The nextToken string that specifies which page of results to return in a paginated response.")
	macie2Cmd.AddCommand(macie2_listAutomatedDiscoveryAccountsCmd)
}
