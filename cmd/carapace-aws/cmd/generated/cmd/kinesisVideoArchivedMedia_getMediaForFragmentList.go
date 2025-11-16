package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisVideoArchivedMedia_getMediaForFragmentListCmd = &cobra.Command{
	Use:   "get-media-for-fragment-list",
	Short: "Gets media for a list of fragments (specified by fragment number) from the archived data in an Amazon Kinesis video stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisVideoArchivedMedia_getMediaForFragmentListCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisVideoArchivedMedia_getMediaForFragmentListCmd).Standalone()

		kinesisVideoArchivedMedia_getMediaForFragmentListCmd.Flags().String("fragments", "", "A list of the numbers of fragments for which to retrieve media.")
		kinesisVideoArchivedMedia_getMediaForFragmentListCmd.Flags().String("stream-arn", "", "The Amazon Resource Name (ARN) of the stream from which to retrieve fragment media.")
		kinesisVideoArchivedMedia_getMediaForFragmentListCmd.Flags().String("stream-name", "", "The name of the stream from which to retrieve fragment media.")
		kinesisVideoArchivedMedia_getMediaForFragmentListCmd.MarkFlagRequired("fragments")
	})
	kinesisVideoArchivedMediaCmd.AddCommand(kinesisVideoArchivedMedia_getMediaForFragmentListCmd)
}
