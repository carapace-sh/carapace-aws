package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexModels_putBotAliasCmd = &cobra.Command{
	Use:   "put-bot-alias",
	Short: "Creates an alias for the specified version of the bot or replaces an alias for the specified bot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexModels_putBotAliasCmd).Standalone()

	lexModels_putBotAliasCmd.Flags().String("bot-name", "", "The name of the bot.")
	lexModels_putBotAliasCmd.Flags().String("bot-version", "", "The version of the bot.")
	lexModels_putBotAliasCmd.Flags().String("checksum", "", "Identifies a specific revision of the `$LATEST` version.")
	lexModels_putBotAliasCmd.Flags().String("conversation-logs", "", "Settings for conversation logs for the alias.")
	lexModels_putBotAliasCmd.Flags().String("description", "", "A description of the alias.")
	lexModels_putBotAliasCmd.Flags().String("name", "", "The name of the alias.")
	lexModels_putBotAliasCmd.Flags().String("tags", "", "A list of tags to add to the bot alias.")
	lexModels_putBotAliasCmd.MarkFlagRequired("bot-name")
	lexModels_putBotAliasCmd.MarkFlagRequired("bot-version")
	lexModels_putBotAliasCmd.MarkFlagRequired("name")
	lexModelsCmd.AddCommand(lexModels_putBotAliasCmd)
}
