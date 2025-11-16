package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getIpamDiscoveredResourceCidrsCmd = &cobra.Command{
	Use:   "get-ipam-discovered-resource-cidrs",
	Short: "Returns the resource CIDRs that are monitored as part of a resource discovery.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getIpamDiscoveredResourceCidrsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_getIpamDiscoveredResourceCidrsCmd).Standalone()

		ec2_getIpamDiscoveredResourceCidrsCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_getIpamDiscoveredResourceCidrsCmd.Flags().String("filters", "", "Filters.")
		ec2_getIpamDiscoveredResourceCidrsCmd.Flags().String("ipam-resource-discovery-id", "", "A resource discovery ID.")
		ec2_getIpamDiscoveredResourceCidrsCmd.Flags().String("max-results", "", "The maximum number of discovered resource CIDRs to return in one page of results.")
		ec2_getIpamDiscoveredResourceCidrsCmd.Flags().String("next-token", "", "Specify the pagination token from a previous request to retrieve the next page of results.")
		ec2_getIpamDiscoveredResourceCidrsCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_getIpamDiscoveredResourceCidrsCmd.Flags().String("resource-region", "", "A resource Region.")
		ec2_getIpamDiscoveredResourceCidrsCmd.MarkFlagRequired("ipam-resource-discovery-id")
		ec2_getIpamDiscoveredResourceCidrsCmd.Flag("no-dry-run").Hidden = true
		ec2_getIpamDiscoveredResourceCidrsCmd.MarkFlagRequired("resource-region")
	})
	ec2Cmd.AddCommand(ec2_getIpamDiscoveredResourceCidrsCmd)
}
