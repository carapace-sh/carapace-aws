package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_describeBotCmd = &cobra.Command{
	Use:   "describe-bot",
	Short: "Provides metadata information about a bot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_describeBotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_describeBotCmd).Standalone()

		lexv2Models_describeBotCmd.Flags().String("bot-id", "", "The unique identifier of the bot to describe.")
		lexv2Models_describeBotCmd.MarkFlagRequired("bot-id")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_describeBotCmd)
}
