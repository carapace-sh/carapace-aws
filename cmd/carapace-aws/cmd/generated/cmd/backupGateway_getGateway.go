package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backupGateway_getGatewayCmd = &cobra.Command{
	Use:   "get-gateway",
	Short: "By providing the ARN (Amazon Resource Name), this API returns the gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backupGateway_getGatewayCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backupGateway_getGatewayCmd).Standalone()

		backupGateway_getGatewayCmd.Flags().String("gateway-arn", "", "The Amazon Resource Name (ARN) of the gateway.")
		backupGateway_getGatewayCmd.MarkFlagRequired("gateway-arn")
	})
	backupGatewayCmd.AddCommand(backupGateway_getGatewayCmd)
}
