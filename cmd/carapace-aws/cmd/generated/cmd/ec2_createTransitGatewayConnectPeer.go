package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createTransitGatewayConnectPeerCmd = &cobra.Command{
	Use:   "create-transit-gateway-connect-peer",
	Short: "Creates a Connect peer for a specified transit gateway Connect attachment between a transit gateway and an appliance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createTransitGatewayConnectPeerCmd).Standalone()

	ec2_createTransitGatewayConnectPeerCmd.Flags().String("bgp-options", "", "The BGP options for the Connect peer.")
	ec2_createTransitGatewayConnectPeerCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createTransitGatewayConnectPeerCmd.Flags().String("inside-cidr-blocks", "", "The range of inside IP addresses that are used for BGP peering.")
	ec2_createTransitGatewayConnectPeerCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createTransitGatewayConnectPeerCmd.Flags().String("peer-address", "", "The peer IP address (GRE outer IP address) on the appliance side of the Connect peer.")
	ec2_createTransitGatewayConnectPeerCmd.Flags().String("tag-specifications", "", "The tags to apply to the Connect peer.")
	ec2_createTransitGatewayConnectPeerCmd.Flags().String("transit-gateway-address", "", "The peer IP address (GRE outer IP address) on the transit gateway side of the Connect peer, which must be specified from a transit gateway CIDR block.")
	ec2_createTransitGatewayConnectPeerCmd.Flags().String("transit-gateway-attachment-id", "", "The ID of the Connect attachment.")
	ec2_createTransitGatewayConnectPeerCmd.MarkFlagRequired("inside-cidr-blocks")
	ec2_createTransitGatewayConnectPeerCmd.Flag("no-dry-run").Hidden = true
	ec2_createTransitGatewayConnectPeerCmd.MarkFlagRequired("peer-address")
	ec2_createTransitGatewayConnectPeerCmd.MarkFlagRequired("transit-gateway-attachment-id")
	ec2Cmd.AddCommand(ec2_createTransitGatewayConnectPeerCmd)
}
