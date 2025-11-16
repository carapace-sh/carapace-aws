package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteNetworkInterfaceCmd = &cobra.Command{
	Use:   "delete-network-interface",
	Short: "Deletes the specified network interface.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteNetworkInterfaceCmd).Standalone()

	ec2_deleteNetworkInterfaceCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteNetworkInterfaceCmd.Flags().String("network-interface-id", "", "The ID of the network interface.")
	ec2_deleteNetworkInterfaceCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteNetworkInterfaceCmd.MarkFlagRequired("network-interface-id")
	ec2_deleteNetworkInterfaceCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_deleteNetworkInterfaceCmd)
}
