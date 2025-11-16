package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_putDeliveryChannelCmd = &cobra.Command{
	Use:   "put-delivery-channel",
	Short: "Creates or updates a delivery channel to deliver configuration information and other compliance information.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_putDeliveryChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_putDeliveryChannelCmd).Standalone()

		config_putDeliveryChannelCmd.Flags().String("delivery-channel", "", "An object for the delivery channel.")
		config_putDeliveryChannelCmd.MarkFlagRequired("delivery-channel")
	})
	configCmd.AddCommand(config_putDeliveryChannelCmd)
}
