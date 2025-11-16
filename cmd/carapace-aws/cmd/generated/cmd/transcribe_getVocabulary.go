package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transcribe_getVocabularyCmd = &cobra.Command{
	Use:   "get-vocabulary",
	Short: "Provides information about the specified custom vocabulary.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transcribe_getVocabularyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(transcribe_getVocabularyCmd).Standalone()

		transcribe_getVocabularyCmd.Flags().String("vocabulary-name", "", "The name of the custom vocabulary you want information about.")
		transcribe_getVocabularyCmd.MarkFlagRequired("vocabulary-name")
	})
	transcribeCmd.AddCommand(transcribe_getVocabularyCmd)
}
