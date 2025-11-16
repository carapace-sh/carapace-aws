package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_createGeofenceCollectionCmd = &cobra.Command{
	Use:   "create-geofence-collection",
	Short: "Creates a geofence collection, which manages and stores geofences.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_createGeofenceCollectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(location_createGeofenceCollectionCmd).Standalone()

		location_createGeofenceCollectionCmd.Flags().String("collection-name", "", "A custom name for the geofence collection.")
		location_createGeofenceCollectionCmd.Flags().String("description", "", "An optional description for the geofence collection.")
		location_createGeofenceCollectionCmd.Flags().String("kms-key-id", "", "A key identifier for an [Amazon Web Services KMS customer managed key](https://docs.aws.amazon.com/kms/latest/developerguide/create-keys.html).")
		location_createGeofenceCollectionCmd.Flags().String("pricing-plan", "", "No longer used.")
		location_createGeofenceCollectionCmd.Flags().String("pricing-plan-data-source", "", "This parameter is no longer used.")
		location_createGeofenceCollectionCmd.Flags().String("tags", "", "Applies one or more tags to the geofence collection.")
		location_createGeofenceCollectionCmd.MarkFlagRequired("collection-name")
	})
	locationCmd.AddCommand(location_createGeofenceCollectionCmd)
}
