package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_describeBotVersionCmd = &cobra.Command{
	Use:   "describe-bot-version",
	Short: "Provides metadata about a version of a bot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_describeBotVersionCmd).Standalone()

	lexv2Models_describeBotVersionCmd.Flags().String("bot-id", "", "The identifier of the bot containing the version to return metadata for.")
	lexv2Models_describeBotVersionCmd.Flags().String("bot-version", "", "The version of the bot to return metadata for.")
	lexv2Models_describeBotVersionCmd.MarkFlagRequired("bot-id")
	lexv2Models_describeBotVersionCmd.MarkFlagRequired("bot-version")
	lexv2ModelsCmd.AddCommand(lexv2Models_describeBotVersionCmd)
}
