package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createVpnConnectionRouteCmd = &cobra.Command{
	Use:   "create-vpn-connection-route",
	Short: "Creates a static route associated with a VPN connection between an existing virtual private gateway and a VPN customer gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createVpnConnectionRouteCmd).Standalone()

	ec2_createVpnConnectionRouteCmd.Flags().String("destination-cidr-block", "", "The CIDR block associated with the local subnet of the customer network.")
	ec2_createVpnConnectionRouteCmd.Flags().String("vpn-connection-id", "", "The ID of the VPN connection.")
	ec2_createVpnConnectionRouteCmd.MarkFlagRequired("destination-cidr-block")
	ec2_createVpnConnectionRouteCmd.MarkFlagRequired("vpn-connection-id")
	ec2Cmd.AddCommand(ec2_createVpnConnectionRouteCmd)
}
