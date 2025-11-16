package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_getWirelessDeviceStatisticsCmd = &cobra.Command{
	Use:   "get-wireless-device-statistics",
	Short: "Gets operating information about a wireless device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_getWirelessDeviceStatisticsCmd).Standalone()

	iotwireless_getWirelessDeviceStatisticsCmd.Flags().String("wireless-device-id", "", "The ID of the wireless device for which to get the data.")
	iotwireless_getWirelessDeviceStatisticsCmd.MarkFlagRequired("wireless-device-id")
	iotwirelessCmd.AddCommand(iotwireless_getWirelessDeviceStatisticsCmd)
}
