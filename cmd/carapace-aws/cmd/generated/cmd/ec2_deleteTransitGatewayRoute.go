package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteTransitGatewayRouteCmd = &cobra.Command{
	Use:   "delete-transit-gateway-route",
	Short: "Deletes the specified route from the specified transit gateway route table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteTransitGatewayRouteCmd).Standalone()

	ec2_deleteTransitGatewayRouteCmd.Flags().String("destination-cidr-block", "", "The CIDR range for the route.")
	ec2_deleteTransitGatewayRouteCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteTransitGatewayRouteCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteTransitGatewayRouteCmd.Flags().String("transit-gateway-route-table-id", "", "The ID of the transit gateway route table.")
	ec2_deleteTransitGatewayRouteCmd.MarkFlagRequired("destination-cidr-block")
	ec2_deleteTransitGatewayRouteCmd.Flag("no-dry-run").Hidden = true
	ec2_deleteTransitGatewayRouteCmd.MarkFlagRequired("transit-gateway-route-table-id")
	ec2Cmd.AddCommand(ec2_deleteTransitGatewayRouteCmd)
}
