package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_setTextMessageSpendLimitOverrideCmd = &cobra.Command{
	Use:   "set-text-message-spend-limit-override",
	Short: "Sets an account level monthly spend limit override for sending text messages.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_setTextMessageSpendLimitOverrideCmd).Standalone()

	pinpointSmsVoiceV2_setTextMessageSpendLimitOverrideCmd.Flags().String("monthly-limit", "", "The new monthly limit to enforce on text messages.")
	pinpointSmsVoiceV2_setTextMessageSpendLimitOverrideCmd.MarkFlagRequired("monthly-limit")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_setTextMessageSpendLimitOverrideCmd)
}
