package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmPricingCalculator_updatePreferencesCmd = &cobra.Command{
	Use:   "update-preferences",
	Short: "Updates the preferences for Pricing Calculator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmPricingCalculator_updatePreferencesCmd).Standalone()

	bcmPricingCalculator_updatePreferencesCmd.Flags().String("management-account-rate-type-selections", "", "The updated preferred rate types for the management account.")
	bcmPricingCalculator_updatePreferencesCmd.Flags().String("member-account-rate-type-selections", "", "The updated preferred rate types for member accounts.")
	bcmPricingCalculator_updatePreferencesCmd.Flags().String("standalone-account-rate-type-selections", "", "The updated preferred rate types for a standalone account.")
	bcmPricingCalculatorCmd.AddCommand(bcmPricingCalculator_updatePreferencesCmd)
}
