package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_updateBotLocaleCmd = &cobra.Command{
	Use:   "update-bot-locale",
	Short: "Updates the settings that a bot has for a specific locale.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_updateBotLocaleCmd).Standalone()

	lexv2Models_updateBotLocaleCmd.Flags().String("bot-id", "", "The unique identifier of the bot that contains the locale.")
	lexv2Models_updateBotLocaleCmd.Flags().String("bot-version", "", "The version of the bot that contains the locale to be updated.")
	lexv2Models_updateBotLocaleCmd.Flags().String("description", "", "The new description of the locale.")
	lexv2Models_updateBotLocaleCmd.Flags().String("generative-aisettings", "", "Contains settings for generative AI features powered by Amazon Bedrock for your bot locale.")
	lexv2Models_updateBotLocaleCmd.Flags().String("locale-id", "", "The identifier of the language and locale to update.")
	lexv2Models_updateBotLocaleCmd.Flags().String("nlu-intent-confidence-threshold", "", "The new confidence threshold where Amazon Lex inserts the `AMAZON.FallbackIntent` and `AMAZON.KendraSearchIntent` intents in the list of possible intents for an utterance.")
	lexv2Models_updateBotLocaleCmd.Flags().String("voice-settings", "", "The new Amazon Polly voice Amazon Lex should use for voice interaction with the user.")
	lexv2Models_updateBotLocaleCmd.MarkFlagRequired("bot-id")
	lexv2Models_updateBotLocaleCmd.MarkFlagRequired("bot-version")
	lexv2Models_updateBotLocaleCmd.MarkFlagRequired("locale-id")
	lexv2Models_updateBotLocaleCmd.MarkFlagRequired("nlu-intent-confidence-threshold")
	lexv2ModelsCmd.AddCommand(lexv2Models_updateBotLocaleCmd)
}
