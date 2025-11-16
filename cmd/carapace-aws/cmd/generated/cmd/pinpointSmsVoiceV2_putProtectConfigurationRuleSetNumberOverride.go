package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_putProtectConfigurationRuleSetNumberOverrideCmd = &cobra.Command{
	Use:   "put-protect-configuration-rule-set-number-override",
	Short: "Create or update a phone number rule override and associate it with a protect configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_putProtectConfigurationRuleSetNumberOverrideCmd).Standalone()

	pinpointSmsVoiceV2_putProtectConfigurationRuleSetNumberOverrideCmd.Flags().String("action", "", "The action for the rule to either block or allow messages to the destination phone number.")
	pinpointSmsVoiceV2_putProtectConfigurationRuleSetNumberOverrideCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	pinpointSmsVoiceV2_putProtectConfigurationRuleSetNumberOverrideCmd.Flags().String("destination-phone-number", "", "The destination phone number in E.164 format.")
	pinpointSmsVoiceV2_putProtectConfigurationRuleSetNumberOverrideCmd.Flags().String("expiration-timestamp", "", "The time the rule will expire at.")
	pinpointSmsVoiceV2_putProtectConfigurationRuleSetNumberOverrideCmd.Flags().String("protect-configuration-id", "", "The unique identifier for the protect configuration.")
	pinpointSmsVoiceV2_putProtectConfigurationRuleSetNumberOverrideCmd.MarkFlagRequired("action")
	pinpointSmsVoiceV2_putProtectConfigurationRuleSetNumberOverrideCmd.MarkFlagRequired("destination-phone-number")
	pinpointSmsVoiceV2_putProtectConfigurationRuleSetNumberOverrideCmd.MarkFlagRequired("protect-configuration-id")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_putProtectConfigurationRuleSetNumberOverrideCmd)
}
