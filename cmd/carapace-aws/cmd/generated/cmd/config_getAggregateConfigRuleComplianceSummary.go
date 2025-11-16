package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_getAggregateConfigRuleComplianceSummaryCmd = &cobra.Command{
	Use:   "get-aggregate-config-rule-compliance-summary",
	Short: "Returns the number of compliant and noncompliant rules for one or more accounts and regions in an aggregator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_getAggregateConfigRuleComplianceSummaryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_getAggregateConfigRuleComplianceSummaryCmd).Standalone()

		config_getAggregateConfigRuleComplianceSummaryCmd.Flags().String("configuration-aggregator-name", "", "The name of the configuration aggregator.")
		config_getAggregateConfigRuleComplianceSummaryCmd.Flags().String("filters", "", "Filters the results based on the ConfigRuleComplianceSummaryFilters object.")
		config_getAggregateConfigRuleComplianceSummaryCmd.Flags().String("group-by-key", "", "Groups the result based on ACCOUNT\\_ID or AWS\\_REGION.")
		config_getAggregateConfigRuleComplianceSummaryCmd.Flags().String("limit", "", "The maximum number of evaluation results returned on each page.")
		config_getAggregateConfigRuleComplianceSummaryCmd.Flags().String("next-token", "", "The `nextToken` string returned on a previous page that you use to get the next page of results in a paginated response.")
		config_getAggregateConfigRuleComplianceSummaryCmd.MarkFlagRequired("configuration-aggregator-name")
	})
	configCmd.AddCommand(config_getAggregateConfigRuleComplianceSummaryCmd)
}
