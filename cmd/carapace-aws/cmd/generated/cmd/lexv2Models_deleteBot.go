package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_deleteBotCmd = &cobra.Command{
	Use:   "delete-bot",
	Short: "Deletes all versions of a bot, including the `Draft` version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_deleteBotCmd).Standalone()

	lexv2Models_deleteBotCmd.Flags().String("bot-id", "", "The identifier of the bot to delete.")
	lexv2Models_deleteBotCmd.Flags().String("skip-resource-in-use-check", "", "By default, Amazon Lex checks if any other resource, such as an alias or bot network, is using the bot version before it is deleted and throws a `ResourceInUseException` exception if the bot is being used by another resource.")
	lexv2Models_deleteBotCmd.MarkFlagRequired("bot-id")
	lexv2ModelsCmd.AddCommand(lexv2Models_deleteBotCmd)
}
