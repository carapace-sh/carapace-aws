package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_registerTransitGatewayMulticastGroupSourcesCmd = &cobra.Command{
	Use:   "register-transit-gateway-multicast-group-sources",
	Short: "Registers sources (network interfaces) with the specified transit gateway multicast group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_registerTransitGatewayMulticastGroupSourcesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_registerTransitGatewayMulticastGroupSourcesCmd).Standalone()

		ec2_registerTransitGatewayMulticastGroupSourcesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_registerTransitGatewayMulticastGroupSourcesCmd.Flags().String("group-ip-address", "", "The IP address assigned to the transit gateway multicast group.")
		ec2_registerTransitGatewayMulticastGroupSourcesCmd.Flags().String("network-interface-ids", "", "The group sources' network interface IDs to register with the transit gateway multicast group.")
		ec2_registerTransitGatewayMulticastGroupSourcesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_registerTransitGatewayMulticastGroupSourcesCmd.Flags().String("transit-gateway-multicast-domain-id", "", "The ID of the transit gateway multicast domain.")
		ec2_registerTransitGatewayMulticastGroupSourcesCmd.MarkFlagRequired("network-interface-ids")
		ec2_registerTransitGatewayMulticastGroupSourcesCmd.Flag("no-dry-run").Hidden = true
		ec2_registerTransitGatewayMulticastGroupSourcesCmd.MarkFlagRequired("transit-gateway-multicast-domain-id")
	})
	ec2Cmd.AddCommand(ec2_registerTransitGatewayMulticastGroupSourcesCmd)
}
