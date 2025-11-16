package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_getOrganizationConfigRuleDetailedStatusCmd = &cobra.Command{
	Use:   "get-organization-config-rule-detailed-status",
	Short: "Returns detailed status for each member account within an organization for a given organization Config rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_getOrganizationConfigRuleDetailedStatusCmd).Standalone()

	config_getOrganizationConfigRuleDetailedStatusCmd.Flags().String("filters", "", "A `StatusDetailFilters` object.")
	config_getOrganizationConfigRuleDetailedStatusCmd.Flags().String("limit", "", "The maximum number of `OrganizationConfigRuleDetailedStatus` returned on each page.")
	config_getOrganizationConfigRuleDetailedStatusCmd.Flags().String("next-token", "", "The `nextToken` string returned on a previous page that you use to get the next page of results in a paginated response.")
	config_getOrganizationConfigRuleDetailedStatusCmd.Flags().String("organization-config-rule-name", "", "The name of your organization Config rule for which you want status details for member accounts.")
	config_getOrganizationConfigRuleDetailedStatusCmd.MarkFlagRequired("organization-config-rule-name")
	configCmd.AddCommand(config_getOrganizationConfigRuleDetailedStatusCmd)
}
