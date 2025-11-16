package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_startBotResourceGenerationCmd = &cobra.Command{
	Use:   "start-bot-resource-generation",
	Short: "Starts a request for the descriptive bot builder to generate a bot locale configuration based on the prompt you provide it.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_startBotResourceGenerationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_startBotResourceGenerationCmd).Standalone()

		lexv2Models_startBotResourceGenerationCmd.Flags().String("bot-id", "", "The unique identifier of the bot for which to generate intents and slot types.")
		lexv2Models_startBotResourceGenerationCmd.Flags().String("bot-version", "", "The version of the bot for which to generate intents and slot types.")
		lexv2Models_startBotResourceGenerationCmd.Flags().String("generation-input-prompt", "", "The prompt to generate intents and slot types for the bot locale.")
		lexv2Models_startBotResourceGenerationCmd.Flags().String("locale-id", "", "The locale of the bot for which to generate intents and slot types.")
		lexv2Models_startBotResourceGenerationCmd.MarkFlagRequired("bot-id")
		lexv2Models_startBotResourceGenerationCmd.MarkFlagRequired("bot-version")
		lexv2Models_startBotResourceGenerationCmd.MarkFlagRequired("generation-input-prompt")
		lexv2Models_startBotResourceGenerationCmd.MarkFlagRequired("locale-id")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_startBotResourceGenerationCmd)
}
