package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transcribe_getMedicalVocabularyCmd = &cobra.Command{
	Use:   "get-medical-vocabulary",
	Short: "Provides information about the specified custom medical vocabulary.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transcribe_getMedicalVocabularyCmd).Standalone()

	transcribe_getMedicalVocabularyCmd.Flags().String("vocabulary-name", "", "The name of the custom medical vocabulary you want information about.")
	transcribe_getMedicalVocabularyCmd.MarkFlagRequired("vocabulary-name")
	transcribeCmd.AddCommand(transcribe_getMedicalVocabularyCmd)
}
