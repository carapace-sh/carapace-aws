package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_deleteDeviceProfileCmd = &cobra.Command{
	Use:   "delete-device-profile",
	Short: "Deletes a device profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_deleteDeviceProfileCmd).Standalone()

	iotwireless_deleteDeviceProfileCmd.Flags().String("id", "", "The ID of the resource to delete.")
	iotwireless_deleteDeviceProfileCmd.MarkFlagRequired("id")
	iotwirelessCmd.AddCommand(iotwireless_deleteDeviceProfileCmd)
}
