package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_describeDeliveryChannelStatusCmd = &cobra.Command{
	Use:   "describe-delivery-channel-status",
	Short: "Returns the current status of the specified delivery channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_describeDeliveryChannelStatusCmd).Standalone()

	config_describeDeliveryChannelStatusCmd.Flags().String("delivery-channel-names", "", "A list of delivery channel names.")
	configCmd.AddCommand(config_describeDeliveryChannelStatusCmd)
}
