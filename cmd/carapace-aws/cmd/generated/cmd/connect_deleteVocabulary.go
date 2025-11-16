package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_deleteVocabularyCmd = &cobra.Command{
	Use:   "delete-vocabulary",
	Short: "Deletes the vocabulary that has the given identifier.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_deleteVocabularyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_deleteVocabularyCmd).Standalone()

		connect_deleteVocabularyCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_deleteVocabularyCmd.Flags().String("vocabulary-id", "", "The identifier of the custom vocabulary.")
		connect_deleteVocabularyCmd.MarkFlagRequired("instance-id")
		connect_deleteVocabularyCmd.MarkFlagRequired("vocabulary-id")
	})
	connectCmd.AddCommand(connect_deleteVocabularyCmd)
}
