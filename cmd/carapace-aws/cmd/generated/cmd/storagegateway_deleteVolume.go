package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_deleteVolumeCmd = &cobra.Command{
	Use:   "delete-volume",
	Short: "Deletes the specified storage volume that you previously created using the [CreateCachediSCSIVolume]() or [CreateStorediSCSIVolume]() API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_deleteVolumeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_deleteVolumeCmd).Standalone()

		storagegateway_deleteVolumeCmd.Flags().String("volume-arn", "", "The Amazon Resource Name (ARN) of the volume.")
		storagegateway_deleteVolumeCmd.MarkFlagRequired("volume-arn")
	})
	storagegatewayCmd.AddCommand(storagegateway_deleteVolumeCmd)
}
