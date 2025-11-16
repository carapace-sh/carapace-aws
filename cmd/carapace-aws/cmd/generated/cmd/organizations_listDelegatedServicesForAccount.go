package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_listDelegatedServicesForAccountCmd = &cobra.Command{
	Use:   "list-delegated-services-for-account",
	Short: "List the Amazon Web Services services for which the specified account is a delegated administrator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_listDelegatedServicesForAccountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(organizations_listDelegatedServicesForAccountCmd).Standalone()

		organizations_listDelegatedServicesForAccountCmd.Flags().String("account-id", "", "The account ID number of a delegated administrator account in the organization.")
		organizations_listDelegatedServicesForAccountCmd.Flags().String("max-results", "", "The total number of results that you want included on each page of the response.")
		organizations_listDelegatedServicesForAccountCmd.Flags().String("next-token", "", "The parameter for receiving additional results if you receive a `NextToken` response in a previous request.")
		organizations_listDelegatedServicesForAccountCmd.MarkFlagRequired("account-id")
	})
	organizationsCmd.AddCommand(organizations_listDelegatedServicesForAccountCmd)
}
