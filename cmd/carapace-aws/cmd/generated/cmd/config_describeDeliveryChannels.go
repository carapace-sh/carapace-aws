package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_describeDeliveryChannelsCmd = &cobra.Command{
	Use:   "describe-delivery-channels",
	Short: "Returns details about the specified delivery channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_describeDeliveryChannelsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_describeDeliveryChannelsCmd).Standalone()

		config_describeDeliveryChannelsCmd.Flags().String("delivery-channel-names", "", "A list of delivery channel names.")
	})
	configCmd.AddCommand(config_describeDeliveryChannelsCmd)
}
