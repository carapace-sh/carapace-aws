package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_getTransitGatewayRouteTableAttachmentCmd = &cobra.Command{
	Use:   "get-transit-gateway-route-table-attachment",
	Short: "Returns information about a transit gateway route table attachment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_getTransitGatewayRouteTableAttachmentCmd).Standalone()

	networkmanager_getTransitGatewayRouteTableAttachmentCmd.Flags().String("attachment-id", "", "The ID of the transit gateway route table attachment.")
	networkmanager_getTransitGatewayRouteTableAttachmentCmd.MarkFlagRequired("attachment-id")
	networkmanagerCmd.AddCommand(networkmanager_getTransitGatewayRouteTableAttachmentCmd)
}
