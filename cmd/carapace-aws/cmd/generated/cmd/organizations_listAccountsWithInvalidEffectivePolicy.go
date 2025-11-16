package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_listAccountsWithInvalidEffectivePolicyCmd = &cobra.Command{
	Use:   "list-accounts-with-invalid-effective-policy",
	Short: "Lists all the accounts in an organization that have invalid effective policies.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_listAccountsWithInvalidEffectivePolicyCmd).Standalone()

	organizations_listAccountsWithInvalidEffectivePolicyCmd.Flags().String("max-results", "", "The total number of results that you want included on each page of the response.")
	organizations_listAccountsWithInvalidEffectivePolicyCmd.Flags().String("next-token", "", "The parameter for receiving additional results if you receive a `NextToken` response in a previous request.")
	organizations_listAccountsWithInvalidEffectivePolicyCmd.Flags().String("policy-type", "", "The type of policy that you want information about.")
	organizations_listAccountsWithInvalidEffectivePolicyCmd.MarkFlagRequired("policy-type")
	organizationsCmd.AddCommand(organizations_listAccountsWithInvalidEffectivePolicyCmd)
}
