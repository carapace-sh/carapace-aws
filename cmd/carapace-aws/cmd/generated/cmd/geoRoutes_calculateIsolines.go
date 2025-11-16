package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var geoRoutes_calculateIsolinesCmd = &cobra.Command{
	Use:   "calculate-isolines",
	Short: "Use the `CalculateIsolines` action to find service areas that can be reached in a given threshold of time, distance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(geoRoutes_calculateIsolinesCmd).Standalone()

	geoRoutes_calculateIsolinesCmd.Flags().String("allow", "", "Features that are allowed while calculating an isoline.")
	geoRoutes_calculateIsolinesCmd.Flags().String("arrival-time", "", "Time of arrival at the destination.")
	geoRoutes_calculateIsolinesCmd.Flags().String("avoid", "", "Features that are avoided while calculating a route.")
	geoRoutes_calculateIsolinesCmd.Flags().String("depart-now", "", "Uses the current time as the time of departure.")
	geoRoutes_calculateIsolinesCmd.Flags().String("departure-time", "", "Time of departure from thr origin.")
	geoRoutes_calculateIsolinesCmd.Flags().String("destination", "", "The final position for the route.")
	geoRoutes_calculateIsolinesCmd.Flags().String("destination-options", "", "Destination related options.")
	geoRoutes_calculateIsolinesCmd.Flags().String("isoline-geometry-format", "", "The format of the returned IsolineGeometry.")
	geoRoutes_calculateIsolinesCmd.Flags().String("isoline-granularity", "", "Defines the granularity of the returned Isoline.")
	geoRoutes_calculateIsolinesCmd.Flags().String("key", "", "Optional: The API key to be used for authorization.")
	geoRoutes_calculateIsolinesCmd.Flags().String("optimize-isoline-for", "", "Specifies the optimization criteria for when calculating an isoline.")
	geoRoutes_calculateIsolinesCmd.Flags().String("optimize-routing-for", "", "Specifies the optimization criteria for calculating a route.")
	geoRoutes_calculateIsolinesCmd.Flags().String("origin", "", "The start position for the route.")
	geoRoutes_calculateIsolinesCmd.Flags().String("origin-options", "", "Origin related options.")
	geoRoutes_calculateIsolinesCmd.Flags().String("thresholds", "", "Threshold to be used for the isoline calculation.")
	geoRoutes_calculateIsolinesCmd.Flags().String("traffic", "", "Traffic related options.")
	geoRoutes_calculateIsolinesCmd.Flags().String("travel-mode", "", "Specifies the mode of transport when calculating a route.")
	geoRoutes_calculateIsolinesCmd.Flags().String("travel-mode-options", "", "Travel mode related options for the provided travel mode.")
	geoRoutes_calculateIsolinesCmd.MarkFlagRequired("thresholds")
	geoRoutesCmd.AddCommand(geoRoutes_calculateIsolinesCmd)
}
