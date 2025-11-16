package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_associateTransitGatewayRouteTableCmd = &cobra.Command{
	Use:   "associate-transit-gateway-route-table",
	Short: "Associates the specified attachment with the specified transit gateway route table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_associateTransitGatewayRouteTableCmd).Standalone()

	ec2_associateTransitGatewayRouteTableCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_associateTransitGatewayRouteTableCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_associateTransitGatewayRouteTableCmd.Flags().String("transit-gateway-attachment-id", "", "The ID of the attachment.")
	ec2_associateTransitGatewayRouteTableCmd.Flags().String("transit-gateway-route-table-id", "", "The ID of the transit gateway route table.")
	ec2_associateTransitGatewayRouteTableCmd.Flag("no-dry-run").Hidden = true
	ec2_associateTransitGatewayRouteTableCmd.MarkFlagRequired("transit-gateway-attachment-id")
	ec2_associateTransitGatewayRouteTableCmd.MarkFlagRequired("transit-gateway-route-table-id")
	ec2Cmd.AddCommand(ec2_associateTransitGatewayRouteTableCmd)
}
