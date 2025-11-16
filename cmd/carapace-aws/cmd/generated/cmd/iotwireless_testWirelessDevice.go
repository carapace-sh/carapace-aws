package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_testWirelessDeviceCmd = &cobra.Command{
	Use:   "test-wireless-device",
	Short: "Simulates a provisioned device by sending an uplink data payload of `Hello`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_testWirelessDeviceCmd).Standalone()

	iotwireless_testWirelessDeviceCmd.Flags().String("id", "", "The ID of the wireless device to test.")
	iotwireless_testWirelessDeviceCmd.MarkFlagRequired("id")
	iotwirelessCmd.AddCommand(iotwireless_testWirelessDeviceCmd)
}
