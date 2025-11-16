package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_deleteTextMessageSpendLimitOverrideCmd = &cobra.Command{
	Use:   "delete-text-message-spend-limit-override",
	Short: "Deletes an account-level monthly spending limit override for sending text messages.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_deleteTextMessageSpendLimitOverrideCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointSmsVoiceV2_deleteTextMessageSpendLimitOverrideCmd).Standalone()

	})
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_deleteTextMessageSpendLimitOverrideCmd)
}
