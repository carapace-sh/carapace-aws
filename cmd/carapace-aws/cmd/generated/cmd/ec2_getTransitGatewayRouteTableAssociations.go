package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getTransitGatewayRouteTableAssociationsCmd = &cobra.Command{
	Use:   "get-transit-gateway-route-table-associations",
	Short: "Gets information about the associations for the specified transit gateway route table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getTransitGatewayRouteTableAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_getTransitGatewayRouteTableAssociationsCmd).Standalone()

		ec2_getTransitGatewayRouteTableAssociationsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getTransitGatewayRouteTableAssociationsCmd.Flags().String("filters", "", "One or more filters.")
		ec2_getTransitGatewayRouteTableAssociationsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		ec2_getTransitGatewayRouteTableAssociationsCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ec2_getTransitGatewayRouteTableAssociationsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getTransitGatewayRouteTableAssociationsCmd.Flags().String("transit-gateway-route-table-id", "", "The ID of the transit gateway route table.")
		ec2_getTransitGatewayRouteTableAssociationsCmd.Flag("no-dry-run").Hidden = true
		ec2_getTransitGatewayRouteTableAssociationsCmd.MarkFlagRequired("transit-gateway-route-table-id")
	})
	ec2Cmd.AddCommand(ec2_getTransitGatewayRouteTableAssociationsCmd)
}
