package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmPricingCalculator_getPreferencesCmd = &cobra.Command{
	Use:   "get-preferences",
	Short: "Retrieves the current preferences for Pricing Calculator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmPricingCalculator_getPreferencesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bcmPricingCalculator_getPreferencesCmd).Standalone()

	})
	bcmPricingCalculatorCmd.AddCommand(bcmPricingCalculator_getPreferencesCmd)
}
