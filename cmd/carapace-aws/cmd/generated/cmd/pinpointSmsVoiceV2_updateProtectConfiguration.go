package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_updateProtectConfigurationCmd = &cobra.Command{
	Use:   "update-protect-configuration",
	Short: "Update the setting for an existing protect configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_updateProtectConfigurationCmd).Standalone()

	pinpointSmsVoiceV2_updateProtectConfigurationCmd.Flags().Bool("deletion-protection-enabled", false, "When set to true deletion protection is enabled.")
	pinpointSmsVoiceV2_updateProtectConfigurationCmd.Flags().Bool("no-deletion-protection-enabled", false, "When set to true deletion protection is enabled.")
	pinpointSmsVoiceV2_updateProtectConfigurationCmd.Flags().String("protect-configuration-id", "", "The unique identifier for the protect configuration.")
	pinpointSmsVoiceV2_updateProtectConfigurationCmd.Flag("no-deletion-protection-enabled").Hidden = true
	pinpointSmsVoiceV2_updateProtectConfigurationCmd.MarkFlagRequired("protect-configuration-id")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_updateProtectConfigurationCmd)
}
