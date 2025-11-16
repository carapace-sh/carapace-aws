package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_disassociateTransitGatewayRouteTableCmd = &cobra.Command{
	Use:   "disassociate-transit-gateway-route-table",
	Short: "Disassociates a resource attachment from a transit gateway route table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_disassociateTransitGatewayRouteTableCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_disassociateTransitGatewayRouteTableCmd).Standalone()

		ec2_disassociateTransitGatewayRouteTableCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_disassociateTransitGatewayRouteTableCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_disassociateTransitGatewayRouteTableCmd.Flags().String("transit-gateway-attachment-id", "", "The ID of the attachment.")
		ec2_disassociateTransitGatewayRouteTableCmd.Flags().String("transit-gateway-route-table-id", "", "The ID of the transit gateway route table.")
		ec2_disassociateTransitGatewayRouteTableCmd.Flag("no-dry-run").Hidden = true
		ec2_disassociateTransitGatewayRouteTableCmd.MarkFlagRequired("transit-gateway-attachment-id")
		ec2_disassociateTransitGatewayRouteTableCmd.MarkFlagRequired("transit-gateway-route-table-id")
	})
	ec2Cmd.AddCommand(ec2_disassociateTransitGatewayRouteTableCmd)
}
