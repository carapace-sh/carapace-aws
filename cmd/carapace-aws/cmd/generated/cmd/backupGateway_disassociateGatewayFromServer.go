package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backupGateway_disassociateGatewayFromServerCmd = &cobra.Command{
	Use:   "disassociate-gateway-from-server",
	Short: "Disassociates a backup gateway from the specified server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backupGateway_disassociateGatewayFromServerCmd).Standalone()

	backupGateway_disassociateGatewayFromServerCmd.Flags().String("gateway-arn", "", "The Amazon Resource Name (ARN) of the gateway to disassociate.")
	backupGateway_disassociateGatewayFromServerCmd.MarkFlagRequired("gateway-arn")
	backupGatewayCmd.AddCommand(backupGateway_disassociateGatewayFromServerCmd)
}
