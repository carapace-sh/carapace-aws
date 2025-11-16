package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisVideoArchivedMedia_getImagesCmd = &cobra.Command{
	Use:   "get-images",
	Short: "Retrieves a list of images corresponding to each timestamp for a given time range, sampling interval, and image format configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisVideoArchivedMedia_getImagesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisVideoArchivedMedia_getImagesCmd).Standalone()

		kinesisVideoArchivedMedia_getImagesCmd.Flags().String("end-timestamp", "", "The end timestamp for the range of images to be generated.")
		kinesisVideoArchivedMedia_getImagesCmd.Flags().String("format", "", "The format that will be used to encode the image.")
		kinesisVideoArchivedMedia_getImagesCmd.Flags().String("format-config", "", "The list of a key-value pair structure that contains extra parameters that can be applied when the image is generated.")
		kinesisVideoArchivedMedia_getImagesCmd.Flags().String("height-pixels", "", "The height of the output image that is used in conjunction with the `WidthPixels` parameter.")
		kinesisVideoArchivedMedia_getImagesCmd.Flags().String("image-selector-type", "", "The origin of the Server or Producer timestamps to use to generate the images.")
		kinesisVideoArchivedMedia_getImagesCmd.Flags().String("max-results", "", "The maximum number of images to be returned by the API.")
		kinesisVideoArchivedMedia_getImagesCmd.Flags().String("next-token", "", "A token that specifies where to start paginating the next set of Images.")
		kinesisVideoArchivedMedia_getImagesCmd.Flags().String("sampling-interval", "", "The time interval in milliseconds (ms) at which the images need to be generated from the stream.")
		kinesisVideoArchivedMedia_getImagesCmd.Flags().String("start-timestamp", "", "The starting point from which the images should be generated.")
		kinesisVideoArchivedMedia_getImagesCmd.Flags().String("stream-arn", "", "The Amazon Resource Name (ARN) of the stream from which to retrieve the images.")
		kinesisVideoArchivedMedia_getImagesCmd.Flags().String("stream-name", "", "The name of the stream from which to retrieve the images.")
		kinesisVideoArchivedMedia_getImagesCmd.Flags().String("width-pixels", "", "The width of the output image that is used in conjunction with the `HeightPixels` parameter.")
		kinesisVideoArchivedMedia_getImagesCmd.MarkFlagRequired("end-timestamp")
		kinesisVideoArchivedMedia_getImagesCmd.MarkFlagRequired("format")
		kinesisVideoArchivedMedia_getImagesCmd.MarkFlagRequired("image-selector-type")
		kinesisVideoArchivedMedia_getImagesCmd.MarkFlagRequired("start-timestamp")
	})
	kinesisVideoArchivedMediaCmd.AddCommand(kinesisVideoArchivedMedia_getImagesCmd)
}
