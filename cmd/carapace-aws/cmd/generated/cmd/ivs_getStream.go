package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivs_getStreamCmd = &cobra.Command{
	Use:   "get-stream",
	Short: "Gets information about the active (live) stream on a specified channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivs_getStreamCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ivs_getStreamCmd).Standalone()

		ivs_getStreamCmd.Flags().String("channel-arn", "", "Channel ARN for stream to be accessed.")
		ivs_getStreamCmd.MarkFlagRequired("channel-arn")
	})
	ivsCmd.AddCommand(ivs_getStreamCmd)
}
