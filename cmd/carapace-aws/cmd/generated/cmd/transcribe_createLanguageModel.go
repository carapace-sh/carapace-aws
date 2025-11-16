package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transcribe_createLanguageModelCmd = &cobra.Command{
	Use:   "create-language-model",
	Short: "Creates a new custom language model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transcribe_createLanguageModelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(transcribe_createLanguageModelCmd).Standalone()

		transcribe_createLanguageModelCmd.Flags().String("base-model-name", "", "The Amazon Transcribe standard language model, or base model, used to create your custom language model.")
		transcribe_createLanguageModelCmd.Flags().String("input-data-config", "", "Contains the Amazon S3 location of the training data you want to use to create a new custom language model, and permissions to access this location.")
		transcribe_createLanguageModelCmd.Flags().String("language-code", "", "The language code that represents the language of your model.")
		transcribe_createLanguageModelCmd.Flags().String("model-name", "", "A unique name, chosen by you, for your custom language model.")
		transcribe_createLanguageModelCmd.Flags().String("tags", "", "Adds one or more custom tags, each in the form of a key:value pair, to a new custom language model at the time you create this new model.")
		transcribe_createLanguageModelCmd.MarkFlagRequired("base-model-name")
		transcribe_createLanguageModelCmd.MarkFlagRequired("input-data-config")
		transcribe_createLanguageModelCmd.MarkFlagRequired("language-code")
		transcribe_createLanguageModelCmd.MarkFlagRequired("model-name")
	})
	transcribeCmd.AddCommand(transcribe_createLanguageModelCmd)
}
