package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediatailor_deletePlaybackConfigurationCmd = &cobra.Command{
	Use:   "delete-playback-configuration",
	Short: "Deletes a playback configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediatailor_deletePlaybackConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediatailor_deletePlaybackConfigurationCmd).Standalone()

		mediatailor_deletePlaybackConfigurationCmd.Flags().String("name", "", "The name of the playback configuration.")
		mediatailor_deletePlaybackConfigurationCmd.MarkFlagRequired("name")
	})
	mediatailorCmd.AddCommand(mediatailor_deletePlaybackConfigurationCmd)
}
