package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getIpamDiscoveredPublicAddressesCmd = &cobra.Command{
	Use:   "get-ipam-discovered-public-addresses",
	Short: "Gets the public IP addresses that have been discovered by IPAM.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getIpamDiscoveredPublicAddressesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_getIpamDiscoveredPublicAddressesCmd).Standalone()

		ec2_getIpamDiscoveredPublicAddressesCmd.Flags().String("address-region", "", "The Amazon Web Services Region for the IP address.")
		ec2_getIpamDiscoveredPublicAddressesCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_getIpamDiscoveredPublicAddressesCmd.Flags().String("filters", "", "Filters.")
		ec2_getIpamDiscoveredPublicAddressesCmd.Flags().String("ipam-resource-discovery-id", "", "An IPAM resource discovery ID.")
		ec2_getIpamDiscoveredPublicAddressesCmd.Flags().String("max-results", "", "The maximum number of IPAM discovered public addresses to return in one page of results.")
		ec2_getIpamDiscoveredPublicAddressesCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ec2_getIpamDiscoveredPublicAddressesCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_getIpamDiscoveredPublicAddressesCmd.MarkFlagRequired("address-region")
		ec2_getIpamDiscoveredPublicAddressesCmd.MarkFlagRequired("ipam-resource-discovery-id")
		ec2_getIpamDiscoveredPublicAddressesCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_getIpamDiscoveredPublicAddressesCmd)
}
