package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_getProtectConfigurationCountryRuleSetCmd = &cobra.Command{
	Use:   "get-protect-configuration-country-rule-set",
	Short: "Retrieve the CountryRuleSet for the specified NumberCapability from a protect configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_getProtectConfigurationCountryRuleSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointSmsVoiceV2_getProtectConfigurationCountryRuleSetCmd).Standalone()

		pinpointSmsVoiceV2_getProtectConfigurationCountryRuleSetCmd.Flags().String("number-capability", "", "The capability type to return the CountryRuleSet for.")
		pinpointSmsVoiceV2_getProtectConfigurationCountryRuleSetCmd.Flags().String("protect-configuration-id", "", "The unique identifier for the protect configuration.")
		pinpointSmsVoiceV2_getProtectConfigurationCountryRuleSetCmd.MarkFlagRequired("number-capability")
		pinpointSmsVoiceV2_getProtectConfigurationCountryRuleSetCmd.MarkFlagRequired("protect-configuration-id")
	})
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_getProtectConfigurationCountryRuleSetCmd)
}
