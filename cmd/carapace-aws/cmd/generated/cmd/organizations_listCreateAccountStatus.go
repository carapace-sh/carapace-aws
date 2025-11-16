package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_listCreateAccountStatusCmd = &cobra.Command{
	Use:   "list-create-account-status",
	Short: "Lists the account creation requests that match the specified status that is currently being tracked for the organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_listCreateAccountStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(organizations_listCreateAccountStatusCmd).Standalone()

		organizations_listCreateAccountStatusCmd.Flags().String("max-results", "", "The total number of results that you want included on each page of the response.")
		organizations_listCreateAccountStatusCmd.Flags().String("next-token", "", "The parameter for receiving additional results if you receive a `NextToken` response in a previous request.")
		organizations_listCreateAccountStatusCmd.Flags().String("states", "", "A list of one or more states that you want included in the response.")
	})
	organizationsCmd.AddCommand(organizations_listCreateAccountStatusCmd)
}
