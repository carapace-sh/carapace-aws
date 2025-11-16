package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_listQuerySuggestionsBlockListsCmd = &cobra.Command{
	Use:   "list-query-suggestions-block-lists",
	Short: "Lists the block lists used for query suggestions for an index.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_listQuerySuggestionsBlockListsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kendra_listQuerySuggestionsBlockListsCmd).Standalone()

		kendra_listQuerySuggestionsBlockListsCmd.Flags().String("index-id", "", "The identifier of the index for a list of all block lists that exist for that index.")
		kendra_listQuerySuggestionsBlockListsCmd.Flags().String("max-results", "", "The maximum number of block lists to return.")
		kendra_listQuerySuggestionsBlockListsCmd.Flags().String("next-token", "", "If the previous response was incomplete (because there is more data to retrieve), Amazon Kendra returns a pagination token in the response.")
		kendra_listQuerySuggestionsBlockListsCmd.MarkFlagRequired("index-id")
	})
	kendraCmd.AddCommand(kendra_listQuerySuggestionsBlockListsCmd)
}
