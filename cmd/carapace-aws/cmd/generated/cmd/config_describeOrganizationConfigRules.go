package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_describeOrganizationConfigRulesCmd = &cobra.Command{
	Use:   "describe-organization-config-rules",
	Short: "Returns a list of organization Config rules.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_describeOrganizationConfigRulesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_describeOrganizationConfigRulesCmd).Standalone()

		config_describeOrganizationConfigRulesCmd.Flags().String("limit", "", "The maximum number of organization Config rules returned on each page.")
		config_describeOrganizationConfigRulesCmd.Flags().String("next-token", "", "The `nextToken` string returned on a previous page that you use to get the next page of results in a paginated response.")
		config_describeOrganizationConfigRulesCmd.Flags().String("organization-config-rule-names", "", "The names of organization Config rules for which you want details.")
	})
	configCmd.AddCommand(config_describeOrganizationConfigRulesCmd)
}
