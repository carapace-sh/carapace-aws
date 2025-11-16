package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var geoRoutes_calculateRouteMatrixCmd = &cobra.Command{
	Use:   "calculate-route-matrix",
	Short: "Use `CalculateRouteMatrix` to compute results for all pairs of Origins to Destinations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(geoRoutes_calculateRouteMatrixCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(geoRoutes_calculateRouteMatrixCmd).Standalone()

		geoRoutes_calculateRouteMatrixCmd.Flags().String("allow", "", "Features that are allowed while calculating a route.")
		geoRoutes_calculateRouteMatrixCmd.Flags().String("avoid", "", "Features that are avoided while calculating a route.")
		geoRoutes_calculateRouteMatrixCmd.Flags().String("depart-now", "", "Uses the current time as the time of departure.")
		geoRoutes_calculateRouteMatrixCmd.Flags().String("departure-time", "", "Time of departure from thr origin.")
		geoRoutes_calculateRouteMatrixCmd.Flags().String("destinations", "", "List of destinations for the route.")
		geoRoutes_calculateRouteMatrixCmd.Flags().String("exclude", "", "Features to be strictly excluded while calculating the route.")
		geoRoutes_calculateRouteMatrixCmd.Flags().String("key", "", "Optional: The API key to be used for authorization.")
		geoRoutes_calculateRouteMatrixCmd.Flags().String("optimize-routing-for", "", "Specifies the optimization criteria for calculating a route.")
		geoRoutes_calculateRouteMatrixCmd.Flags().String("origins", "", "The position in longitude and latitude for the origin.")
		geoRoutes_calculateRouteMatrixCmd.Flags().String("routing-boundary", "", "Boundary within which the matrix is to be calculated.")
		geoRoutes_calculateRouteMatrixCmd.Flags().String("traffic", "", "Traffic related options.")
		geoRoutes_calculateRouteMatrixCmd.Flags().String("travel-mode", "", "Specifies the mode of transport when calculating a route.")
		geoRoutes_calculateRouteMatrixCmd.Flags().String("travel-mode-options", "", "Travel mode related options for the provided travel mode.")
		geoRoutes_calculateRouteMatrixCmd.MarkFlagRequired("destinations")
		geoRoutes_calculateRouteMatrixCmd.MarkFlagRequired("origins")
		geoRoutes_calculateRouteMatrixCmd.MarkFlagRequired("routing-boundary")
	})
	geoRoutesCmd.AddCommand(geoRoutes_calculateRouteMatrixCmd)
}
