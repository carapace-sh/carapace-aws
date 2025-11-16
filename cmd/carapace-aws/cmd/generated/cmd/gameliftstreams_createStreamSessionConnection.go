package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gameliftstreams_createStreamSessionConnectionCmd = &cobra.Command{
	Use:   "create-stream-session-connection",
	Short: "Enables clients to reconnect to a stream session while preserving all session state and data in the disconnected session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gameliftstreams_createStreamSessionConnectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gameliftstreams_createStreamSessionConnectionCmd).Standalone()

		gameliftstreams_createStreamSessionConnectionCmd.Flags().String("client-token", "", "A unique identifier that represents a client request.")
		gameliftstreams_createStreamSessionConnectionCmd.Flags().String("identifier", "", "[Amazon Resource Name (ARN)](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference-arns.html) or ID that uniquely identifies the stream group resource.")
		gameliftstreams_createStreamSessionConnectionCmd.Flags().String("signal-request", "", "A WebRTC ICE offer string to use when initializing a WebRTC connection.")
		gameliftstreams_createStreamSessionConnectionCmd.Flags().String("stream-session-identifier", "", "[Amazon Resource Name (ARN)](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference-arns.html) or ID that uniquely identifies the stream session resource.")
		gameliftstreams_createStreamSessionConnectionCmd.MarkFlagRequired("identifier")
		gameliftstreams_createStreamSessionConnectionCmd.MarkFlagRequired("signal-request")
		gameliftstreams_createStreamSessionConnectionCmd.MarkFlagRequired("stream-session-identifier")
	})
	gameliftstreamsCmd.AddCommand(gameliftstreams_createStreamSessionConnectionCmd)
}
