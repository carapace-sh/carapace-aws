package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_disassociateProtectConfigurationCmd = &cobra.Command{
	Use:   "disassociate-protect-configuration",
	Short: "Disassociate a protect configuration from a configuration set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_disassociateProtectConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointSmsVoiceV2_disassociateProtectConfigurationCmd).Standalone()

		pinpointSmsVoiceV2_disassociateProtectConfigurationCmd.Flags().String("configuration-set-name", "", "The name of the ConfigurationSet.")
		pinpointSmsVoiceV2_disassociateProtectConfigurationCmd.Flags().String("protect-configuration-id", "", "The unique identifier for the protect configuration.")
		pinpointSmsVoiceV2_disassociateProtectConfigurationCmd.MarkFlagRequired("configuration-set-name")
		pinpointSmsVoiceV2_disassociateProtectConfigurationCmd.MarkFlagRequired("protect-configuration-id")
	})
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_disassociateProtectConfigurationCmd)
}
