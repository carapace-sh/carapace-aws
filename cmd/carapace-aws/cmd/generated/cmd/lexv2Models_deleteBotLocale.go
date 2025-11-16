package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_deleteBotLocaleCmd = &cobra.Command{
	Use:   "delete-bot-locale",
	Short: "Removes a locale from a bot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_deleteBotLocaleCmd).Standalone()

	lexv2Models_deleteBotLocaleCmd.Flags().String("bot-id", "", "The unique identifier of the bot that contains the locale.")
	lexv2Models_deleteBotLocaleCmd.Flags().String("bot-version", "", "The version of the bot that contains the locale.")
	lexv2Models_deleteBotLocaleCmd.Flags().String("locale-id", "", "The identifier of the language and locale that will be deleted.")
	lexv2Models_deleteBotLocaleCmd.MarkFlagRequired("bot-id")
	lexv2Models_deleteBotLocaleCmd.MarkFlagRequired("bot-version")
	lexv2Models_deleteBotLocaleCmd.MarkFlagRequired("locale-id")
	lexv2ModelsCmd.AddCommand(lexv2Models_deleteBotLocaleCmd)
}
