package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_calculateRouteCmd = &cobra.Command{
	Use:   "calculate-route",
	Short: "This operation is no longer current and may be deprecated in the future.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_calculateRouteCmd).Standalone()

	location_calculateRouteCmd.Flags().String("arrival-time", "", "Specifies the desired time of arrival.")
	location_calculateRouteCmd.Flags().String("calculator-name", "", "The name of the route calculator resource that you want to use to calculate the route.")
	location_calculateRouteCmd.Flags().String("car-mode-options", "", "Specifies route preferences when traveling by `Car`, such as avoiding routes that use ferries or tolls.")
	location_calculateRouteCmd.Flags().String("depart-now", "", "Sets the time of departure as the current time.")
	location_calculateRouteCmd.Flags().String("departure-position", "", "The start position for the route.")
	location_calculateRouteCmd.Flags().String("departure-time", "", "Specifies the desired time of departure.")
	location_calculateRouteCmd.Flags().String("destination-position", "", "The finish position for the route.")
	location_calculateRouteCmd.Flags().String("distance-unit", "", "Set the unit system to specify the distance.")
	location_calculateRouteCmd.Flags().String("include-leg-geometry", "", "Set to include the geometry details in the result for each path between a pair of positions.")
	location_calculateRouteCmd.Flags().String("key", "", "The optional [API key](https://docs.aws.amazon.com/location/previous/developerguide/using-apikeys.html) to authorize the request.")
	location_calculateRouteCmd.Flags().String("optimize-for", "", "Specifies the distance to optimize for when calculating a route.")
	location_calculateRouteCmd.Flags().String("travel-mode", "", "Specifies the mode of transport when calculating a route.")
	location_calculateRouteCmd.Flags().String("truck-mode-options", "", "Specifies route preferences when traveling by `Truck`, such as avoiding routes that use ferries or tolls, and truck specifications to consider when choosing an optimal road.")
	location_calculateRouteCmd.Flags().String("waypoint-positions", "", "Specifies an ordered list of up to 23 intermediate positions to include along a route between the departure position and destination position.")
	location_calculateRouteCmd.MarkFlagRequired("calculator-name")
	location_calculateRouteCmd.MarkFlagRequired("departure-position")
	location_calculateRouteCmd.MarkFlagRequired("destination-position")
	locationCmd.AddCommand(location_calculateRouteCmd)
}
