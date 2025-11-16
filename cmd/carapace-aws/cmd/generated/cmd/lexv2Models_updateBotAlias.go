package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_updateBotAliasCmd = &cobra.Command{
	Use:   "update-bot-alias",
	Short: "Updates the configuration of an existing bot alias.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_updateBotAliasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_updateBotAliasCmd).Standalone()

		lexv2Models_updateBotAliasCmd.Flags().String("bot-alias-id", "", "The unique identifier of the bot alias.")
		lexv2Models_updateBotAliasCmd.Flags().String("bot-alias-locale-settings", "", "The new Lambda functions to use in each locale for the bot alias.")
		lexv2Models_updateBotAliasCmd.Flags().String("bot-alias-name", "", "The new name to assign to the bot alias.")
		lexv2Models_updateBotAliasCmd.Flags().String("bot-id", "", "The identifier of the bot with the updated alias.")
		lexv2Models_updateBotAliasCmd.Flags().String("bot-version", "", "The new bot version to assign to the bot alias.")
		lexv2Models_updateBotAliasCmd.Flags().String("conversation-log-settings", "", "The new settings for storing conversation logs in Amazon CloudWatch Logs and Amazon S3 buckets.")
		lexv2Models_updateBotAliasCmd.Flags().String("description", "", "The new description to assign to the bot alias.")
		lexv2Models_updateBotAliasCmd.Flags().String("sentiment-analysis-settings", "", "")
		lexv2Models_updateBotAliasCmd.MarkFlagRequired("bot-alias-id")
		lexv2Models_updateBotAliasCmd.MarkFlagRequired("bot-alias-name")
		lexv2Models_updateBotAliasCmd.MarkFlagRequired("bot-id")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_updateBotAliasCmd)
}
