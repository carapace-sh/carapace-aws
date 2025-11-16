package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_buildBotLocaleCmd = &cobra.Command{
	Use:   "build-bot-locale",
	Short: "Builds a bot, its intents, and its slot types into a specific locale.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_buildBotLocaleCmd).Standalone()

	lexv2Models_buildBotLocaleCmd.Flags().String("bot-id", "", "The identifier of the bot to build.")
	lexv2Models_buildBotLocaleCmd.Flags().String("bot-version", "", "The version of the bot to build.")
	lexv2Models_buildBotLocaleCmd.Flags().String("locale-id", "", "The identifier of the language and locale that the bot will be used in.")
	lexv2Models_buildBotLocaleCmd.MarkFlagRequired("bot-id")
	lexv2Models_buildBotLocaleCmd.MarkFlagRequired("bot-version")
	lexv2Models_buildBotLocaleCmd.MarkFlagRequired("locale-id")
	lexv2ModelsCmd.AddCommand(lexv2Models_buildBotLocaleCmd)
}
