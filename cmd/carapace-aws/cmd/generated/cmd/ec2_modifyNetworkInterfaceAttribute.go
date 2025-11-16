package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyNetworkInterfaceAttributeCmd = &cobra.Command{
	Use:   "modify-network-interface-attribute",
	Short: "Modifies the specified network interface attribute.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyNetworkInterfaceAttributeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_modifyNetworkInterfaceAttributeCmd).Standalone()

		ec2_modifyNetworkInterfaceAttributeCmd.Flags().Bool("associate-public-ip-address", false, "Indicates whether to assign a public IPv4 address to a network interface.")
		ec2_modifyNetworkInterfaceAttributeCmd.Flags().String("associated-subnet-ids", "", "A list of subnet IDs to associate with the network interface.")
		ec2_modifyNetworkInterfaceAttributeCmd.Flags().String("attachment", "", "Information about the interface attachment.")
		ec2_modifyNetworkInterfaceAttributeCmd.Flags().String("connection-tracking-specification", "", "A connection tracking specification.")
		ec2_modifyNetworkInterfaceAttributeCmd.Flags().String("description", "", "A description for the network interface.")
		ec2_modifyNetworkInterfaceAttributeCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyNetworkInterfaceAttributeCmd.Flags().String("ena-srd-specification", "", "Updates the ENA Express configuration for the network interface that’s attached to the instance.")
		ec2_modifyNetworkInterfaceAttributeCmd.Flags().Bool("enable-primary-ipv6", false, "If you’re modifying a network interface in a dual-stack or IPv6-only subnet, you have the option to assign a primary IPv6 IP address.")
		ec2_modifyNetworkInterfaceAttributeCmd.Flags().String("groups", "", "Changes the security groups for the network interface.")
		ec2_modifyNetworkInterfaceAttributeCmd.Flags().String("network-interface-id", "", "The ID of the network interface.")
		ec2_modifyNetworkInterfaceAttributeCmd.Flags().Bool("no-associate-public-ip-address", false, "Indicates whether to assign a public IPv4 address to a network interface.")
		ec2_modifyNetworkInterfaceAttributeCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyNetworkInterfaceAttributeCmd.Flags().Bool("no-enable-primary-ipv6", false, "If you’re modifying a network interface in a dual-stack or IPv6-only subnet, you have the option to assign a primary IPv6 IP address.")
		ec2_modifyNetworkInterfaceAttributeCmd.Flags().String("source-dest-check", "", "Enable or disable source/destination checks, which ensure that the instance is either the source or the destination of any traffic that it receives.")
		ec2_modifyNetworkInterfaceAttributeCmd.MarkFlagRequired("network-interface-id")
		ec2_modifyNetworkInterfaceAttributeCmd.Flag("no-associate-public-ip-address").Hidden = true
		ec2_modifyNetworkInterfaceAttributeCmd.Flag("no-dry-run").Hidden = true
		ec2_modifyNetworkInterfaceAttributeCmd.Flag("no-enable-primary-ipv6").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_modifyNetworkInterfaceAttributeCmd)
}
