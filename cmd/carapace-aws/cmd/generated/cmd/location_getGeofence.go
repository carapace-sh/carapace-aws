package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_getGeofenceCmd = &cobra.Command{
	Use:   "get-geofence",
	Short: "Retrieves the geofence details from a geofence collection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_getGeofenceCmd).Standalone()

	location_getGeofenceCmd.Flags().String("collection-name", "", "The geofence collection storing the target geofence.")
	location_getGeofenceCmd.Flags().String("geofence-id", "", "The geofence you're retrieving details for.")
	location_getGeofenceCmd.MarkFlagRequired("collection-name")
	location_getGeofenceCmd.MarkFlagRequired("geofence-id")
	locationCmd.AddCommand(location_getGeofenceCmd)
}
