package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_putGeofenceCmd = &cobra.Command{
	Use:   "put-geofence",
	Short: "Stores a geofence geometry in a given geofence collection, or updates the geometry of an existing geofence if a geofence ID is included in the request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_putGeofenceCmd).Standalone()

	location_putGeofenceCmd.Flags().String("collection-name", "", "The geofence collection to store the geofence in.")
	location_putGeofenceCmd.Flags().String("geofence-id", "", "An identifier for the geofence.")
	location_putGeofenceCmd.Flags().String("geofence-properties", "", "Associates one of more properties with the geofence.")
	location_putGeofenceCmd.Flags().String("geometry", "", "Contains the details to specify the position of the geofence.")
	location_putGeofenceCmd.MarkFlagRequired("collection-name")
	location_putGeofenceCmd.MarkFlagRequired("geofence-id")
	location_putGeofenceCmd.MarkFlagRequired("geometry")
	locationCmd.AddCommand(location_putGeofenceCmd)
}
