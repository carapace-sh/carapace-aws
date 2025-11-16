package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_createDirectConnectGatewayCmd = &cobra.Command{
	Use:   "create-direct-connect-gateway",
	Short: "Creates a Direct Connect gateway, which is an intermediate object that enables you to connect a set of virtual interfaces and virtual private gateways.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_createDirectConnectGatewayCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(directconnect_createDirectConnectGatewayCmd).Standalone()

		directconnect_createDirectConnectGatewayCmd.Flags().String("amazon-side-asn", "", "The autonomous system number (ASN) for Border Gateway Protocol (BGP) to be configured on the Amazon side of the connection.")
		directconnect_createDirectConnectGatewayCmd.Flags().String("direct-connect-gateway-name", "", "The name of the Direct Connect gateway.")
		directconnect_createDirectConnectGatewayCmd.Flags().String("tags", "", "The key-value pair tags associated with the request.")
		directconnect_createDirectConnectGatewayCmd.MarkFlagRequired("direct-connect-gateway-name")
	})
	directconnectCmd.AddCommand(directconnect_createDirectConnectGatewayCmd)
}
