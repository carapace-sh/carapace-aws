package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_shutdownGatewayCmd = &cobra.Command{
	Use:   "shutdown-gateway",
	Short: "Shuts down a Tape Gateway or Volume Gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_shutdownGatewayCmd).Standalone()

	storagegateway_shutdownGatewayCmd.Flags().String("gateway-arn", "", "")
	storagegateway_shutdownGatewayCmd.MarkFlagRequired("gateway-arn")
	storagegatewayCmd.AddCommand(storagegateway_shutdownGatewayCmd)
}
