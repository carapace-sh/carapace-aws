package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisvideo_describeImageGenerationConfigurationCmd = &cobra.Command{
	Use:   "describe-image-generation-configuration",
	Short: "Gets the `ImageGenerationConfiguration` for a given Kinesis video stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisvideo_describeImageGenerationConfigurationCmd).Standalone()

	kinesisvideo_describeImageGenerationConfigurationCmd.Flags().String("stream-arn", "", "The Amazon Resource Name (ARN) of the Kinesis video stream from which to retrieve the image generation configuration.")
	kinesisvideo_describeImageGenerationConfigurationCmd.Flags().String("stream-name", "", "The name of the stream from which to retrieve the image generation configuration.")
	kinesisvideoCmd.AddCommand(kinesisvideo_describeImageGenerationConfigurationCmd)
}
