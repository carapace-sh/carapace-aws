package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_putOrganizationConfigRuleCmd = &cobra.Command{
	Use:   "put-organization-config-rule",
	Short: "Adds or updates an Config rule for your entire organization to evaluate if your Amazon Web Services resources comply with your desired configurations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_putOrganizationConfigRuleCmd).Standalone()

	config_putOrganizationConfigRuleCmd.Flags().String("excluded-accounts", "", "A comma-separated list of accounts that you want to exclude from an organization Config rule.")
	config_putOrganizationConfigRuleCmd.Flags().String("organization-config-rule-name", "", "The name that you assign to an organization Config rule.")
	config_putOrganizationConfigRuleCmd.Flags().String("organization-custom-policy-rule-metadata", "", "An `OrganizationCustomPolicyRuleMetadata` object.")
	config_putOrganizationConfigRuleCmd.Flags().String("organization-custom-rule-metadata", "", "An `OrganizationCustomRuleMetadata` object.")
	config_putOrganizationConfigRuleCmd.Flags().String("organization-managed-rule-metadata", "", "An `OrganizationManagedRuleMetadata` object.")
	config_putOrganizationConfigRuleCmd.MarkFlagRequired("organization-config-rule-name")
	configCmd.AddCommand(config_putOrganizationConfigRuleCmd)
}
