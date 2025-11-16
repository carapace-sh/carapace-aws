package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_deleteDeliveryChannelCmd = &cobra.Command{
	Use:   "delete-delivery-channel",
	Short: "Deletes the delivery channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_deleteDeliveryChannelCmd).Standalone()

	config_deleteDeliveryChannelCmd.Flags().String("delivery-channel-name", "", "The name of the delivery channel that you want to delete.")
	config_deleteDeliveryChannelCmd.MarkFlagRequired("delivery-channel-name")
	configCmd.AddCommand(config_deleteDeliveryChannelCmd)
}
