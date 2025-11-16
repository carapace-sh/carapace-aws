package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_batchDeleteGeofenceCmd = &cobra.Command{
	Use:   "batch-delete-geofence",
	Short: "Deletes a batch of geofences from a geofence collection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_batchDeleteGeofenceCmd).Standalone()

	location_batchDeleteGeofenceCmd.Flags().String("collection-name", "", "The geofence collection storing the geofences to be deleted.")
	location_batchDeleteGeofenceCmd.Flags().String("geofence-ids", "", "The batch of geofences to be deleted.")
	location_batchDeleteGeofenceCmd.MarkFlagRequired("collection-name")
	location_batchDeleteGeofenceCmd.MarkFlagRequired("geofence-ids")
	locationCmd.AddCommand(location_batchDeleteGeofenceCmd)
}
