package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_terminateClientVpnConnectionsCmd = &cobra.Command{
	Use:   "terminate-client-vpn-connections",
	Short: "Terminates active Client VPN endpoint connections.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_terminateClientVpnConnectionsCmd).Standalone()

	ec2_terminateClientVpnConnectionsCmd.Flags().String("client-vpn-endpoint-id", "", "The ID of the Client VPN endpoint to which the client is connected.")
	ec2_terminateClientVpnConnectionsCmd.Flags().String("connection-id", "", "The ID of the client connection to be terminated.")
	ec2_terminateClientVpnConnectionsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_terminateClientVpnConnectionsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_terminateClientVpnConnectionsCmd.Flags().String("username", "", "The name of the user who initiated the connection.")
	ec2_terminateClientVpnConnectionsCmd.MarkFlagRequired("client-vpn-endpoint-id")
	ec2_terminateClientVpnConnectionsCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_terminateClientVpnConnectionsCmd)
}
