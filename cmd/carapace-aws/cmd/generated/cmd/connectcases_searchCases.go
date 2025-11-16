package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcases_searchCasesCmd = &cobra.Command{
	Use:   "search-cases",
	Short: "Searches for cases within their associated Cases domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcases_searchCasesCmd).Standalone()

	connectcases_searchCasesCmd.Flags().String("domain-id", "", "The unique identifier of the Cases domain.")
	connectcases_searchCasesCmd.Flags().String("fields", "", "The list of field identifiers to be returned as part of the response.")
	connectcases_searchCasesCmd.Flags().String("filter", "", "A list of filter objects.")
	connectcases_searchCasesCmd.Flags().String("max-results", "", "The maximum number of cases to return.")
	connectcases_searchCasesCmd.Flags().String("next-token", "", "The token for the next set of results.")
	connectcases_searchCasesCmd.Flags().String("search-term", "", "A word or phrase used to perform a quick search.")
	connectcases_searchCasesCmd.Flags().String("sorts", "", "A list of sorts where each sort specifies a field and their sort order to be applied to the results.")
	connectcases_searchCasesCmd.MarkFlagRequired("domain-id")
	connectcasesCmd.AddCommand(connectcases_searchCasesCmd)
}
