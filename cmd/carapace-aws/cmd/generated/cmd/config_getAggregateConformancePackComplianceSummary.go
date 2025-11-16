package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_getAggregateConformancePackComplianceSummaryCmd = &cobra.Command{
	Use:   "get-aggregate-conformance-pack-compliance-summary",
	Short: "Returns the count of compliant and noncompliant conformance packs across all Amazon Web Services accounts and Amazon Web Services Regions in an aggregator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_getAggregateConformancePackComplianceSummaryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_getAggregateConformancePackComplianceSummaryCmd).Standalone()

		config_getAggregateConformancePackComplianceSummaryCmd.Flags().String("configuration-aggregator-name", "", "The name of the configuration aggregator.")
		config_getAggregateConformancePackComplianceSummaryCmd.Flags().String("filters", "", "Filters the results based on the `AggregateConformancePackComplianceSummaryFilters` object.")
		config_getAggregateConformancePackComplianceSummaryCmd.Flags().String("group-by-key", "", "Groups the result based on Amazon Web Services account ID or Amazon Web Services Region.")
		config_getAggregateConformancePackComplianceSummaryCmd.Flags().String("limit", "", "The maximum number of results returned on each page.")
		config_getAggregateConformancePackComplianceSummaryCmd.Flags().String("next-token", "", "The `nextToken` string returned on a previous page that you use to get the next page of results in a paginated response.")
		config_getAggregateConformancePackComplianceSummaryCmd.MarkFlagRequired("configuration-aggregator-name")
	})
	configCmd.AddCommand(config_getAggregateConformancePackComplianceSummaryCmd)
}
