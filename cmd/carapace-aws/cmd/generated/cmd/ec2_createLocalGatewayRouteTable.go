package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createLocalGatewayRouteTableCmd = &cobra.Command{
	Use:   "create-local-gateway-route-table",
	Short: "Creates a local gateway route table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createLocalGatewayRouteTableCmd).Standalone()

	ec2_createLocalGatewayRouteTableCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createLocalGatewayRouteTableCmd.Flags().String("local-gateway-id", "", "The ID of the local gateway.")
	ec2_createLocalGatewayRouteTableCmd.Flags().String("mode", "", "The mode of the local gateway route table.")
	ec2_createLocalGatewayRouteTableCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createLocalGatewayRouteTableCmd.Flags().String("tag-specifications", "", "The tags assigned to the local gateway route table.")
	ec2_createLocalGatewayRouteTableCmd.MarkFlagRequired("local-gateway-id")
	ec2_createLocalGatewayRouteTableCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_createLocalGatewayRouteTableCmd)
}
