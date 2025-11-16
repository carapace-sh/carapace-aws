package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_batchEvaluateGeofencesCmd = &cobra.Command{
	Use:   "batch-evaluate-geofences",
	Short: "Evaluates device positions against the geofence geometries from a given geofence collection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_batchEvaluateGeofencesCmd).Standalone()

	location_batchEvaluateGeofencesCmd.Flags().String("collection-name", "", "The geofence collection used in evaluating the position of devices against its geofences.")
	location_batchEvaluateGeofencesCmd.Flags().String("device-position-updates", "", "Contains device details for each device to be evaluated against the given geofence collection.")
	location_batchEvaluateGeofencesCmd.MarkFlagRequired("collection-name")
	location_batchEvaluateGeofencesCmd.MarkFlagRequired("device-position-updates")
	locationCmd.AddCommand(location_batchEvaluateGeofencesCmd)
}
