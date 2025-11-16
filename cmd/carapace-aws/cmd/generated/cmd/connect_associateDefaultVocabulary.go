package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_associateDefaultVocabularyCmd = &cobra.Command{
	Use:   "associate-default-vocabulary",
	Short: "Associates an existing vocabulary as the default.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_associateDefaultVocabularyCmd).Standalone()

	connect_associateDefaultVocabularyCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_associateDefaultVocabularyCmd.Flags().String("language-code", "", "The language code of the vocabulary entries.")
	connect_associateDefaultVocabularyCmd.Flags().String("vocabulary-id", "", "The identifier of the custom vocabulary.")
	connect_associateDefaultVocabularyCmd.MarkFlagRequired("instance-id")
	connect_associateDefaultVocabularyCmd.MarkFlagRequired("language-code")
	connectCmd.AddCommand(connect_associateDefaultVocabularyCmd)
}
