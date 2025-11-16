package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisVideoArchivedMedia_getHlsstreamingSessionUrlCmd = &cobra.Command{
	Use:   "get-hlsstreaming-session-url",
	Short: "Retrieves an HTTP Live Streaming (HLS) URL for the stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisVideoArchivedMedia_getHlsstreamingSessionUrlCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisVideoArchivedMedia_getHlsstreamingSessionUrlCmd).Standalone()

		kinesisVideoArchivedMedia_getHlsstreamingSessionUrlCmd.Flags().String("container-format", "", "Specifies which format should be used for packaging the media.")
		kinesisVideoArchivedMedia_getHlsstreamingSessionUrlCmd.Flags().String("discontinuity-mode", "", "Specifies when flags marking discontinuities between fragments are added to the media playlists.")
		kinesisVideoArchivedMedia_getHlsstreamingSessionUrlCmd.Flags().String("display-fragment-timestamp", "", "Specifies when the fragment start timestamps should be included in the HLS media playlist.")
		kinesisVideoArchivedMedia_getHlsstreamingSessionUrlCmd.Flags().String("expires", "", "The time in seconds until the requested session expires.")
		kinesisVideoArchivedMedia_getHlsstreamingSessionUrlCmd.Flags().String("hlsfragment-selector", "", "The time range of the requested fragment and the source of the timestamps.")
		kinesisVideoArchivedMedia_getHlsstreamingSessionUrlCmd.Flags().String("max-media-playlist-fragment-results", "", "The maximum number of fragments that are returned in the HLS media playlists.")
		kinesisVideoArchivedMedia_getHlsstreamingSessionUrlCmd.Flags().String("playback-mode", "", "Whether to retrieve live, live replay, or archived, on-demand data.")
		kinesisVideoArchivedMedia_getHlsstreamingSessionUrlCmd.Flags().String("stream-arn", "", "The Amazon Resource Name (ARN) of the stream for which to retrieve the HLS master playlist URL.")
		kinesisVideoArchivedMedia_getHlsstreamingSessionUrlCmd.Flags().String("stream-name", "", "The name of the stream for which to retrieve the HLS master playlist URL.")
	})
	kinesisVideoArchivedMediaCmd.AddCommand(kinesisVideoArchivedMedia_getHlsstreamingSessionUrlCmd)
}
