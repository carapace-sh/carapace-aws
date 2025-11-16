package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rtbfabric_createResponderGatewayCmd = &cobra.Command{
	Use:   "create-responder-gateway",
	Short: "Creates a responder gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rtbfabric_createResponderGatewayCmd).Standalone()

	rtbfabric_createResponderGatewayCmd.Flags().String("client-token", "", "The unique client token.")
	rtbfabric_createResponderGatewayCmd.Flags().String("description", "", "An optional description for the responder gateway.")
	rtbfabric_createResponderGatewayCmd.Flags().String("domain-name", "", "The domain name for the responder gateway.")
	rtbfabric_createResponderGatewayCmd.Flags().String("managed-endpoint-configuration", "", "The configuration for the managed endpoint.")
	rtbfabric_createResponderGatewayCmd.Flags().String("port", "", "The networking port to use.")
	rtbfabric_createResponderGatewayCmd.Flags().String("protocol", "", "The networking protocol to use.")
	rtbfabric_createResponderGatewayCmd.Flags().String("security-group-ids", "", "The unique identifiers of the security groups.")
	rtbfabric_createResponderGatewayCmd.Flags().String("subnet-ids", "", "The unique identifiers of the subnets.")
	rtbfabric_createResponderGatewayCmd.Flags().String("tags", "", "A map of the key-value pairs of the tag or tags to assign to the resource.")
	rtbfabric_createResponderGatewayCmd.Flags().String("trust-store-configuration", "", "The configuration of the trust store.")
	rtbfabric_createResponderGatewayCmd.Flags().String("vpc-id", "", "The unique identifier of the Virtual Private Cloud (VPC).")
	rtbfabric_createResponderGatewayCmd.MarkFlagRequired("client-token")
	rtbfabric_createResponderGatewayCmd.MarkFlagRequired("port")
	rtbfabric_createResponderGatewayCmd.MarkFlagRequired("protocol")
	rtbfabric_createResponderGatewayCmd.MarkFlagRequired("security-group-ids")
	rtbfabric_createResponderGatewayCmd.MarkFlagRequired("subnet-ids")
	rtbfabric_createResponderGatewayCmd.MarkFlagRequired("vpc-id")
	rtbfabricCmd.AddCommand(rtbfabric_createResponderGatewayCmd)
}
