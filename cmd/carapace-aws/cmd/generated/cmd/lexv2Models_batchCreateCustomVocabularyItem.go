package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_batchCreateCustomVocabularyItemCmd = &cobra.Command{
	Use:   "batch-create-custom-vocabulary-item",
	Short: "Create a batch of custom vocabulary items for a given bot locale's custom vocabulary.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_batchCreateCustomVocabularyItemCmd).Standalone()

	lexv2Models_batchCreateCustomVocabularyItemCmd.Flags().String("bot-id", "", "The identifier of the bot associated with this custom vocabulary.")
	lexv2Models_batchCreateCustomVocabularyItemCmd.Flags().String("bot-version", "", "The identifier of the version of the bot associated with this custom vocabulary.")
	lexv2Models_batchCreateCustomVocabularyItemCmd.Flags().String("custom-vocabulary-item-list", "", "A list of new custom vocabulary items.")
	lexv2Models_batchCreateCustomVocabularyItemCmd.Flags().String("locale-id", "", "The identifier of the language and locale where this custom vocabulary is used.")
	lexv2Models_batchCreateCustomVocabularyItemCmd.MarkFlagRequired("bot-id")
	lexv2Models_batchCreateCustomVocabularyItemCmd.MarkFlagRequired("bot-version")
	lexv2Models_batchCreateCustomVocabularyItemCmd.MarkFlagRequired("custom-vocabulary-item-list")
	lexv2Models_batchCreateCustomVocabularyItemCmd.MarkFlagRequired("locale-id")
	lexv2ModelsCmd.AddCommand(lexv2Models_batchCreateCustomVocabularyItemCmd)
}
