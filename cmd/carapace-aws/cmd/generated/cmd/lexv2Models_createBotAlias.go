package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_createBotAliasCmd = &cobra.Command{
	Use:   "create-bot-alias",
	Short: "Creates an alias for the specified version of a bot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_createBotAliasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_createBotAliasCmd).Standalone()

		lexv2Models_createBotAliasCmd.Flags().String("bot-alias-locale-settings", "", "Maps configuration information to a specific locale.")
		lexv2Models_createBotAliasCmd.Flags().String("bot-alias-name", "", "The alias to create.")
		lexv2Models_createBotAliasCmd.Flags().String("bot-id", "", "The unique identifier of the bot that the alias applies to.")
		lexv2Models_createBotAliasCmd.Flags().String("bot-version", "", "The version of the bot that this alias points to.")
		lexv2Models_createBotAliasCmd.Flags().String("conversation-log-settings", "", "Specifies whether Amazon Lex logs text and audio for a conversation with the bot.")
		lexv2Models_createBotAliasCmd.Flags().String("description", "", "A description of the alias.")
		lexv2Models_createBotAliasCmd.Flags().String("sentiment-analysis-settings", "", "")
		lexv2Models_createBotAliasCmd.Flags().String("tags", "", "A list of tags to add to the bot alias.")
		lexv2Models_createBotAliasCmd.MarkFlagRequired("bot-alias-name")
		lexv2Models_createBotAliasCmd.MarkFlagRequired("bot-id")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_createBotAliasCmd)
}
