package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_listGeofencesCmd = &cobra.Command{
	Use:   "list-geofences",
	Short: "Lists geofences stored in a given geofence collection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_listGeofencesCmd).Standalone()

	location_listGeofencesCmd.Flags().String("collection-name", "", "The name of the geofence collection storing the list of geofences.")
	location_listGeofencesCmd.Flags().String("max-results", "", "An optional limit for the number of geofences returned in a single call.")
	location_listGeofencesCmd.Flags().String("next-token", "", "The pagination token specifying which page of results to return in the response.")
	location_listGeofencesCmd.MarkFlagRequired("collection-name")
	locationCmd.AddCommand(location_listGeofencesCmd)
}
