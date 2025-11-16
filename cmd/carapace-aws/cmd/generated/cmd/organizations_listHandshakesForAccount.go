package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_listHandshakesForAccountCmd = &cobra.Command{
	Use:   "list-handshakes-for-account",
	Short: "Lists the current handshakes that are associated with the account of the requesting user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_listHandshakesForAccountCmd).Standalone()

	organizations_listHandshakesForAccountCmd.Flags().String("filter", "", "Filters the handshakes that you want included in the response.")
	organizations_listHandshakesForAccountCmd.Flags().String("max-results", "", "The total number of results that you want included on each page of the response.")
	organizations_listHandshakesForAccountCmd.Flags().String("next-token", "", "The parameter for receiving additional results if you receive a `NextToken` response in a previous request.")
	organizationsCmd.AddCommand(organizations_listHandshakesForAccountCmd)
}
