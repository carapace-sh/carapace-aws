package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_updateGatewaySoftwareNowCmd = &cobra.Command{
	Use:   "update-gateway-software-now",
	Short: "Updates the gateway virtual machine (VM) software.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_updateGatewaySoftwareNowCmd).Standalone()

	storagegateway_updateGatewaySoftwareNowCmd.Flags().String("gateway-arn", "", "")
	storagegateway_updateGatewaySoftwareNowCmd.MarkFlagRequired("gateway-arn")
	storagegatewayCmd.AddCommand(storagegateway_updateGatewaySoftwareNowCmd)
}
