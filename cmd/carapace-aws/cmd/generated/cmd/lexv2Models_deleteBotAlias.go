package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_deleteBotAliasCmd = &cobra.Command{
	Use:   "delete-bot-alias",
	Short: "Deletes the specified bot alias.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_deleteBotAliasCmd).Standalone()

	lexv2Models_deleteBotAliasCmd.Flags().String("bot-alias-id", "", "The unique identifier of the bot alias to delete.")
	lexv2Models_deleteBotAliasCmd.Flags().String("bot-id", "", "The unique identifier of the bot associated with the alias to delete.")
	lexv2Models_deleteBotAliasCmd.Flags().String("skip-resource-in-use-check", "", "By default, Amazon Lex checks if any other resource, such as a bot network, is using the bot alias before it is deleted and throws a `ResourceInUseException` exception if the alias is being used by another resource.")
	lexv2Models_deleteBotAliasCmd.MarkFlagRequired("bot-alias-id")
	lexv2Models_deleteBotAliasCmd.MarkFlagRequired("bot-id")
	lexv2ModelsCmd.AddCommand(lexv2Models_deleteBotAliasCmd)
}
