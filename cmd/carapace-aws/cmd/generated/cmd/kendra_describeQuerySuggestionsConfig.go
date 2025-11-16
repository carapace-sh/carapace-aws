package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_describeQuerySuggestionsConfigCmd = &cobra.Command{
	Use:   "describe-query-suggestions-config",
	Short: "Gets information on the settings of query suggestions for an index.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_describeQuerySuggestionsConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kendra_describeQuerySuggestionsConfigCmd).Standalone()

		kendra_describeQuerySuggestionsConfigCmd.Flags().String("index-id", "", "The identifier of the index with query suggestions that you want to get information on.")
		kendra_describeQuerySuggestionsConfigCmd.MarkFlagRequired("index-id")
	})
	kendraCmd.AddCommand(kendra_describeQuerySuggestionsConfigCmd)
}
