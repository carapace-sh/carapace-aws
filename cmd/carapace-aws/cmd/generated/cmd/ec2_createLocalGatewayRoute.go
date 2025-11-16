package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createLocalGatewayRouteCmd = &cobra.Command{
	Use:   "create-local-gateway-route",
	Short: "Creates a static route for the specified local gateway route table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createLocalGatewayRouteCmd).Standalone()

	ec2_createLocalGatewayRouteCmd.Flags().String("destination-cidr-block", "", "The CIDR range used for destination matches.")
	ec2_createLocalGatewayRouteCmd.Flags().String("destination-prefix-list-id", "", "The ID of the prefix list.")
	ec2_createLocalGatewayRouteCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createLocalGatewayRouteCmd.Flags().String("local-gateway-route-table-id", "", "The ID of the local gateway route table.")
	ec2_createLocalGatewayRouteCmd.Flags().String("local-gateway-virtual-interface-group-id", "", "The ID of the virtual interface group.")
	ec2_createLocalGatewayRouteCmd.Flags().String("network-interface-id", "", "The ID of the network interface.")
	ec2_createLocalGatewayRouteCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createLocalGatewayRouteCmd.MarkFlagRequired("local-gateway-route-table-id")
	ec2_createLocalGatewayRouteCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_createLocalGatewayRouteCmd)
}
