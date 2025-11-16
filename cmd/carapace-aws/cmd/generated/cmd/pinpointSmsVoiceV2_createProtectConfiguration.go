package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_createProtectConfigurationCmd = &cobra.Command{
	Use:   "create-protect-configuration",
	Short: "Create a new protect configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_createProtectConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointSmsVoiceV2_createProtectConfigurationCmd).Standalone()

		pinpointSmsVoiceV2_createProtectConfigurationCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		pinpointSmsVoiceV2_createProtectConfigurationCmd.Flags().Bool("deletion-protection-enabled", false, "When set to true deletion protection is enabled.")
		pinpointSmsVoiceV2_createProtectConfigurationCmd.Flags().Bool("no-deletion-protection-enabled", false, "When set to true deletion protection is enabled.")
		pinpointSmsVoiceV2_createProtectConfigurationCmd.Flags().String("tags", "", "An array of key and value pair tags that are associated with the resource.")
		pinpointSmsVoiceV2_createProtectConfigurationCmd.Flag("no-deletion-protection-enabled").Hidden = true
	})
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_createProtectConfigurationCmd)
}
