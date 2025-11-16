package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var geoRoutes_snapToRoadsCmd = &cobra.Command{
	Use:   "snap-to-roads",
	Short: "`SnapToRoads` matches GPS trace to roads most likely traveled on.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(geoRoutes_snapToRoadsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(geoRoutes_snapToRoadsCmd).Standalone()

		geoRoutes_snapToRoadsCmd.Flags().String("key", "", "Optional: The API key to be used for authorization.")
		geoRoutes_snapToRoadsCmd.Flags().String("snap-radius", "", "The radius around the provided tracepoint that is considered for snapping.")
		geoRoutes_snapToRoadsCmd.Flags().String("snapped-geometry-format", "", "Chooses what the returned SnappedGeometry format should be.")
		geoRoutes_snapToRoadsCmd.Flags().String("trace-points", "", "List of trace points to be snapped onto the road network.")
		geoRoutes_snapToRoadsCmd.Flags().String("travel-mode", "", "Specifies the mode of transport when calculating a route.")
		geoRoutes_snapToRoadsCmd.Flags().String("travel-mode-options", "", "Travel mode related options for the provided travel mode.")
		geoRoutes_snapToRoadsCmd.MarkFlagRequired("trace-points")
	})
	geoRoutesCmd.AddCommand(geoRoutes_snapToRoadsCmd)
}
