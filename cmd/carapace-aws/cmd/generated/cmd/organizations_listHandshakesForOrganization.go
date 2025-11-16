package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_listHandshakesForOrganizationCmd = &cobra.Command{
	Use:   "list-handshakes-for-organization",
	Short: "Lists the handshakes that are associated with the organization that the requesting user is part of.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_listHandshakesForOrganizationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(organizations_listHandshakesForOrganizationCmd).Standalone()

		organizations_listHandshakesForOrganizationCmd.Flags().String("filter", "", "A filter of the handshakes that you want included in the response.")
		organizations_listHandshakesForOrganizationCmd.Flags().String("max-results", "", "The total number of results that you want included on each page of the response.")
		organizations_listHandshakesForOrganizationCmd.Flags().String("next-token", "", "The parameter for receiving additional results if you receive a `NextToken` response in a previous request.")
	})
	organizationsCmd.AddCommand(organizations_listHandshakesForOrganizationCmd)
}
