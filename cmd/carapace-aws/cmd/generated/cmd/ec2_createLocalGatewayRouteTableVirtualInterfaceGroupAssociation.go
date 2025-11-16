package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createLocalGatewayRouteTableVirtualInterfaceGroupAssociationCmd = &cobra.Command{
	Use:   "create-local-gateway-route-table-virtual-interface-group-association",
	Short: "Creates a local gateway route table virtual interface group association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createLocalGatewayRouteTableVirtualInterfaceGroupAssociationCmd).Standalone()

	ec2_createLocalGatewayRouteTableVirtualInterfaceGroupAssociationCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createLocalGatewayRouteTableVirtualInterfaceGroupAssociationCmd.Flags().String("local-gateway-route-table-id", "", "The ID of the local gateway route table.")
	ec2_createLocalGatewayRouteTableVirtualInterfaceGroupAssociationCmd.Flags().String("local-gateway-virtual-interface-group-id", "", "The ID of the local gateway route table virtual interface group association.")
	ec2_createLocalGatewayRouteTableVirtualInterfaceGroupAssociationCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createLocalGatewayRouteTableVirtualInterfaceGroupAssociationCmd.Flags().String("tag-specifications", "", "The tags assigned to the local gateway route table virtual interface group association.")
	ec2_createLocalGatewayRouteTableVirtualInterfaceGroupAssociationCmd.MarkFlagRequired("local-gateway-route-table-id")
	ec2_createLocalGatewayRouteTableVirtualInterfaceGroupAssociationCmd.MarkFlagRequired("local-gateway-virtual-interface-group-id")
	ec2_createLocalGatewayRouteTableVirtualInterfaceGroupAssociationCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_createLocalGatewayRouteTableVirtualInterfaceGroupAssociationCmd)
}
