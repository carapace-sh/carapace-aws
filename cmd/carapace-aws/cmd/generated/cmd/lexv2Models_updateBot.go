package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_updateBotCmd = &cobra.Command{
	Use:   "update-bot",
	Short: "Updates the configuration of an existing bot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_updateBotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_updateBotCmd).Standalone()

		lexv2Models_updateBotCmd.Flags().String("bot-id", "", "The unique identifier of the bot to update.")
		lexv2Models_updateBotCmd.Flags().String("bot-members", "", "The list of bot members in the network associated with the update action.")
		lexv2Models_updateBotCmd.Flags().String("bot-name", "", "The new name of the bot.")
		lexv2Models_updateBotCmd.Flags().String("bot-type", "", "The type of the bot to be updated.")
		lexv2Models_updateBotCmd.Flags().String("data-privacy", "", "Provides information on additional privacy protections Amazon Lex should use with the bot's data.")
		lexv2Models_updateBotCmd.Flags().String("description", "", "A description of the bot.")
		lexv2Models_updateBotCmd.Flags().String("error-log-settings", "", "Allows you to modify how Amazon Lex logs errors during bot interactions, including destinations for error logs and the types of errors to be captured.")
		lexv2Models_updateBotCmd.Flags().String("idle-session-ttlin-seconds", "", "The time, in seconds, that Amazon Lex should keep information about a user's conversation with the bot.")
		lexv2Models_updateBotCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of an IAM role that has permissions to access the bot.")
		lexv2Models_updateBotCmd.MarkFlagRequired("bot-id")
		lexv2Models_updateBotCmd.MarkFlagRequired("bot-name")
		lexv2Models_updateBotCmd.MarkFlagRequired("data-privacy")
		lexv2Models_updateBotCmd.MarkFlagRequired("idle-session-ttlin-seconds")
		lexv2Models_updateBotCmd.MarkFlagRequired("role-arn")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_updateBotCmd)
}
