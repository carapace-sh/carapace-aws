package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_listAccountsInAccountPoolCmd = &cobra.Command{
	Use:   "list-accounts-in-account-pool",
	Short: "Lists the accounts in the specified account pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_listAccountsInAccountPoolCmd).Standalone()

	datazone_listAccountsInAccountPoolCmd.Flags().String("domain-identifier", "", "The ID of the domain in which the accounts in the specified account pool are to be listed.")
	datazone_listAccountsInAccountPoolCmd.Flags().String("identifier", "", "The ID of the account pool whose accounts are to be listed.")
	datazone_listAccountsInAccountPoolCmd.Flags().String("max-results", "", "The maximum number of accounts to return in a single call to ListAccountsInAccountPool.")
	datazone_listAccountsInAccountPoolCmd.Flags().String("next-token", "", "When the number of accounts is greater than the default value for the MaxResults parameter, or if you explicitly specify a value for MaxResults that is less than the number of accounts, the response includes a pagination token named NextToken.")
	datazone_listAccountsInAccountPoolCmd.MarkFlagRequired("domain-identifier")
	datazone_listAccountsInAccountPoolCmd.MarkFlagRequired("identifier")
	datazoneCmd.AddCommand(datazone_listAccountsInAccountPoolCmd)
}
