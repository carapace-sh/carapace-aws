package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisVideoArchivedMedia_getClipCmd = &cobra.Command{
	Use:   "get-clip",
	Short: "Downloads an MP4 file (clip) containing the archived, on-demand media from the specified video stream over the specified time range.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisVideoArchivedMedia_getClipCmd).Standalone()

	kinesisVideoArchivedMedia_getClipCmd.Flags().String("clip-fragment-selector", "", "The time range of the requested clip and the source of the timestamps.")
	kinesisVideoArchivedMedia_getClipCmd.Flags().String("stream-arn", "", "The Amazon Resource Name (ARN) of the stream for which to retrieve the media clip.")
	kinesisVideoArchivedMedia_getClipCmd.Flags().String("stream-name", "", "The name of the stream for which to retrieve the media clip.")
	kinesisVideoArchivedMedia_getClipCmd.MarkFlagRequired("clip-fragment-selector")
	kinesisVideoArchivedMediaCmd.AddCommand(kinesisVideoArchivedMedia_getClipCmd)
}
