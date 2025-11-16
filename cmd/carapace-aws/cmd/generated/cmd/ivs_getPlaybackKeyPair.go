package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivs_getPlaybackKeyPairCmd = &cobra.Command{
	Use:   "get-playback-key-pair",
	Short: "Gets a specified playback authorization key pair and returns the `arn` and `fingerprint`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivs_getPlaybackKeyPairCmd).Standalone()

	ivs_getPlaybackKeyPairCmd.Flags().String("arn", "", "ARN of the key pair to be returned.")
	ivs_getPlaybackKeyPairCmd.MarkFlagRequired("arn")
	ivsCmd.AddCommand(ivs_getPlaybackKeyPairCmd)
}
