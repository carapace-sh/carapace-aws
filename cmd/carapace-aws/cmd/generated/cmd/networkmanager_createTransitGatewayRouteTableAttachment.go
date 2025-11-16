package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_createTransitGatewayRouteTableAttachmentCmd = &cobra.Command{
	Use:   "create-transit-gateway-route-table-attachment",
	Short: "Creates a transit gateway route table attachment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_createTransitGatewayRouteTableAttachmentCmd).Standalone()

	networkmanager_createTransitGatewayRouteTableAttachmentCmd.Flags().String("client-token", "", "The client token associated with the request.")
	networkmanager_createTransitGatewayRouteTableAttachmentCmd.Flags().String("peering-id", "", "The ID of the peer for the")
	networkmanager_createTransitGatewayRouteTableAttachmentCmd.Flags().String("tags", "", "The list of key-value tags associated with the request.")
	networkmanager_createTransitGatewayRouteTableAttachmentCmd.Flags().String("transit-gateway-route-table-arn", "", "The ARN of the transit gateway route table for the attachment request.")
	networkmanager_createTransitGatewayRouteTableAttachmentCmd.MarkFlagRequired("peering-id")
	networkmanager_createTransitGatewayRouteTableAttachmentCmd.MarkFlagRequired("transit-gateway-route-table-arn")
	networkmanagerCmd.AddCommand(networkmanager_createTransitGatewayRouteTableAttachmentCmd)
}
