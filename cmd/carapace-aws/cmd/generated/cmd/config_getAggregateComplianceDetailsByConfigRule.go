package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_getAggregateComplianceDetailsByConfigRuleCmd = &cobra.Command{
	Use:   "get-aggregate-compliance-details-by-config-rule",
	Short: "Returns the evaluation results for the specified Config rule for a specific resource in a rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_getAggregateComplianceDetailsByConfigRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_getAggregateComplianceDetailsByConfigRuleCmd).Standalone()

		config_getAggregateComplianceDetailsByConfigRuleCmd.Flags().String("account-id", "", "The 12-digit account ID of the source account.")
		config_getAggregateComplianceDetailsByConfigRuleCmd.Flags().String("aws-region", "", "The source region from where the data is aggregated.")
		config_getAggregateComplianceDetailsByConfigRuleCmd.Flags().String("compliance-type", "", "The resource compliance status.")
		config_getAggregateComplianceDetailsByConfigRuleCmd.Flags().String("config-rule-name", "", "The name of the Config rule for which you want compliance information.")
		config_getAggregateComplianceDetailsByConfigRuleCmd.Flags().String("configuration-aggregator-name", "", "The name of the configuration aggregator.")
		config_getAggregateComplianceDetailsByConfigRuleCmd.Flags().String("limit", "", "The maximum number of evaluation results returned on each page.")
		config_getAggregateComplianceDetailsByConfigRuleCmd.Flags().String("next-token", "", "The `nextToken` string returned on a previous page that you use to get the next page of results in a paginated response.")
		config_getAggregateComplianceDetailsByConfigRuleCmd.MarkFlagRequired("account-id")
		config_getAggregateComplianceDetailsByConfigRuleCmd.MarkFlagRequired("aws-region")
		config_getAggregateComplianceDetailsByConfigRuleCmd.MarkFlagRequired("config-rule-name")
		config_getAggregateComplianceDetailsByConfigRuleCmd.MarkFlagRequired("configuration-aggregator-name")
	})
	configCmd.AddCommand(config_getAggregateComplianceDetailsByConfigRuleCmd)
}
