package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_describeRouteCalculatorCmd = &cobra.Command{
	Use:   "describe-route-calculator",
	Short: "This operation is no longer current and may be deprecated in the future.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_describeRouteCalculatorCmd).Standalone()

	location_describeRouteCalculatorCmd.Flags().String("calculator-name", "", "The name of the route calculator resource.")
	location_describeRouteCalculatorCmd.MarkFlagRequired("calculator-name")
	locationCmd.AddCommand(location_describeRouteCalculatorCmd)
}
