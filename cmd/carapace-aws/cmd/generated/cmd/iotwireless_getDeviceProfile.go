package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_getDeviceProfileCmd = &cobra.Command{
	Use:   "get-device-profile",
	Short: "Gets information about a device profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_getDeviceProfileCmd).Standalone()

	iotwireless_getDeviceProfileCmd.Flags().String("id", "", "The ID of the resource to get.")
	iotwireless_getDeviceProfileCmd.MarkFlagRequired("id")
	iotwirelessCmd.AddCommand(iotwireless_getDeviceProfileCmd)
}
