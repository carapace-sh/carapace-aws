package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_updateQuerySuggestionsConfigCmd = &cobra.Command{
	Use:   "update-query-suggestions-config",
	Short: "Updates the settings of query suggestions for an index.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_updateQuerySuggestionsConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kendra_updateQuerySuggestionsConfigCmd).Standalone()

		kendra_updateQuerySuggestionsConfigCmd.Flags().String("attribute-suggestions-config", "", "Configuration information for the document fields/attributes that you want to base query suggestions on.")
		kendra_updateQuerySuggestionsConfigCmd.Flags().String("include-queries-without-user-information", "", "`TRUE` to include queries without user information (i.e. all queries, irrespective of the user), otherwise `FALSE` to only include queries with user information.")
		kendra_updateQuerySuggestionsConfigCmd.Flags().String("index-id", "", "The identifier of the index with query suggestions you want to update.")
		kendra_updateQuerySuggestionsConfigCmd.Flags().String("minimum-number-of-querying-users", "", "The minimum number of unique users who must search a query in order for the query to be eligible to suggest to your users.")
		kendra_updateQuerySuggestionsConfigCmd.Flags().String("minimum-query-count", "", "The the minimum number of times a query must be searched in order to be eligible to suggest to your users.")
		kendra_updateQuerySuggestionsConfigCmd.Flags().String("mode", "", "Set the mode to `ENABLED` or `LEARN_ONLY`.")
		kendra_updateQuerySuggestionsConfigCmd.Flags().String("query-log-look-back-window-in-days", "", "How recent your queries are in your query log time window.")
		kendra_updateQuerySuggestionsConfigCmd.MarkFlagRequired("index-id")
	})
	kendraCmd.AddCommand(kendra_updateQuerySuggestionsConfigCmd)
}
