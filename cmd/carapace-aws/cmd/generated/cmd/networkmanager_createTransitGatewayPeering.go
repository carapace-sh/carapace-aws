package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_createTransitGatewayPeeringCmd = &cobra.Command{
	Use:   "create-transit-gateway-peering",
	Short: "Creates a transit gateway peering connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_createTransitGatewayPeeringCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_createTransitGatewayPeeringCmd).Standalone()

		networkmanager_createTransitGatewayPeeringCmd.Flags().String("client-token", "", "The client token associated with the request.")
		networkmanager_createTransitGatewayPeeringCmd.Flags().String("core-network-id", "", "The ID of a core network.")
		networkmanager_createTransitGatewayPeeringCmd.Flags().String("tags", "", "The list of key-value tags associated with the request.")
		networkmanager_createTransitGatewayPeeringCmd.Flags().String("transit-gateway-arn", "", "The ARN of the transit gateway for the peering request.")
		networkmanager_createTransitGatewayPeeringCmd.MarkFlagRequired("core-network-id")
		networkmanager_createTransitGatewayPeeringCmd.MarkFlagRequired("transit-gateway-arn")
	})
	networkmanagerCmd.AddCommand(networkmanager_createTransitGatewayPeeringCmd)
}
