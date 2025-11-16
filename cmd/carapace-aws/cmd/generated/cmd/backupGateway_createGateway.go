package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backupGateway_createGatewayCmd = &cobra.Command{
	Use:   "create-gateway",
	Short: "Creates a backup gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backupGateway_createGatewayCmd).Standalone()

	backupGateway_createGatewayCmd.Flags().String("activation-key", "", "The activation key of the created gateway.")
	backupGateway_createGatewayCmd.Flags().String("gateway-display-name", "", "The display name of the created gateway.")
	backupGateway_createGatewayCmd.Flags().String("gateway-type", "", "The type of created gateway.")
	backupGateway_createGatewayCmd.Flags().String("tags", "", "A list of up to 50 tags to assign to the gateway.")
	backupGateway_createGatewayCmd.MarkFlagRequired("activation-key")
	backupGateway_createGatewayCmd.MarkFlagRequired("gateway-display-name")
	backupGateway_createGatewayCmd.MarkFlagRequired("gateway-type")
	backupGatewayCmd.AddCommand(backupGateway_createGatewayCmd)
}
