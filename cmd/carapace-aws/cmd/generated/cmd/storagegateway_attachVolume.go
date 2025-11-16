package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_attachVolumeCmd = &cobra.Command{
	Use:   "attach-volume",
	Short: "Connects a volume to an iSCSI connection and then attaches the volume to the specified gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_attachVolumeCmd).Standalone()

	storagegateway_attachVolumeCmd.Flags().String("disk-id", "", "The unique device ID or other distinguishing data that identifies the local disk used to create the volume.")
	storagegateway_attachVolumeCmd.Flags().String("gateway-arn", "", "The Amazon Resource Name (ARN) of the gateway that you want to attach the volume to.")
	storagegateway_attachVolumeCmd.Flags().String("network-interface-id", "", "The network interface of the gateway on which to expose the iSCSI target.")
	storagegateway_attachVolumeCmd.Flags().String("target-name", "", "The name of the iSCSI target used by an initiator to connect to a volume and used as a suffix for the target ARN.")
	storagegateway_attachVolumeCmd.Flags().String("volume-arn", "", "The Amazon Resource Name (ARN) of the volume to attach to the specified gateway.")
	storagegateway_attachVolumeCmd.MarkFlagRequired("gateway-arn")
	storagegateway_attachVolumeCmd.MarkFlagRequired("network-interface-id")
	storagegateway_attachVolumeCmd.MarkFlagRequired("volume-arn")
	storagegatewayCmd.AddCommand(storagegateway_attachVolumeCmd)
}
