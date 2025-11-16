package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backupGateway_associateGatewayToServerCmd = &cobra.Command{
	Use:   "associate-gateway-to-server",
	Short: "Associates a backup gateway with your server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backupGateway_associateGatewayToServerCmd).Standalone()

	backupGateway_associateGatewayToServerCmd.Flags().String("gateway-arn", "", "The Amazon Resource Name (ARN) of the gateway.")
	backupGateway_associateGatewayToServerCmd.Flags().String("server-arn", "", "The Amazon Resource Name (ARN) of the server that hosts your virtual machines.")
	backupGateway_associateGatewayToServerCmd.MarkFlagRequired("gateway-arn")
	backupGateway_associateGatewayToServerCmd.MarkFlagRequired("server-arn")
	backupGatewayCmd.AddCommand(backupGateway_associateGatewayToServerCmd)
}
