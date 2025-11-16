package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transcribe_listVocabularyFiltersCmd = &cobra.Command{
	Use:   "list-vocabulary-filters",
	Short: "Provides a list of custom vocabulary filters that match the specified criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transcribe_listVocabularyFiltersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(transcribe_listVocabularyFiltersCmd).Standalone()

		transcribe_listVocabularyFiltersCmd.Flags().String("max-results", "", "The maximum number of custom vocabulary filters to return in each page of results.")
		transcribe_listVocabularyFiltersCmd.Flags().String("name-contains", "", "Returns only the custom vocabulary filters that contain the specified string.")
		transcribe_listVocabularyFiltersCmd.Flags().String("next-token", "", "If your `ListVocabularyFilters` request returns more results than can be displayed, `NextToken` is displayed in the response with an associated string.")
	})
	transcribeCmd.AddCommand(transcribe_listVocabularyFiltersCmd)
}
