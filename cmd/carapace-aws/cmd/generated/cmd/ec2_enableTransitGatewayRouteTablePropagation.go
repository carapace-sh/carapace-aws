package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_enableTransitGatewayRouteTablePropagationCmd = &cobra.Command{
	Use:   "enable-transit-gateway-route-table-propagation",
	Short: "Enables the specified attachment to propagate routes to the specified propagation route table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_enableTransitGatewayRouteTablePropagationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_enableTransitGatewayRouteTablePropagationCmd).Standalone()

		ec2_enableTransitGatewayRouteTablePropagationCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_enableTransitGatewayRouteTablePropagationCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_enableTransitGatewayRouteTablePropagationCmd.Flags().String("transit-gateway-attachment-id", "", "The ID of the attachment.")
		ec2_enableTransitGatewayRouteTablePropagationCmd.Flags().String("transit-gateway-route-table-announcement-id", "", "The ID of the transit gateway route table announcement.")
		ec2_enableTransitGatewayRouteTablePropagationCmd.Flags().String("transit-gateway-route-table-id", "", "The ID of the propagation route table.")
		ec2_enableTransitGatewayRouteTablePropagationCmd.Flag("no-dry-run").Hidden = true
		ec2_enableTransitGatewayRouteTablePropagationCmd.MarkFlagRequired("transit-gateway-route-table-id")
	})
	ec2Cmd.AddCommand(ec2_enableTransitGatewayRouteTablePropagationCmd)
}
