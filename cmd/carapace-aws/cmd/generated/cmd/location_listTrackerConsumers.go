package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_listTrackerConsumersCmd = &cobra.Command{
	Use:   "list-tracker-consumers",
	Short: "Lists geofence collections currently associated to the given tracker resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_listTrackerConsumersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(location_listTrackerConsumersCmd).Standalone()

		location_listTrackerConsumersCmd.Flags().String("max-results", "", "An optional limit for the number of resources returned in a single call.")
		location_listTrackerConsumersCmd.Flags().String("next-token", "", "The pagination token specifying which page of results to return in the response.")
		location_listTrackerConsumersCmd.Flags().String("tracker-name", "", "The tracker resource whose associated geofence collections you want to list.")
		location_listTrackerConsumersCmd.MarkFlagRequired("tracker-name")
	})
	locationCmd.AddCommand(location_listTrackerConsumersCmd)
}
