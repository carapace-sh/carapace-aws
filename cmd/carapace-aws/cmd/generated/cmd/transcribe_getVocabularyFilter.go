package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transcribe_getVocabularyFilterCmd = &cobra.Command{
	Use:   "get-vocabulary-filter",
	Short: "Provides information about the specified custom vocabulary filter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transcribe_getVocabularyFilterCmd).Standalone()

	transcribe_getVocabularyFilterCmd.Flags().String("vocabulary-filter-name", "", "The name of the custom vocabulary filter you want information about.")
	transcribe_getVocabularyFilterCmd.MarkFlagRequired("vocabulary-filter-name")
	transcribeCmd.AddCommand(transcribe_getVocabularyFilterCmd)
}
