package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_listAccountsCmd = &cobra.Command{
	Use:   "list-accounts",
	Short: "Lists all the accounts in the organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_listAccountsCmd).Standalone()

	organizations_listAccountsCmd.Flags().String("max-results", "", "The total number of results that you want included on each page of the response.")
	organizations_listAccountsCmd.Flags().String("next-token", "", "The parameter for receiving additional results if you receive a `NextToken` response in a previous request.")
	organizationsCmd.AddCommand(organizations_listAccountsCmd)
}
