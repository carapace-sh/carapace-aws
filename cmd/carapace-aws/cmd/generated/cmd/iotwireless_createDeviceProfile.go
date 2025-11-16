package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_createDeviceProfileCmd = &cobra.Command{
	Use:   "create-device-profile",
	Short: "Creates a new device profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_createDeviceProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_createDeviceProfileCmd).Standalone()

		iotwireless_createDeviceProfileCmd.Flags().String("client-request-token", "", "Each resource must have a unique client request token.")
		iotwireless_createDeviceProfileCmd.Flags().String("lo-ra-wan", "", "The device profile information to use to create the device profile.")
		iotwireless_createDeviceProfileCmd.Flags().String("name", "", "The name of the new resource.")
		iotwireless_createDeviceProfileCmd.Flags().String("sidewalk", "", "The Sidewalk-related information for creating the Sidewalk device profile.")
		iotwireless_createDeviceProfileCmd.Flags().String("tags", "", "The tags to attach to the new device profile.")
	})
	iotwirelessCmd.AddCommand(iotwireless_createDeviceProfileCmd)
}
