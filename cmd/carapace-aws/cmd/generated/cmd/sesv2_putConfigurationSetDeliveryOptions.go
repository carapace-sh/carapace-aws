package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_putConfigurationSetDeliveryOptionsCmd = &cobra.Command{
	Use:   "put-configuration-set-delivery-options",
	Short: "Associate a configuration set with a dedicated IP pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_putConfigurationSetDeliveryOptionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_putConfigurationSetDeliveryOptionsCmd).Standalone()

		sesv2_putConfigurationSetDeliveryOptionsCmd.Flags().String("configuration-set-name", "", "The name of the configuration set to associate with a dedicated IP pool.")
		sesv2_putConfigurationSetDeliveryOptionsCmd.Flags().String("max-delivery-seconds", "", "The maximum amount of time, in seconds, that Amazon SES API v2 will attempt delivery of email.")
		sesv2_putConfigurationSetDeliveryOptionsCmd.Flags().String("sending-pool-name", "", "The name of the dedicated IP pool to associate with the configuration set.")
		sesv2_putConfigurationSetDeliveryOptionsCmd.Flags().String("tls-policy", "", "Specifies whether messages that use the configuration set are required to use Transport Layer Security (TLS).")
		sesv2_putConfigurationSetDeliveryOptionsCmd.MarkFlagRequired("configuration-set-name")
	})
	sesv2Cmd.AddCommand(sesv2_putConfigurationSetDeliveryOptionsCmd)
}
