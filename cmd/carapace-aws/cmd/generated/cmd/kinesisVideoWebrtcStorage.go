package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisVideoWebrtcStorageCmd = &cobra.Command{
	Use:   "kinesis-video-webrtc-storage",
	Short: "webrtc",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisVideoWebrtcStorageCmd).Standalone()

	rootCmd.AddCommand(kinesisVideoWebrtcStorageCmd)
}
