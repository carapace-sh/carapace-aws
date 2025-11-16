package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_deleteDeviceCmd = &cobra.Command{
	Use:   "delete-device",
	Short: "Deletes an existing device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_deleteDeviceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_deleteDeviceCmd).Standalone()

		networkmanager_deleteDeviceCmd.Flags().String("device-id", "", "The ID of the device.")
		networkmanager_deleteDeviceCmd.Flags().String("global-network-id", "", "The ID of the global network.")
		networkmanager_deleteDeviceCmd.MarkFlagRequired("device-id")
		networkmanager_deleteDeviceCmd.MarkFlagRequired("global-network-id")
	})
	networkmanagerCmd.AddCommand(networkmanager_deleteDeviceCmd)
}
