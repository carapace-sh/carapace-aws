package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisvideo_describeMediaStorageConfigurationCmd = &cobra.Command{
	Use:   "describe-media-storage-configuration",
	Short: "Returns the most current information about the channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisvideo_describeMediaStorageConfigurationCmd).Standalone()

	kinesisvideo_describeMediaStorageConfigurationCmd.Flags().String("channel-arn", "", "The Amazon Resource Name (ARN) of the channel.")
	kinesisvideo_describeMediaStorageConfigurationCmd.Flags().String("channel-name", "", "The name of the channel.")
	kinesisvideoCmd.AddCommand(kinesisvideo_describeMediaStorageConfigurationCmd)
}
