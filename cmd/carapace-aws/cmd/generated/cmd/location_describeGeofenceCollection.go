package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_describeGeofenceCollectionCmd = &cobra.Command{
	Use:   "describe-geofence-collection",
	Short: "Retrieves the geofence collection details.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_describeGeofenceCollectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(location_describeGeofenceCollectionCmd).Standalone()

		location_describeGeofenceCollectionCmd.Flags().String("collection-name", "", "The name of the geofence collection.")
		location_describeGeofenceCollectionCmd.MarkFlagRequired("collection-name")
	})
	locationCmd.AddCommand(location_describeGeofenceCollectionCmd)
}
