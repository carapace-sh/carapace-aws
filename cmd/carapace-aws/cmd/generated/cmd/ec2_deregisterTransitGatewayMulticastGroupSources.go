package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deregisterTransitGatewayMulticastGroupSourcesCmd = &cobra.Command{
	Use:   "deregister-transit-gateway-multicast-group-sources",
	Short: "Deregisters the specified sources (network interfaces) from the transit gateway multicast group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deregisterTransitGatewayMulticastGroupSourcesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deregisterTransitGatewayMulticastGroupSourcesCmd).Standalone()

		ec2_deregisterTransitGatewayMulticastGroupSourcesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deregisterTransitGatewayMulticastGroupSourcesCmd.Flags().String("group-ip-address", "", "The IP address assigned to the transit gateway multicast group.")
		ec2_deregisterTransitGatewayMulticastGroupSourcesCmd.Flags().String("network-interface-ids", "", "The IDs of the group sources' network interfaces.")
		ec2_deregisterTransitGatewayMulticastGroupSourcesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deregisterTransitGatewayMulticastGroupSourcesCmd.Flags().String("transit-gateway-multicast-domain-id", "", "The ID of the transit gateway multicast domain.")
		ec2_deregisterTransitGatewayMulticastGroupSourcesCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_deregisterTransitGatewayMulticastGroupSourcesCmd)
}
