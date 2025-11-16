package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_deleteWirelessDeviceImportTaskCmd = &cobra.Command{
	Use:   "delete-wireless-device-import-task",
	Short: "Delete an import task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_deleteWirelessDeviceImportTaskCmd).Standalone()

	iotwireless_deleteWirelessDeviceImportTaskCmd.Flags().String("id", "", "The unique identifier of the import task to be deleted.")
	iotwireless_deleteWirelessDeviceImportTaskCmd.MarkFlagRequired("id")
	iotwirelessCmd.AddCommand(iotwireless_deleteWirelessDeviceImportTaskCmd)
}
