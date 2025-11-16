package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteNetworkInterfacePermissionCmd = &cobra.Command{
	Use:   "delete-network-interface-permission",
	Short: "Deletes a permission for a network interface.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteNetworkInterfacePermissionCmd).Standalone()

	ec2_deleteNetworkInterfacePermissionCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteNetworkInterfacePermissionCmd.Flags().Bool("force", false, "Specify `true` to remove the permission even if the network interface is attached to an instance.")
	ec2_deleteNetworkInterfacePermissionCmd.Flags().String("network-interface-permission-id", "", "The ID of the network interface permission.")
	ec2_deleteNetworkInterfacePermissionCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteNetworkInterfacePermissionCmd.Flags().Bool("no-force", false, "Specify `true` to remove the permission even if the network interface is attached to an instance.")
	ec2_deleteNetworkInterfacePermissionCmd.MarkFlagRequired("network-interface-permission-id")
	ec2_deleteNetworkInterfacePermissionCmd.Flag("no-dry-run").Hidden = true
	ec2_deleteNetworkInterfacePermissionCmd.Flag("no-force").Hidden = true
	ec2Cmd.AddCommand(ec2_deleteNetworkInterfacePermissionCmd)
}
