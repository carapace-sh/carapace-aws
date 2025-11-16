package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_resetNetworkInterfaceAttributeCmd = &cobra.Command{
	Use:   "reset-network-interface-attribute",
	Short: "Resets a network interface attribute.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_resetNetworkInterfaceAttributeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_resetNetworkInterfaceAttributeCmd).Standalone()

		ec2_resetNetworkInterfaceAttributeCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_resetNetworkInterfaceAttributeCmd.Flags().String("network-interface-id", "", "The ID of the network interface.")
		ec2_resetNetworkInterfaceAttributeCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_resetNetworkInterfaceAttributeCmd.Flags().String("source-dest-check", "", "The source/destination checking attribute.")
		ec2_resetNetworkInterfaceAttributeCmd.MarkFlagRequired("network-interface-id")
		ec2_resetNetworkInterfaceAttributeCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_resetNetworkInterfaceAttributeCmd)
}
