package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_deleteWirelessDeviceCmd = &cobra.Command{
	Use:   "delete-wireless-device",
	Short: "Deletes a wireless device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_deleteWirelessDeviceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_deleteWirelessDeviceCmd).Standalone()

		iotwireless_deleteWirelessDeviceCmd.Flags().String("id", "", "The ID of the resource to delete.")
		iotwireless_deleteWirelessDeviceCmd.MarkFlagRequired("id")
	})
	iotwirelessCmd.AddCommand(iotwireless_deleteWirelessDeviceCmd)
}
