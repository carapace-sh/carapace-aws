package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_forecastGeofenceEventsCmd = &cobra.Command{
	Use:   "forecast-geofence-events",
	Short: "This action forecasts future geofence events that are likely to occur within a specified time horizon if a device continues moving at its current speed.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_forecastGeofenceEventsCmd).Standalone()

	location_forecastGeofenceEventsCmd.Flags().String("collection-name", "", "The name of the geofence collection.")
	location_forecastGeofenceEventsCmd.Flags().String("device-state", "", "Represents the device's state, including its current position and speed.")
	location_forecastGeofenceEventsCmd.Flags().String("distance-unit", "", "The distance unit used for the `NearestDistance` property returned in a forecasted event.")
	location_forecastGeofenceEventsCmd.Flags().String("max-results", "", "An optional limit for the number of resources returned in a single call.")
	location_forecastGeofenceEventsCmd.Flags().String("next-token", "", "The pagination token specifying which page of results to return in the response.")
	location_forecastGeofenceEventsCmd.Flags().String("speed-unit", "", "The speed unit for the device captured by the device state.")
	location_forecastGeofenceEventsCmd.Flags().String("time-horizon-minutes", "", "The forward-looking time window for forecasting, specified in minutes.")
	location_forecastGeofenceEventsCmd.MarkFlagRequired("collection-name")
	location_forecastGeofenceEventsCmd.MarkFlagRequired("device-state")
	locationCmd.AddCommand(location_forecastGeofenceEventsCmd)
}
