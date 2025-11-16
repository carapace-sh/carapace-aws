package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_clearQuerySuggestionsCmd = &cobra.Command{
	Use:   "clear-query-suggestions",
	Short: "Clears existing query suggestions from an index.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_clearQuerySuggestionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kendra_clearQuerySuggestionsCmd).Standalone()

		kendra_clearQuerySuggestionsCmd.Flags().String("index-id", "", "The identifier of the index you want to clear query suggestions from.")
		kendra_clearQuerySuggestionsCmd.MarkFlagRequired("index-id")
	})
	kendraCmd.AddCommand(kendra_clearQuerySuggestionsCmd)
}
