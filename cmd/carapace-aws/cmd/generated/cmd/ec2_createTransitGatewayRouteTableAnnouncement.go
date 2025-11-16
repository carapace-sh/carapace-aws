package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createTransitGatewayRouteTableAnnouncementCmd = &cobra.Command{
	Use:   "create-transit-gateway-route-table-announcement",
	Short: "Advertises a new transit gateway route table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createTransitGatewayRouteTableAnnouncementCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_createTransitGatewayRouteTableAnnouncementCmd).Standalone()

		ec2_createTransitGatewayRouteTableAnnouncementCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createTransitGatewayRouteTableAnnouncementCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createTransitGatewayRouteTableAnnouncementCmd.Flags().String("peering-attachment-id", "", "The ID of the peering attachment.")
		ec2_createTransitGatewayRouteTableAnnouncementCmd.Flags().String("tag-specifications", "", "The tags specifications applied to the transit gateway route table announcement.")
		ec2_createTransitGatewayRouteTableAnnouncementCmd.Flags().String("transit-gateway-route-table-id", "", "The ID of the transit gateway route table.")
		ec2_createTransitGatewayRouteTableAnnouncementCmd.Flag("no-dry-run").Hidden = true
		ec2_createTransitGatewayRouteTableAnnouncementCmd.MarkFlagRequired("peering-attachment-id")
		ec2_createTransitGatewayRouteTableAnnouncementCmd.MarkFlagRequired("transit-gateway-route-table-id")
	})
	ec2Cmd.AddCommand(ec2_createTransitGatewayRouteTableAnnouncementCmd)
}
