package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createTransitGatewayRouteCmd = &cobra.Command{
	Use:   "create-transit-gateway-route",
	Short: "Creates a static route for the specified transit gateway route table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createTransitGatewayRouteCmd).Standalone()

	ec2_createTransitGatewayRouteCmd.Flags().Bool("blackhole", false, "Indicates whether to drop traffic that matches this route.")
	ec2_createTransitGatewayRouteCmd.Flags().String("destination-cidr-block", "", "The CIDR range used for destination matches.")
	ec2_createTransitGatewayRouteCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createTransitGatewayRouteCmd.Flags().Bool("no-blackhole", false, "Indicates whether to drop traffic that matches this route.")
	ec2_createTransitGatewayRouteCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createTransitGatewayRouteCmd.Flags().String("transit-gateway-attachment-id", "", "The ID of the attachment.")
	ec2_createTransitGatewayRouteCmd.Flags().String("transit-gateway-route-table-id", "", "The ID of the transit gateway route table.")
	ec2_createTransitGatewayRouteCmd.MarkFlagRequired("destination-cidr-block")
	ec2_createTransitGatewayRouteCmd.Flag("no-blackhole").Hidden = true
	ec2_createTransitGatewayRouteCmd.Flag("no-dry-run").Hidden = true
	ec2_createTransitGatewayRouteCmd.MarkFlagRequired("transit-gateway-route-table-id")
	ec2Cmd.AddCommand(ec2_createTransitGatewayRouteCmd)
}
