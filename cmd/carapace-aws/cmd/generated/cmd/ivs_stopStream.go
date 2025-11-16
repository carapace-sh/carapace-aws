package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivs_stopStreamCmd = &cobra.Command{
	Use:   "stop-stream",
	Short: "Disconnects the incoming RTMPS stream for the specified channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivs_stopStreamCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ivs_stopStreamCmd).Standalone()

		ivs_stopStreamCmd.Flags().String("channel-arn", "", "ARN of the channel for which the stream is to be stopped.")
		ivs_stopStreamCmd.MarkFlagRequired("channel-arn")
	})
	ivsCmd.AddCommand(ivs_stopStreamCmd)
}
