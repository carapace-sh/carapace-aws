package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisvideo_createStreamCmd = &cobra.Command{
	Use:   "create-stream",
	Short: "Creates a new Kinesis video stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisvideo_createStreamCmd).Standalone()

	kinesisvideo_createStreamCmd.Flags().String("data-retention-in-hours", "", "The number of hours that you want to retain the data in the stream.")
	kinesisvideo_createStreamCmd.Flags().String("device-name", "", "The name of the device that is writing to the stream.")
	kinesisvideo_createStreamCmd.Flags().String("kms-key-id", "", "The ID of the Key Management Service (KMS) key that you want Kinesis Video Streams to use to encrypt stream data.")
	kinesisvideo_createStreamCmd.Flags().String("media-type", "", "The media type of the stream.")
	kinesisvideo_createStreamCmd.Flags().String("stream-name", "", "A name for the stream that you are creating.")
	kinesisvideo_createStreamCmd.Flags().String("tags", "", "A list of tags to associate with the specified stream.")
	kinesisvideo_createStreamCmd.MarkFlagRequired("stream-name")
	kinesisvideoCmd.AddCommand(kinesisvideo_createStreamCmd)
}
