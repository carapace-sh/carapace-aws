package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_setVoiceMessageSpendLimitOverrideCmd = &cobra.Command{
	Use:   "set-voice-message-spend-limit-override",
	Short: "Sets an account level monthly spend limit override for sending voice messages.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_setVoiceMessageSpendLimitOverrideCmd).Standalone()

	pinpointSmsVoiceV2_setVoiceMessageSpendLimitOverrideCmd.Flags().String("monthly-limit", "", "The new monthly limit to enforce on voice messages.")
	pinpointSmsVoiceV2_setVoiceMessageSpendLimitOverrideCmd.MarkFlagRequired("monthly-limit")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_setVoiceMessageSpendLimitOverrideCmd)
}
