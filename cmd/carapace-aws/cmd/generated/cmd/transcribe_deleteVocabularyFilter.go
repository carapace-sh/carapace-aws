package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transcribe_deleteVocabularyFilterCmd = &cobra.Command{
	Use:   "delete-vocabulary-filter",
	Short: "Deletes a custom vocabulary filter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transcribe_deleteVocabularyFilterCmd).Standalone()

	transcribe_deleteVocabularyFilterCmd.Flags().String("vocabulary-filter-name", "", "The name of the custom vocabulary filter you want to delete.")
	transcribe_deleteVocabularyFilterCmd.MarkFlagRequired("vocabulary-filter-name")
	transcribeCmd.AddCommand(transcribe_deleteVocabularyFilterCmd)
}
