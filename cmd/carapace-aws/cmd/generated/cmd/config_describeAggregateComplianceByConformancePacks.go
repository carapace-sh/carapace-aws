package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_describeAggregateComplianceByConformancePacksCmd = &cobra.Command{
	Use:   "describe-aggregate-compliance-by-conformance-packs",
	Short: "Returns a list of the existing and deleted conformance packs and their associated compliance status with the count of compliant and noncompliant Config rules within each conformance pack.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_describeAggregateComplianceByConformancePacksCmd).Standalone()

	config_describeAggregateComplianceByConformancePacksCmd.Flags().String("configuration-aggregator-name", "", "The name of the configuration aggregator.")
	config_describeAggregateComplianceByConformancePacksCmd.Flags().String("filters", "", "Filters the result by `AggregateConformancePackComplianceFilters` object.")
	config_describeAggregateComplianceByConformancePacksCmd.Flags().String("limit", "", "The maximum number of conformance packs compliance details returned on each page.")
	config_describeAggregateComplianceByConformancePacksCmd.Flags().String("next-token", "", "The `nextToken` string returned on a previous page that you use to get the next page of results in a paginated response.")
	config_describeAggregateComplianceByConformancePacksCmd.MarkFlagRequired("configuration-aggregator-name")
	configCmd.AddCommand(config_describeAggregateComplianceByConformancePacksCmd)
}
