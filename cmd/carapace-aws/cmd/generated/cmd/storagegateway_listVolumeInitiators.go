package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_listVolumeInitiatorsCmd = &cobra.Command{
	Use:   "list-volume-initiators",
	Short: "Lists iSCSI initiators that are connected to a volume.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_listVolumeInitiatorsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_listVolumeInitiatorsCmd).Standalone()

		storagegateway_listVolumeInitiatorsCmd.Flags().String("volume-arn", "", "The Amazon Resource Name (ARN) of the volume.")
		storagegateway_listVolumeInitiatorsCmd.MarkFlagRequired("volume-arn")
	})
	storagegatewayCmd.AddCommand(storagegateway_listVolumeInitiatorsCmd)
}
