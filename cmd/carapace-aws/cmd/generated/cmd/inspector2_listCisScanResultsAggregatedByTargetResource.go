package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_listCisScanResultsAggregatedByTargetResourceCmd = &cobra.Command{
	Use:   "list-cis-scan-results-aggregated-by-target-resource",
	Short: "Lists scan results aggregated by a target resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_listCisScanResultsAggregatedByTargetResourceCmd).Standalone()

	inspector2_listCisScanResultsAggregatedByTargetResourceCmd.Flags().String("filter-criteria", "", "The filter criteria.")
	inspector2_listCisScanResultsAggregatedByTargetResourceCmd.Flags().String("max-results", "", "The maximum number of scan results aggregated by a target resource to be returned in a single page of results.")
	inspector2_listCisScanResultsAggregatedByTargetResourceCmd.Flags().String("next-token", "", "The pagination token from a previous request that's used to retrieve the next page of results.")
	inspector2_listCisScanResultsAggregatedByTargetResourceCmd.Flags().String("scan-arn", "", "The scan ARN.")
	inspector2_listCisScanResultsAggregatedByTargetResourceCmd.Flags().String("sort-by", "", "The sort by order.")
	inspector2_listCisScanResultsAggregatedByTargetResourceCmd.Flags().String("sort-order", "", "The sort order.")
	inspector2_listCisScanResultsAggregatedByTargetResourceCmd.MarkFlagRequired("scan-arn")
	inspector2Cmd.AddCommand(inspector2_listCisScanResultsAggregatedByTargetResourceCmd)
}
