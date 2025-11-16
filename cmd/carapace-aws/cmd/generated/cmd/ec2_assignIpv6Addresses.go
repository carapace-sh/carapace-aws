package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_assignIpv6AddressesCmd = &cobra.Command{
	Use:   "assign-ipv6-addresses",
	Short: "Assigns the specified IPv6 addresses to the specified network interface.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_assignIpv6AddressesCmd).Standalone()

	ec2_assignIpv6AddressesCmd.Flags().String("ipv6-address-count", "", "The number of additional IPv6 addresses to assign to the network interface.")
	ec2_assignIpv6AddressesCmd.Flags().String("ipv6-addresses", "", "The IPv6 addresses to be assigned to the network interface.")
	ec2_assignIpv6AddressesCmd.Flags().String("ipv6-prefix-count", "", "The number of IPv6 prefixes that Amazon Web Services automatically assigns to the network interface.")
	ec2_assignIpv6AddressesCmd.Flags().String("ipv6-prefixes", "", "One or more IPv6 prefixes assigned to the network interface.")
	ec2_assignIpv6AddressesCmd.Flags().String("network-interface-id", "", "The ID of the network interface.")
	ec2_assignIpv6AddressesCmd.MarkFlagRequired("network-interface-id")
	ec2Cmd.AddCommand(ec2_assignIpv6AddressesCmd)
}
