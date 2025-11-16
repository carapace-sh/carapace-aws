package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteLocalGatewayRouteTableCmd = &cobra.Command{
	Use:   "delete-local-gateway-route-table",
	Short: "Deletes a local gateway route table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteLocalGatewayRouteTableCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteLocalGatewayRouteTableCmd).Standalone()

		ec2_deleteLocalGatewayRouteTableCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteLocalGatewayRouteTableCmd.Flags().String("local-gateway-route-table-id", "", "The ID of the local gateway route table.")
		ec2_deleteLocalGatewayRouteTableCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteLocalGatewayRouteTableCmd.MarkFlagRequired("local-gateway-route-table-id")
		ec2_deleteLocalGatewayRouteTableCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_deleteLocalGatewayRouteTableCmd)
}
