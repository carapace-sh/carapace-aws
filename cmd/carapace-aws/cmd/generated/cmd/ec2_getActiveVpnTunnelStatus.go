package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getActiveVpnTunnelStatusCmd = &cobra.Command{
	Use:   "get-active-vpn-tunnel-status",
	Short: "Returns the currently negotiated security parameters for an active VPN tunnel, including IKE version, DH groups, encryption algorithms, and integrity algorithms.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getActiveVpnTunnelStatusCmd).Standalone()

	ec2_getActiveVpnTunnelStatusCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request.")
	ec2_getActiveVpnTunnelStatusCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request.")
	ec2_getActiveVpnTunnelStatusCmd.Flags().String("vpn-connection-id", "", "The ID of the VPN connection for which to retrieve the active tunnel status.")
	ec2_getActiveVpnTunnelStatusCmd.Flags().String("vpn-tunnel-outside-ip-address", "", "The external IP address of the VPN tunnel for which to retrieve the active status.")
	ec2_getActiveVpnTunnelStatusCmd.Flag("no-dry-run").Hidden = true
	ec2_getActiveVpnTunnelStatusCmd.MarkFlagRequired("vpn-connection-id")
	ec2_getActiveVpnTunnelStatusCmd.MarkFlagRequired("vpn-tunnel-outside-ip-address")
	ec2Cmd.AddCommand(ec2_getActiveVpnTunnelStatusCmd)
}
