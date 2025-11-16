package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_setAccountDefaultProtectConfigurationCmd = &cobra.Command{
	Use:   "set-account-default-protect-configuration",
	Short: "Set a protect configuration as your account default.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_setAccountDefaultProtectConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointSmsVoiceV2_setAccountDefaultProtectConfigurationCmd).Standalone()

		pinpointSmsVoiceV2_setAccountDefaultProtectConfigurationCmd.Flags().String("protect-configuration-id", "", "The unique identifier for the protect configuration.")
		pinpointSmsVoiceV2_setAccountDefaultProtectConfigurationCmd.MarkFlagRequired("protect-configuration-id")
	})
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_setAccountDefaultProtectConfigurationCmd)
}
