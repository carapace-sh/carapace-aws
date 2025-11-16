package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_updateDeviceCmd = &cobra.Command{
	Use:   "update-device",
	Short: "Updates the details for an existing device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_updateDeviceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_updateDeviceCmd).Standalone()

		networkmanager_updateDeviceCmd.Flags().String("awslocation", "", "The Amazon Web Services location of the device, if applicable.")
		networkmanager_updateDeviceCmd.Flags().String("description", "", "A description of the device.")
		networkmanager_updateDeviceCmd.Flags().String("device-id", "", "The ID of the device.")
		networkmanager_updateDeviceCmd.Flags().String("global-network-id", "", "The ID of the global network.")
		networkmanager_updateDeviceCmd.Flags().String("location", "", "")
		networkmanager_updateDeviceCmd.Flags().String("model", "", "The model of the device.")
		networkmanager_updateDeviceCmd.Flags().String("serial-number", "", "The serial number of the device.")
		networkmanager_updateDeviceCmd.Flags().String("site-id", "", "The ID of the site.")
		networkmanager_updateDeviceCmd.Flags().String("type", "", "The type of the device.")
		networkmanager_updateDeviceCmd.Flags().String("vendor", "", "The vendor of the device.")
		networkmanager_updateDeviceCmd.MarkFlagRequired("device-id")
		networkmanager_updateDeviceCmd.MarkFlagRequired("global-network-id")
	})
	networkmanagerCmd.AddCommand(networkmanager_updateDeviceCmd)
}
