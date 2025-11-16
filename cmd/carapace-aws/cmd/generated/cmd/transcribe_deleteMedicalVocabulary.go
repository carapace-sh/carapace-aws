package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transcribe_deleteMedicalVocabularyCmd = &cobra.Command{
	Use:   "delete-medical-vocabulary",
	Short: "Deletes a custom medical vocabulary.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transcribe_deleteMedicalVocabularyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(transcribe_deleteMedicalVocabularyCmd).Standalone()

		transcribe_deleteMedicalVocabularyCmd.Flags().String("vocabulary-name", "", "The name of the custom medical vocabulary you want to delete.")
		transcribe_deleteMedicalVocabularyCmd.MarkFlagRequired("vocabulary-name")
	})
	transcribeCmd.AddCommand(transcribe_deleteMedicalVocabularyCmd)
}
