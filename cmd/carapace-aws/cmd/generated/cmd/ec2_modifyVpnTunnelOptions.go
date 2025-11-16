package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyVpnTunnelOptionsCmd = &cobra.Command{
	Use:   "modify-vpn-tunnel-options",
	Short: "Modifies the options for a VPN tunnel in an Amazon Web Services Site-to-Site VPN connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyVpnTunnelOptionsCmd).Standalone()

	ec2_modifyVpnTunnelOptionsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_modifyVpnTunnelOptionsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_modifyVpnTunnelOptionsCmd.Flags().Bool("no-skip-tunnel-replacement", false, "Choose whether or not to trigger immediate tunnel replacement.")
	ec2_modifyVpnTunnelOptionsCmd.Flags().String("pre-shared-key-storage", "", "Specifies the storage mode for the pre-shared key (PSK).")
	ec2_modifyVpnTunnelOptionsCmd.Flags().Bool("skip-tunnel-replacement", false, "Choose whether or not to trigger immediate tunnel replacement.")
	ec2_modifyVpnTunnelOptionsCmd.Flags().String("tunnel-options", "", "The tunnel options to modify.")
	ec2_modifyVpnTunnelOptionsCmd.Flags().String("vpn-connection-id", "", "The ID of the Amazon Web Services Site-to-Site VPN connection.")
	ec2_modifyVpnTunnelOptionsCmd.Flags().String("vpn-tunnel-outside-ip-address", "", "The external IP address of the VPN tunnel.")
	ec2_modifyVpnTunnelOptionsCmd.Flag("no-dry-run").Hidden = true
	ec2_modifyVpnTunnelOptionsCmd.Flag("no-skip-tunnel-replacement").Hidden = true
	ec2_modifyVpnTunnelOptionsCmd.MarkFlagRequired("tunnel-options")
	ec2_modifyVpnTunnelOptionsCmd.MarkFlagRequired("vpn-connection-id")
	ec2_modifyVpnTunnelOptionsCmd.MarkFlagRequired("vpn-tunnel-outside-ip-address")
	ec2Cmd.AddCommand(ec2_modifyVpnTunnelOptionsCmd)
}
