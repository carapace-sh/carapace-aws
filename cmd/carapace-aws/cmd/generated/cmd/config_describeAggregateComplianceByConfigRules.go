package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_describeAggregateComplianceByConfigRulesCmd = &cobra.Command{
	Use:   "describe-aggregate-compliance-by-config-rules",
	Short: "Returns a list of compliant and noncompliant rules with the number of resources for compliant and noncompliant rules.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_describeAggregateComplianceByConfigRulesCmd).Standalone()

	config_describeAggregateComplianceByConfigRulesCmd.Flags().String("configuration-aggregator-name", "", "The name of the configuration aggregator.")
	config_describeAggregateComplianceByConfigRulesCmd.Flags().String("filters", "", "Filters the results by ConfigRuleComplianceFilters object.")
	config_describeAggregateComplianceByConfigRulesCmd.Flags().String("limit", "", "The maximum number of evaluation results returned on each page.")
	config_describeAggregateComplianceByConfigRulesCmd.Flags().String("next-token", "", "The `nextToken` string returned on a previous page that you use to get the next page of results in a paginated response.")
	config_describeAggregateComplianceByConfigRulesCmd.MarkFlagRequired("configuration-aggregator-name")
	configCmd.AddCommand(config_describeAggregateComplianceByConfigRulesCmd)
}
