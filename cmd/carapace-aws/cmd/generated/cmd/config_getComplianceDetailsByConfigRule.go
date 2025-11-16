package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_getComplianceDetailsByConfigRuleCmd = &cobra.Command{
	Use:   "get-compliance-details-by-config-rule",
	Short: "Returns the evaluation results for the specified Config rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_getComplianceDetailsByConfigRuleCmd).Standalone()

	config_getComplianceDetailsByConfigRuleCmd.Flags().String("compliance-types", "", "Filters the results by compliance.")
	config_getComplianceDetailsByConfigRuleCmd.Flags().String("config-rule-name", "", "The name of the Config rule for which you want compliance information.")
	config_getComplianceDetailsByConfigRuleCmd.Flags().String("limit", "", "The maximum number of evaluation results returned on each page.")
	config_getComplianceDetailsByConfigRuleCmd.Flags().String("next-token", "", "The `nextToken` string returned on a previous page that you use to get the next page of results in a paginated response.")
	config_getComplianceDetailsByConfigRuleCmd.MarkFlagRequired("config-rule-name")
	configCmd.AddCommand(config_getComplianceDetailsByConfigRuleCmd)
}
