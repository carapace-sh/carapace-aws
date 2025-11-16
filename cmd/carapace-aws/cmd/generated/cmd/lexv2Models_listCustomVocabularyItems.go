package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_listCustomVocabularyItemsCmd = &cobra.Command{
	Use:   "list-custom-vocabulary-items",
	Short: "Paginated list of custom vocabulary items for a given bot locale's custom vocabulary.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_listCustomVocabularyItemsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_listCustomVocabularyItemsCmd).Standalone()

		lexv2Models_listCustomVocabularyItemsCmd.Flags().String("bot-id", "", "The identifier of the version of the bot associated with this custom vocabulary.")
		lexv2Models_listCustomVocabularyItemsCmd.Flags().String("bot-version", "", "The bot version of the bot to the list custom vocabulary request.")
		lexv2Models_listCustomVocabularyItemsCmd.Flags().String("locale-id", "", "The identifier of the language and locale where this custom vocabulary is used.")
		lexv2Models_listCustomVocabularyItemsCmd.Flags().String("max-results", "", "The maximum number of items returned by the list operation.")
		lexv2Models_listCustomVocabularyItemsCmd.Flags().String("next-token", "", "The nextToken identifier to the list custom vocabulary request.")
		lexv2Models_listCustomVocabularyItemsCmd.MarkFlagRequired("bot-id")
		lexv2Models_listCustomVocabularyItemsCmd.MarkFlagRequired("bot-version")
		lexv2Models_listCustomVocabularyItemsCmd.MarkFlagRequired("locale-id")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_listCustomVocabularyItemsCmd)
}
