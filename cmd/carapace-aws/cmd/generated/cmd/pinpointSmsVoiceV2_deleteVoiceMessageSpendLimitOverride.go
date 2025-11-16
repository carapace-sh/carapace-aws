package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_deleteVoiceMessageSpendLimitOverrideCmd = &cobra.Command{
	Use:   "delete-voice-message-spend-limit-override",
	Short: "Deletes an account level monthly spend limit override for sending voice messages.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_deleteVoiceMessageSpendLimitOverrideCmd).Standalone()

	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_deleteVoiceMessageSpendLimitOverrideCmd)
}
