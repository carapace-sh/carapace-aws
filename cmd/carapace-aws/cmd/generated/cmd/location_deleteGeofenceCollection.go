package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_deleteGeofenceCollectionCmd = &cobra.Command{
	Use:   "delete-geofence-collection",
	Short: "Deletes a geofence collection from your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_deleteGeofenceCollectionCmd).Standalone()

	location_deleteGeofenceCollectionCmd.Flags().String("collection-name", "", "The name of the geofence collection to be deleted.")
	location_deleteGeofenceCollectionCmd.MarkFlagRequired("collection-name")
	locationCmd.AddCommand(location_deleteGeofenceCollectionCmd)
}
