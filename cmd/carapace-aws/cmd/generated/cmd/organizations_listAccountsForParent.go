package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_listAccountsForParentCmd = &cobra.Command{
	Use:   "list-accounts-for-parent",
	Short: "Lists the accounts in an organization that are contained by the specified target root or organizational unit (OU).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_listAccountsForParentCmd).Standalone()

	organizations_listAccountsForParentCmd.Flags().String("max-results", "", "The total number of results that you want included on each page of the response.")
	organizations_listAccountsForParentCmd.Flags().String("next-token", "", "The parameter for receiving additional results if you receive a `NextToken` response in a previous request.")
	organizations_listAccountsForParentCmd.Flags().String("parent-id", "", "The unique identifier (ID) for the parent root or organization unit (OU) whose accounts you want to list.")
	organizations_listAccountsForParentCmd.MarkFlagRequired("parent-id")
	organizationsCmd.AddCommand(organizations_listAccountsForParentCmd)
}
