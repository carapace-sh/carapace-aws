package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_getTransitGatewayPeeringCmd = &cobra.Command{
	Use:   "get-transit-gateway-peering",
	Short: "Returns information about a transit gateway peer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_getTransitGatewayPeeringCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_getTransitGatewayPeeringCmd).Standalone()

		networkmanager_getTransitGatewayPeeringCmd.Flags().String("peering-id", "", "The ID of the peering request.")
		networkmanager_getTransitGatewayPeeringCmd.MarkFlagRequired("peering-id")
	})
	networkmanagerCmd.AddCommand(networkmanager_getTransitGatewayPeeringCmd)
}
