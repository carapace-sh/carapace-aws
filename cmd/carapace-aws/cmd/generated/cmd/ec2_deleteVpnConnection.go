package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteVpnConnectionCmd = &cobra.Command{
	Use:   "delete-vpn-connection",
	Short: "Deletes the specified VPN connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteVpnConnectionCmd).Standalone()

	ec2_deleteVpnConnectionCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteVpnConnectionCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteVpnConnectionCmd.Flags().String("vpn-connection-id", "", "The ID of the VPN connection.")
	ec2_deleteVpnConnectionCmd.Flag("no-dry-run").Hidden = true
	ec2_deleteVpnConnectionCmd.MarkFlagRequired("vpn-connection-id")
	ec2Cmd.AddCommand(ec2_deleteVpnConnectionCmd)
}
