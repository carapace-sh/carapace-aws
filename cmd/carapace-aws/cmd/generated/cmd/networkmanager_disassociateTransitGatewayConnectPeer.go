package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_disassociateTransitGatewayConnectPeerCmd = &cobra.Command{
	Use:   "disassociate-transit-gateway-connect-peer",
	Short: "Disassociates a transit gateway Connect peer from a device and link.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_disassociateTransitGatewayConnectPeerCmd).Standalone()

	networkmanager_disassociateTransitGatewayConnectPeerCmd.Flags().String("global-network-id", "", "The ID of the global network.")
	networkmanager_disassociateTransitGatewayConnectPeerCmd.Flags().String("transit-gateway-connect-peer-arn", "", "The Amazon Resource Name (ARN) of the transit gateway Connect peer.")
	networkmanager_disassociateTransitGatewayConnectPeerCmd.MarkFlagRequired("global-network-id")
	networkmanager_disassociateTransitGatewayConnectPeerCmd.MarkFlagRequired("transit-gateway-connect-peer-arn")
	networkmanagerCmd.AddCommand(networkmanager_disassociateTransitGatewayConnectPeerCmd)
}
