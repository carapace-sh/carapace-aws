package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_updateVolumeCmd = &cobra.Command{
	Use:   "update-volume",
	Short: "Updates the configuration of an Amazon FSx for NetApp ONTAP or Amazon FSx for OpenZFS volume.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_updateVolumeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fsx_updateVolumeCmd).Standalone()

		fsx_updateVolumeCmd.Flags().String("client-request-token", "", "")
		fsx_updateVolumeCmd.Flags().String("name", "", "The name of the OpenZFS volume.")
		fsx_updateVolumeCmd.Flags().String("ontap-configuration", "", "The configuration of the ONTAP volume that you are updating.")
		fsx_updateVolumeCmd.Flags().String("open-zfsconfiguration", "", "The configuration of the OpenZFS volume that you are updating.")
		fsx_updateVolumeCmd.Flags().String("volume-id", "", "The ID of the volume that you want to update, in the format `fsvol-0123456789abcdef0`.")
		fsx_updateVolumeCmd.MarkFlagRequired("volume-id")
	})
	fsxCmd.AddCommand(fsx_updateVolumeCmd)
}
