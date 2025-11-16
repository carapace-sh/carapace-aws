package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_getWirelessDeviceImportTaskCmd = &cobra.Command{
	Use:   "get-wireless-device-import-task",
	Short: "Get information about an import task and count of device onboarding summary information for the import task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_getWirelessDeviceImportTaskCmd).Standalone()

	iotwireless_getWirelessDeviceImportTaskCmd.Flags().String("id", "", "The identifier of the import task for which information is requested.")
	iotwireless_getWirelessDeviceImportTaskCmd.MarkFlagRequired("id")
	iotwirelessCmd.AddCommand(iotwireless_getWirelessDeviceImportTaskCmd)
}
