package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_startWirelessDeviceImportTaskCmd = &cobra.Command{
	Use:   "start-wireless-device-import-task",
	Short: "Start import task for provisioning Sidewalk devices in bulk using an S3 CSV file.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_startWirelessDeviceImportTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_startWirelessDeviceImportTaskCmd).Standalone()

		iotwireless_startWirelessDeviceImportTaskCmd.Flags().String("client-request-token", "", "")
		iotwireless_startWirelessDeviceImportTaskCmd.Flags().String("destination-name", "", "The name of the Sidewalk destination that describes the IoT rule to route messages from the devices in the import task that are onboarded to AWS IoT Wireless.")
		iotwireless_startWirelessDeviceImportTaskCmd.Flags().String("positioning", "", "The integration status of the Device Location feature for Sidewalk devices.")
		iotwireless_startWirelessDeviceImportTaskCmd.Flags().String("sidewalk", "", "The Sidewalk-related parameters for importing wireless devices that need to be provisioned in bulk.")
		iotwireless_startWirelessDeviceImportTaskCmd.Flags().String("tags", "", "")
		iotwireless_startWirelessDeviceImportTaskCmd.MarkFlagRequired("destination-name")
		iotwireless_startWirelessDeviceImportTaskCmd.MarkFlagRequired("sidewalk")
	})
	iotwirelessCmd.AddCommand(iotwireless_startWirelessDeviceImportTaskCmd)
}
