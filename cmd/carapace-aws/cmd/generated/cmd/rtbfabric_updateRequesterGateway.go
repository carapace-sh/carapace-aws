package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rtbfabric_updateRequesterGatewayCmd = &cobra.Command{
	Use:   "update-requester-gateway",
	Short: "Updates a requester gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rtbfabric_updateRequesterGatewayCmd).Standalone()

	rtbfabric_updateRequesterGatewayCmd.Flags().String("client-token", "", "The unique client token.")
	rtbfabric_updateRequesterGatewayCmd.Flags().String("description", "", "An optional description for the requester gateway.")
	rtbfabric_updateRequesterGatewayCmd.Flags().String("gateway-id", "", "The unique identifier of the gateway.")
	rtbfabric_updateRequesterGatewayCmd.MarkFlagRequired("client-token")
	rtbfabric_updateRequesterGatewayCmd.MarkFlagRequired("gateway-id")
	rtbfabricCmd.AddCommand(rtbfabric_updateRequesterGatewayCmd)
}
