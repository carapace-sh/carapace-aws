package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transcribe_createMedicalVocabularyCmd = &cobra.Command{
	Use:   "create-medical-vocabulary",
	Short: "Creates a new custom medical vocabulary.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transcribe_createMedicalVocabularyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(transcribe_createMedicalVocabularyCmd).Standalone()

		transcribe_createMedicalVocabularyCmd.Flags().String("language-code", "", "The language code that represents the language of the entries in your custom vocabulary.")
		transcribe_createMedicalVocabularyCmd.Flags().String("tags", "", "Adds one or more custom tags, each in the form of a key:value pair, to a new custom medical vocabulary at the time you create this new custom vocabulary.")
		transcribe_createMedicalVocabularyCmd.Flags().String("vocabulary-file-uri", "", "The Amazon S3 location (URI) of the text file that contains your custom medical vocabulary.")
		transcribe_createMedicalVocabularyCmd.Flags().String("vocabulary-name", "", "A unique name, chosen by you, for your new custom medical vocabulary.")
		transcribe_createMedicalVocabularyCmd.MarkFlagRequired("language-code")
		transcribe_createMedicalVocabularyCmd.MarkFlagRequired("vocabulary-file-uri")
		transcribe_createMedicalVocabularyCmd.MarkFlagRequired("vocabulary-name")
	})
	transcribeCmd.AddCommand(transcribe_createMedicalVocabularyCmd)
}
