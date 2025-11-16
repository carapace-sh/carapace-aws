package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_deleteRouteCalculatorCmd = &cobra.Command{
	Use:   "delete-route-calculator",
	Short: "This operation is no longer current and may be deprecated in the future.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_deleteRouteCalculatorCmd).Standalone()

	location_deleteRouteCalculatorCmd.Flags().String("calculator-name", "", "The name of the route calculator resource to be deleted.")
	location_deleteRouteCalculatorCmd.MarkFlagRequired("calculator-name")
	locationCmd.AddCommand(location_deleteRouteCalculatorCmd)
}
