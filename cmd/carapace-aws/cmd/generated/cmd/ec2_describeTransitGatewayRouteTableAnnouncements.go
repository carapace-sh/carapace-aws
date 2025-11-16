package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeTransitGatewayRouteTableAnnouncementsCmd = &cobra.Command{
	Use:   "describe-transit-gateway-route-table-announcements",
	Short: "Describes one or more transit gateway route table advertisements.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeTransitGatewayRouteTableAnnouncementsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeTransitGatewayRouteTableAnnouncementsCmd).Standalone()

		ec2_describeTransitGatewayRouteTableAnnouncementsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeTransitGatewayRouteTableAnnouncementsCmd.Flags().String("filters", "", "The filters associated with the transit gateway policy table.")
		ec2_describeTransitGatewayRouteTableAnnouncementsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		ec2_describeTransitGatewayRouteTableAnnouncementsCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ec2_describeTransitGatewayRouteTableAnnouncementsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeTransitGatewayRouteTableAnnouncementsCmd.Flags().String("transit-gateway-route-table-announcement-ids", "", "The IDs of the transit gateway route tables that are being advertised.")
		ec2_describeTransitGatewayRouteTableAnnouncementsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeTransitGatewayRouteTableAnnouncementsCmd)
}
