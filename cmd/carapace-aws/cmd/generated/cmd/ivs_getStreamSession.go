package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivs_getStreamSessionCmd = &cobra.Command{
	Use:   "get-stream-session",
	Short: "Gets metadata on a specified stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivs_getStreamSessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ivs_getStreamSessionCmd).Standalone()

		ivs_getStreamSessionCmd.Flags().String("channel-arn", "", "ARN of the channel resource")
		ivs_getStreamSessionCmd.Flags().String("stream-id", "", "Unique identifier for a live or previously live stream in the specified channel.")
		ivs_getStreamSessionCmd.MarkFlagRequired("channel-arn")
	})
	ivsCmd.AddCommand(ivs_getStreamSessionCmd)
}
