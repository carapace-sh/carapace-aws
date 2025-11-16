package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_disableTransitGatewayRouteTablePropagationCmd = &cobra.Command{
	Use:   "disable-transit-gateway-route-table-propagation",
	Short: "Disables the specified resource attachment from propagating routes to the specified propagation route table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_disableTransitGatewayRouteTablePropagationCmd).Standalone()

	ec2_disableTransitGatewayRouteTablePropagationCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_disableTransitGatewayRouteTablePropagationCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_disableTransitGatewayRouteTablePropagationCmd.Flags().String("transit-gateway-attachment-id", "", "The ID of the attachment.")
	ec2_disableTransitGatewayRouteTablePropagationCmd.Flags().String("transit-gateway-route-table-announcement-id", "", "The ID of the route table announcement.")
	ec2_disableTransitGatewayRouteTablePropagationCmd.Flags().String("transit-gateway-route-table-id", "", "The ID of the propagation route table.")
	ec2_disableTransitGatewayRouteTablePropagationCmd.Flag("no-dry-run").Hidden = true
	ec2_disableTransitGatewayRouteTablePropagationCmd.MarkFlagRequired("transit-gateway-route-table-id")
	ec2Cmd.AddCommand(ec2_disableTransitGatewayRouteTablePropagationCmd)
}
