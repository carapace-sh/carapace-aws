package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getVpnTunnelReplacementStatusCmd = &cobra.Command{
	Use:   "get-vpn-tunnel-replacement-status",
	Short: "Get details of available tunnel endpoint maintenance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getVpnTunnelReplacementStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_getVpnTunnelReplacementStatusCmd).Standalone()

		ec2_getVpnTunnelReplacementStatusCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getVpnTunnelReplacementStatusCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getVpnTunnelReplacementStatusCmd.Flags().String("vpn-connection-id", "", "The ID of the Site-to-Site VPN connection.")
		ec2_getVpnTunnelReplacementStatusCmd.Flags().String("vpn-tunnel-outside-ip-address", "", "The external IP address of the VPN tunnel.")
		ec2_getVpnTunnelReplacementStatusCmd.Flag("no-dry-run").Hidden = true
		ec2_getVpnTunnelReplacementStatusCmd.MarkFlagRequired("vpn-connection-id")
		ec2_getVpnTunnelReplacementStatusCmd.MarkFlagRequired("vpn-tunnel-outside-ip-address")
	})
	ec2Cmd.AddCommand(ec2_getVpnTunnelReplacementStatusCmd)
}
