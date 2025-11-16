package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_listProtectConfigurationRuleSetNumberOverridesCmd = &cobra.Command{
	Use:   "list-protect-configuration-rule-set-number-overrides",
	Short: "Retrieve all of the protect configuration rule set number overrides that match the filters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_listProtectConfigurationRuleSetNumberOverridesCmd).Standalone()

	pinpointSmsVoiceV2_listProtectConfigurationRuleSetNumberOverridesCmd.Flags().String("filters", "", "An array of ProtectConfigurationRuleSetNumberOverrideFilterItem objects to filter the results.")
	pinpointSmsVoiceV2_listProtectConfigurationRuleSetNumberOverridesCmd.Flags().String("max-results", "", "The maximum number of results to return per each request.")
	pinpointSmsVoiceV2_listProtectConfigurationRuleSetNumberOverridesCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
	pinpointSmsVoiceV2_listProtectConfigurationRuleSetNumberOverridesCmd.Flags().String("protect-configuration-id", "", "The unique identifier for the protect configuration.")
	pinpointSmsVoiceV2_listProtectConfigurationRuleSetNumberOverridesCmd.MarkFlagRequired("protect-configuration-id")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_listProtectConfigurationRuleSetNumberOverridesCmd)
}
