package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_describeCustomVocabularyMetadataCmd = &cobra.Command{
	Use:   "describe-custom-vocabulary-metadata",
	Short: "Provides metadata information about a custom vocabulary.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_describeCustomVocabularyMetadataCmd).Standalone()

	lexv2Models_describeCustomVocabularyMetadataCmd.Flags().String("bot-id", "", "The unique identifier of the bot that contains the custom vocabulary.")
	lexv2Models_describeCustomVocabularyMetadataCmd.Flags().String("bot-version", "", "The bot version of the bot to return metadata for.")
	lexv2Models_describeCustomVocabularyMetadataCmd.Flags().String("locale-id", "", "The locale to return the custom vocabulary information for.")
	lexv2Models_describeCustomVocabularyMetadataCmd.MarkFlagRequired("bot-id")
	lexv2Models_describeCustomVocabularyMetadataCmd.MarkFlagRequired("bot-version")
	lexv2Models_describeCustomVocabularyMetadataCmd.MarkFlagRequired("locale-id")
	lexv2ModelsCmd.AddCommand(lexv2Models_describeCustomVocabularyMetadataCmd)
}
