package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_setMediaMessageSpendLimitOverrideCmd = &cobra.Command{
	Use:   "set-media-message-spend-limit-override",
	Short: "Sets an account level monthly spend limit override for sending MMS messages.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_setMediaMessageSpendLimitOverrideCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointSmsVoiceV2_setMediaMessageSpendLimitOverrideCmd).Standalone()

		pinpointSmsVoiceV2_setMediaMessageSpendLimitOverrideCmd.Flags().String("monthly-limit", "", "The new monthly limit to enforce on text messages.")
		pinpointSmsVoiceV2_setMediaMessageSpendLimitOverrideCmd.MarkFlagRequired("monthly-limit")
	})
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_setMediaMessageSpendLimitOverrideCmd)
}
