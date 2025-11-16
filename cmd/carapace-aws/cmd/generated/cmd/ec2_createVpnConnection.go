package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createVpnConnectionCmd = &cobra.Command{
	Use:   "create-vpn-connection",
	Short: "Creates a VPN connection between an existing virtual private gateway or transit gateway and a customer gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createVpnConnectionCmd).Standalone()

	ec2_createVpnConnectionCmd.Flags().String("customer-gateway-id", "", "The ID of the customer gateway.")
	ec2_createVpnConnectionCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createVpnConnectionCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createVpnConnectionCmd.Flags().String("options", "", "The options for the VPN connection.")
	ec2_createVpnConnectionCmd.Flags().String("pre-shared-key-storage", "", "Specifies the storage mode for the pre-shared key (PSK).")
	ec2_createVpnConnectionCmd.Flags().String("tag-specifications", "", "The tags to apply to the VPN connection.")
	ec2_createVpnConnectionCmd.Flags().String("transit-gateway-id", "", "The ID of the transit gateway.")
	ec2_createVpnConnectionCmd.Flags().String("type", "", "The type of VPN connection (`ipsec.1`).")
	ec2_createVpnConnectionCmd.Flags().String("vpn-gateway-id", "", "The ID of the virtual private gateway.")
	ec2_createVpnConnectionCmd.MarkFlagRequired("customer-gateway-id")
	ec2_createVpnConnectionCmd.Flag("no-dry-run").Hidden = true
	ec2_createVpnConnectionCmd.MarkFlagRequired("type")
	ec2Cmd.AddCommand(ec2_createVpnConnectionCmd)
}
