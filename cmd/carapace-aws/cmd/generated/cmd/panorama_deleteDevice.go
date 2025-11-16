package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var panorama_deleteDeviceCmd = &cobra.Command{
	Use:   "delete-device",
	Short: "Deletes a device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(panorama_deleteDeviceCmd).Standalone()

	panorama_deleteDeviceCmd.Flags().String("device-id", "", "The device's ID.")
	panorama_deleteDeviceCmd.MarkFlagRequired("device-id")
	panoramaCmd.AddCommand(panorama_deleteDeviceCmd)
}
