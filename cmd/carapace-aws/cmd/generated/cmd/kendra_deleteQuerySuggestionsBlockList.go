package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_deleteQuerySuggestionsBlockListCmd = &cobra.Command{
	Use:   "delete-query-suggestions-block-list",
	Short: "Deletes a block list used for query suggestions for an index.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_deleteQuerySuggestionsBlockListCmd).Standalone()

	kendra_deleteQuerySuggestionsBlockListCmd.Flags().String("id", "", "The identifier of the block list you want to delete.")
	kendra_deleteQuerySuggestionsBlockListCmd.Flags().String("index-id", "", "The identifier of the index for the block list.")
	kendra_deleteQuerySuggestionsBlockListCmd.MarkFlagRequired("id")
	kendra_deleteQuerySuggestionsBlockListCmd.MarkFlagRequired("index-id")
	kendraCmd.AddCommand(kendra_deleteQuerySuggestionsBlockListCmd)
}
