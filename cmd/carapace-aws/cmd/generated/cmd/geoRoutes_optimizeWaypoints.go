package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var geoRoutes_optimizeWaypointsCmd = &cobra.Command{
	Use:   "optimize-waypoints",
	Short: "`OptimizeWaypoints` calculates the optimal order to travel between a set of waypoints to minimize either the travel time or the distance travelled during the journey, based on road network restrictions and the traffic pattern data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(geoRoutes_optimizeWaypointsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(geoRoutes_optimizeWaypointsCmd).Standalone()

		geoRoutes_optimizeWaypointsCmd.Flags().String("avoid", "", "Features that are avoided.")
		geoRoutes_optimizeWaypointsCmd.Flags().String("clustering", "", "Clustering allows you to specify how nearby waypoints can be clustered to improve the optimized sequence.")
		geoRoutes_optimizeWaypointsCmd.Flags().String("departure-time", "", "Departure time from the waypoint.")
		geoRoutes_optimizeWaypointsCmd.Flags().String("destination", "", "The final position for the route in the World Geodetic System (WGS 84) format: `[longitude, latitude]`.")
		geoRoutes_optimizeWaypointsCmd.Flags().String("destination-options", "", "Destination related options.")
		geoRoutes_optimizeWaypointsCmd.Flags().String("driver", "", "Driver related options.")
		geoRoutes_optimizeWaypointsCmd.Flags().String("exclude", "", "Features to be strictly excluded while calculating the route.")
		geoRoutes_optimizeWaypointsCmd.Flags().String("key", "", "Optional: The API key to be used for authorization.")
		geoRoutes_optimizeWaypointsCmd.Flags().String("optimize-sequencing-for", "", "Specifies the optimization criteria for the calculated sequence.")
		geoRoutes_optimizeWaypointsCmd.Flags().String("origin", "", "The start position for the route.")
		geoRoutes_optimizeWaypointsCmd.Flags().String("origin-options", "", "Origin related options.")
		geoRoutes_optimizeWaypointsCmd.Flags().String("traffic", "", "Traffic-related options.")
		geoRoutes_optimizeWaypointsCmd.Flags().String("travel-mode", "", "Specifies the mode of transport when calculating a route.")
		geoRoutes_optimizeWaypointsCmd.Flags().String("travel-mode-options", "", "Travel mode related options for the provided travel mode.")
		geoRoutes_optimizeWaypointsCmd.Flags().String("waypoints", "", "List of waypoints between the `Origin` and `Destination`.")
		geoRoutes_optimizeWaypointsCmd.MarkFlagRequired("origin")
	})
	geoRoutesCmd.AddCommand(geoRoutes_optimizeWaypointsCmd)
}
