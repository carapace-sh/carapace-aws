package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_deleteProtectConfigurationRuleSetNumberOverrideCmd = &cobra.Command{
	Use:   "delete-protect-configuration-rule-set-number-override",
	Short: "Permanently delete the protect configuration rule set number override.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_deleteProtectConfigurationRuleSetNumberOverrideCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointSmsVoiceV2_deleteProtectConfigurationRuleSetNumberOverrideCmd).Standalone()

		pinpointSmsVoiceV2_deleteProtectConfigurationRuleSetNumberOverrideCmd.Flags().String("destination-phone-number", "", "The destination phone number in E.164 format.")
		pinpointSmsVoiceV2_deleteProtectConfigurationRuleSetNumberOverrideCmd.Flags().String("protect-configuration-id", "", "The unique identifier for the protect configuration.")
		pinpointSmsVoiceV2_deleteProtectConfigurationRuleSetNumberOverrideCmd.MarkFlagRequired("destination-phone-number")
		pinpointSmsVoiceV2_deleteProtectConfigurationRuleSetNumberOverrideCmd.MarkFlagRequired("protect-configuration-id")
	})
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_deleteProtectConfigurationRuleSetNumberOverrideCmd)
}
