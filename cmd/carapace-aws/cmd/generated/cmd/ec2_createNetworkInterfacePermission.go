package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createNetworkInterfacePermissionCmd = &cobra.Command{
	Use:   "create-network-interface-permission",
	Short: "Grants an Amazon Web Services-authorized account permission to attach the specified network interface to an instance in their account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createNetworkInterfacePermissionCmd).Standalone()

	ec2_createNetworkInterfacePermissionCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID.")
	ec2_createNetworkInterfacePermissionCmd.Flags().String("aws-service", "", "The Amazon Web Services service.")
	ec2_createNetworkInterfacePermissionCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createNetworkInterfacePermissionCmd.Flags().String("network-interface-id", "", "The ID of the network interface.")
	ec2_createNetworkInterfacePermissionCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createNetworkInterfacePermissionCmd.Flags().String("permission", "", "The type of permission to grant.")
	ec2_createNetworkInterfacePermissionCmd.MarkFlagRequired("network-interface-id")
	ec2_createNetworkInterfacePermissionCmd.Flag("no-dry-run").Hidden = true
	ec2_createNetworkInterfacePermissionCmd.MarkFlagRequired("permission")
	ec2Cmd.AddCommand(ec2_createNetworkInterfacePermissionCmd)
}
