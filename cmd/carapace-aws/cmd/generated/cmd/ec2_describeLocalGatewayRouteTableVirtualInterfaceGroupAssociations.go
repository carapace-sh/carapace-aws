package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsCmd = &cobra.Command{
	Use:   "describe-local-gateway-route-table-virtual-interface-group-associations",
	Short: "Describes the associations between virtual interface groups and local gateway route tables.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsCmd).Standalone()

		ec2_describeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsCmd.Flags().String("filters", "", "One or more filters.")
		ec2_describeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsCmd.Flags().String("local-gateway-route-table-virtual-interface-group-association-ids", "", "The IDs of the associations.")
		ec2_describeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		ec2_describeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ec2_describeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsCmd)
}
