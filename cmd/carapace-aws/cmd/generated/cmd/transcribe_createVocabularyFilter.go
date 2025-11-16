package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transcribe_createVocabularyFilterCmd = &cobra.Command{
	Use:   "create-vocabulary-filter",
	Short: "Creates a new custom vocabulary filter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transcribe_createVocabularyFilterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(transcribe_createVocabularyFilterCmd).Standalone()

		transcribe_createVocabularyFilterCmd.Flags().String("data-access-role-arn", "", "The Amazon Resource Name (ARN) of an IAM role that has permissions to access the Amazon S3 bucket that contains your input files (in this case, your custom vocabulary filter).")
		transcribe_createVocabularyFilterCmd.Flags().String("language-code", "", "The language code that represents the language of the entries in your vocabulary filter.")
		transcribe_createVocabularyFilterCmd.Flags().String("tags", "", "Adds one or more custom tags, each in the form of a key:value pair, to a new custom vocabulary filter at the time you create this new vocabulary filter.")
		transcribe_createVocabularyFilterCmd.Flags().String("vocabulary-filter-file-uri", "", "The Amazon S3 location of the text file that contains your custom vocabulary filter terms.")
		transcribe_createVocabularyFilterCmd.Flags().String("vocabulary-filter-name", "", "A unique name, chosen by you, for your new custom vocabulary filter.")
		transcribe_createVocabularyFilterCmd.Flags().String("words", "", "Use this parameter if you want to create your custom vocabulary filter by including all desired terms, as comma-separated values, within your request.")
		transcribe_createVocabularyFilterCmd.MarkFlagRequired("language-code")
		transcribe_createVocabularyFilterCmd.MarkFlagRequired("vocabulary-filter-name")
	})
	transcribeCmd.AddCommand(transcribe_createVocabularyFilterCmd)
}
