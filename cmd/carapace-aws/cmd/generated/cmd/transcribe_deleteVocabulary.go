package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transcribe_deleteVocabularyCmd = &cobra.Command{
	Use:   "delete-vocabulary",
	Short: "Deletes a custom vocabulary.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transcribe_deleteVocabularyCmd).Standalone()

	transcribe_deleteVocabularyCmd.Flags().String("vocabulary-name", "", "The name of the custom vocabulary you want to delete.")
	transcribe_deleteVocabularyCmd.MarkFlagRequired("vocabulary-name")
	transcribeCmd.AddCommand(transcribe_deleteVocabularyCmd)
}
