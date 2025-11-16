package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyVpnConnectionOptionsCmd = &cobra.Command{
	Use:   "modify-vpn-connection-options",
	Short: "Modifies the connection options for your Site-to-Site VPN connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyVpnConnectionOptionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_modifyVpnConnectionOptionsCmd).Standalone()

		ec2_modifyVpnConnectionOptionsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyVpnConnectionOptionsCmd.Flags().String("local-ipv4-network-cidr", "", "The IPv4 CIDR on the customer gateway (on-premises) side of the VPN connection.")
		ec2_modifyVpnConnectionOptionsCmd.Flags().String("local-ipv6-network-cidr", "", "The IPv6 CIDR on the customer gateway (on-premises) side of the VPN connection.")
		ec2_modifyVpnConnectionOptionsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyVpnConnectionOptionsCmd.Flags().String("remote-ipv4-network-cidr", "", "The IPv4 CIDR on the Amazon Web Services side of the VPN connection.")
		ec2_modifyVpnConnectionOptionsCmd.Flags().String("remote-ipv6-network-cidr", "", "The IPv6 CIDR on the Amazon Web Services side of the VPN connection.")
		ec2_modifyVpnConnectionOptionsCmd.Flags().String("vpn-connection-id", "", "The ID of the Site-to-Site VPN connection.")
		ec2_modifyVpnConnectionOptionsCmd.Flag("no-dry-run").Hidden = true
		ec2_modifyVpnConnectionOptionsCmd.MarkFlagRequired("vpn-connection-id")
	})
	ec2Cmd.AddCommand(ec2_modifyVpnConnectionOptionsCmd)
}
