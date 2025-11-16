package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_updateRouteCalculatorCmd = &cobra.Command{
	Use:   "update-route-calculator",
	Short: "This operation is no longer current and may be deprecated in the future.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_updateRouteCalculatorCmd).Standalone()

	location_updateRouteCalculatorCmd.Flags().String("calculator-name", "", "The name of the route calculator resource to update.")
	location_updateRouteCalculatorCmd.Flags().String("description", "", "Updates the description for the route calculator resource.")
	location_updateRouteCalculatorCmd.Flags().String("pricing-plan", "", "No longer used.")
	location_updateRouteCalculatorCmd.MarkFlagRequired("calculator-name")
	locationCmd.AddCommand(location_updateRouteCalculatorCmd)
}
