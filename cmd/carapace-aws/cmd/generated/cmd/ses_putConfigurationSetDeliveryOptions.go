package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_putConfigurationSetDeliveryOptionsCmd = &cobra.Command{
	Use:   "put-configuration-set-delivery-options",
	Short: "Adds or updates the delivery options for a configuration set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_putConfigurationSetDeliveryOptionsCmd).Standalone()

	ses_putConfigurationSetDeliveryOptionsCmd.Flags().String("configuration-set-name", "", "The name of the configuration set.")
	ses_putConfigurationSetDeliveryOptionsCmd.Flags().String("delivery-options", "", "Specifies whether messages that use the configuration set are required to use Transport Layer Security (TLS).")
	ses_putConfigurationSetDeliveryOptionsCmd.MarkFlagRequired("configuration-set-name")
	sesCmd.AddCommand(ses_putConfigurationSetDeliveryOptionsCmd)
}
