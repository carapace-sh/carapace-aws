package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisVideoArchivedMedia_getDashstreamingSessionUrlCmd = &cobra.Command{
	Use:   "get-dashstreaming-session-url",
	Short: "Retrieves an MPEG Dynamic Adaptive Streaming over HTTP (DASH) URL for the stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisVideoArchivedMedia_getDashstreamingSessionUrlCmd).Standalone()

	kinesisVideoArchivedMedia_getDashstreamingSessionUrlCmd.Flags().String("dashfragment-selector", "", "The time range of the requested fragment and the source of the timestamps.")
	kinesisVideoArchivedMedia_getDashstreamingSessionUrlCmd.Flags().String("display-fragment-number", "", "Fragments are identified in the manifest file based on their sequence number in the session.")
	kinesisVideoArchivedMedia_getDashstreamingSessionUrlCmd.Flags().String("display-fragment-timestamp", "", "Per the MPEG-DASH specification, the wall-clock time of fragments in the manifest file can be derived using attributes in the manifest itself.")
	kinesisVideoArchivedMedia_getDashstreamingSessionUrlCmd.Flags().String("expires", "", "The time in seconds until the requested session expires.")
	kinesisVideoArchivedMedia_getDashstreamingSessionUrlCmd.Flags().String("max-manifest-fragment-results", "", "The maximum number of fragments that are returned in the MPEG-DASH manifest.")
	kinesisVideoArchivedMedia_getDashstreamingSessionUrlCmd.Flags().String("playback-mode", "", "Whether to retrieve live, live replay, or archived, on-demand data.")
	kinesisVideoArchivedMedia_getDashstreamingSessionUrlCmd.Flags().String("stream-arn", "", "The Amazon Resource Name (ARN) of the stream for which to retrieve the MPEG-DASH manifest URL.")
	kinesisVideoArchivedMedia_getDashstreamingSessionUrlCmd.Flags().String("stream-name", "", "The name of the stream for which to retrieve the MPEG-DASH manifest URL.")
	kinesisVideoArchivedMediaCmd.AddCommand(kinesisVideoArchivedMedia_getDashstreamingSessionUrlCmd)
}
