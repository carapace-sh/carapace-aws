package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_calculateRouteMatrixCmd = &cobra.Command{
	Use:   "calculate-route-matrix",
	Short: "This operation is no longer current and may be deprecated in the future.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_calculateRouteMatrixCmd).Standalone()

	location_calculateRouteMatrixCmd.Flags().String("calculator-name", "", "The name of the route calculator resource that you want to use to calculate the route matrix.")
	location_calculateRouteMatrixCmd.Flags().String("car-mode-options", "", "Specifies route preferences when traveling by `Car`, such as avoiding routes that use ferries or tolls.")
	location_calculateRouteMatrixCmd.Flags().String("depart-now", "", "Sets the time of departure as the current time.")
	location_calculateRouteMatrixCmd.Flags().String("departure-positions", "", "The list of departure (origin) positions for the route matrix.")
	location_calculateRouteMatrixCmd.Flags().String("departure-time", "", "Specifies the desired time of departure.")
	location_calculateRouteMatrixCmd.Flags().String("destination-positions", "", "The list of destination positions for the route matrix.")
	location_calculateRouteMatrixCmd.Flags().String("distance-unit", "", "Set the unit system to specify the distance.")
	location_calculateRouteMatrixCmd.Flags().String("key", "", "The optional [API key](https://docs.aws.amazon.com/location/previous/developerguide/using-apikeys.html) to authorize the request.")
	location_calculateRouteMatrixCmd.Flags().String("travel-mode", "", "Specifies the mode of transport when calculating a route.")
	location_calculateRouteMatrixCmd.Flags().String("truck-mode-options", "", "Specifies route preferences when traveling by `Truck`, such as avoiding routes that use ferries or tolls, and truck specifications to consider when choosing an optimal road.")
	location_calculateRouteMatrixCmd.MarkFlagRequired("calculator-name")
	location_calculateRouteMatrixCmd.MarkFlagRequired("departure-positions")
	location_calculateRouteMatrixCmd.MarkFlagRequired("destination-positions")
	locationCmd.AddCommand(location_calculateRouteMatrixCmd)
}
