package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_replaceTransitGatewayRouteCmd = &cobra.Command{
	Use:   "replace-transit-gateway-route",
	Short: "Replaces the specified route in the specified transit gateway route table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_replaceTransitGatewayRouteCmd).Standalone()

	ec2_replaceTransitGatewayRouteCmd.Flags().Bool("blackhole", false, "Indicates whether traffic matching this route is to be dropped.")
	ec2_replaceTransitGatewayRouteCmd.Flags().String("destination-cidr-block", "", "The CIDR range used for the destination match.")
	ec2_replaceTransitGatewayRouteCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_replaceTransitGatewayRouteCmd.Flags().Bool("no-blackhole", false, "Indicates whether traffic matching this route is to be dropped.")
	ec2_replaceTransitGatewayRouteCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_replaceTransitGatewayRouteCmd.Flags().String("transit-gateway-attachment-id", "", "The ID of the attachment.")
	ec2_replaceTransitGatewayRouteCmd.Flags().String("transit-gateway-route-table-id", "", "The ID of the route table.")
	ec2_replaceTransitGatewayRouteCmd.MarkFlagRequired("destination-cidr-block")
	ec2_replaceTransitGatewayRouteCmd.Flag("no-blackhole").Hidden = true
	ec2_replaceTransitGatewayRouteCmd.Flag("no-dry-run").Hidden = true
	ec2_replaceTransitGatewayRouteCmd.MarkFlagRequired("transit-gateway-route-table-id")
	ec2Cmd.AddCommand(ec2_replaceTransitGatewayRouteCmd)
}
