package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_getCustomerGatewayAssociationsCmd = &cobra.Command{
	Use:   "get-customer-gateway-associations",
	Short: "Gets the association information for customer gateways that are associated with devices and links in your global network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_getCustomerGatewayAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_getCustomerGatewayAssociationsCmd).Standalone()

		networkmanager_getCustomerGatewayAssociationsCmd.Flags().String("customer-gateway-arns", "", "One or more customer gateway Amazon Resource Names (ARNs).")
		networkmanager_getCustomerGatewayAssociationsCmd.Flags().String("global-network-id", "", "The ID of the global network.")
		networkmanager_getCustomerGatewayAssociationsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		networkmanager_getCustomerGatewayAssociationsCmd.Flags().String("next-token", "", "The token for the next page of results.")
		networkmanager_getCustomerGatewayAssociationsCmd.MarkFlagRequired("global-network-id")
	})
	networkmanagerCmd.AddCommand(networkmanager_getCustomerGatewayAssociationsCmd)
}
