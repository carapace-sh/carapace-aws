package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_disableGatewayCmd = &cobra.Command{
	Use:   "disable-gateway",
	Short: "Disables a tape gateway when the gateway is no longer functioning.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_disableGatewayCmd).Standalone()

	storagegateway_disableGatewayCmd.Flags().String("gateway-arn", "", "")
	storagegateway_disableGatewayCmd.MarkFlagRequired("gateway-arn")
	storagegatewayCmd.AddCommand(storagegateway_disableGatewayCmd)
}
