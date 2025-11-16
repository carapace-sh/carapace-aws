package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_listDelegatedAdministratorsCmd = &cobra.Command{
	Use:   "list-delegated-administrators",
	Short: "Lists the Amazon Web Services accounts that are designated as delegated administrators in this organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_listDelegatedAdministratorsCmd).Standalone()

	organizations_listDelegatedAdministratorsCmd.Flags().String("max-results", "", "The total number of results that you want included on each page of the response.")
	organizations_listDelegatedAdministratorsCmd.Flags().String("next-token", "", "The parameter for receiving additional results if you receive a `NextToken` response in a previous request.")
	organizations_listDelegatedAdministratorsCmd.Flags().String("service-principal", "", "Specifies a service principal name.")
	organizationsCmd.AddCommand(organizations_listDelegatedAdministratorsCmd)
}
