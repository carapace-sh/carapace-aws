package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediatailor_listLiveSourcesCmd = &cobra.Command{
	Use:   "list-live-sources",
	Short: "Lists the live sources contained in a source location.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediatailor_listLiveSourcesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediatailor_listLiveSourcesCmd).Standalone()

		mediatailor_listLiveSourcesCmd.Flags().String("max-results", "", "The maximum number of live sources that you want MediaTailor to return in response to the current request.")
		mediatailor_listLiveSourcesCmd.Flags().String("next-token", "", "Pagination token returned by the list request when results exceed the maximum allowed.")
		mediatailor_listLiveSourcesCmd.Flags().String("source-location-name", "", "The name of the source location associated with this Live Sources list.")
		mediatailor_listLiveSourcesCmd.MarkFlagRequired("source-location-name")
	})
	mediatailorCmd.AddCommand(mediatailor_listLiveSourcesCmd)
}
