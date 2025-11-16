package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fms_listDiscoveredResourcesCmd = &cobra.Command{
	Use:   "list-discovered-resources",
	Short: "Returns an array of resources in the organization's accounts that are available to be associated with a resource set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fms_listDiscoveredResourcesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fms_listDiscoveredResourcesCmd).Standalone()

		fms_listDiscoveredResourcesCmd.Flags().String("max-results", "", "The maximum number of objects that you want Firewall Manager to return for this request.")
		fms_listDiscoveredResourcesCmd.Flags().String("member-account-ids", "", "The Amazon Web Services account IDs to discover resources in.")
		fms_listDiscoveredResourcesCmd.Flags().String("next-token", "", "When you request a list of objects with a `MaxResults` setting, if the number of objects that are still available for retrieval exceeds the maximum you requested, Firewall Manager returns a `NextToken` value in the response.")
		fms_listDiscoveredResourcesCmd.Flags().String("resource-type", "", "The type of resources to discover.")
		fms_listDiscoveredResourcesCmd.MarkFlagRequired("member-account-ids")
		fms_listDiscoveredResourcesCmd.MarkFlagRequired("resource-type")
	})
	fmsCmd.AddCommand(fms_listDiscoveredResourcesCmd)
}
