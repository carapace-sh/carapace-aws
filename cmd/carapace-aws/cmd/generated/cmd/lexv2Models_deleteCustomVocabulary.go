package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_deleteCustomVocabularyCmd = &cobra.Command{
	Use:   "delete-custom-vocabulary",
	Short: "Removes a custom vocabulary from the specified locale in the specified bot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_deleteCustomVocabularyCmd).Standalone()

	lexv2Models_deleteCustomVocabularyCmd.Flags().String("bot-id", "", "The unique identifier of the bot to remove the custom vocabulary from.")
	lexv2Models_deleteCustomVocabularyCmd.Flags().String("bot-version", "", "The version of the bot to remove the custom vocabulary from.")
	lexv2Models_deleteCustomVocabularyCmd.Flags().String("locale-id", "", "The locale identifier for the locale that contains the custom vocabulary to remove.")
	lexv2Models_deleteCustomVocabularyCmd.MarkFlagRequired("bot-id")
	lexv2Models_deleteCustomVocabularyCmd.MarkFlagRequired("bot-version")
	lexv2Models_deleteCustomVocabularyCmd.MarkFlagRequired("locale-id")
	lexv2ModelsCmd.AddCommand(lexv2Models_deleteCustomVocabularyCmd)
}
