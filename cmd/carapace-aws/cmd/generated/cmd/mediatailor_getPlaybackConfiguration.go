package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediatailor_getPlaybackConfigurationCmd = &cobra.Command{
	Use:   "get-playback-configuration",
	Short: "Retrieves a playback configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediatailor_getPlaybackConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediatailor_getPlaybackConfigurationCmd).Standalone()

		mediatailor_getPlaybackConfigurationCmd.Flags().String("name", "", "The identifier for the playback configuration.")
		mediatailor_getPlaybackConfigurationCmd.MarkFlagRequired("name")
	})
	mediatailorCmd.AddCommand(mediatailor_getPlaybackConfigurationCmd)
}
