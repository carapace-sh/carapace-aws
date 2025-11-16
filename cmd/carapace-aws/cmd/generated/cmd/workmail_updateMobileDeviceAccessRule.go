package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_updateMobileDeviceAccessRuleCmd = &cobra.Command{
	Use:   "update-mobile-device-access-rule",
	Short: "Updates a mobile device access rule for the specified WorkMail organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_updateMobileDeviceAccessRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workmail_updateMobileDeviceAccessRuleCmd).Standalone()

		workmail_updateMobileDeviceAccessRuleCmd.Flags().String("description", "", "The updated rule description.")
		workmail_updateMobileDeviceAccessRuleCmd.Flags().String("device-models", "", "Device models that the updated rule will match.")
		workmail_updateMobileDeviceAccessRuleCmd.Flags().String("device-operating-systems", "", "Device operating systems that the updated rule will match.")
		workmail_updateMobileDeviceAccessRuleCmd.Flags().String("device-types", "", "Device types that the updated rule will match.")
		workmail_updateMobileDeviceAccessRuleCmd.Flags().String("device-user-agents", "", "User agents that the updated rule will match.")
		workmail_updateMobileDeviceAccessRuleCmd.Flags().String("effect", "", "The effect of the rule when it matches.")
		workmail_updateMobileDeviceAccessRuleCmd.Flags().String("mobile-device-access-rule-id", "", "The identifier of the rule to be updated.")
		workmail_updateMobileDeviceAccessRuleCmd.Flags().String("name", "", "The updated rule name.")
		workmail_updateMobileDeviceAccessRuleCmd.Flags().String("not-device-models", "", "Device models that the updated rule **will not** match.")
		workmail_updateMobileDeviceAccessRuleCmd.Flags().String("not-device-operating-systems", "", "Device operating systems that the updated rule **will not** match.")
		workmail_updateMobileDeviceAccessRuleCmd.Flags().String("not-device-types", "", "Device types that the updated rule **will not** match.")
		workmail_updateMobileDeviceAccessRuleCmd.Flags().String("not-device-user-agents", "", "User agents that the updated rule **will not** match.")
		workmail_updateMobileDeviceAccessRuleCmd.Flags().String("organization-id", "", "The WorkMail organization under which the rule will be updated.")
		workmail_updateMobileDeviceAccessRuleCmd.MarkFlagRequired("effect")
		workmail_updateMobileDeviceAccessRuleCmd.MarkFlagRequired("mobile-device-access-rule-id")
		workmail_updateMobileDeviceAccessRuleCmd.MarkFlagRequired("name")
		workmail_updateMobileDeviceAccessRuleCmd.MarkFlagRequired("organization-id")
	})
	workmailCmd.AddCommand(workmail_updateMobileDeviceAccessRuleCmd)
}
