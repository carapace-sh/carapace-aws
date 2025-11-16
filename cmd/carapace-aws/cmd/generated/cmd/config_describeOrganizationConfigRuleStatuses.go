package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_describeOrganizationConfigRuleStatusesCmd = &cobra.Command{
	Use:   "describe-organization-config-rule-statuses",
	Short: "Provides organization Config rule deployment status for an organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_describeOrganizationConfigRuleStatusesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_describeOrganizationConfigRuleStatusesCmd).Standalone()

		config_describeOrganizationConfigRuleStatusesCmd.Flags().String("limit", "", "The maximum number of `OrganizationConfigRuleStatuses` returned on each page.")
		config_describeOrganizationConfigRuleStatusesCmd.Flags().String("next-token", "", "The `nextToken` string returned on a previous page that you use to get the next page of results in a paginated response.")
		config_describeOrganizationConfigRuleStatusesCmd.Flags().String("organization-config-rule-names", "", "The names of organization Config rules for which you want status details.")
	})
	configCmd.AddCommand(config_describeOrganizationConfigRuleStatusesCmd)
}
