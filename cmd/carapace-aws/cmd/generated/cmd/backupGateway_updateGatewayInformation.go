package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backupGateway_updateGatewayInformationCmd = &cobra.Command{
	Use:   "update-gateway-information",
	Short: "Updates a gateway's name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backupGateway_updateGatewayInformationCmd).Standalone()

	backupGateway_updateGatewayInformationCmd.Flags().String("gateway-arn", "", "The Amazon Resource Name (ARN) of the gateway to update.")
	backupGateway_updateGatewayInformationCmd.Flags().String("gateway-display-name", "", "The updated display name of the gateway.")
	backupGateway_updateGatewayInformationCmd.MarkFlagRequired("gateway-arn")
	backupGatewayCmd.AddCommand(backupGateway_updateGatewayInformationCmd)
}
