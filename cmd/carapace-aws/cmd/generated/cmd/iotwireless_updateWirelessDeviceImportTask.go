package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_updateWirelessDeviceImportTaskCmd = &cobra.Command{
	Use:   "update-wireless-device-import-task",
	Short: "Update an import task to add more devices to the task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_updateWirelessDeviceImportTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_updateWirelessDeviceImportTaskCmd).Standalone()

		iotwireless_updateWirelessDeviceImportTaskCmd.Flags().String("id", "", "The identifier of the import task to be updated.")
		iotwireless_updateWirelessDeviceImportTaskCmd.Flags().String("sidewalk", "", "The Sidewalk-related parameters of the import task to be updated.")
		iotwireless_updateWirelessDeviceImportTaskCmd.MarkFlagRequired("id")
		iotwireless_updateWirelessDeviceImportTaskCmd.MarkFlagRequired("sidewalk")
	})
	iotwirelessCmd.AddCommand(iotwireless_updateWirelessDeviceImportTaskCmd)
}
