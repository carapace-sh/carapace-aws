package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_listGeofenceCollectionsCmd = &cobra.Command{
	Use:   "list-geofence-collections",
	Short: "Lists geofence collections in your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_listGeofenceCollectionsCmd).Standalone()

	location_listGeofenceCollectionsCmd.Flags().String("max-results", "", "An optional limit for the number of resources returned in a single call.")
	location_listGeofenceCollectionsCmd.Flags().String("next-token", "", "The pagination token specifying which page of results to return in the response.")
	locationCmd.AddCommand(location_listGeofenceCollectionsCmd)
}
