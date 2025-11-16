package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_listWirelessDevicesCmd = &cobra.Command{
	Use:   "list-wireless-devices",
	Short: "Lists the wireless devices registered to your AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_listWirelessDevicesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_listWirelessDevicesCmd).Standalone()

		iotwireless_listWirelessDevicesCmd.Flags().String("destination-name", "", "A filter to list only the wireless devices that use as uplink destination.")
		iotwireless_listWirelessDevicesCmd.Flags().String("device-profile-id", "", "A filter to list only the wireless devices that use this device profile.")
		iotwireless_listWirelessDevicesCmd.Flags().String("fuota-task-id", "", "")
		iotwireless_listWirelessDevicesCmd.Flags().String("max-results", "", "The maximum number of results to return in this operation.")
		iotwireless_listWirelessDevicesCmd.Flags().String("multicast-group-id", "", "")
		iotwireless_listWirelessDevicesCmd.Flags().String("next-token", "", "To retrieve the next set of results, the `nextToken` value from a previous response; otherwise **null** to receive the first set of results.")
		iotwireless_listWirelessDevicesCmd.Flags().String("service-profile-id", "", "A filter to list only the wireless devices that use this service profile.")
		iotwireless_listWirelessDevicesCmd.Flags().String("wireless-device-type", "", "A filter to list only the wireless devices that use this wireless device type.")
	})
	iotwirelessCmd.AddCommand(iotwireless_listWirelessDevicesCmd)
}
