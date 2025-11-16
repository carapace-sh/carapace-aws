package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisvideo_listStreamsCmd = &cobra.Command{
	Use:   "list-streams",
	Short: "Returns an array of `StreamInfo` objects.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisvideo_listStreamsCmd).Standalone()

	kinesisvideo_listStreamsCmd.Flags().String("max-results", "", "The maximum number of streams to return in the response.")
	kinesisvideo_listStreamsCmd.Flags().String("next-token", "", "If you specify this parameter, when the result of a `ListStreams` operation is truncated, the call returns the `NextToken` in the response.")
	kinesisvideo_listStreamsCmd.Flags().String("stream-name-condition", "", "Optional: Returns only streams that satisfy a specific condition.")
	kinesisvideoCmd.AddCommand(kinesisvideo_listStreamsCmd)
}
