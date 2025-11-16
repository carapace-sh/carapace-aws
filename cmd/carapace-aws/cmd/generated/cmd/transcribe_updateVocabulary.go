package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transcribe_updateVocabularyCmd = &cobra.Command{
	Use:   "update-vocabulary",
	Short: "Updates an existing custom vocabulary with new values.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transcribe_updateVocabularyCmd).Standalone()

	transcribe_updateVocabularyCmd.Flags().String("data-access-role-arn", "", "The Amazon Resource Name (ARN) of an IAM role that has permissions to access the Amazon S3 bucket that contains your input files (in this case, your custom vocabulary).")
	transcribe_updateVocabularyCmd.Flags().String("language-code", "", "The language code that represents the language of the entries in the custom vocabulary you want to update.")
	transcribe_updateVocabularyCmd.Flags().String("phrases", "", "Use this parameter if you want to update your custom vocabulary by including all desired terms, as comma-separated values, within your request.")
	transcribe_updateVocabularyCmd.Flags().String("vocabulary-file-uri", "", "The Amazon S3 location of the text file that contains your custom vocabulary.")
	transcribe_updateVocabularyCmd.Flags().String("vocabulary-name", "", "The name of the custom vocabulary you want to update.")
	transcribe_updateVocabularyCmd.MarkFlagRequired("language-code")
	transcribe_updateVocabularyCmd.MarkFlagRequired("vocabulary-name")
	transcribeCmd.AddCommand(transcribe_updateVocabularyCmd)
}
