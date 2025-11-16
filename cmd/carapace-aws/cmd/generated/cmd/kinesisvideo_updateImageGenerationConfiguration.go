package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisvideo_updateImageGenerationConfigurationCmd = &cobra.Command{
	Use:   "update-image-generation-configuration",
	Short: "Updates the `StreamInfo` and `ImageProcessingConfiguration` fields.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisvideo_updateImageGenerationConfigurationCmd).Standalone()

	kinesisvideo_updateImageGenerationConfigurationCmd.Flags().String("image-generation-configuration", "", "The structure that contains the information required for the KVS images delivery.")
	kinesisvideo_updateImageGenerationConfigurationCmd.Flags().String("stream-arn", "", "The Amazon Resource Name (ARN) of the Kinesis video stream from where you want to update the image generation configuration.")
	kinesisvideo_updateImageGenerationConfigurationCmd.Flags().String("stream-name", "", "The name of the stream from which to update the image generation configuration.")
	kinesisvideoCmd.AddCommand(kinesisvideo_updateImageGenerationConfigurationCmd)
}
