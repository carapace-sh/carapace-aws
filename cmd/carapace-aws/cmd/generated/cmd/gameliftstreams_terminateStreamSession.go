package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gameliftstreams_terminateStreamSessionCmd = &cobra.Command{
	Use:   "terminate-stream-session",
	Short: "Permanently terminates an active stream session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gameliftstreams_terminateStreamSessionCmd).Standalone()

	gameliftstreams_terminateStreamSessionCmd.Flags().String("identifier", "", "[Amazon Resource Name (ARN)](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference-arns.html) or ID that uniquely identifies the stream group resource.")
	gameliftstreams_terminateStreamSessionCmd.Flags().String("stream-session-identifier", "", "[Amazon Resource Name (ARN)](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference-arns.html) or ID that uniquely identifies the stream session resource.")
	gameliftstreams_terminateStreamSessionCmd.MarkFlagRequired("identifier")
	gameliftstreams_terminateStreamSessionCmd.MarkFlagRequired("stream-session-identifier")
	gameliftstreamsCmd.AddCommand(gameliftstreams_terminateStreamSessionCmd)
}
