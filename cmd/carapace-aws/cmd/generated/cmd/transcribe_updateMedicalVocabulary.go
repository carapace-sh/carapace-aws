package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transcribe_updateMedicalVocabularyCmd = &cobra.Command{
	Use:   "update-medical-vocabulary",
	Short: "Updates an existing custom medical vocabulary with new values.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transcribe_updateMedicalVocabularyCmd).Standalone()

	transcribe_updateMedicalVocabularyCmd.Flags().String("language-code", "", "The language code that represents the language of the entries in the custom vocabulary you want to update.")
	transcribe_updateMedicalVocabularyCmd.Flags().String("vocabulary-file-uri", "", "The Amazon S3 location of the text file that contains your custom medical vocabulary.")
	transcribe_updateMedicalVocabularyCmd.Flags().String("vocabulary-name", "", "The name of the custom medical vocabulary you want to update.")
	transcribe_updateMedicalVocabularyCmd.MarkFlagRequired("language-code")
	transcribe_updateMedicalVocabularyCmd.MarkFlagRequired("vocabulary-file-uri")
	transcribe_updateMedicalVocabularyCmd.MarkFlagRequired("vocabulary-name")
	transcribeCmd.AddCommand(transcribe_updateMedicalVocabularyCmd)
}
