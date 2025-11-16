package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisVideoMedia_getMediaCmd = &cobra.Command{
	Use:   "get-media",
	Short: "Use this API to retrieve media content from a Kinesis video stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisVideoMedia_getMediaCmd).Standalone()

	kinesisVideoMedia_getMediaCmd.Flags().String("start-selector", "", "Identifies the starting chunk to get from the specified stream.")
	kinesisVideoMedia_getMediaCmd.Flags().String("stream-arn", "", "The ARN of the stream from where you want to get the media content.")
	kinesisVideoMedia_getMediaCmd.Flags().String("stream-name", "", "The Kinesis video stream name from where you want to get the media content.")
	kinesisVideoMedia_getMediaCmd.MarkFlagRequired("start-selector")
	kinesisVideoMediaCmd.AddCommand(kinesisVideoMedia_getMediaCmd)
}
