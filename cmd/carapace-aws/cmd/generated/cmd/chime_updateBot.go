package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_updateBotCmd = &cobra.Command{
	Use:   "update-bot",
	Short: "Updates the status of the specified bot, such as starting or stopping the bot from running in your Amazon Chime Enterprise account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_updateBotCmd).Standalone()

	chime_updateBotCmd.Flags().String("account-id", "", "The Amazon Chime account ID.")
	chime_updateBotCmd.Flags().String("bot-id", "", "The bot ID.")
	chime_updateBotCmd.Flags().String("disabled", "", "When true, stops the specified bot from running in your account.")
	chime_updateBotCmd.MarkFlagRequired("account-id")
	chime_updateBotCmd.MarkFlagRequired("bot-id")
	chimeCmd.AddCommand(chime_updateBotCmd)
}
