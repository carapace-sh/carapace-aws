package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteLocalGatewayVirtualInterfaceCmd = &cobra.Command{
	Use:   "delete-local-gateway-virtual-interface",
	Short: "Deletes the specified local gateway virtual interface.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteLocalGatewayVirtualInterfaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteLocalGatewayVirtualInterfaceCmd).Standalone()

		ec2_deleteLocalGatewayVirtualInterfaceCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteLocalGatewayVirtualInterfaceCmd.Flags().String("local-gateway-virtual-interface-id", "", "The ID of the local virtual interface to delete.")
		ec2_deleteLocalGatewayVirtualInterfaceCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteLocalGatewayVirtualInterfaceCmd.MarkFlagRequired("local-gateway-virtual-interface-id")
		ec2_deleteLocalGatewayVirtualInterfaceCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_deleteLocalGatewayVirtualInterfaceCmd)
}
