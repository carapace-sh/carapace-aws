package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_describeBotLocaleCmd = &cobra.Command{
	Use:   "describe-bot-locale",
	Short: "Describes the settings that a bot has for a specific locale.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_describeBotLocaleCmd).Standalone()

	lexv2Models_describeBotLocaleCmd.Flags().String("bot-id", "", "The identifier of the bot associated with the locale.")
	lexv2Models_describeBotLocaleCmd.Flags().String("bot-version", "", "The version of the bot associated with the locale.")
	lexv2Models_describeBotLocaleCmd.Flags().String("locale-id", "", "The unique identifier of the locale to describe.")
	lexv2Models_describeBotLocaleCmd.MarkFlagRequired("bot-id")
	lexv2Models_describeBotLocaleCmd.MarkFlagRequired("bot-version")
	lexv2Models_describeBotLocaleCmd.MarkFlagRequired("locale-id")
	lexv2ModelsCmd.AddCommand(lexv2Models_describeBotLocaleCmd)
}
