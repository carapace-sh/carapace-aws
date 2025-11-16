package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transcribe_describeLanguageModelCmd = &cobra.Command{
	Use:   "describe-language-model",
	Short: "Provides information about the specified custom language model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transcribe_describeLanguageModelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(transcribe_describeLanguageModelCmd).Standalone()

		transcribe_describeLanguageModelCmd.Flags().String("model-name", "", "The name of the custom language model you want information about.")
		transcribe_describeLanguageModelCmd.MarkFlagRequired("model-name")
	})
	transcribeCmd.AddCommand(transcribe_describeLanguageModelCmd)
}
