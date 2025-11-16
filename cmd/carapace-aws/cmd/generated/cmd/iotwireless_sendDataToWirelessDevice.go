package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_sendDataToWirelessDeviceCmd = &cobra.Command{
	Use:   "send-data-to-wireless-device",
	Short: "Sends a decrypted application data frame to a device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_sendDataToWirelessDeviceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_sendDataToWirelessDeviceCmd).Standalone()

		iotwireless_sendDataToWirelessDeviceCmd.Flags().String("id", "", "The ID of the wireless device to receive the data.")
		iotwireless_sendDataToWirelessDeviceCmd.Flags().String("payload-data", "", "")
		iotwireless_sendDataToWirelessDeviceCmd.Flags().String("transmit-mode", "", "The transmit mode to use to send data to the wireless device.")
		iotwireless_sendDataToWirelessDeviceCmd.Flags().String("wireless-metadata", "", "Metadata about the message request.")
		iotwireless_sendDataToWirelessDeviceCmd.MarkFlagRequired("id")
		iotwireless_sendDataToWirelessDeviceCmd.MarkFlagRequired("payload-data")
		iotwireless_sendDataToWirelessDeviceCmd.MarkFlagRequired("transmit-mode")
	})
	iotwirelessCmd.AddCommand(iotwireless_sendDataToWirelessDeviceCmd)
}
