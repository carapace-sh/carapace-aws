package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediatailor_listPlaybackConfigurationsCmd = &cobra.Command{
	Use:   "list-playback-configurations",
	Short: "Retrieves existing playback configurations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediatailor_listPlaybackConfigurationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediatailor_listPlaybackConfigurationsCmd).Standalone()

		mediatailor_listPlaybackConfigurationsCmd.Flags().String("max-results", "", "The maximum number of playback configurations that you want MediaTailor to return in response to the current request.")
		mediatailor_listPlaybackConfigurationsCmd.Flags().String("next-token", "", "Pagination token returned by the list request when results exceed the maximum allowed.")
	})
	mediatailorCmd.AddCommand(mediatailor_listPlaybackConfigurationsCmd)
}
