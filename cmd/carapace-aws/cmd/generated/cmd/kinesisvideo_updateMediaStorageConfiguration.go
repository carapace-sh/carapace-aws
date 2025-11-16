package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisvideo_updateMediaStorageConfigurationCmd = &cobra.Command{
	Use:   "update-media-storage-configuration",
	Short: "Associates a `SignalingChannel` to a stream to store the media.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisvideo_updateMediaStorageConfigurationCmd).Standalone()

	kinesisvideo_updateMediaStorageConfigurationCmd.Flags().String("channel-arn", "", "The Amazon Resource Name (ARN) of the channel.")
	kinesisvideo_updateMediaStorageConfigurationCmd.Flags().String("media-storage-configuration", "", "A structure that encapsulates, or contains, the media storage configuration properties.")
	kinesisvideo_updateMediaStorageConfigurationCmd.MarkFlagRequired("channel-arn")
	kinesisvideo_updateMediaStorageConfigurationCmd.MarkFlagRequired("media-storage-configuration")
	kinesisvideoCmd.AddCommand(kinesisvideo_updateMediaStorageConfigurationCmd)
}
