package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_listResourceScanRelatedResourcesCmd = &cobra.Command{
	Use:   "list-resource-scan-related-resources",
	Short: "Lists the related resources for a list of resources from a resource scan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_listResourceScanRelatedResourcesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_listResourceScanRelatedResourcesCmd).Standalone()

		cloudformation_listResourceScanRelatedResourcesCmd.Flags().String("max-results", "", "If the number of available results exceeds this maximum, the response includes a `NextToken` value that you can use for the `NextToken` parameter to get the next set of results.")
		cloudformation_listResourceScanRelatedResourcesCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
		cloudformation_listResourceScanRelatedResourcesCmd.Flags().String("resource-scan-id", "", "The Amazon Resource Name (ARN) of the resource scan.")
		cloudformation_listResourceScanRelatedResourcesCmd.Flags().String("resources", "", "The list of resources for which you want to get the related resources.")
		cloudformation_listResourceScanRelatedResourcesCmd.MarkFlagRequired("resource-scan-id")
		cloudformation_listResourceScanRelatedResourcesCmd.MarkFlagRequired("resources")
	})
	cloudformationCmd.AddCommand(cloudformation_listResourceScanRelatedResourcesCmd)
}
