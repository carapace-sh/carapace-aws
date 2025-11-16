package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_deleteMobileDeviceAccessRuleCmd = &cobra.Command{
	Use:   "delete-mobile-device-access-rule",
	Short: "Deletes a mobile device access rule for the specified WorkMail organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_deleteMobileDeviceAccessRuleCmd).Standalone()

	workmail_deleteMobileDeviceAccessRuleCmd.Flags().String("mobile-device-access-rule-id", "", "The identifier of the rule to be deleted.")
	workmail_deleteMobileDeviceAccessRuleCmd.Flags().String("organization-id", "", "The WorkMail organization under which the rule will be deleted.")
	workmail_deleteMobileDeviceAccessRuleCmd.MarkFlagRequired("mobile-device-access-rule-id")
	workmail_deleteMobileDeviceAccessRuleCmd.MarkFlagRequired("organization-id")
	workmailCmd.AddCommand(workmail_deleteMobileDeviceAccessRuleCmd)
}
