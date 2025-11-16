package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transcribe_createVocabularyCmd = &cobra.Command{
	Use:   "create-vocabulary",
	Short: "Creates a new custom vocabulary.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transcribe_createVocabularyCmd).Standalone()

	transcribe_createVocabularyCmd.Flags().String("data-access-role-arn", "", "The Amazon Resource Name (ARN) of an IAM role that has permissions to access the Amazon S3 bucket that contains your input files (in this case, your custom vocabulary).")
	transcribe_createVocabularyCmd.Flags().String("language-code", "", "The language code that represents the language of the entries in your custom vocabulary.")
	transcribe_createVocabularyCmd.Flags().String("phrases", "", "Use this parameter if you want to create your custom vocabulary by including all desired terms, as comma-separated values, within your request.")
	transcribe_createVocabularyCmd.Flags().String("tags", "", "Adds one or more custom tags, each in the form of a key:value pair, to a new custom vocabulary at the time you create this new custom vocabulary.")
	transcribe_createVocabularyCmd.Flags().String("vocabulary-file-uri", "", "The Amazon S3 location of the text file that contains your custom vocabulary.")
	transcribe_createVocabularyCmd.Flags().String("vocabulary-name", "", "A unique name, chosen by you, for your new custom vocabulary.")
	transcribe_createVocabularyCmd.MarkFlagRequired("language-code")
	transcribe_createVocabularyCmd.MarkFlagRequired("vocabulary-name")
	transcribeCmd.AddCommand(transcribe_createVocabularyCmd)
}
