package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_assignPrivateIpAddressesCmd = &cobra.Command{
	Use:   "assign-private-ip-addresses",
	Short: "Assigns the specified secondary private IP addresses to the specified network interface.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_assignPrivateIpAddressesCmd).Standalone()

	ec2_assignPrivateIpAddressesCmd.Flags().Bool("allow-reassignment", false, "Indicates whether to allow an IP address that is already assigned to another network interface or instance to be reassigned to the specified network interface.")
	ec2_assignPrivateIpAddressesCmd.Flags().String("ipv4-prefix-count", "", "The number of IPv4 prefixes that Amazon Web Services automatically assigns to the network interface.")
	ec2_assignPrivateIpAddressesCmd.Flags().String("ipv4-prefixes", "", "One or more IPv4 prefixes assigned to the network interface.")
	ec2_assignPrivateIpAddressesCmd.Flags().String("network-interface-id", "", "The ID of the network interface.")
	ec2_assignPrivateIpAddressesCmd.Flags().Bool("no-allow-reassignment", false, "Indicates whether to allow an IP address that is already assigned to another network interface or instance to be reassigned to the specified network interface.")
	ec2_assignPrivateIpAddressesCmd.Flags().String("private-ip-addresses", "", "The IP addresses to be assigned as a secondary private IP address to the network interface.")
	ec2_assignPrivateIpAddressesCmd.Flags().String("secondary-private-ip-address-count", "", "The number of secondary IP addresses to assign to the network interface.")
	ec2_assignPrivateIpAddressesCmd.MarkFlagRequired("network-interface-id")
	ec2_assignPrivateIpAddressesCmd.Flag("no-allow-reassignment").Hidden = true
	ec2Cmd.AddCommand(ec2_assignPrivateIpAddressesCmd)
}
