package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_createDeviceCmd = &cobra.Command{
	Use:   "create-device",
	Short: "Creates a new device in a global network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_createDeviceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_createDeviceCmd).Standalone()

		networkmanager_createDeviceCmd.Flags().String("awslocation", "", "The Amazon Web Services location of the device, if applicable.")
		networkmanager_createDeviceCmd.Flags().String("description", "", "A description of the device.")
		networkmanager_createDeviceCmd.Flags().String("global-network-id", "", "The ID of the global network.")
		networkmanager_createDeviceCmd.Flags().String("location", "", "The location of the device.")
		networkmanager_createDeviceCmd.Flags().String("model", "", "The model of the device.")
		networkmanager_createDeviceCmd.Flags().String("serial-number", "", "The serial number of the device.")
		networkmanager_createDeviceCmd.Flags().String("site-id", "", "The ID of the site.")
		networkmanager_createDeviceCmd.Flags().String("tags", "", "The tags to apply to the resource during creation.")
		networkmanager_createDeviceCmd.Flags().String("type", "", "The type of the device.")
		networkmanager_createDeviceCmd.Flags().String("vendor", "", "The vendor of the device.")
		networkmanager_createDeviceCmd.MarkFlagRequired("global-network-id")
	})
	networkmanagerCmd.AddCommand(networkmanager_createDeviceCmd)
}
