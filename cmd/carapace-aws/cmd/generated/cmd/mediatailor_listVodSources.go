package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediatailor_listVodSourcesCmd = &cobra.Command{
	Use:   "list-vod-sources",
	Short: "Lists the VOD sources contained in a source location.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediatailor_listVodSourcesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediatailor_listVodSourcesCmd).Standalone()

		mediatailor_listVodSourcesCmd.Flags().String("max-results", "", "The maximum number of VOD sources that you want MediaTailor to return in response to the current request.")
		mediatailor_listVodSourcesCmd.Flags().String("next-token", "", "Pagination token returned by the list request when results exceed the maximum allowed.")
		mediatailor_listVodSourcesCmd.Flags().String("source-location-name", "", "The name of the source location associated with this VOD Source list.")
		mediatailor_listVodSourcesCmd.MarkFlagRequired("source-location-name")
	})
	mediatailorCmd.AddCommand(mediatailor_listVodSourcesCmd)
}
