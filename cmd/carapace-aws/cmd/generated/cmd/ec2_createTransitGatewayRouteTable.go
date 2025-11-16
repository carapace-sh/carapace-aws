package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createTransitGatewayRouteTableCmd = &cobra.Command{
	Use:   "create-transit-gateway-route-table",
	Short: "Creates a route table for the specified transit gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createTransitGatewayRouteTableCmd).Standalone()

	ec2_createTransitGatewayRouteTableCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createTransitGatewayRouteTableCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createTransitGatewayRouteTableCmd.Flags().String("tag-specifications", "", "The tags to apply to the transit gateway route table.")
	ec2_createTransitGatewayRouteTableCmd.Flags().String("transit-gateway-id", "", "The ID of the transit gateway.")
	ec2_createTransitGatewayRouteTableCmd.Flag("no-dry-run").Hidden = true
	ec2_createTransitGatewayRouteTableCmd.MarkFlagRequired("transit-gateway-id")
	ec2Cmd.AddCommand(ec2_createTransitGatewayRouteTableCmd)
}
