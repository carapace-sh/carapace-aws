package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backupGateway_updateGatewaySoftwareNowCmd = &cobra.Command{
	Use:   "update-gateway-software-now",
	Short: "Updates the gateway virtual machine (VM) software.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backupGateway_updateGatewaySoftwareNowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backupGateway_updateGatewaySoftwareNowCmd).Standalone()

		backupGateway_updateGatewaySoftwareNowCmd.Flags().String("gateway-arn", "", "The Amazon Resource Name (ARN) of the gateway to be updated.")
		backupGateway_updateGatewaySoftwareNowCmd.MarkFlagRequired("gateway-arn")
	})
	backupGatewayCmd.AddCommand(backupGateway_updateGatewaySoftwareNowCmd)
}
