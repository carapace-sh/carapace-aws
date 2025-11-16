package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_createMobileDeviceAccessRuleCmd = &cobra.Command{
	Use:   "create-mobile-device-access-rule",
	Short: "Creates a new mobile device access rule for the specified WorkMail organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_createMobileDeviceAccessRuleCmd).Standalone()

	workmail_createMobileDeviceAccessRuleCmd.Flags().String("client-token", "", "The idempotency token for the client request.")
	workmail_createMobileDeviceAccessRuleCmd.Flags().String("description", "", "The rule description.")
	workmail_createMobileDeviceAccessRuleCmd.Flags().String("device-models", "", "Device models that the rule will match.")
	workmail_createMobileDeviceAccessRuleCmd.Flags().String("device-operating-systems", "", "Device operating systems that the rule will match.")
	workmail_createMobileDeviceAccessRuleCmd.Flags().String("device-types", "", "Device types that the rule will match.")
	workmail_createMobileDeviceAccessRuleCmd.Flags().String("device-user-agents", "", "Device user agents that the rule will match.")
	workmail_createMobileDeviceAccessRuleCmd.Flags().String("effect", "", "The effect of the rule when it matches.")
	workmail_createMobileDeviceAccessRuleCmd.Flags().String("name", "", "The rule name.")
	workmail_createMobileDeviceAccessRuleCmd.Flags().String("not-device-models", "", "Device models that the rule **will not** match.")
	workmail_createMobileDeviceAccessRuleCmd.Flags().String("not-device-operating-systems", "", "Device operating systems that the rule **will not** match.")
	workmail_createMobileDeviceAccessRuleCmd.Flags().String("not-device-types", "", "Device types that the rule **will not** match.")
	workmail_createMobileDeviceAccessRuleCmd.Flags().String("not-device-user-agents", "", "Device user agents that the rule **will not** match.")
	workmail_createMobileDeviceAccessRuleCmd.Flags().String("organization-id", "", "The WorkMail organization under which the rule will be created.")
	workmail_createMobileDeviceAccessRuleCmd.MarkFlagRequired("effect")
	workmail_createMobileDeviceAccessRuleCmd.MarkFlagRequired("name")
	workmail_createMobileDeviceAccessRuleCmd.MarkFlagRequired("organization-id")
	workmailCmd.AddCommand(workmail_createMobileDeviceAccessRuleCmd)
}
