package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_deleteAccountDefaultProtectConfigurationCmd = &cobra.Command{
	Use:   "delete-account-default-protect-configuration",
	Short: "Removes the current account default protect configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_deleteAccountDefaultProtectConfigurationCmd).Standalone()

	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_deleteAccountDefaultProtectConfigurationCmd)
}
