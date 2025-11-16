package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_associateProtectConfigurationCmd = &cobra.Command{
	Use:   "associate-protect-configuration",
	Short: "Associate a protect configuration with a configuration set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_associateProtectConfigurationCmd).Standalone()

	pinpointSmsVoiceV2_associateProtectConfigurationCmd.Flags().String("configuration-set-name", "", "The name of the ConfigurationSet.")
	pinpointSmsVoiceV2_associateProtectConfigurationCmd.Flags().String("protect-configuration-id", "", "The unique identifier for the protect configuration.")
	pinpointSmsVoiceV2_associateProtectConfigurationCmd.MarkFlagRequired("configuration-set-name")
	pinpointSmsVoiceV2_associateProtectConfigurationCmd.MarkFlagRequired("protect-configuration-id")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_associateProtectConfigurationCmd)
}
