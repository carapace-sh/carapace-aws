package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_batchPutGeofenceCmd = &cobra.Command{
	Use:   "batch-put-geofence",
	Short: "A batch request for storing geofence geometries into a given geofence collection, or updates the geometry of an existing geofence if a geofence ID is included in the request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_batchPutGeofenceCmd).Standalone()

	location_batchPutGeofenceCmd.Flags().String("collection-name", "", "The geofence collection storing the geofences.")
	location_batchPutGeofenceCmd.Flags().String("entries", "", "The batch of geofences to be stored in a geofence collection.")
	location_batchPutGeofenceCmd.MarkFlagRequired("collection-name")
	location_batchPutGeofenceCmd.MarkFlagRequired("entries")
	locationCmd.AddCommand(location_batchPutGeofenceCmd)
}
