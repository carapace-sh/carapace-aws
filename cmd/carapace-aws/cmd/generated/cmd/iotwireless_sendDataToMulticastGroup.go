package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_sendDataToMulticastGroupCmd = &cobra.Command{
	Use:   "send-data-to-multicast-group",
	Short: "Sends the specified data to a multicast group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_sendDataToMulticastGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_sendDataToMulticastGroupCmd).Standalone()

		iotwireless_sendDataToMulticastGroupCmd.Flags().String("id", "", "")
		iotwireless_sendDataToMulticastGroupCmd.Flags().String("payload-data", "", "")
		iotwireless_sendDataToMulticastGroupCmd.Flags().String("wireless-metadata", "", "")
		iotwireless_sendDataToMulticastGroupCmd.MarkFlagRequired("id")
		iotwireless_sendDataToMulticastGroupCmd.MarkFlagRequired("payload-data")
		iotwireless_sendDataToMulticastGroupCmd.MarkFlagRequired("wireless-metadata")
	})
	iotwirelessCmd.AddCommand(iotwireless_sendDataToMulticastGroupCmd)
}
