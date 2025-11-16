package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteLocalGatewayVirtualInterfaceGroupCmd = &cobra.Command{
	Use:   "delete-local-gateway-virtual-interface-group",
	Short: "Delete the specified local gateway interface group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteLocalGatewayVirtualInterfaceGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteLocalGatewayVirtualInterfaceGroupCmd).Standalone()

		ec2_deleteLocalGatewayVirtualInterfaceGroupCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteLocalGatewayVirtualInterfaceGroupCmd.Flags().String("local-gateway-virtual-interface-group-id", "", "The ID of the local gateway virtual interface group to delete.")
		ec2_deleteLocalGatewayVirtualInterfaceGroupCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteLocalGatewayVirtualInterfaceGroupCmd.MarkFlagRequired("local-gateway-virtual-interface-group-id")
		ec2_deleteLocalGatewayVirtualInterfaceGroupCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_deleteLocalGatewayVirtualInterfaceGroupCmd)
}
