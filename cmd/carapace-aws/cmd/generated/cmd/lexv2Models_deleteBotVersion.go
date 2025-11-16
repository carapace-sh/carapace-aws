package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_deleteBotVersionCmd = &cobra.Command{
	Use:   "delete-bot-version",
	Short: "Deletes a specific version of a bot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_deleteBotVersionCmd).Standalone()

	lexv2Models_deleteBotVersionCmd.Flags().String("bot-id", "", "The identifier of the bot that contains the version.")
	lexv2Models_deleteBotVersionCmd.Flags().String("bot-version", "", "The version of the bot to delete.")
	lexv2Models_deleteBotVersionCmd.Flags().String("skip-resource-in-use-check", "", "By default, Amazon Lex checks if any other resource, such as an alias or bot network, is using the bot version before it is deleted and throws a `ResourceInUseException` exception if the version is being used by another resource.")
	lexv2Models_deleteBotVersionCmd.MarkFlagRequired("bot-id")
	lexv2Models_deleteBotVersionCmd.MarkFlagRequired("bot-version")
	lexv2ModelsCmd.AddCommand(lexv2Models_deleteBotVersionCmd)
}
