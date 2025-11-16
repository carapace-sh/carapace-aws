package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_getConformancePackComplianceDetailsCmd = &cobra.Command{
	Use:   "get-conformance-pack-compliance-details",
	Short: "Returns compliance details of a conformance pack for all Amazon Web Services resources that are monitered by conformance pack.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_getConformancePackComplianceDetailsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_getConformancePackComplianceDetailsCmd).Standalone()

		config_getConformancePackComplianceDetailsCmd.Flags().String("conformance-pack-name", "", "Name of the conformance pack.")
		config_getConformancePackComplianceDetailsCmd.Flags().String("filters", "", "A `ConformancePackEvaluationFilters` object.")
		config_getConformancePackComplianceDetailsCmd.Flags().String("limit", "", "The maximum number of evaluation results returned on each page.")
		config_getConformancePackComplianceDetailsCmd.Flags().String("next-token", "", "The `nextToken` string returned in a previous request that you use to request the next page of results in a paginated response.")
		config_getConformancePackComplianceDetailsCmd.MarkFlagRequired("conformance-pack-name")
	})
	configCmd.AddCommand(config_getConformancePackComplianceDetailsCmd)
}
