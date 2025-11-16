package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transcribe_deleteLanguageModelCmd = &cobra.Command{
	Use:   "delete-language-model",
	Short: "Deletes a custom language model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transcribe_deleteLanguageModelCmd).Standalone()

	transcribe_deleteLanguageModelCmd.Flags().String("model-name", "", "The name of the custom language model you want to delete.")
	transcribe_deleteLanguageModelCmd.MarkFlagRequired("model-name")
	transcribeCmd.AddCommand(transcribe_deleteLanguageModelCmd)
}
