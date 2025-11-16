package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteTransitGatewayRouteTableAnnouncementCmd = &cobra.Command{
	Use:   "delete-transit-gateway-route-table-announcement",
	Short: "Advertises to the transit gateway that a transit gateway route table is deleted.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteTransitGatewayRouteTableAnnouncementCmd).Standalone()

	ec2_deleteTransitGatewayRouteTableAnnouncementCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteTransitGatewayRouteTableAnnouncementCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteTransitGatewayRouteTableAnnouncementCmd.Flags().String("transit-gateway-route-table-announcement-id", "", "The transit gateway route table ID that's being deleted.")
	ec2_deleteTransitGatewayRouteTableAnnouncementCmd.Flag("no-dry-run").Hidden = true
	ec2_deleteTransitGatewayRouteTableAnnouncementCmd.MarkFlagRequired("transit-gateway-route-table-announcement-id")
	ec2Cmd.AddCommand(ec2_deleteTransitGatewayRouteTableAnnouncementCmd)
}
