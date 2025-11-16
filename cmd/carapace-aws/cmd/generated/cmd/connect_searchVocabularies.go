package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_searchVocabulariesCmd = &cobra.Command{
	Use:   "search-vocabularies",
	Short: "Searches for vocabularies within a specific Amazon Connect instance using `State`, `NameStartsWith`, and `LanguageCode`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_searchVocabulariesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_searchVocabulariesCmd).Standalone()

		connect_searchVocabulariesCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_searchVocabulariesCmd.Flags().String("language-code", "", "The language code of the vocabulary entries.")
		connect_searchVocabulariesCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		connect_searchVocabulariesCmd.Flags().String("name-starts-with", "", "The starting pattern of the name of the vocabulary.")
		connect_searchVocabulariesCmd.Flags().String("next-token", "", "The token for the next set of results.")
		connect_searchVocabulariesCmd.Flags().String("state", "", "The current state of the custom vocabulary.")
		connect_searchVocabulariesCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_searchVocabulariesCmd)
}
