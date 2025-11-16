package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyVpnConnectionCmd = &cobra.Command{
	Use:   "modify-vpn-connection",
	Short: "Modifies the customer gateway or the target gateway of an Amazon Web Services Site-to-Site VPN connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyVpnConnectionCmd).Standalone()

	ec2_modifyVpnConnectionCmd.Flags().String("customer-gateway-id", "", "The ID of the customer gateway at your end of the VPN connection.")
	ec2_modifyVpnConnectionCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_modifyVpnConnectionCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_modifyVpnConnectionCmd.Flags().String("transit-gateway-id", "", "The ID of the transit gateway.")
	ec2_modifyVpnConnectionCmd.Flags().String("vpn-connection-id", "", "The ID of the VPN connection.")
	ec2_modifyVpnConnectionCmd.Flags().String("vpn-gateway-id", "", "The ID of the virtual private gateway at the Amazon Web Services side of the VPN connection.")
	ec2_modifyVpnConnectionCmd.Flag("no-dry-run").Hidden = true
	ec2_modifyVpnConnectionCmd.MarkFlagRequired("vpn-connection-id")
	ec2Cmd.AddCommand(ec2_modifyVpnConnectionCmd)
}
