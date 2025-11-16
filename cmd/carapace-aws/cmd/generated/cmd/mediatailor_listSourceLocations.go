package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediatailor_listSourceLocationsCmd = &cobra.Command{
	Use:   "list-source-locations",
	Short: "Lists the source locations for a channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediatailor_listSourceLocationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediatailor_listSourceLocationsCmd).Standalone()

		mediatailor_listSourceLocationsCmd.Flags().String("max-results", "", "The maximum number of source locations that you want MediaTailor to return in response to the current request.")
		mediatailor_listSourceLocationsCmd.Flags().String("next-token", "", "Pagination token returned by the list request when results exceed the maximum allowed.")
	})
	mediatailorCmd.AddCommand(mediatailor_listSourceLocationsCmd)
}
