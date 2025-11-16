package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_listDevicesForWirelessDeviceImportTaskCmd = &cobra.Command{
	Use:   "list-devices-for-wireless-device-import-task",
	Short: "List the Sidewalk devices in an import task and their onboarding status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_listDevicesForWirelessDeviceImportTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_listDevicesForWirelessDeviceImportTaskCmd).Standalone()

		iotwireless_listDevicesForWirelessDeviceImportTaskCmd.Flags().String("id", "", "The identifier of the import task for which wireless devices are listed.")
		iotwireless_listDevicesForWirelessDeviceImportTaskCmd.Flags().String("max-results", "", "")
		iotwireless_listDevicesForWirelessDeviceImportTaskCmd.Flags().String("next-token", "", "To retrieve the next set of results, the `nextToken` value from a previous response; otherwise `null` to receive the first set of results.")
		iotwireless_listDevicesForWirelessDeviceImportTaskCmd.Flags().String("status", "", "The status of the devices in the import task.")
		iotwireless_listDevicesForWirelessDeviceImportTaskCmd.MarkFlagRequired("id")
	})
	iotwirelessCmd.AddCommand(iotwireless_listDevicesForWirelessDeviceImportTaskCmd)
}
