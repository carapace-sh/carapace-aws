package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_describeConformancePackComplianceCmd = &cobra.Command{
	Use:   "describe-conformance-pack-compliance",
	Short: "Returns compliance details for each rule in that conformance pack.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_describeConformancePackComplianceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_describeConformancePackComplianceCmd).Standalone()

		config_describeConformancePackComplianceCmd.Flags().String("conformance-pack-name", "", "Name of the conformance pack.")
		config_describeConformancePackComplianceCmd.Flags().String("filters", "", "A `ConformancePackComplianceFilters` object.")
		config_describeConformancePackComplianceCmd.Flags().String("limit", "", "The maximum number of Config rules within a conformance pack are returned on each page.")
		config_describeConformancePackComplianceCmd.Flags().String("next-token", "", "The `nextToken` string returned in a previous request that you use to request the next page of results in a paginated response.")
		config_describeConformancePackComplianceCmd.MarkFlagRequired("conformance-pack-name")
	})
	configCmd.AddCommand(config_describeConformancePackComplianceCmd)
}
