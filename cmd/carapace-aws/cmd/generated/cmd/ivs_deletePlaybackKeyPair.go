package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivs_deletePlaybackKeyPairCmd = &cobra.Command{
	Use:   "delete-playback-key-pair",
	Short: "Deletes a specified authorization key pair.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivs_deletePlaybackKeyPairCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ivs_deletePlaybackKeyPairCmd).Standalone()

		ivs_deletePlaybackKeyPairCmd.Flags().String("arn", "", "ARN of the key pair to be deleted.")
		ivs_deletePlaybackKeyPairCmd.MarkFlagRequired("arn")
	})
	ivsCmd.AddCommand(ivs_deletePlaybackKeyPairCmd)
}
