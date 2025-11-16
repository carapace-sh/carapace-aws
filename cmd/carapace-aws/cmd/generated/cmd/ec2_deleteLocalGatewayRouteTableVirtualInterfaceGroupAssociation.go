package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteLocalGatewayRouteTableVirtualInterfaceGroupAssociationCmd = &cobra.Command{
	Use:   "delete-local-gateway-route-table-virtual-interface-group-association",
	Short: "Deletes a local gateway route table virtual interface group association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteLocalGatewayRouteTableVirtualInterfaceGroupAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteLocalGatewayRouteTableVirtualInterfaceGroupAssociationCmd).Standalone()

		ec2_deleteLocalGatewayRouteTableVirtualInterfaceGroupAssociationCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteLocalGatewayRouteTableVirtualInterfaceGroupAssociationCmd.Flags().String("local-gateway-route-table-virtual-interface-group-association-id", "", "The ID of the local gateway route table virtual interface group association.")
		ec2_deleteLocalGatewayRouteTableVirtualInterfaceGroupAssociationCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteLocalGatewayRouteTableVirtualInterfaceGroupAssociationCmd.MarkFlagRequired("local-gateway-route-table-virtual-interface-group-association-id")
		ec2_deleteLocalGatewayRouteTableVirtualInterfaceGroupAssociationCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_deleteLocalGatewayRouteTableVirtualInterfaceGroupAssociationCmd)
}
