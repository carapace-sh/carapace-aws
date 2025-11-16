package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivs_listPlaybackKeyPairsCmd = &cobra.Command{
	Use:   "list-playback-key-pairs",
	Short: "Gets summary information about playback key pairs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivs_listPlaybackKeyPairsCmd).Standalone()

	ivs_listPlaybackKeyPairsCmd.Flags().String("max-results", "", "Maximum number of key pairs to return.")
	ivs_listPlaybackKeyPairsCmd.Flags().String("next-token", "", "The first key pair to retrieve.")
	ivsCmd.AddCommand(ivs_listPlaybackKeyPairsCmd)
}
