package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourcegroupstaggingapi_getComplianceSummaryCmd = &cobra.Command{
	Use:   "get-compliance-summary",
	Short: "Returns a table that shows counts of resources that are noncompliant with their tag policies.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourcegroupstaggingapi_getComplianceSummaryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resourcegroupstaggingapi_getComplianceSummaryCmd).Standalone()

		resourcegroupstaggingapi_getComplianceSummaryCmd.Flags().String("group-by", "", "Specifies a list of attributes to group the counts of noncompliant resources by.")
		resourcegroupstaggingapi_getComplianceSummaryCmd.Flags().String("max-results", "", "Specifies the maximum number of results to be returned in each page.")
		resourcegroupstaggingapi_getComplianceSummaryCmd.Flags().String("pagination-token", "", "Specifies a `PaginationToken` response value from a previous request to indicate that you want the next page of results.")
		resourcegroupstaggingapi_getComplianceSummaryCmd.Flags().String("region-filters", "", "Specifies a list of Amazon Web Services Regions to limit the output to.")
		resourcegroupstaggingapi_getComplianceSummaryCmd.Flags().String("resource-type-filters", "", "Specifies that you want the response to include information for only resources of the specified types.")
		resourcegroupstaggingapi_getComplianceSummaryCmd.Flags().String("tag-key-filters", "", "Specifies that you want the response to include information for only resources that have tags with the specified tag keys.")
		resourcegroupstaggingapi_getComplianceSummaryCmd.Flags().String("target-id-filters", "", "Specifies target identifiers (usually, specific account IDs) to limit the output by.")
	})
	resourcegroupstaggingapiCmd.AddCommand(resourcegroupstaggingapi_getComplianceSummaryCmd)
}
