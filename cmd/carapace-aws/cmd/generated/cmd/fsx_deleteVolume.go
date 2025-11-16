package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_deleteVolumeCmd = &cobra.Command{
	Use:   "delete-volume",
	Short: "Deletes an Amazon FSx for NetApp ONTAP or Amazon FSx for OpenZFS volume.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_deleteVolumeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fsx_deleteVolumeCmd).Standalone()

		fsx_deleteVolumeCmd.Flags().String("client-request-token", "", "")
		fsx_deleteVolumeCmd.Flags().String("ontap-configuration", "", "For Amazon FSx for ONTAP volumes, specify whether to take a final backup of the volume and apply tags to the backup.")
		fsx_deleteVolumeCmd.Flags().String("open-zfsconfiguration", "", "For Amazon FSx for OpenZFS volumes, specify whether to delete all child volumes and snapshots.")
		fsx_deleteVolumeCmd.Flags().String("volume-id", "", "The ID of the volume that you are deleting.")
		fsx_deleteVolumeCmd.MarkFlagRequired("volume-id")
	})
	fsxCmd.AddCommand(fsx_deleteVolumeCmd)
}
