package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gameliftstreams_getStreamSessionCmd = &cobra.Command{
	Use:   "get-stream-session",
	Short: "Retrieves properties for a Amazon GameLift Streams stream session resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gameliftstreams_getStreamSessionCmd).Standalone()

	gameliftstreams_getStreamSessionCmd.Flags().String("identifier", "", "The stream group that runs this stream session.")
	gameliftstreams_getStreamSessionCmd.Flags().String("stream-session-identifier", "", "An [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference-arns.html) or ID that uniquely identifies the stream session resource.")
	gameliftstreams_getStreamSessionCmd.MarkFlagRequired("identifier")
	gameliftstreams_getStreamSessionCmd.MarkFlagRequired("stream-session-identifier")
	gameliftstreamsCmd.AddCommand(gameliftstreams_getStreamSessionCmd)
}
