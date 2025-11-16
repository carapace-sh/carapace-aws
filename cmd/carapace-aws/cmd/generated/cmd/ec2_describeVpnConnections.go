package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeVpnConnectionsCmd = &cobra.Command{
	Use:   "describe-vpn-connections",
	Short: "Describes one or more of your VPN connections.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeVpnConnectionsCmd).Standalone()

	ec2_describeVpnConnectionsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeVpnConnectionsCmd.Flags().String("filters", "", "One or more filters.")
	ec2_describeVpnConnectionsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeVpnConnectionsCmd.Flags().String("vpn-connection-ids", "", "One or more VPN connection IDs.")
	ec2_describeVpnConnectionsCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeVpnConnectionsCmd)
}
