package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createNetworkInterfaceCmd = &cobra.Command{
	Use:   "create-network-interface",
	Short: "Creates a network interface in the specified subnet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createNetworkInterfaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_createNetworkInterfaceCmd).Standalone()

		ec2_createNetworkInterfaceCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		ec2_createNetworkInterfaceCmd.Flags().String("connection-tracking-specification", "", "A connection tracking specification for the network interface.")
		ec2_createNetworkInterfaceCmd.Flags().String("description", "", "A description for the network interface.")
		ec2_createNetworkInterfaceCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createNetworkInterfaceCmd.Flags().Bool("enable-primary-ipv6", false, "If you’re creating a network interface in a dual-stack or IPv6-only subnet, you have the option to assign a primary IPv6 IP address.")
		ec2_createNetworkInterfaceCmd.Flags().String("groups", "", "The IDs of the security groups.")
		ec2_createNetworkInterfaceCmd.Flags().String("interface-type", "", "The type of network interface.")
		ec2_createNetworkInterfaceCmd.Flags().String("ipv4-prefix-count", "", "The number of IPv4 prefixes that Amazon Web Services automatically assigns to the network interface.")
		ec2_createNetworkInterfaceCmd.Flags().String("ipv4-prefixes", "", "The IPv4 prefixes assigned to the network interface.")
		ec2_createNetworkInterfaceCmd.Flags().String("ipv6-address-count", "", "The number of IPv6 addresses to assign to a network interface.")
		ec2_createNetworkInterfaceCmd.Flags().String("ipv6-addresses", "", "The IPv6 addresses from the IPv6 CIDR block range of your subnet.")
		ec2_createNetworkInterfaceCmd.Flags().String("ipv6-prefix-count", "", "The number of IPv6 prefixes that Amazon Web Services automatically assigns to the network interface.")
		ec2_createNetworkInterfaceCmd.Flags().String("ipv6-prefixes", "", "The IPv6 prefixes assigned to the network interface.")
		ec2_createNetworkInterfaceCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createNetworkInterfaceCmd.Flags().Bool("no-enable-primary-ipv6", false, "If you’re creating a network interface in a dual-stack or IPv6-only subnet, you have the option to assign a primary IPv6 IP address.")
		ec2_createNetworkInterfaceCmd.Flags().String("operator", "", "Reserved for internal use.")
		ec2_createNetworkInterfaceCmd.Flags().String("private-ip-address", "", "The primary private IPv4 address of the network interface.")
		ec2_createNetworkInterfaceCmd.Flags().String("private-ip-addresses", "", "The private IPv4 addresses.")
		ec2_createNetworkInterfaceCmd.Flags().String("secondary-private-ip-address-count", "", "The number of secondary private IPv4 addresses to assign to a network interface.")
		ec2_createNetworkInterfaceCmd.Flags().String("subnet-id", "", "The ID of the subnet to associate with the network interface.")
		ec2_createNetworkInterfaceCmd.Flags().String("tag-specifications", "", "The tags to apply to the new network interface.")
		ec2_createNetworkInterfaceCmd.Flag("no-dry-run").Hidden = true
		ec2_createNetworkInterfaceCmd.Flag("no-enable-primary-ipv6").Hidden = true
		ec2_createNetworkInterfaceCmd.MarkFlagRequired("subnet-id")
	})
	ec2Cmd.AddCommand(ec2_createNetworkInterfaceCmd)
}
