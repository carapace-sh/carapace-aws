package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexModels_deleteBotAliasCmd = &cobra.Command{
	Use:   "delete-bot-alias",
	Short: "Deletes an alias for the specified bot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexModels_deleteBotAliasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexModels_deleteBotAliasCmd).Standalone()

		lexModels_deleteBotAliasCmd.Flags().String("bot-name", "", "The name of the bot that the alias points to.")
		lexModels_deleteBotAliasCmd.Flags().String("name", "", "The name of the alias to delete.")
		lexModels_deleteBotAliasCmd.MarkFlagRequired("bot-name")
		lexModels_deleteBotAliasCmd.MarkFlagRequired("name")
	})
	lexModelsCmd.AddCommand(lexModels_deleteBotAliasCmd)
}
