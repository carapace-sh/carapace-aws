package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rtbfabric_deleteRequesterGatewayCmd = &cobra.Command{
	Use:   "delete-requester-gateway",
	Short: "Deletes a requester gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rtbfabric_deleteRequesterGatewayCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rtbfabric_deleteRequesterGatewayCmd).Standalone()

		rtbfabric_deleteRequesterGatewayCmd.Flags().String("gateway-id", "", "The unique identifier of the gateway.")
		rtbfabric_deleteRequesterGatewayCmd.MarkFlagRequired("gateway-id")
	})
	rtbfabricCmd.AddCommand(rtbfabric_deleteRequesterGatewayCmd)
}
