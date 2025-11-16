package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rtbfabric_getRequesterGatewayCmd = &cobra.Command{
	Use:   "get-requester-gateway",
	Short: "Retrieves information about a requester gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rtbfabric_getRequesterGatewayCmd).Standalone()

	rtbfabric_getRequesterGatewayCmd.Flags().String("gateway-id", "", "The unique identifier of the gateway.")
	rtbfabric_getRequesterGatewayCmd.MarkFlagRequired("gateway-id")
	rtbfabricCmd.AddCommand(rtbfabric_getRequesterGatewayCmd)
}
