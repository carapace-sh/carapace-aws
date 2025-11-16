package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_unassignIpv6AddressesCmd = &cobra.Command{
	Use:   "unassign-ipv6-addresses",
	Short: "Unassigns the specified IPv6 addresses or Prefix Delegation prefixes from a network interface.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_unassignIpv6AddressesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_unassignIpv6AddressesCmd).Standalone()

		ec2_unassignIpv6AddressesCmd.Flags().String("ipv6-addresses", "", "The IPv6 addresses to unassign from the network interface.")
		ec2_unassignIpv6AddressesCmd.Flags().String("ipv6-prefixes", "", "The IPv6 prefixes to unassign from the network interface.")
		ec2_unassignIpv6AddressesCmd.Flags().String("network-interface-id", "", "The ID of the network interface.")
		ec2_unassignIpv6AddressesCmd.MarkFlagRequired("network-interface-id")
	})
	ec2Cmd.AddCommand(ec2_unassignIpv6AddressesCmd)
}
