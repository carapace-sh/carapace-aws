package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_createRouteCalculatorCmd = &cobra.Command{
	Use:   "create-route-calculator",
	Short: "This operation is no longer current and may be deprecated in the future.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_createRouteCalculatorCmd).Standalone()

	location_createRouteCalculatorCmd.Flags().String("calculator-name", "", "The name of the route calculator resource.")
	location_createRouteCalculatorCmd.Flags().String("data-source", "", "Specifies the data provider of traffic and road network data.")
	location_createRouteCalculatorCmd.Flags().String("description", "", "The optional description for the route calculator resource.")
	location_createRouteCalculatorCmd.Flags().String("pricing-plan", "", "No longer used.")
	location_createRouteCalculatorCmd.Flags().String("tags", "", "Applies one or more tags to the route calculator resource.")
	location_createRouteCalculatorCmd.MarkFlagRequired("calculator-name")
	location_createRouteCalculatorCmd.MarkFlagRequired("data-source")
	locationCmd.AddCommand(location_createRouteCalculatorCmd)
}
