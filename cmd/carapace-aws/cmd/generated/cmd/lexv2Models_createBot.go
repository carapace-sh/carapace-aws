package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_createBotCmd = &cobra.Command{
	Use:   "create-bot",
	Short: "Creates an Amazon Lex conversational bot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_createBotCmd).Standalone()

	lexv2Models_createBotCmd.Flags().String("bot-members", "", "The list of bot members in a network to be created.")
	lexv2Models_createBotCmd.Flags().String("bot-name", "", "The name of the bot.")
	lexv2Models_createBotCmd.Flags().String("bot-tags", "", "A list of tags to add to the bot.")
	lexv2Models_createBotCmd.Flags().String("bot-type", "", "The type of a bot to create.")
	lexv2Models_createBotCmd.Flags().String("data-privacy", "", "Provides information on additional privacy protections Amazon Lex should use with the bot's data.")
	lexv2Models_createBotCmd.Flags().String("description", "", "A description of the bot.")
	lexv2Models_createBotCmd.Flags().String("error-log-settings", "", "Specifies the configuration for error logging during bot creation.")
	lexv2Models_createBotCmd.Flags().String("idle-session-ttlin-seconds", "", "The time, in seconds, that Amazon Lex should keep information about a user's conversation with the bot.")
	lexv2Models_createBotCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of an IAM role that has permission to access the bot.")
	lexv2Models_createBotCmd.Flags().String("test-bot-alias-tags", "", "A list of tags to add to the test alias for a bot.")
	lexv2Models_createBotCmd.MarkFlagRequired("bot-name")
	lexv2Models_createBotCmd.MarkFlagRequired("data-privacy")
	lexv2Models_createBotCmd.MarkFlagRequired("idle-session-ttlin-seconds")
	lexv2Models_createBotCmd.MarkFlagRequired("role-arn")
	lexv2ModelsCmd.AddCommand(lexv2Models_createBotCmd)
}
