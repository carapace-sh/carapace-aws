package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_batchUpdateCustomVocabularyItemCmd = &cobra.Command{
	Use:   "batch-update-custom-vocabulary-item",
	Short: "Update a batch of custom vocabulary items for a given bot locale's custom vocabulary.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_batchUpdateCustomVocabularyItemCmd).Standalone()

	lexv2Models_batchUpdateCustomVocabularyItemCmd.Flags().String("bot-id", "", "The identifier of the bot associated with this custom vocabulary")
	lexv2Models_batchUpdateCustomVocabularyItemCmd.Flags().String("bot-version", "", "The identifier of the version of the bot associated with this custom vocabulary.")
	lexv2Models_batchUpdateCustomVocabularyItemCmd.Flags().String("custom-vocabulary-item-list", "", "A list of custom vocabulary items with updated fields.")
	lexv2Models_batchUpdateCustomVocabularyItemCmd.Flags().String("locale-id", "", "The identifier of the language and locale where this custom vocabulary is used.")
	lexv2Models_batchUpdateCustomVocabularyItemCmd.MarkFlagRequired("bot-id")
	lexv2Models_batchUpdateCustomVocabularyItemCmd.MarkFlagRequired("bot-version")
	lexv2Models_batchUpdateCustomVocabularyItemCmd.MarkFlagRequired("custom-vocabulary-item-list")
	lexv2Models_batchUpdateCustomVocabularyItemCmd.MarkFlagRequired("locale-id")
	lexv2ModelsCmd.AddCommand(lexv2Models_batchUpdateCustomVocabularyItemCmd)
}
