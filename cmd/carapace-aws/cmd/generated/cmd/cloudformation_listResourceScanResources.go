package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_listResourceScanResourcesCmd = &cobra.Command{
	Use:   "list-resource-scan-resources",
	Short: "Lists the resources from a resource scan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_listResourceScanResourcesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_listResourceScanResourcesCmd).Standalone()

		cloudformation_listResourceScanResourcesCmd.Flags().String("max-results", "", "If the number of available results exceeds this maximum, the response includes a `NextToken` value that you can use for the `NextToken` parameter to get the next set of results.")
		cloudformation_listResourceScanResourcesCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
		cloudformation_listResourceScanResourcesCmd.Flags().String("resource-identifier", "", "If specified, the returned resources will have the specified resource identifier (or one of them in the case where the resource has multiple identifiers).")
		cloudformation_listResourceScanResourcesCmd.Flags().String("resource-scan-id", "", "The Amazon Resource Name (ARN) of the resource scan.")
		cloudformation_listResourceScanResourcesCmd.Flags().String("resource-type-prefix", "", "If specified, the returned resources will be of any of the resource types with the specified prefix.")
		cloudformation_listResourceScanResourcesCmd.Flags().String("tag-key", "", "If specified, the returned resources will have a matching tag key.")
		cloudformation_listResourceScanResourcesCmd.Flags().String("tag-value", "", "If specified, the returned resources will have a matching tag value.")
		cloudformation_listResourceScanResourcesCmd.MarkFlagRequired("resource-scan-id")
	})
	cloudformationCmd.AddCommand(cloudformation_listResourceScanResourcesCmd)
}
