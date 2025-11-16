package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transcribe_updateVocabularyFilterCmd = &cobra.Command{
	Use:   "update-vocabulary-filter",
	Short: "Updates an existing custom vocabulary filter with a new list of words.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transcribe_updateVocabularyFilterCmd).Standalone()

	transcribe_updateVocabularyFilterCmd.Flags().String("data-access-role-arn", "", "The Amazon Resource Name (ARN) of an IAM role that has permissions to access the Amazon S3 bucket that contains your input files (in this case, your custom vocabulary filter).")
	transcribe_updateVocabularyFilterCmd.Flags().String("vocabulary-filter-file-uri", "", "The Amazon S3 location of the text file that contains your custom vocabulary filter terms.")
	transcribe_updateVocabularyFilterCmd.Flags().String("vocabulary-filter-name", "", "The name of the custom vocabulary filter you want to update.")
	transcribe_updateVocabularyFilterCmd.Flags().String("words", "", "Use this parameter if you want to update your custom vocabulary filter by including all desired terms, as comma-separated values, within your request.")
	transcribe_updateVocabularyFilterCmd.MarkFlagRequired("vocabulary-filter-name")
	transcribeCmd.AddCommand(transcribe_updateVocabularyFilterCmd)
}
