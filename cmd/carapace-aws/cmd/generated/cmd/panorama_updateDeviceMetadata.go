package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var panorama_updateDeviceMetadataCmd = &cobra.Command{
	Use:   "update-device-metadata",
	Short: "Updates a device's metadata.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(panorama_updateDeviceMetadataCmd).Standalone()

	panorama_updateDeviceMetadataCmd.Flags().String("description", "", "A description for the device.")
	panorama_updateDeviceMetadataCmd.Flags().String("device-id", "", "The device's ID.")
	panorama_updateDeviceMetadataCmd.MarkFlagRequired("device-id")
	panoramaCmd.AddCommand(panorama_updateDeviceMetadataCmd)
}
