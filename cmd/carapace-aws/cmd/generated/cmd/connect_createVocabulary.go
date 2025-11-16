package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_createVocabularyCmd = &cobra.Command{
	Use:   "create-vocabulary",
	Short: "Creates a custom vocabulary associated with your Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_createVocabularyCmd).Standalone()

	connect_createVocabularyCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	connect_createVocabularyCmd.Flags().String("content", "", "The content of the custom vocabulary in plain-text format with a table of values.")
	connect_createVocabularyCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_createVocabularyCmd.Flags().String("language-code", "", "The language code of the vocabulary entries.")
	connect_createVocabularyCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
	connect_createVocabularyCmd.Flags().String("vocabulary-name", "", "A unique name of the custom vocabulary.")
	connect_createVocabularyCmd.MarkFlagRequired("content")
	connect_createVocabularyCmd.MarkFlagRequired("instance-id")
	connect_createVocabularyCmd.MarkFlagRequired("language-code")
	connect_createVocabularyCmd.MarkFlagRequired("vocabulary-name")
	connectCmd.AddCommand(connect_createVocabularyCmd)
}
