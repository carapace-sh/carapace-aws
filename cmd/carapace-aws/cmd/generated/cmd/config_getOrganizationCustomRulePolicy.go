package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_getOrganizationCustomRulePolicyCmd = &cobra.Command{
	Use:   "get-organization-custom-rule-policy",
	Short: "Returns the policy definition containing the logic for your organization Config Custom Policy rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_getOrganizationCustomRulePolicyCmd).Standalone()

	config_getOrganizationCustomRulePolicyCmd.Flags().String("organization-config-rule-name", "", "The name of your organization Config Custom Policy rule.")
	config_getOrganizationCustomRulePolicyCmd.MarkFlagRequired("organization-config-rule-name")
	configCmd.AddCommand(config_getOrganizationCustomRulePolicyCmd)
}
