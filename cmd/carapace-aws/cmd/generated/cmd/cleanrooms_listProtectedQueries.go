package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_listProtectedQueriesCmd = &cobra.Command{
	Use:   "list-protected-queries",
	Short: "Lists protected queries, sorted by the most recent query.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_listProtectedQueriesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanrooms_listProtectedQueriesCmd).Standalone()

		cleanrooms_listProtectedQueriesCmd.Flags().String("max-results", "", "The maximum number of results that are returned for an API request call.")
		cleanrooms_listProtectedQueriesCmd.Flags().String("membership-identifier", "", "The identifier for the membership in the collaboration.")
		cleanrooms_listProtectedQueriesCmd.Flags().String("next-token", "", "The pagination token that's used to fetch the next set of results.")
		cleanrooms_listProtectedQueriesCmd.Flags().String("status", "", "A filter on the status of the protected query.")
		cleanrooms_listProtectedQueriesCmd.MarkFlagRequired("membership-identifier")
	})
	cleanroomsCmd.AddCommand(cleanrooms_listProtectedQueriesCmd)
}
