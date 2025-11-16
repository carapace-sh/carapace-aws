package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmPricingCalculatorCmd = &cobra.Command{
	Use:   "bcm-pricing-calculator",
	Short: "You can use the Pricing Calculator API to programmatically create estimates for your planned cloud use.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmPricingCalculatorCmd).Standalone()

	rootCmd.AddCommand(bcmPricingCalculatorCmd)
}
