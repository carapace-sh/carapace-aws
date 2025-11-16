package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_deleteMediaMessageSpendLimitOverrideCmd = &cobra.Command{
	Use:   "delete-media-message-spend-limit-override",
	Short: "Deletes an account-level monthly spending limit override for sending multimedia messages (MMS).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_deleteMediaMessageSpendLimitOverrideCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointSmsVoiceV2_deleteMediaMessageSpendLimitOverrideCmd).Standalone()

	})
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_deleteMediaMessageSpendLimitOverrideCmd)
}
