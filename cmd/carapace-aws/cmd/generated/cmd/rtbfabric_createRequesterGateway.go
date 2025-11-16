package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rtbfabric_createRequesterGatewayCmd = &cobra.Command{
	Use:   "create-requester-gateway",
	Short: "Creates a requester gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rtbfabric_createRequesterGatewayCmd).Standalone()

	rtbfabric_createRequesterGatewayCmd.Flags().String("client-token", "", "The unique client token.")
	rtbfabric_createRequesterGatewayCmd.Flags().String("description", "", "An optional description for the requester gateway.")
	rtbfabric_createRequesterGatewayCmd.Flags().String("security-group-ids", "", "The unique identifiers of the security groups.")
	rtbfabric_createRequesterGatewayCmd.Flags().String("subnet-ids", "", "The unique identifiers of the subnets.")
	rtbfabric_createRequesterGatewayCmd.Flags().String("tags", "", "A map of the key-value pairs of the tag or tags to assign to the resource.")
	rtbfabric_createRequesterGatewayCmd.Flags().String("vpc-id", "", "The unique identifier of the Virtual Private Cloud (VPC).")
	rtbfabric_createRequesterGatewayCmd.MarkFlagRequired("client-token")
	rtbfabric_createRequesterGatewayCmd.MarkFlagRequired("security-group-ids")
	rtbfabric_createRequesterGatewayCmd.MarkFlagRequired("subnet-ids")
	rtbfabric_createRequesterGatewayCmd.MarkFlagRequired("vpc-id")
	rtbfabricCmd.AddCommand(rtbfabric_createRequesterGatewayCmd)
}
