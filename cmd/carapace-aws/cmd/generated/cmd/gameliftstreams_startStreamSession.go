package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gameliftstreams_startStreamSessionCmd = &cobra.Command{
	Use:   "start-stream-session",
	Short: "This action initiates a new stream session and outputs connection information that clients can use to access the stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gameliftstreams_startStreamSessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gameliftstreams_startStreamSessionCmd).Standalone()

		gameliftstreams_startStreamSessionCmd.Flags().String("additional-environment-variables", "", "A set of options that you can use to control the stream session runtime environment, expressed as a set of key-value pairs.")
		gameliftstreams_startStreamSessionCmd.Flags().String("additional-launch-args", "", "A list of CLI arguments that are sent to the streaming server when a stream session launches.")
		gameliftstreams_startStreamSessionCmd.Flags().String("application-identifier", "", "An [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference-arns.html) or ID that uniquely identifies the application resource.")
		gameliftstreams_startStreamSessionCmd.Flags().String("client-token", "", "A unique identifier that represents a client request.")
		gameliftstreams_startStreamSessionCmd.Flags().String("connection-timeout-seconds", "", "Length of time (in seconds) that Amazon GameLift Streams should wait for a client to connect or reconnect to the stream session.")
		gameliftstreams_startStreamSessionCmd.Flags().String("description", "", "A human-readable label for the stream session.")
		gameliftstreams_startStreamSessionCmd.Flags().String("identifier", "", "The stream group to run this stream session with.")
		gameliftstreams_startStreamSessionCmd.Flags().String("locations", "", "A list of locations, in order of priority, where you want Amazon GameLift Streams to start a stream from.")
		gameliftstreams_startStreamSessionCmd.Flags().String("protocol", "", "The data transport protocol to use for the stream session.")
		gameliftstreams_startStreamSessionCmd.Flags().String("session-length-seconds", "", "The maximum duration of a session.")
		gameliftstreams_startStreamSessionCmd.Flags().String("signal-request", "", "A WebRTC ICE offer string to use when initializing a WebRTC connection.")
		gameliftstreams_startStreamSessionCmd.Flags().String("user-id", "", "An opaque, unique identifier for an end-user, defined by the developer.")
		gameliftstreams_startStreamSessionCmd.MarkFlagRequired("application-identifier")
		gameliftstreams_startStreamSessionCmd.MarkFlagRequired("identifier")
		gameliftstreams_startStreamSessionCmd.MarkFlagRequired("protocol")
		gameliftstreams_startStreamSessionCmd.MarkFlagRequired("signal-request")
	})
	gameliftstreamsCmd.AddCommand(gameliftstreams_startStreamSessionCmd)
}
