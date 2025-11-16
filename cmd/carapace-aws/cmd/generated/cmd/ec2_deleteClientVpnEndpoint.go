package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteClientVpnEndpointCmd = &cobra.Command{
	Use:   "delete-client-vpn-endpoint",
	Short: "Deletes the specified Client VPN endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteClientVpnEndpointCmd).Standalone()

	ec2_deleteClientVpnEndpointCmd.Flags().String("client-vpn-endpoint-id", "", "The ID of the Client VPN to be deleted.")
	ec2_deleteClientVpnEndpointCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteClientVpnEndpointCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteClientVpnEndpointCmd.MarkFlagRequired("client-vpn-endpoint-id")
	ec2_deleteClientVpnEndpointCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_deleteClientVpnEndpointCmd)
}
