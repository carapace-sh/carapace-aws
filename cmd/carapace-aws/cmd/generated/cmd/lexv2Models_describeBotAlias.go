package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_describeBotAliasCmd = &cobra.Command{
	Use:   "describe-bot-alias",
	Short: "Get information about a specific bot alias.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_describeBotAliasCmd).Standalone()

	lexv2Models_describeBotAliasCmd.Flags().String("bot-alias-id", "", "The identifier of the bot alias to describe.")
	lexv2Models_describeBotAliasCmd.Flags().String("bot-id", "", "The identifier of the bot associated with the bot alias to describe.")
	lexv2Models_describeBotAliasCmd.MarkFlagRequired("bot-alias-id")
	lexv2Models_describeBotAliasCmd.MarkFlagRequired("bot-id")
	lexv2ModelsCmd.AddCommand(lexv2Models_describeBotAliasCmd)
}
