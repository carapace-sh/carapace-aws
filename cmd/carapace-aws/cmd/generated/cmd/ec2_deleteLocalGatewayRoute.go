package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteLocalGatewayRouteCmd = &cobra.Command{
	Use:   "delete-local-gateway-route",
	Short: "Deletes the specified route from the specified local gateway route table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteLocalGatewayRouteCmd).Standalone()

	ec2_deleteLocalGatewayRouteCmd.Flags().String("destination-cidr-block", "", "The CIDR range for the route.")
	ec2_deleteLocalGatewayRouteCmd.Flags().String("destination-prefix-list-id", "", "Use a prefix list in place of `DestinationCidrBlock`.")
	ec2_deleteLocalGatewayRouteCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteLocalGatewayRouteCmd.Flags().String("local-gateway-route-table-id", "", "The ID of the local gateway route table.")
	ec2_deleteLocalGatewayRouteCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteLocalGatewayRouteCmd.MarkFlagRequired("local-gateway-route-table-id")
	ec2_deleteLocalGatewayRouteCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_deleteLocalGatewayRouteCmd)
}
