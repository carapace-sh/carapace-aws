package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_listCisScanResultsAggregatedByChecksCmd = &cobra.Command{
	Use:   "list-cis-scan-results-aggregated-by-checks",
	Short: "Lists scan results aggregated by checks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_listCisScanResultsAggregatedByChecksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector2_listCisScanResultsAggregatedByChecksCmd).Standalone()

		inspector2_listCisScanResultsAggregatedByChecksCmd.Flags().String("filter-criteria", "", "The filter criteria.")
		inspector2_listCisScanResultsAggregatedByChecksCmd.Flags().String("max-results", "", "The maximum number of scan results aggregated by checks to be returned in a single page of results.")
		inspector2_listCisScanResultsAggregatedByChecksCmd.Flags().String("next-token", "", "The pagination token from a previous request that's used to retrieve the next page of results.")
		inspector2_listCisScanResultsAggregatedByChecksCmd.Flags().String("scan-arn", "", "The scan ARN.")
		inspector2_listCisScanResultsAggregatedByChecksCmd.Flags().String("sort-by", "", "The sort by order.")
		inspector2_listCisScanResultsAggregatedByChecksCmd.Flags().String("sort-order", "", "The sort order.")
		inspector2_listCisScanResultsAggregatedByChecksCmd.MarkFlagRequired("scan-arn")
	})
	inspector2Cmd.AddCommand(inspector2_listCisScanResultsAggregatedByChecksCmd)
}
