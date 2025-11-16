package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_getCisScanResultDetailsCmd = &cobra.Command{
	Use:   "get-cis-scan-result-details",
	Short: "Retrieves CIS scan result details.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_getCisScanResultDetailsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector2_getCisScanResultDetailsCmd).Standalone()

		inspector2_getCisScanResultDetailsCmd.Flags().String("account-id", "", "The account ID.")
		inspector2_getCisScanResultDetailsCmd.Flags().String("filter-criteria", "", "The filter criteria.")
		inspector2_getCisScanResultDetailsCmd.Flags().String("max-results", "", "The maximum number of CIS scan result details to be returned in a single page of results.")
		inspector2_getCisScanResultDetailsCmd.Flags().String("next-token", "", "The pagination token from a previous request that's used to retrieve the next page of results.")
		inspector2_getCisScanResultDetailsCmd.Flags().String("scan-arn", "", "The scan ARN.")
		inspector2_getCisScanResultDetailsCmd.Flags().String("sort-by", "", "The sort by order.")
		inspector2_getCisScanResultDetailsCmd.Flags().String("sort-order", "", "The sort order.")
		inspector2_getCisScanResultDetailsCmd.Flags().String("target-resource-id", "", "The target resource ID.")
		inspector2_getCisScanResultDetailsCmd.MarkFlagRequired("account-id")
		inspector2_getCisScanResultDetailsCmd.MarkFlagRequired("scan-arn")
		inspector2_getCisScanResultDetailsCmd.MarkFlagRequired("target-resource-id")
	})
	inspector2Cmd.AddCommand(inspector2_getCisScanResultDetailsCmd)
}
