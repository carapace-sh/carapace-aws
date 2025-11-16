package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var geoRoutes_calculateRoutesCmd = &cobra.Command{
	Use:   "calculate-routes",
	Short: "`CalculateRoutes` computes routes given the following required parameters: `Origin` and `Destination`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(geoRoutes_calculateRoutesCmd).Standalone()

	geoRoutes_calculateRoutesCmd.Flags().String("allow", "", "Features that are allowed while calculating a route.")
	geoRoutes_calculateRoutesCmd.Flags().String("arrival-time", "", "Time of arrival at the destination.")
	geoRoutes_calculateRoutesCmd.Flags().String("avoid", "", "Features that are avoided while calculating a route.")
	geoRoutes_calculateRoutesCmd.Flags().String("depart-now", "", "Uses the current time as the time of departure.")
	geoRoutes_calculateRoutesCmd.Flags().String("departure-time", "", "Time of departure from thr origin.")
	geoRoutes_calculateRoutesCmd.Flags().String("destination", "", "The final position for the route.")
	geoRoutes_calculateRoutesCmd.Flags().String("destination-options", "", "Destination related options.")
	geoRoutes_calculateRoutesCmd.Flags().String("driver", "", "Driver related options.")
	geoRoutes_calculateRoutesCmd.Flags().String("exclude", "", "Features to be strictly excluded while calculating the route.")
	geoRoutes_calculateRoutesCmd.Flags().String("instructions-measurement-system", "", "Measurement system to be used for instructions within steps in the response.")
	geoRoutes_calculateRoutesCmd.Flags().String("key", "", "Optional: The API key to be used for authorization.")
	geoRoutes_calculateRoutesCmd.Flags().String("languages", "", "List of languages for instructions within steps in the response.")
	geoRoutes_calculateRoutesCmd.Flags().String("leg-additional-features", "", "A list of optional additional parameters such as timezone that can be requested for each result.")
	geoRoutes_calculateRoutesCmd.Flags().String("leg-geometry-format", "", "Specifies the format of the geometry returned for each leg of the route.")
	geoRoutes_calculateRoutesCmd.Flags().String("max-alternatives", "", "Maximum number of alternative routes to be provided in the response, if available.")
	geoRoutes_calculateRoutesCmd.Flags().String("optimize-routing-for", "", "Specifies the optimization criteria for calculating a route.")
	geoRoutes_calculateRoutesCmd.Flags().String("origin", "", "The start position for the route.")
	geoRoutes_calculateRoutesCmd.Flags().String("origin-options", "", "Origin related options.")
	geoRoutes_calculateRoutesCmd.Flags().String("span-additional-features", "", "A list of optional features such as SpeedLimit that can be requested for a Span.")
	geoRoutes_calculateRoutesCmd.Flags().String("tolls", "", "Toll related options.")
	geoRoutes_calculateRoutesCmd.Flags().String("traffic", "", "Traffic related options.")
	geoRoutes_calculateRoutesCmd.Flags().String("travel-mode", "", "Specifies the mode of transport when calculating a route.")
	geoRoutes_calculateRoutesCmd.Flags().String("travel-mode-options", "", "Travel mode related options for the provided travel mode.")
	geoRoutes_calculateRoutesCmd.Flags().String("travel-step-type", "", "Type of step returned by the response.")
	geoRoutes_calculateRoutesCmd.Flags().String("waypoints", "", "List of waypoints between the Origin and Destination.")
	geoRoutes_calculateRoutesCmd.MarkFlagRequired("destination")
	geoRoutes_calculateRoutesCmd.MarkFlagRequired("origin")
	geoRoutesCmd.AddCommand(geoRoutes_calculateRoutesCmd)
}
