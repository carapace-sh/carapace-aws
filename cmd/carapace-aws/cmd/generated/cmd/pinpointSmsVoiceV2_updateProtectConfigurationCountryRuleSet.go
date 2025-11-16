package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_updateProtectConfigurationCountryRuleSetCmd = &cobra.Command{
	Use:   "update-protect-configuration-country-rule-set",
	Short: "Update a country rule set to `ALLOW`, `BLOCK`, `MONITOR`, or `FILTER` messages to be sent to the specified destination counties.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_updateProtectConfigurationCountryRuleSetCmd).Standalone()

	pinpointSmsVoiceV2_updateProtectConfigurationCountryRuleSetCmd.Flags().String("country-rule-set-updates", "", "A map of ProtectConfigurationCountryRuleSetInformation objects that contain the details for the requested NumberCapability.")
	pinpointSmsVoiceV2_updateProtectConfigurationCountryRuleSetCmd.Flags().String("number-capability", "", "The number capability to apply the CountryRuleSetUpdates updates to.")
	pinpointSmsVoiceV2_updateProtectConfigurationCountryRuleSetCmd.Flags().String("protect-configuration-id", "", "The unique identifier for the protect configuration.")
	pinpointSmsVoiceV2_updateProtectConfigurationCountryRuleSetCmd.MarkFlagRequired("country-rule-set-updates")
	pinpointSmsVoiceV2_updateProtectConfigurationCountryRuleSetCmd.MarkFlagRequired("number-capability")
	pinpointSmsVoiceV2_updateProtectConfigurationCountryRuleSetCmd.MarkFlagRequired("protect-configuration-id")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_updateProtectConfigurationCountryRuleSetCmd)
}
