package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_startSingleWirelessDeviceImportTaskCmd = &cobra.Command{
	Use:   "start-single-wireless-device-import-task",
	Short: "Start import task for a single wireless device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_startSingleWirelessDeviceImportTaskCmd).Standalone()

	iotwireless_startSingleWirelessDeviceImportTaskCmd.Flags().String("client-request-token", "", "")
	iotwireless_startSingleWirelessDeviceImportTaskCmd.Flags().String("destination-name", "", "The name of the Sidewalk destination that describes the IoT rule to route messages from the device in the import task that will be onboarded to AWS IoT Wireless.")
	iotwireless_startSingleWirelessDeviceImportTaskCmd.Flags().String("device-name", "", "The name of the wireless device for which an import task is being started.")
	iotwireless_startSingleWirelessDeviceImportTaskCmd.Flags().String("positioning", "", "The integration status of the Device Location feature for Sidewalk devices.")
	iotwireless_startSingleWirelessDeviceImportTaskCmd.Flags().String("sidewalk", "", "The Sidewalk-related parameters for importing a single wireless device.")
	iotwireless_startSingleWirelessDeviceImportTaskCmd.Flags().String("tags", "", "")
	iotwireless_startSingleWirelessDeviceImportTaskCmd.MarkFlagRequired("destination-name")
	iotwireless_startSingleWirelessDeviceImportTaskCmd.MarkFlagRequired("sidewalk")
	iotwirelessCmd.AddCommand(iotwireless_startSingleWirelessDeviceImportTaskCmd)
}
