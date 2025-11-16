package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_associateTransitGatewayConnectPeerCmd = &cobra.Command{
	Use:   "associate-transit-gateway-connect-peer",
	Short: "Associates a transit gateway Connect peer with a device, and optionally, with a link.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_associateTransitGatewayConnectPeerCmd).Standalone()

	networkmanager_associateTransitGatewayConnectPeerCmd.Flags().String("device-id", "", "The ID of the device.")
	networkmanager_associateTransitGatewayConnectPeerCmd.Flags().String("global-network-id", "", "The ID of the global network.")
	networkmanager_associateTransitGatewayConnectPeerCmd.Flags().String("link-id", "", "The ID of the link.")
	networkmanager_associateTransitGatewayConnectPeerCmd.Flags().String("transit-gateway-connect-peer-arn", "", "The Amazon Resource Name (ARN) of the Connect peer.")
	networkmanager_associateTransitGatewayConnectPeerCmd.MarkFlagRequired("device-id")
	networkmanager_associateTransitGatewayConnectPeerCmd.MarkFlagRequired("global-network-id")
	networkmanager_associateTransitGatewayConnectPeerCmd.MarkFlagRequired("transit-gateway-connect-peer-arn")
	networkmanagerCmd.AddCommand(networkmanager_associateTransitGatewayConnectPeerCmd)
}
