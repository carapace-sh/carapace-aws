package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_replaceVpnTunnelCmd = &cobra.Command{
	Use:   "replace-vpn-tunnel",
	Short: "Trigger replacement of specified VPN tunnel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_replaceVpnTunnelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_replaceVpnTunnelCmd).Standalone()

		ec2_replaceVpnTunnelCmd.Flags().Bool("apply-pending-maintenance", false, "Trigger pending tunnel endpoint maintenance.")
		ec2_replaceVpnTunnelCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_replaceVpnTunnelCmd.Flags().Bool("no-apply-pending-maintenance", false, "Trigger pending tunnel endpoint maintenance.")
		ec2_replaceVpnTunnelCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_replaceVpnTunnelCmd.Flags().String("vpn-connection-id", "", "The ID of the Site-to-Site VPN connection.")
		ec2_replaceVpnTunnelCmd.Flags().String("vpn-tunnel-outside-ip-address", "", "The external IP address of the VPN tunnel.")
		ec2_replaceVpnTunnelCmd.Flag("no-apply-pending-maintenance").Hidden = true
		ec2_replaceVpnTunnelCmd.Flag("no-dry-run").Hidden = true
		ec2_replaceVpnTunnelCmd.MarkFlagRequired("vpn-connection-id")
		ec2_replaceVpnTunnelCmd.MarkFlagRequired("vpn-tunnel-outside-ip-address")
	})
	ec2Cmd.AddCommand(ec2_replaceVpnTunnelCmd)
}
