package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_batchDeleteCustomVocabularyItemCmd = &cobra.Command{
	Use:   "batch-delete-custom-vocabulary-item",
	Short: "Delete a batch of custom vocabulary items for a given bot locale's custom vocabulary.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_batchDeleteCustomVocabularyItemCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_batchDeleteCustomVocabularyItemCmd).Standalone()

		lexv2Models_batchDeleteCustomVocabularyItemCmd.Flags().String("bot-id", "", "The identifier of the bot associated with this custom vocabulary.")
		lexv2Models_batchDeleteCustomVocabularyItemCmd.Flags().String("bot-version", "", "The identifier of the version of the bot associated with this custom vocabulary.")
		lexv2Models_batchDeleteCustomVocabularyItemCmd.Flags().String("custom-vocabulary-item-list", "", "A list of custom vocabulary items requested to be deleted.")
		lexv2Models_batchDeleteCustomVocabularyItemCmd.Flags().String("locale-id", "", "The identifier of the language and locale where this custom vocabulary is used.")
		lexv2Models_batchDeleteCustomVocabularyItemCmd.MarkFlagRequired("bot-id")
		lexv2Models_batchDeleteCustomVocabularyItemCmd.MarkFlagRequired("bot-version")
		lexv2Models_batchDeleteCustomVocabularyItemCmd.MarkFlagRequired("custom-vocabulary-item-list")
		lexv2Models_batchDeleteCustomVocabularyItemCmd.MarkFlagRequired("locale-id")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_batchDeleteCustomVocabularyItemCmd)
}
