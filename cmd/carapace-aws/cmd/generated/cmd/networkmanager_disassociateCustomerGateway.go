package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_disassociateCustomerGatewayCmd = &cobra.Command{
	Use:   "disassociate-customer-gateway",
	Short: "Disassociates a customer gateway from a device and a link.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_disassociateCustomerGatewayCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_disassociateCustomerGatewayCmd).Standalone()

		networkmanager_disassociateCustomerGatewayCmd.Flags().String("customer-gateway-arn", "", "The Amazon Resource Name (ARN) of the customer gateway.")
		networkmanager_disassociateCustomerGatewayCmd.Flags().String("global-network-id", "", "The ID of the global network.")
		networkmanager_disassociateCustomerGatewayCmd.MarkFlagRequired("customer-gateway-arn")
		networkmanager_disassociateCustomerGatewayCmd.MarkFlagRequired("global-network-id")
	})
	networkmanagerCmd.AddCommand(networkmanager_disassociateCustomerGatewayCmd)
}
