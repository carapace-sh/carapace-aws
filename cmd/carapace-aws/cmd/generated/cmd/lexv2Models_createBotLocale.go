package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_createBotLocaleCmd = &cobra.Command{
	Use:   "create-bot-locale",
	Short: "Creates a locale in the bot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_createBotLocaleCmd).Standalone()

	lexv2Models_createBotLocaleCmd.Flags().String("bot-id", "", "The identifier of the bot to create the locale for.")
	lexv2Models_createBotLocaleCmd.Flags().String("bot-version", "", "The version of the bot to create the locale for.")
	lexv2Models_createBotLocaleCmd.Flags().String("description", "", "A description of the bot locale.")
	lexv2Models_createBotLocaleCmd.Flags().String("generative-aisettings", "", "")
	lexv2Models_createBotLocaleCmd.Flags().String("locale-id", "", "The identifier of the language and locale that the bot will be used in.")
	lexv2Models_createBotLocaleCmd.Flags().String("nlu-intent-confidence-threshold", "", "Determines the threshold where Amazon Lex will insert the `AMAZON.FallbackIntent`, `AMAZON.KendraSearchIntent`, or both when returning alternative intents.")
	lexv2Models_createBotLocaleCmd.Flags().String("voice-settings", "", "The Amazon Polly voice ID that Amazon Lex uses for voice interaction with the user.")
	lexv2Models_createBotLocaleCmd.MarkFlagRequired("bot-id")
	lexv2Models_createBotLocaleCmd.MarkFlagRequired("bot-version")
	lexv2Models_createBotLocaleCmd.MarkFlagRequired("locale-id")
	lexv2Models_createBotLocaleCmd.MarkFlagRequired("nlu-intent-confidence-threshold")
	lexv2ModelsCmd.AddCommand(lexv2Models_createBotLocaleCmd)
}
