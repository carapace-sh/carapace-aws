package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rtbfabric_getResponderGatewayCmd = &cobra.Command{
	Use:   "get-responder-gateway",
	Short: "Retrieves information about a responder gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rtbfabric_getResponderGatewayCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rtbfabric_getResponderGatewayCmd).Standalone()

		rtbfabric_getResponderGatewayCmd.Flags().String("gateway-id", "", "The unique identifier of the gateway.")
		rtbfabric_getResponderGatewayCmd.MarkFlagRequired("gateway-id")
	})
	rtbfabricCmd.AddCommand(rtbfabric_getResponderGatewayCmd)
}
