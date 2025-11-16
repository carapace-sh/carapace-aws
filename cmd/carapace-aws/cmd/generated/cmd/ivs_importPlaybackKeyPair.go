package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivs_importPlaybackKeyPairCmd = &cobra.Command{
	Use:   "import-playback-key-pair",
	Short: "Imports the public portion of a new key pair and returns its `arn` and `fingerprint`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivs_importPlaybackKeyPairCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ivs_importPlaybackKeyPairCmd).Standalone()

		ivs_importPlaybackKeyPairCmd.Flags().String("name", "", "Playback-key-pair name.")
		ivs_importPlaybackKeyPairCmd.Flags().String("public-key-material", "", "The public portion of a customer-generated key pair.")
		ivs_importPlaybackKeyPairCmd.Flags().String("tags", "", "Any tags provided with the request are added to the playback key pair tags.")
		ivs_importPlaybackKeyPairCmd.MarkFlagRequired("public-key-material")
	})
	ivsCmd.AddCommand(ivs_importPlaybackKeyPairCmd)
}
