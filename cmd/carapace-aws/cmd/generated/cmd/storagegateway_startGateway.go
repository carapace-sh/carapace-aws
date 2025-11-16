package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_startGatewayCmd = &cobra.Command{
	Use:   "start-gateway",
	Short: "Starts a gateway that you previously shut down (see [ShutdownGateway]()).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_startGatewayCmd).Standalone()

	storagegateway_startGatewayCmd.Flags().String("gateway-arn", "", "")
	storagegateway_startGatewayCmd.MarkFlagRequired("gateway-arn")
	storagegatewayCmd.AddCommand(storagegateway_startGatewayCmd)
}
