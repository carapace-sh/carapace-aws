package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_deregisterTransitGatewayCmd = &cobra.Command{
	Use:   "deregister-transit-gateway",
	Short: "Deregisters a transit gateway from your global network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_deregisterTransitGatewayCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_deregisterTransitGatewayCmd).Standalone()

		networkmanager_deregisterTransitGatewayCmd.Flags().String("global-network-id", "", "The ID of the global network.")
		networkmanager_deregisterTransitGatewayCmd.Flags().String("transit-gateway-arn", "", "The Amazon Resource Name (ARN) of the transit gateway.")
		networkmanager_deregisterTransitGatewayCmd.MarkFlagRequired("global-network-id")
		networkmanager_deregisterTransitGatewayCmd.MarkFlagRequired("transit-gateway-arn")
	})
	networkmanagerCmd.AddCommand(networkmanager_deregisterTransitGatewayCmd)
}
