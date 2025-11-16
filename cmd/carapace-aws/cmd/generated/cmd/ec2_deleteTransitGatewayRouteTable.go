package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteTransitGatewayRouteTableCmd = &cobra.Command{
	Use:   "delete-transit-gateway-route-table",
	Short: "Deletes the specified transit gateway route table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteTransitGatewayRouteTableCmd).Standalone()

	ec2_deleteTransitGatewayRouteTableCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteTransitGatewayRouteTableCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteTransitGatewayRouteTableCmd.Flags().String("transit-gateway-route-table-id", "", "The ID of the transit gateway route table.")
	ec2_deleteTransitGatewayRouteTableCmd.Flag("no-dry-run").Hidden = true
	ec2_deleteTransitGatewayRouteTableCmd.MarkFlagRequired("transit-gateway-route-table-id")
	ec2Cmd.AddCommand(ec2_deleteTransitGatewayRouteTableCmd)
}
