package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backupGateway_deleteGatewayCmd = &cobra.Command{
	Use:   "delete-gateway",
	Short: "Deletes a backup gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backupGateway_deleteGatewayCmd).Standalone()

	backupGateway_deleteGatewayCmd.Flags().String("gateway-arn", "", "The Amazon Resource Name (ARN) of the gateway to delete.")
	backupGateway_deleteGatewayCmd.MarkFlagRequired("gateway-arn")
	backupGatewayCmd.AddCommand(backupGateway_deleteGatewayCmd)
}
