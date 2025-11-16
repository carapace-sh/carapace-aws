package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_updateGeofenceCollectionCmd = &cobra.Command{
	Use:   "update-geofence-collection",
	Short: "Updates the specified properties of a given geofence collection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_updateGeofenceCollectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(location_updateGeofenceCollectionCmd).Standalone()

		location_updateGeofenceCollectionCmd.Flags().String("collection-name", "", "The name of the geofence collection to update.")
		location_updateGeofenceCollectionCmd.Flags().String("description", "", "Updates the description for the geofence collection.")
		location_updateGeofenceCollectionCmd.Flags().String("pricing-plan", "", "No longer used.")
		location_updateGeofenceCollectionCmd.Flags().String("pricing-plan-data-source", "", "This parameter is no longer used.")
		location_updateGeofenceCollectionCmd.MarkFlagRequired("collection-name")
	})
	locationCmd.AddCommand(location_updateGeofenceCollectionCmd)
}
