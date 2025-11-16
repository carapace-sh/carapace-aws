package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeNetworkInterfaceAttributeCmd = &cobra.Command{
	Use:   "describe-network-interface-attribute",
	Short: "Describes a network interface attribute.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeNetworkInterfaceAttributeCmd).Standalone()

	ec2_describeNetworkInterfaceAttributeCmd.Flags().String("attribute", "", "The attribute of the network interface.")
	ec2_describeNetworkInterfaceAttributeCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeNetworkInterfaceAttributeCmd.Flags().String("network-interface-id", "", "The ID of the network interface.")
	ec2_describeNetworkInterfaceAttributeCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeNetworkInterfaceAttributeCmd.MarkFlagRequired("network-interface-id")
	ec2_describeNetworkInterfaceAttributeCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeNetworkInterfaceAttributeCmd)
}
