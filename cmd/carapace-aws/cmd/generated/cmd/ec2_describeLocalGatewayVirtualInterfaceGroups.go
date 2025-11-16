package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeLocalGatewayVirtualInterfaceGroupsCmd = &cobra.Command{
	Use:   "describe-local-gateway-virtual-interface-groups",
	Short: "Describes the specified local gateway virtual interface groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeLocalGatewayVirtualInterfaceGroupsCmd).Standalone()

	ec2_describeLocalGatewayVirtualInterfaceGroupsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeLocalGatewayVirtualInterfaceGroupsCmd.Flags().String("filters", "", "One or more filters.")
	ec2_describeLocalGatewayVirtualInterfaceGroupsCmd.Flags().String("local-gateway-virtual-interface-group-ids", "", "The IDs of the virtual interface groups.")
	ec2_describeLocalGatewayVirtualInterfaceGroupsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	ec2_describeLocalGatewayVirtualInterfaceGroupsCmd.Flags().String("next-token", "", "The token for the next page of results.")
	ec2_describeLocalGatewayVirtualInterfaceGroupsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeLocalGatewayVirtualInterfaceGroupsCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeLocalGatewayVirtualInterfaceGroupsCmd)
}
