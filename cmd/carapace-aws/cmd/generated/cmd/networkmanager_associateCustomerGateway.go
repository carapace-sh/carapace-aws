package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_associateCustomerGatewayCmd = &cobra.Command{
	Use:   "associate-customer-gateway",
	Short: "Associates a customer gateway with a device and optionally, with a link.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_associateCustomerGatewayCmd).Standalone()

	networkmanager_associateCustomerGatewayCmd.Flags().String("customer-gateway-arn", "", "The Amazon Resource Name (ARN) of the customer gateway.")
	networkmanager_associateCustomerGatewayCmd.Flags().String("device-id", "", "The ID of the device.")
	networkmanager_associateCustomerGatewayCmd.Flags().String("global-network-id", "", "The ID of the global network.")
	networkmanager_associateCustomerGatewayCmd.Flags().String("link-id", "", "The ID of the link.")
	networkmanager_associateCustomerGatewayCmd.MarkFlagRequired("customer-gateway-arn")
	networkmanager_associateCustomerGatewayCmd.MarkFlagRequired("device-id")
	networkmanager_associateCustomerGatewayCmd.MarkFlagRequired("global-network-id")
	networkmanagerCmd.AddCommand(networkmanager_associateCustomerGatewayCmd)
}
