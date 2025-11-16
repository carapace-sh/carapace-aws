package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_registerTransitGatewayCmd = &cobra.Command{
	Use:   "register-transit-gateway",
	Short: "Registers a transit gateway in your global network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_registerTransitGatewayCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_registerTransitGatewayCmd).Standalone()

		networkmanager_registerTransitGatewayCmd.Flags().String("global-network-id", "", "The ID of the global network.")
		networkmanager_registerTransitGatewayCmd.Flags().String("transit-gateway-arn", "", "The Amazon Resource Name (ARN) of the transit gateway.")
		networkmanager_registerTransitGatewayCmd.MarkFlagRequired("global-network-id")
		networkmanager_registerTransitGatewayCmd.MarkFlagRequired("transit-gateway-arn")
	})
	networkmanagerCmd.AddCommand(networkmanager_registerTransitGatewayCmd)
}
