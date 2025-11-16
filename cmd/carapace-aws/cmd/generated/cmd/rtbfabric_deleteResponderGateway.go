package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rtbfabric_deleteResponderGatewayCmd = &cobra.Command{
	Use:   "delete-responder-gateway",
	Short: "Deletes a responder gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rtbfabric_deleteResponderGatewayCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rtbfabric_deleteResponderGatewayCmd).Standalone()

		rtbfabric_deleteResponderGatewayCmd.Flags().String("gateway-id", "", "The unique identifier of the gateway.")
		rtbfabric_deleteResponderGatewayCmd.MarkFlagRequired("gateway-id")
	})
	rtbfabricCmd.AddCommand(rtbfabric_deleteResponderGatewayCmd)
}
