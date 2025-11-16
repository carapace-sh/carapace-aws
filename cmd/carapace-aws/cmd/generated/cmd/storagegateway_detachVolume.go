package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_detachVolumeCmd = &cobra.Command{
	Use:   "detach-volume",
	Short: "Disconnects a volume from an iSCSI connection and then detaches the volume from the specified gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_detachVolumeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_detachVolumeCmd).Standalone()

		storagegateway_detachVolumeCmd.Flags().Bool("force-detach", false, "Set to `true` to forcibly remove the iSCSI connection of the target volume and detach the volume.")
		storagegateway_detachVolumeCmd.Flags().Bool("no-force-detach", false, "Set to `true` to forcibly remove the iSCSI connection of the target volume and detach the volume.")
		storagegateway_detachVolumeCmd.Flags().String("volume-arn", "", "The Amazon Resource Name (ARN) of the volume to detach from the gateway.")
		storagegateway_detachVolumeCmd.Flag("no-force-detach").Hidden = true
		storagegateway_detachVolumeCmd.MarkFlagRequired("volume-arn")
	})
	storagegatewayCmd.AddCommand(storagegateway_detachVolumeCmd)
}
