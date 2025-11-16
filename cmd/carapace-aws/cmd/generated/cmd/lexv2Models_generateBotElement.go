package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_generateBotElementCmd = &cobra.Command{
	Use:   "generate-bot-element",
	Short: "Generates sample utterances for an intent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_generateBotElementCmd).Standalone()

	lexv2Models_generateBotElementCmd.Flags().String("bot-id", "", "The bot unique Id for the bot request to generate utterances.")
	lexv2Models_generateBotElementCmd.Flags().String("bot-version", "", "The bot version for the bot request to generate utterances.")
	lexv2Models_generateBotElementCmd.Flags().String("intent-id", "", "The intent unique Id for the bot request to generate utterances.")
	lexv2Models_generateBotElementCmd.Flags().String("locale-id", "", "The unique locale Id for the bot request to generate utterances.")
	lexv2Models_generateBotElementCmd.MarkFlagRequired("bot-id")
	lexv2Models_generateBotElementCmd.MarkFlagRequired("bot-version")
	lexv2Models_generateBotElementCmd.MarkFlagRequired("intent-id")
	lexv2Models_generateBotElementCmd.MarkFlagRequired("locale-id")
	lexv2ModelsCmd.AddCommand(lexv2Models_generateBotElementCmd)
}
