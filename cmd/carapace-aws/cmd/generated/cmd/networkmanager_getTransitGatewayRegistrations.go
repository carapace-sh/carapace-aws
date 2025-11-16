package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_getTransitGatewayRegistrationsCmd = &cobra.Command{
	Use:   "get-transit-gateway-registrations",
	Short: "Gets information about the transit gateway registrations in a specified global network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_getTransitGatewayRegistrationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_getTransitGatewayRegistrationsCmd).Standalone()

		networkmanager_getTransitGatewayRegistrationsCmd.Flags().String("global-network-id", "", "The ID of the global network.")
		networkmanager_getTransitGatewayRegistrationsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		networkmanager_getTransitGatewayRegistrationsCmd.Flags().String("next-token", "", "The token for the next page of results.")
		networkmanager_getTransitGatewayRegistrationsCmd.Flags().String("transit-gateway-arns", "", "The Amazon Resource Names (ARNs) of one or more transit gateways.")
		networkmanager_getTransitGatewayRegistrationsCmd.MarkFlagRequired("global-network-id")
	})
	networkmanagerCmd.AddCommand(networkmanager_getTransitGatewayRegistrationsCmd)
}
