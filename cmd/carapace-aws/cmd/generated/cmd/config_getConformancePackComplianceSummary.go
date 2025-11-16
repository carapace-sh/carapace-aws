package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_getConformancePackComplianceSummaryCmd = &cobra.Command{
	Use:   "get-conformance-pack-compliance-summary",
	Short: "Returns compliance details for the conformance pack based on the cumulative compliance results of all the rules in that conformance pack.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_getConformancePackComplianceSummaryCmd).Standalone()

	config_getConformancePackComplianceSummaryCmd.Flags().String("conformance-pack-names", "", "Names of conformance packs.")
	config_getConformancePackComplianceSummaryCmd.Flags().String("limit", "", "The maximum number of conformance packs returned on each page.")
	config_getConformancePackComplianceSummaryCmd.Flags().String("next-token", "", "The nextToken string returned on a previous page that you use to get the next page of results in a paginated response.")
	config_getConformancePackComplianceSummaryCmd.MarkFlagRequired("conformance-pack-names")
	configCmd.AddCommand(config_getConformancePackComplianceSummaryCmd)
}
