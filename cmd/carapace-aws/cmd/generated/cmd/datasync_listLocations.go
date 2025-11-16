package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_listLocationsCmd = &cobra.Command{
	Use:   "list-locations",
	Short: "Returns a list of source and destination locations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_listLocationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datasync_listLocationsCmd).Standalone()

		datasync_listLocationsCmd.Flags().String("filters", "", "You can use API filters to narrow down the list of resources returned by `ListLocations`.")
		datasync_listLocationsCmd.Flags().String("max-results", "", "The maximum number of locations to return.")
		datasync_listLocationsCmd.Flags().String("next-token", "", "An opaque string that indicates the position at which to begin the next list of locations.")
	})
	datasyncCmd.AddCommand(datasync_listLocationsCmd)
}
