package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointEmail_putConfigurationSetDeliveryOptionsCmd = &cobra.Command{
	Use:   "put-configuration-set-delivery-options",
	Short: "Associate a configuration set with a dedicated IP pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointEmail_putConfigurationSetDeliveryOptionsCmd).Standalone()

	pinpointEmail_putConfigurationSetDeliveryOptionsCmd.Flags().String("configuration-set-name", "", "The name of the configuration set that you want to associate with a dedicated IP pool.")
	pinpointEmail_putConfigurationSetDeliveryOptionsCmd.Flags().String("sending-pool-name", "", "The name of the dedicated IP pool that you want to associate with the configuration set.")
	pinpointEmail_putConfigurationSetDeliveryOptionsCmd.Flags().String("tls-policy", "", "Specifies whether messages that use the configuration set are required to use Transport Layer Security (TLS).")
	pinpointEmail_putConfigurationSetDeliveryOptionsCmd.MarkFlagRequired("configuration-set-name")
	pinpointEmailCmd.AddCommand(pinpointEmail_putConfigurationSetDeliveryOptionsCmd)
}
