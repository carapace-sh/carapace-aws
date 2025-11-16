package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteVpnConnectionRouteCmd = &cobra.Command{
	Use:   "delete-vpn-connection-route",
	Short: "Deletes the specified static route associated with a VPN connection between an existing virtual private gateway and a VPN customer gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteVpnConnectionRouteCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteVpnConnectionRouteCmd).Standalone()

		ec2_deleteVpnConnectionRouteCmd.Flags().String("destination-cidr-block", "", "The CIDR block associated with the local subnet of the customer network.")
		ec2_deleteVpnConnectionRouteCmd.Flags().String("vpn-connection-id", "", "The ID of the VPN connection.")
		ec2_deleteVpnConnectionRouteCmd.MarkFlagRequired("destination-cidr-block")
		ec2_deleteVpnConnectionRouteCmd.MarkFlagRequired("vpn-connection-id")
	})
	ec2Cmd.AddCommand(ec2_deleteVpnConnectionRouteCmd)
}
