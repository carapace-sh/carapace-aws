package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_describeQuerySuggestionsBlockListCmd = &cobra.Command{
	Use:   "describe-query-suggestions-block-list",
	Short: "Gets information about a block list used for query suggestions for an index.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_describeQuerySuggestionsBlockListCmd).Standalone()

	kendra_describeQuerySuggestionsBlockListCmd.Flags().String("id", "", "The identifier of the block list you want to get information on.")
	kendra_describeQuerySuggestionsBlockListCmd.Flags().String("index-id", "", "The identifier of the index for the block list.")
	kendra_describeQuerySuggestionsBlockListCmd.MarkFlagRequired("id")
	kendra_describeQuerySuggestionsBlockListCmd.MarkFlagRequired("index-id")
	kendraCmd.AddCommand(kendra_describeQuerySuggestionsBlockListCmd)
}
