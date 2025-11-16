package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyLocalGatewayRouteCmd = &cobra.Command{
	Use:   "modify-local-gateway-route",
	Short: "Modifies the specified local gateway route.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyLocalGatewayRouteCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_modifyLocalGatewayRouteCmd).Standalone()

		ec2_modifyLocalGatewayRouteCmd.Flags().String("destination-cidr-block", "", "The CIDR block used for destination matches.")
		ec2_modifyLocalGatewayRouteCmd.Flags().String("destination-prefix-list-id", "", "The ID of the prefix list.")
		ec2_modifyLocalGatewayRouteCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyLocalGatewayRouteCmd.Flags().String("local-gateway-route-table-id", "", "The ID of the local gateway route table.")
		ec2_modifyLocalGatewayRouteCmd.Flags().String("local-gateway-virtual-interface-group-id", "", "The ID of the virtual interface group.")
		ec2_modifyLocalGatewayRouteCmd.Flags().String("network-interface-id", "", "The ID of the network interface.")
		ec2_modifyLocalGatewayRouteCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyLocalGatewayRouteCmd.MarkFlagRequired("local-gateway-route-table-id")
		ec2_modifyLocalGatewayRouteCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_modifyLocalGatewayRouteCmd)
}
