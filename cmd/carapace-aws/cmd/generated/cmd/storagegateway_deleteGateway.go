package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_deleteGatewayCmd = &cobra.Command{
	Use:   "delete-gateway",
	Short: "Deletes a gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_deleteGatewayCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_deleteGatewayCmd).Standalone()

		storagegateway_deleteGatewayCmd.Flags().String("gateway-arn", "", "")
		storagegateway_deleteGatewayCmd.MarkFlagRequired("gateway-arn")
	})
	storagegatewayCmd.AddCommand(storagegateway_deleteGatewayCmd)
}
