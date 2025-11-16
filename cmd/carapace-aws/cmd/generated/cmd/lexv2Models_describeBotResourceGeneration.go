package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_describeBotResourceGenerationCmd = &cobra.Command{
	Use:   "describe-bot-resource-generation",
	Short: "Returns information about a request to generate a bot through natural language description, made through the `StartBotResource` API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_describeBotResourceGenerationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_describeBotResourceGenerationCmd).Standalone()

		lexv2Models_describeBotResourceGenerationCmd.Flags().String("bot-id", "", "The unique identifier of the bot for which to return the generation details.")
		lexv2Models_describeBotResourceGenerationCmd.Flags().String("bot-version", "", "The version of the bot for which to return the generation details.")
		lexv2Models_describeBotResourceGenerationCmd.Flags().String("generation-id", "", "The unique identifier of the generation request for which to return the generation details.")
		lexv2Models_describeBotResourceGenerationCmd.Flags().String("locale-id", "", "The locale of the bot for which to return the generation details.")
		lexv2Models_describeBotResourceGenerationCmd.MarkFlagRequired("bot-id")
		lexv2Models_describeBotResourceGenerationCmd.MarkFlagRequired("bot-version")
		lexv2Models_describeBotResourceGenerationCmd.MarkFlagRequired("generation-id")
		lexv2Models_describeBotResourceGenerationCmd.MarkFlagRequired("locale-id")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_describeBotResourceGenerationCmd)
}
