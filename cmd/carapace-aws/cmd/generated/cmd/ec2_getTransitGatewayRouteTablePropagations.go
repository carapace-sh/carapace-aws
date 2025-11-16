package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getTransitGatewayRouteTablePropagationsCmd = &cobra.Command{
	Use:   "get-transit-gateway-route-table-propagations",
	Short: "Gets information about the route table propagations for the specified transit gateway route table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getTransitGatewayRouteTablePropagationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_getTransitGatewayRouteTablePropagationsCmd).Standalone()

		ec2_getTransitGatewayRouteTablePropagationsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getTransitGatewayRouteTablePropagationsCmd.Flags().String("filters", "", "One or more filters.")
		ec2_getTransitGatewayRouteTablePropagationsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		ec2_getTransitGatewayRouteTablePropagationsCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ec2_getTransitGatewayRouteTablePropagationsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getTransitGatewayRouteTablePropagationsCmd.Flags().String("transit-gateway-route-table-id", "", "The ID of the transit gateway route table.")
		ec2_getTransitGatewayRouteTablePropagationsCmd.Flag("no-dry-run").Hidden = true
		ec2_getTransitGatewayRouteTablePropagationsCmd.MarkFlagRequired("transit-gateway-route-table-id")
	})
	ec2Cmd.AddCommand(ec2_getTransitGatewayRouteTablePropagationsCmd)
}
