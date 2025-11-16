package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_activateGatewayCmd = &cobra.Command{
	Use:   "activate-gateway",
	Short: "Activates the gateway you previously deployed on your host.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_activateGatewayCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_activateGatewayCmd).Standalone()

		storagegateway_activateGatewayCmd.Flags().String("activation-key", "", "Your gateway activation key.")
		storagegateway_activateGatewayCmd.Flags().String("gateway-name", "", "The name you configured for your gateway.")
		storagegateway_activateGatewayCmd.Flags().String("gateway-region", "", "A value that indicates the Amazon Web Services Region where you want to store your data.")
		storagegateway_activateGatewayCmd.Flags().String("gateway-timezone", "", "A value that indicates the time zone you want to set for the gateway.")
		storagegateway_activateGatewayCmd.Flags().String("gateway-type", "", "A value that defines the type of gateway to activate.")
		storagegateway_activateGatewayCmd.Flags().String("medium-changer-type", "", "The value that indicates the type of medium changer to use for tape gateway.")
		storagegateway_activateGatewayCmd.Flags().String("tags", "", "A list of up to 50 tags that you can assign to the gateway.")
		storagegateway_activateGatewayCmd.Flags().String("tape-drive-type", "", "The value that indicates the type of tape drive to use for tape gateway.")
		storagegateway_activateGatewayCmd.MarkFlagRequired("activation-key")
		storagegateway_activateGatewayCmd.MarkFlagRequired("gateway-name")
		storagegateway_activateGatewayCmd.MarkFlagRequired("gateway-region")
		storagegateway_activateGatewayCmd.MarkFlagRequired("gateway-timezone")
	})
	storagegatewayCmd.AddCommand(storagegateway_activateGatewayCmd)
}
