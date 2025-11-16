package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rtbfabric_updateResponderGatewayCmd = &cobra.Command{
	Use:   "update-responder-gateway",
	Short: "Updates a responder gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rtbfabric_updateResponderGatewayCmd).Standalone()

	rtbfabric_updateResponderGatewayCmd.Flags().String("client-token", "", "The unique client token.")
	rtbfabric_updateResponderGatewayCmd.Flags().String("description", "", "An optional description for the responder gateway.")
	rtbfabric_updateResponderGatewayCmd.Flags().String("domain-name", "", "The domain name for the responder gateway.")
	rtbfabric_updateResponderGatewayCmd.Flags().String("gateway-id", "", "The unique identifier of the gateway.")
	rtbfabric_updateResponderGatewayCmd.Flags().String("managed-endpoint-configuration", "", "The configuration for the managed endpoint.")
	rtbfabric_updateResponderGatewayCmd.Flags().String("port", "", "The networking port to use.")
	rtbfabric_updateResponderGatewayCmd.Flags().String("protocol", "", "The networking protocol to use.")
	rtbfabric_updateResponderGatewayCmd.Flags().String("trust-store-configuration", "", "The configuration of the trust store.")
	rtbfabric_updateResponderGatewayCmd.MarkFlagRequired("client-token")
	rtbfabric_updateResponderGatewayCmd.MarkFlagRequired("gateway-id")
	rtbfabric_updateResponderGatewayCmd.MarkFlagRequired("port")
	rtbfabric_updateResponderGatewayCmd.MarkFlagRequired("protocol")
	rtbfabricCmd.AddCommand(rtbfabric_updateResponderGatewayCmd)
}
