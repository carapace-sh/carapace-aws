package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_unassignPrivateIpAddressesCmd = &cobra.Command{
	Use:   "unassign-private-ip-addresses",
	Short: "Unassigns the specified secondary private IP addresses or IPv4 Prefix Delegation prefixes from a network interface.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_unassignPrivateIpAddressesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_unassignPrivateIpAddressesCmd).Standalone()

		ec2_unassignPrivateIpAddressesCmd.Flags().String("ipv4-prefixes", "", "The IPv4 prefixes to unassign from the network interface.")
		ec2_unassignPrivateIpAddressesCmd.Flags().String("network-interface-id", "", "The ID of the network interface.")
		ec2_unassignPrivateIpAddressesCmd.Flags().String("private-ip-addresses", "", "The secondary private IP addresses to unassign from the network interface.")
		ec2_unassignPrivateIpAddressesCmd.MarkFlagRequired("network-interface-id")
	})
	ec2Cmd.AddCommand(ec2_unassignPrivateIpAddressesCmd)
}
