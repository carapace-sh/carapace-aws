package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_getQuerySuggestionsCmd = &cobra.Command{
	Use:   "get-query-suggestions",
	Short: "Fetches the queries that are suggested to your users.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_getQuerySuggestionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kendra_getQuerySuggestionsCmd).Standalone()

		kendra_getQuerySuggestionsCmd.Flags().String("attribute-suggestions-config", "", "Configuration information for the document fields/attributes that you want to base query suggestions on.")
		kendra_getQuerySuggestionsCmd.Flags().String("index-id", "", "The identifier of the index you want to get query suggestions from.")
		kendra_getQuerySuggestionsCmd.Flags().String("max-suggestions-count", "", "The maximum number of query suggestions you want to show to your users.")
		kendra_getQuerySuggestionsCmd.Flags().String("query-text", "", "The text of a user's query to generate query suggestions.")
		kendra_getQuerySuggestionsCmd.Flags().String("suggestion-types", "", "The suggestions type to base query suggestions on.")
		kendra_getQuerySuggestionsCmd.MarkFlagRequired("index-id")
		kendra_getQuerySuggestionsCmd.MarkFlagRequired("query-text")
	})
	kendraCmd.AddCommand(kendra_getQuerySuggestionsCmd)
}
