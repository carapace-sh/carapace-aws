package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyVpnTunnelCertificateCmd = &cobra.Command{
	Use:   "modify-vpn-tunnel-certificate",
	Short: "Modifies the VPN tunnel endpoint certificate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyVpnTunnelCertificateCmd).Standalone()

	ec2_modifyVpnTunnelCertificateCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_modifyVpnTunnelCertificateCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_modifyVpnTunnelCertificateCmd.Flags().String("vpn-connection-id", "", "The ID of the Amazon Web Services Site-to-Site VPN connection.")
	ec2_modifyVpnTunnelCertificateCmd.Flags().String("vpn-tunnel-outside-ip-address", "", "The external IP address of the VPN tunnel.")
	ec2_modifyVpnTunnelCertificateCmd.Flag("no-dry-run").Hidden = true
	ec2_modifyVpnTunnelCertificateCmd.MarkFlagRequired("vpn-connection-id")
	ec2_modifyVpnTunnelCertificateCmd.MarkFlagRequired("vpn-tunnel-outside-ip-address")
	ec2Cmd.AddCommand(ec2_modifyVpnTunnelCertificateCmd)
}
